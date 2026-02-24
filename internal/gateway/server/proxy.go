/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	httptransport "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/genproto/googleapis/api/httpbody"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

// ObjectStoreProxy handles reverse proxying to the objectstore service using service discovery.
// It acts as a BFF (Backend for Frontend) adapter for file operations.
type ObjectStoreProxy struct {
	app       *runtime.App
	discovery registry.Discovery
	logger    *log.Helper

	// gRPC Clients
	objClient  objectstore.ObjectStoreServiceClient
	fileClient filemanager.FileManagerServiceClient

	// Simple cache for endpoints to avoid querying registry on every request
	endpoints  []string
	mu         sync.RWMutex
	lastUpdate time.Time
}

// NewObjectStoreProxy creates a new ObjectStoreProxy.
// It declares its dependency on ClientMiddlewareProvider, which will be injected by wire.
func NewObjectStoreProxy(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*ObjectStoreProxy, error) {
	// Create gRPC client for ObjectStore using the injected middleware provider
	objConn, err := grpcclient.NewConn(app, bootstrap, "objectstore", middlewareProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create objectstore client: %w", err)
	}
	objClient := objectstore.NewObjectStoreServiceClient(objConn)

	// Create gRPC client for FileManager using the injected middleware provider
	fileConn, err := grpcclient.NewConn(app, bootstrap, "filemanager", middlewareProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create filemanager client: %w", err)
	}
	fileClient := filemanager.NewFileManagerServiceClient(fileConn)

	p := &ObjectStoreProxy{
		app:        app,
		logger:     log.NewHelper(log.With(app.Logger(), "module", "gateway.proxy")),
		objClient:  objClient,
		fileClient: fileClient,
	}

	return p, nil
}

// RegisterHandlers registers all HTTP handlers for the ObjectStoreProxy.
func (p *ObjectStoreProxy) RegisterHandlers(srv *httptransport.Server) {
	r := srv.Route("/")

	// Register upload handler - POST /storage/upload
	r.POST("/storage/upload", p.handleUpload)

	// Register download handler - GET /storage/download/{id}
	r.GET("/storage/download/{id}", p.handleDownload)

	// Register delete handler - DELETE /storage/delete/{id}
	r.DELETE("/storage/delete/{id}", p.handleDelete)
}

// handleUpload handles multipart/form-data file upload.
func (p *ObjectStoreProxy) handleUpload(ctx httptransport.Context) error {
	// Parse multipart form manually since we need to handle file upload
	err := ctx.Request().ParseMultipartForm(32 << 20)
	if err != nil {
		return err
	}

	// Extract file and metadata
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

	// Use middleware chain to ensure auth and other middleware are applied
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		uploadReq := req.(*filemanager.UploadFileRequest)
		return p.fileClient.UploadFile(ctx, uploadReq)
	})

	uploadReq := &filemanager.UploadFileRequest{
		Name:        header.Filename,
		ContentType: header.Header.Get("Content-Type"),
		Data:        fileBytes,
		Visibility:  visibility,
	}

	out, err := h(ctx, uploadReq)
	if err != nil {
		return err
	}

	resp := out.(*filemanager.UploadFileResponse)

	// Return response
	result := &UploadFileResponse{
		ID:   resp.FileMetadata.Id,
		Name: resp.FileMetadata.Name,
		Size: resp.FileMetadata.Size,
		URL:  fmt.Sprintf("/storage/download/%d", resp.FileMetadata.Id),
	}

	return ctx.Result(200, result)
}

// handleDownload handles file download with query parameter based behavior.
// - GET /storage/download/{id}?metadata=true → returns file metadata (GetFileResponse)
// - GET /storage/download/{id} → returns file content (binary)
func (p *ObjectStoreProxy) handleDownload(ctx httptransport.Context) error {
	var in DownloadFileRequest
	if err := ctx.BindVars(&in); err != nil {
		return err
	}
	if err := ctx.BindQuery(&in); err != nil {
		return err
	}

	// Check if metadata is requested via query parameter
	metadataOnly := ctx.Request().URL.Query().Get("metadata") == "true"

	// Use middleware chain to ensure auth and other middleware are applied
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		downloadReq := req.(*DownloadFileRequest)

		// Get file metadata
		fileResp, err := p.fileClient.GetFile(ctx, &filemanager.GetFileRequest{Id: downloadReq.ID})
		if err != nil {
			return nil, err
		}

		// If metadata is requested, return it directly
		if metadataOnly {
			// Rewrite the download URL to point to our proxy instead of the internal service
			if fileResp.FileMetadata != nil {
				fileResp.DownloadUrl = fmt.Sprintf("/storage/download/%d", fileResp.FileMetadata.Id)
			}
			return fileResp, nil
		}

		// Otherwise, return file content (HttpBody)
		objectID := fileResp.FileMetadata.ObjectId
		downReq := &objectstore.DownloadObjectRequest{
			Id: objectID,
		}

		return p.objClient.DownloadObject(ctx, downReq)
	})

	out, err := h(ctx, &in)
	if err != nil {
		return err
	}

	// Handle different response types
	if metadataOnly {
		// Return JSON metadata
		return ctx.Result(200, out)
	}

	// Return binary file content from HttpBody
	body := out.(*httpbody.HttpBody)
	// Use the Data and ContentType from HttpBody
	ctx.Response().Header().Set("Content-Type", body.ContentType)
	_, err = ctx.Response().Write(body.Data)
	return err
}

// handleDelete handles file deletion.
func (p *ObjectStoreProxy) handleDelete(ctx httptransport.Context) error {
	var in DeleteFileRequest
	if err := ctx.BindVars(&in); err != nil {
		return err
	}
	if err := ctx.BindQuery(&in); err != nil {
		return err
	}

	// Use middleware chain to ensure auth and other middleware are applied
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		deleteReq := req.(*DeleteFileRequest)
		_, err := p.fileClient.DeleteFile(ctx, &filemanager.DeleteFileRequest{Id: deleteReq.ID})
		if err != nil {
			return nil, err
		}

		// Return response
		result := &DeleteFileResponse{
			Status: "deleted",
		}

		return result, nil
	})

	out, err := h(ctx, &in)
	if err != nil {
		return err
	}

	result := out.(*DeleteFileResponse)
	return ctx.Result(200, result)
}

// UploadFileRequest represents the upload request structure.
type UploadFileRequest struct {
	Name        string `form:"name"`
	ContentType string `form:"content_type"`
	Data        []byte `form:"data"`
	Visibility  string `form:"visibility"`
}

// UploadFileResponse represents the upload response structure.
type UploadFileResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	URL  string `json:"url"`
}

// DownloadFileRequest represents the download request structure.
type DownloadFileRequest struct {
	ID int64 `path:"id" json:"id" form:"id"`
}

// DownloadFileResponse represents the download response structure.
type DownloadFileResponse struct {
	Data        []byte `json:"data"`
	ContentType string `json:"content_type"`
	Name        string `json:"name"`
}

// DeleteFileRequest represents the delete request structure.
type DeleteFileRequest struct {
	ID int64 `path:"id" json:"id" form:"id"`
}

// DeleteFileResponse represents the delete response structure.
type DeleteFileResponse struct {
	Status string `json:"status"`
}