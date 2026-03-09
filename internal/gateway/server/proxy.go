/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	httptransport "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/metadata"

	"github.com/origadmin/runtime"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/middleware"
	runtimehttp "github.com/origadmin/runtime/service/transport/http"
	"origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/objectstore"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

type ObjectStoreProxy struct {
	app           *runtime.App
	logger        *log.Helper
	objClient     objectstore.ObjectStoreServiceClient
	fileClient    filemanager.FileManagerServiceClient
	objHTTPClient *httptransport.Client
}

// NewObjectStoreProxy creates a new proxy for object store and file manager services.
func NewObjectStoreProxy(app *runtime.App, bootstrap *confpb.Bootstrap) (*ObjectStoreProxy, error) {
	objConn, err := grpcclient.NewConn(app, bootstrap, "objectstore")
	if err != nil {
		return nil, fmt.Errorf("failed to create objectstore grpc client: %w", err)
	}
	objClient := objectstore.NewObjectStoreServiceClient(objConn)

	fileConn, err := grpcclient.NewConn(app, bootstrap, "filemanager")
	if err != nil {
		return nil, fmt.Errorf("failed to create filemanager grpc client: %w", err)
	}
	fileClient := filemanager.NewFileManagerServiceClient(fileConn)

	var clientConfig *transportv1.Client
	name := "objectstore"
	convention := fmt.Sprintf("origadmin.service.%s.client.http", name)

	clients := bootstrap.GetClients()
	if clients != nil {
		for _, cli := range clients.Configs {
			if cli.GetHttp() == nil {
				continue
			}
			if cli.Name == name || cli.Name == convention {
				clientConfig = cli
				break
			}
		}
	}

	if clientConfig == nil {
		return nil, fmt.Errorf("HTTP client config not found for service: %s", name)
	}

	discoveries, err := app.Discoveries()
	if err != nil {
		return nil, fmt.Errorf("failed to get discoveries: %w", err)
	}

	// Get client middlewares from container (ClientScope)
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithScope(runtime.ClientScope))
	mws, err := middleware.GetMiddlewares(app.Context(), h)
	if err != nil {
		return nil, fmt.Errorf("failed to get client middlewares: %w", err)
	}

	endpoint := clientConfig.GetHttp().GetEndpoint()
	objHTTPClient, err := runtimehttp.NewClient(
		app.Context(),
		clientConfig.GetHttp(),
		&runtimehttp.ClientOptions{
			Discoveries:       discoveries,
			ClientMiddlewares: mws,
			ClientOptions: []httptransport.ClientOption{
				httptransport.WithEndpoint(endpoint),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &ObjectStoreProxy{
		app: app, logger: log.NewHelper(log.With(app.Logger(), "module", "gateway.proxy")),
		objClient:     objClient,
		fileClient:    fileClient,
		objHTTPClient: objHTTPClient,
	}, nil
}

func (p *ObjectStoreProxy) RegisterHandlers(srv *httptransport.Server) {
	r := srv.Route("/")

	r.POST("/storage/upload", p.handleUpload)
	r.GET("/storage/download/{id}", p.handleDownload)
	r.DELETE("/storage/delete/{id}", p.handleDelete)
	r.POST("/storage/multipart/initiate", p.handleInitiateMultipart)
	r.GET("/storage/multipart/{upload_id}/parts", p.handleListParts)
	r.PUT("/storage/multipart/{upload_id}/parts/{part_number}", p.handleUploadPart)
	r.POST("/storage/multipart/{upload_id}/complete", p.handleCompleteMultipart)
	r.DELETE("/storage/multipart/{upload_id}", p.handleAbortMultipart)

	p.logger.Infof("[Debug] Walking handlers to verify registration:")
	srv.WalkHandle(func(method, path string, _ http.HandlerFunc) {
		p.logger.Infof("Registered http handler: %s %s", method, path)
	})
}

func (p *ObjectStoreProxy) handleUploadPart(ctx httptransport.Context) error {
	uID, pNumStr := ctx.Vars().Get("upload_id"), ctx.Vars().Get("part_number")
	pNum, err := strconv.Atoi(pNumStr)
	if err != nil || pNum <= 0 {
		return ctx.Result(http.StatusBadRequest, kerrors.BadRequest("INVALID_PART_NUMBER", "part_number must be greater than 0"))
	}
	p.logger.Infof("[Debug] handleUploadPart start: ID=%s, Part=%d", uID, pNum)

	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.fileClient.GetMultipartUploadUrl(ctx, req.(*filemanager.GetMultipartUploadUrlRequest))
	})
	uResp, err := h(ctx, &filemanager.GetMultipartUploadUrlRequest{UploadId: uID, PartNumber: int32(pNum)})
	if err != nil {
		return ctx.Result(http.StatusInternalServerError, err)
	}
	target := uResp.(*filemanager.GetMultipartUploadUrlResponse).UploadUrl

	// Resolve discovery:/// to http:// for proxying
	finalURL := target
	if u, err := url.Parse(target); err == nil && u.Scheme == "discovery" {
		u.Scheme = "http"
		if u.Host == "" {
			// Handle discovery:///service.name/path where service.name is part of the path
			path := u.Path
			if strings.HasPrefix(path, "//") {
				path = strings.TrimPrefix(path, "//")
			}
			if strings.HasPrefix(path, "/") {
				ps := strings.SplitN(strings.TrimPrefix(path, "/"), "/", 2)
				u.Host = ps[0]
				if len(ps) > 1 {
					u.Path = "/" + ps[1]
				} else {
					u.Path = ""
				}
			}
		}
		finalURL = u.String()
	}
	p.logger.Infof("[Debug] Proxying to: %s", finalURL)

	req, err := http.NewRequestWithContext(ctx, "PUT", finalURL, ctx.Request().Body)
	if err != nil {
		return ctx.Result(http.StatusInternalServerError, err)
	}
	req.ContentLength = ctx.Request().ContentLength

	// CRITICAL: Pass through ALL headers from incoming request (Auth, Content-Type, etc.)
	for k, vv := range ctx.Request().Header {
		if k == "Host" || k == "Content-Length" || k == "Connection" {
			continue
		}
		for _, v := range vv {
			req.Header.Add(k, v)
		}
	}
	// Explicitly ensure Content-Type is preserved
	if ct := ctx.Request().Header.Get("Content-Type"); ct != "" {
		req.Header.Set("Content-Type", ct)
	} else if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	// Use system discovery-aware client
	resp, err := p.objHTTPClient.Do(req)
	if err != nil {
		return ctx.Result(http.StatusInternalServerError, fmt.Errorf("backend call failed: %w", err))
	}
	defer resp.Body.Close()

	// Drain incoming body
	_, _ = io.Copy(io.Discard, ctx.Request().Body)

	// Copy headers back to client (important for ETag)
	for k, vv := range resp.Header {
		if strings.HasPrefix(k, "Access-Control-") || k == "Content-Length" {
			continue
		}
		for _, v := range vv {
			ctx.Response().Header().Add(k, v)
		}
	}
	// Force expose ETag header for browser JS
	ctx.Response().Header().Set("Access-Control-Expose-Headers", "ETag")

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return ctx.Result(resp.StatusCode, kerrors.New(resp.StatusCode, "BACKEND_ERROR", string(body)))
	}

	return ctx.Result(resp.StatusCode, nil)
}

func (p *ObjectStoreProxy) handleListParts(ctx httptransport.Context) error {
	uID := ctx.Vars().Get("upload_id")
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.objClient.ListParts(ctx, req.(*objectstore.ListPartsRequest))
	})
	resp, err := h(ctx, &objectstore.ListPartsRequest{UploadId: uID})
	if err != nil {
		return err
	}
	return ctx.Result(200, resp)
}

func (p *ObjectStoreProxy) handleAbortMultipart(ctx httptransport.Context) error {
	uID := ctx.Vars().Get("upload_id")
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.objClient.AbortMultipartUpload(ctx, req.(*objectstore.AbortMultipartUploadRequest))
	})
	_, err := h(ctx, &objectstore.AbortMultipartUploadRequest{UploadId: uID})
	return err
}

func (p *ObjectStoreProxy) handleUpload(ctx httptransport.Context) error {
	_ = ctx.Request().ParseMultipartForm(32 << 20)
	file, header, err := ctx.Request().FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()
	fb, _ := io.ReadAll(file)
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.fileClient.UploadFile(ctx, req.(*filemanager.UploadFileRequest))
	})
	out, err := h(ctx, &filemanager.UploadFileRequest{Name: header.Filename, Data: fb, Visibility: "private"})
	if err != nil {
		return err
	}
	resp := out.(*filemanager.UploadFileResponse)
	return ctx.Result(200, &UploadFileResponse{ID: resp.FileMetadata.Id, Name: resp.FileMetadata.Name, Size: resp.FileMetadata.Size})
}

func (p *ObjectStoreProxy) handleInitiateMultipart(ctx httptransport.Context) error {
	var in struct {
		Name        string
		Size        int64
		ContentType string
		Visibility  string
	}
	if err := ctx.Bind(&in); err != nil {
		return err
	}
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-file-visibility", in.Visibility)
		return p.fileClient.InitiateMultipartUpload(ctx, req.(*filemanager.InitiateMultipartUploadRequest))
	})
	resp, err := h(ctx, &filemanager.InitiateMultipartUploadRequest{Name: in.Name, Size: in.Size, ContentType: in.ContentType, Visibility: in.Visibility})
	if err != nil {
		return err
	}
	return ctx.Result(200, resp)
}

func (p *ObjectStoreProxy) handleCompleteMultipart(ctx httptransport.Context) error {
	uID := ctx.Vars().Get("upload_id")
	var in struct {
		Parts []struct {
			PartNumber int32  `json:"part_number"`
			PartNum    int32  `json:"partNumber"` // Support camelCase from JS
			Etag       string `json:"etag"`
		} `json:"parts"`
		Sha256 string `json:"sha256"`
	}
	if err := ctx.Bind(&in); err != nil {
		return ctx.Result(http.StatusBadRequest, kerrors.BadRequest("INVALID_REQUEST", err.Error()))
	}

	if len(in.Parts) == 0 {
		return ctx.Result(http.StatusBadRequest, kerrors.BadRequest("EMPTY_PARTS", "multipart upload requires at least one part"))
	}

	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-file-sha256", in.Sha256)
		return p.fileClient.CompleteMultipartUpload(ctx, req.(*filemanager.CompleteMultipartUploadRequest))
	})
	protoParts := make([]*filemanager.PartInfo, len(in.Parts))
	for i, part := range in.Parts {
		pNum := part.PartNumber
		if pNum == 0 {
			pNum = part.PartNum
		}
		protoParts[i] = &filemanager.PartInfo{PartNumber: pNum, Etag: part.Etag}
	}
	resp, err := h(ctx, &filemanager.CompleteMultipartUploadRequest{
		UploadId: uID,
		Parts:    protoParts,
		Sha256:   in.Sha256,
	})
	if err != nil {
		return err
	}
	return ctx.Result(200, resp)
}

type downloadCarrier struct {
	Body *httpbody.HttpBody
	Meta *filemanager.GetFileResponse
}

func (p *ObjectStoreProxy) handleDownload(ctx httptransport.Context) error {
	id, _ := strconv.ParseInt(ctx.Vars().Get("id"), 10, 64)
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		fileResp, err := p.fileClient.GetFile(ctx, &filemanager.GetFileRequest{Id: id})
		if err != nil {
			return nil, err
		}
		body, err := p.objClient.DownloadObject(ctx, &objectstore.DownloadObjectRequest{Id: fileResp.FileMetadata.ObjectId})
		return &downloadCarrier{Body: body, Meta: fileResp}, err
	})
	out, err := h(ctx, nil)
	if err != nil {
		return err
	}
	data := out.(*downloadCarrier)
	ct := data.Body.ContentType
	if ct == "" || ct == "application/octet-stream" {
		ct = data.Meta.FileMetadata.MimeType
	}
	ctx.Response().Header().Set("Content-Type", ct)
	_, err = ctx.Response().Write(data.Body.Data)
	return err
}

func (p *ObjectStoreProxy) handleDelete(ctx httptransport.Context) error {
	id, _ := strconv.ParseInt(ctx.Vars().Get("id"), 10, 64)
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		_, err := p.fileClient.DeleteFile(ctx, &filemanager.DeleteFileRequest{Id: id})
		return &DeleteFileResponse{Status: "deleted"}, err
	})
	out, err := h(ctx, nil)
	if err != nil {
		return err
	}
	return ctx.Result(200, out)
}

type UploadFileResponse struct {
	ID   int64
	Name string
	Size int64
}
type DeleteFileResponse struct{ Status string }
