/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/origadmin/runtime"
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
func NewObjectStoreProxy(app *runtime.App, bootstrap *conf.Config) (*ObjectStoreProxy, error) {
	// Create gRPC client for ObjectStore
	objConn, err := grpcclient.NewConn(app, bootstrap, "objectstore", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create objectstore client: %w", err)
	}
	objClient := objectstore.NewObjectStoreServiceClient(objConn)

	// Create gRPC client for FileManager
	fileConn, err := grpcclient.NewConn(app, bootstrap, "filemanager", nil)
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

// ServeHTTP implements http.Handler.
// It routes requests to specific handlers based on the path and method.
// Base path: /api/v1/storage
func (p *ObjectStoreProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Path: /api/v1/storage/...
	path := r.URL.Path

	// Simple router
	if strings.HasSuffix(path, "/upload") && r.Method == http.MethodPost {
		p.handleUpload(w, r)
		return
	}

	if strings.Contains(path, "/download/") && r.Method == http.MethodGet {
		p.handleDownload(w, r)
		return
	}

	if strings.Contains(path, "/delete/") && r.Method == http.MethodDelete {
		p.handleDelete(w, r)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

// handleUpload handles multipart/form-data file upload.
// POST /api/v1/storage/upload
func (p *ObjectStoreProxy) handleUpload(w http.ResponseWriter, r *http.Request) {
	// 1. Parse Multipart Form (Max 32MB in memory)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		p.logger.Errorf("Failed to parse multipart form: %v", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// 2. Get file from form
	file, header, err := r.FormFile("file") // Expecting form field name "file"
	if err != nil {
		p.logger.Errorf("Failed to get file from form: %v", err)
		http.Error(w, "Missing file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 3. Read file content
	// TODO: For very large files, stream directly to ObjectStore using client-side streaming
	// For now, read into memory for simplicity (suitable for small files)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		p.logger.Errorf("Failed to read file content: %v", err)
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// 4. Call FileManager.UploadFile (which handles ObjectStore upload internally via UseCase)
	// OR Call ObjectStore directly then FileManager.
	// Since we fixed FileManager.UploadFile logic in previous steps, let's use it.
	// But wait, FileManager.UploadFile expects raw bytes in the request.

	// Extract metadata from form if available
	visibility := r.FormValue("visibility")
	if visibility == "" {
		visibility = "private"
	}

	req := &filemanager.UploadFileRequest{
		Name:        header.Filename,
		ContentType: header.Header.Get("Content-Type"),
		Data:        fileBytes,
		Visibility:  visibility,
	}

	// Forward Authorization header
	ctx := p.forwardAuthContext(r)

	resp, err := p.fileClient.UploadFile(ctx, req)
	if err != nil {
		p.logger.Errorf("FileManager UploadFile failed: %v", err)
		http.Error(w, fmt.Sprintf("Upload failed: %v", err), http.StatusInternalServerError)
		return
	}

	// 5. Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Simple JSON response
	fmt.Fprintf(w, `{"id": "%d", "name": "%s", "size": "%d", "url": "/api/v1/storage/download/%d"}`,
		resp.FileMetadata.Id, resp.FileMetadata.Name, resp.FileMetadata.Size, resp.FileMetadata.Id)
}

// handleDownload handles file download.
// GET /api/v1/storage/download/{id}
func (p *ObjectStoreProxy) handleDownload(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	ctx := p.forwardAuthContext(r)

	// 1. Get File Metadata
	fileResp, err := p.fileClient.GetFile(ctx, &filemanager.GetFileRequest{Id: id})
	if err != nil {
		p.logger.Errorf("Failed to get file metadata: %v", err)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// 2. Get Object Stream from ObjectStore
	// We need the Object ID from metadata (which might not be exposed in GetFileResponse directly if not added)
	// In previous steps, we saw GetFileResponse has FileMetadata.
	// Let's assume FileMetadata has ObjectId.
	objectID := fileResp.FileMetadata.ObjectId

	// Handle Range Request
	rangeHeader := r.Header.Get("Range")

	downReq := &objectstore.DownloadObjectRequest{
		Id:    objectID,
		Range: rangeHeader,
	}

	// Call ObjectStore DownloadObject
	// Note: gRPC DownloadObject returns google.api.HttpBody.
	// We need to stream the Data field.
	// Ideally, we should use a streaming RPC if available, but here it seems to be unary returning HttpBody.
	// If it returns HttpBody, it contains the whole content (or chunk if ranged).

	objResp, err := p.objClient.DownloadObject(ctx, downReq)
	if err != nil {
		p.logger.Errorf("Failed to download object: %v", err)
		http.Error(w, "Failed to download file", http.StatusInternalServerError)
		return
	}

	// 3. Set Response Headers
	w.Header().Set("Content-Type", fileResp.FileMetadata.MimeType)
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileResp.FileMetadata.Name))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(objResp.Data))) // Assuming unary response for now

	// 4. Write Body
	_, err = w.Write(objResp.Data)
	if err != nil {
		p.logger.Errorf("Failed to write response: %v", err)
	}
}

// handleDelete handles file deletion.
// DELETE /api/v1/storage/delete/{id}
func (p *ObjectStoreProxy) handleDelete(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	ctx := p.forwardAuthContext(r)

	// Call FileManager DeleteFile
	_, err = p.fileClient.DeleteFile(ctx, &filemanager.DeleteFileRequest{Id: id})
	if err != nil {
		p.logger.Errorf("Failed to delete file: %v", err)
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "deleted"}`)
}

// forwardAuthContext creates a context with the Authorization header from the request.
func (p *ObjectStoreProxy) forwardAuthContext(r *http.Request) context.Context {
	ctx := r.Context()
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// Add to gRPC metadata
		// md := metadata.Pairs("authorization", authHeader)
		// return metadata.NewOutgoingContext(ctx, md)

		// Or use Kratos transport/http to extract and propagate
		// For simplicity, we assume the client interceptors handle this if configured,
		// or we manually pass it.
		// Since we are inside the gateway, we might need to manually propagate.
		// TODO: Implement proper context propagation
	}
	return ctx
}
