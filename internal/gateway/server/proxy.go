/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/registry"
	httptransport "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/metadata"

	"github.com/origadmin/runtime"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/container"
	"origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

// ObjectStoreProxy handles reverse proxying to the objectstore service using service discovery.
type ObjectStoreProxy struct {
	app       *runtime.App
	discovery registry.Discovery
	logger    *log.Helper

	// gRPC Clients
	objClient  objectstore.ObjectStoreServiceClient
	fileClient filemanager.FileManagerServiceClient

	// HTTP Client for OBS using discovery and middleware
	objHttpClient *httptransport.Client
}

// NewObjectStoreProxy creates a new ObjectStoreProxy.
func NewObjectStoreProxy(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*ObjectStoreProxy, error) {
	objConn, err := grpcclient.NewConn(app, bootstrap, "objectstore", middlewareProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create objectstore grpc client: %w", err)
	}
	objClient := objectstore.NewObjectStoreServiceClient(objConn)

	fileConn, err := grpcclient.NewConn(app, bootstrap, "filemanager", middlewareProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create filemanager grpc client: %w", err)
	}
	fileClient := filemanager.NewFileManagerServiceClient(fileConn)

	var clientConfig *transportv1.Client
	name := "objectstore"
	convention := fmt.Sprintf("origadmin.service.%s.client.http", name)

	clients := bootstrap.Clients()
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

	registryProvider, err := app.RegistryProvider()
	if err != nil {
		return nil, err
	}
	discoveries, err := registryProvider.Discoveries()
	if err != nil {
		return nil, err
	}

	mws, err := middlewareProvider.ClientMiddlewares()
	if err != nil {
		return nil, err
	}

	customTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   100,
		MaxConnsPerHost:       100,
	}

	objHttpClient, err := httptransport.NewClient(
		context.Background(),
		httptransport.WithEndpoint(clientConfig.GetHttp().GetEndpoint()),
		httptransport.WithDiscovery(discoveries[clientConfig.GetHttp().GetDiscoveryName()]),
		httptransport.WithMiddleware(middleware.Chain(mapsToSlice(mws)...)),
		httptransport.WithTimeout(60*time.Second),
		httptransport.WithTransport(customTransport),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create objectstore http client: %w", err)
	}

	p := &ObjectStoreProxy{
		app:           app,
		logger:        log.NewHelper(log.With(app.Logger(), "module", "gateway.proxy")),
		objClient:     objClient,
		fileClient:    fileClient,
		objHttpClient: objHttpClient,
	}

	return p, nil
}

func mapsToSlice(m map[string]middleware.Middleware) []middleware.Middleware {
	s := make([]middleware.Middleware, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// RegisterHandlers registers all HTTP handlers for the ObjectStoreProxy.
func (p *ObjectStoreProxy) RegisterHandlers(srv *httptransport.Server) {
	r := srv.Route("/")

	// Standard CRUD
	r.POST("/storage/upload", p.handleUpload)
	r.GET("/storage/download/{id}", p.handleDownload)
	r.DELETE("/storage/delete/{id}", p.handleDelete)

	// Multipart upload
	r.POST("/storage/multipart/initiate", p.handleInitiateMultipart)
	r.GET("/storage/multipart/{upload_id}/parts", p.handleListParts)
	r.PUT("/storage/multipart/{upload_id}/parts/{part_number}", p.handleUploadPart)
	r.POST("/storage/multipart/{upload_id}/complete", p.handleCompleteMultipart)
	r.DELETE("/storage/multipart/{upload_id}", p.handleAbortMultipart)
}

func (p *ObjectStoreProxy) handleListParts(ctx httptransport.Context) error {
	uploadID := ctx.Vars().Get("upload_id")
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.objClient.ListParts(ctx, req.(*objectstore.ListPartsRequest))
	})
	resp, err := h(ctx, &objectstore.ListPartsRequest{UploadId: uploadID})
	if err != nil {
		return err
	}
	return ctx.Result(200, resp)
}

func (p *ObjectStoreProxy) handleAbortMultipart(ctx httptransport.Context) error {
	uploadID := ctx.Vars().Get("upload_id")
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.objClient.AbortMultipartUpload(ctx, req.(*objectstore.AbortMultipartUploadRequest))
	})
	_, err := h(ctx, &objectstore.AbortMultipartUploadRequest{UploadId: uploadID})
	return err
}

func (p *ObjectStoreProxy) handleUpload(ctx httptransport.Context) error {
	err := ctx.Request().ParseMultipartForm(32 << 20)
	if err != nil {
		return err
	}
	file, header, err := ctx.Request().FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	visibility := ctx.Request().FormValue("visibility")
	if visibility == "" {
		visibility = "private"
	}

	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.fileClient.UploadFile(ctx, req.(*filemanager.UploadFileRequest))
	})
	uploadReq := &filemanager.UploadFileRequest{
		Name: header.Filename, ContentType: header.Header.Get("Content-Type"), Data: fileBytes, Visibility: visibility,
	}
	out, err := h(ctx, uploadReq)
	if err != nil {
		return err
	}
	resp := out.(*filemanager.UploadFileResponse)
	return ctx.Result(200, &UploadFileResponse{
		ID: resp.FileMetadata.Id, Name: resp.FileMetadata.Name, Size: resp.FileMetadata.Size,
		URL: fmt.Sprintf("/api/v1/storage/download/%d", resp.FileMetadata.Id),
	})
}

func (p *ObjectStoreProxy) handleInitiateMultipart(ctx httptransport.Context) error {
	var in struct {
		Name        string `json:"name"`
		Size        int64  `json:"size"`
		ContentType string `json:"content_type"`
		Visibility  string `json:"visibility"`
	}
	if err := ctx.Bind(&in); err != nil {
		return err
	}
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-file-visibility", in.Visibility)
		return p.fileClient.InitiateMultipartUpload(ctx, req.(*filemanager.InitiateMultipartUploadRequest))
	})
	resp, err := h(ctx, &filemanager.InitiateMultipartUploadRequest{
		Name: in.Name, Size: in.Size, ContentType: in.ContentType, Visibility: in.Visibility,
	})
	if err != nil {
		return err
	}
	return ctx.Result(200, resp)
}

func (p *ObjectStoreProxy) handleUploadPart(ctx httptransport.Context) error {
	uploadID := ctx.Vars().Get("upload_id")
	partNum, _ := strconv.Atoi(ctx.Vars().Get("part_number"))
	contentLength := ctx.Request().ContentLength
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return p.fileClient.GetMultipartUploadUrl(ctx, req.(*filemanager.GetMultipartUploadUrlRequest))
	})
	urlResp, err := h(ctx, &filemanager.GetMultipartUploadUrlRequest{UploadId: uploadID, PartNumber: int32(partNum)})
	if err != nil {
		return err
	}
	targetPath := urlResp.(*filemanager.GetMultipartUploadUrlResponse).UploadUrl
	req, _ := http.NewRequestWithContext(ctx, "PUT", targetPath, ctx.Request().Body)
	req.ContentLength = contentLength
	if ct := ctx.Request().Header.Get("Content-Type"); ct != "" {
		req.Header.Set("Content-Type", ct)
	} else {
		req.Header.Set("Content-Type", "application/octet-stream")
	}
	resp, err := p.objHttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("OBS upload failed (%d)", resp.StatusCode)
	}
	if etag := resp.Header.Get("ETag"); etag != "" {
		ctx.Response().Header().Set("ETag", etag)
	}
	return ctx.Result(resp.StatusCode, nil)
}

func (p *ObjectStoreProxy) handleCompleteMultipart(ctx httptransport.Context) error {
	uploadID := ctx.Vars().Get("upload_id")
	var in struct {
		Parts []struct {
			PartNumber int32  `json:"part_number"`
			Etag       string `json:"etag"`
		} `json:"parts"`
		Sha256 string `json:"sha256"`
	}
	if err := ctx.Bind(&in); err != nil {
		return err
	}
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-file-sha256", in.Sha256)
		return p.fileClient.CompleteMultipartUpload(ctx, req.(*filemanager.CompleteMultipartUploadRequest))
	})
	protoParts := make([]*filemanager.PartInfo, len(in.Parts))
	for i, part := range in.Parts {
		protoParts[i] = &filemanager.PartInfo{PartNumber: part.PartNumber, Etag: part.Etag}
	}
	resp, err := h(ctx, &filemanager.CompleteMultipartUploadRequest{UploadId: uploadID, Parts: protoParts})
	if err != nil {
		return err
	}
	return ctx.Result(200, resp)
}

func (p *ObjectStoreProxy) handleDownload(ctx httptransport.Context) error {
	id, _ := strconv.ParseInt(ctx.Vars().Get("id"), 10, 64)
	metadataOnly := ctx.Request().URL.Query().Get("metadata") == "true"
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		fileResp, err := p.fileClient.GetFile(ctx, &filemanager.GetFileRequest{Id: id})
		if err != nil {
			return nil, err
		}
		if metadataOnly {
			if fileResp.FileMetadata != nil {
				fileResp.DownloadUrl = fmt.Sprintf("/api/v1/storage/download/%d", fileResp.FileMetadata.Id)
			}
			return fileResp, nil
		}
		// FETCH WITH CONTENT TYPE
		body, err := p.objClient.DownloadObject(ctx, &objectstore.DownloadObjectRequest{Id: fileResp.FileMetadata.ObjectId})
		if err != nil {
			return nil, err
		}

		// Map the metadata back so we can use it in the outer scope
		return struct {
			Body *httpbody.HttpBody
			Meta *filemanager.GetFileResponse
		}{Body: body, Meta: fileResp}, nil
	})
	out, err := h(ctx, nil)
	if err != nil {
		return err
	}
	if metadataOnly {
		return ctx.Result(200, out)
	}

	data := out.(struct {
		Body *httpbody.HttpBody
		Meta *filemanager.GetFileResponse
	})

	// FIX: Force Content-Type from database/objectstore metadata
	contentType := data.Body.ContentType
	if contentType == "" || contentType == "application/octet-stream" {
		if data.Meta.FileMetadata.MimeType != "" {
			contentType = data.Meta.FileMetadata.MimeType
		}
	}

	ctx.Response().Header().Set("Content-Type", contentType)
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
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	URL  string `json:"url"`
}
type DeleteFileResponse struct {
	Status string `json:"status"`
}
