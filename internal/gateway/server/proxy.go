/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/golang-jwt/jwt/v5"
	"github.com/origadmin/runtime"
	"origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

// ObjectStoreProxy handles reverse proxying to the objectstore service using service discovery.
type ObjectStoreProxy struct {
	app       *runtime.App
	discovery registry.Discovery
	logger    *log.Helper
	client    objectstore.ObjectStoreServiceHTTPServer
	fileMgr   filemanager.FileManagerServiceHTTPServer

	// JWT secret for generating access tokens
	jwtSecret string

	// Simple cache for endpoints to avoid querying registry on every request
	endpoints  []string
	mu         sync.RWMutex
	lastUpdate time.Time
}

// NewObjectStoreProxy creates a new ObjectStoreProxy.
func NewObjectStoreProxy(app *runtime.App, bootstrap *conf.Config) (*ObjectStoreProxy, error) {
	regProvider, err := app.RegistryProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to get registry provider: %w", err)
	}

	// Get all available discovery clients from the provider.
	discoveries, err := regProvider.Discoveries()
	if err != nil {
		return nil, fmt.Errorf("failed to get discoveries map: %w", err)
	}

	// Get the name of the default discovery client from the configuration.
	defaultDiscoveryName := bootstrap.GetBootstrap().GetDefaultDiscovery()
	if defaultDiscoveryName == "" {
		return nil, fmt.Errorf("no default discovery configured")
	}

	// Get the correct discovery instance from the map.
	discovery, ok := discoveries[defaultDiscoveryName]
	if !ok {
		return nil, fmt.Errorf("default discovery '%s' not found in available discoveries", defaultDiscoveryName)
	}

	// Create gRPC client and bridge to HTTP for ObjectStore (without middleware for proxy)
	conn, err := grpcclient.NewConn(app, bootstrap, "objectstore", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create objectstore client: %w", err)
	}
	client := objectstore.NewObjectStoreServiceGRPC2HTTP(conn)

	// Create gRPC client and bridge to HTTP for FileManager (without middleware for proxy)
	fileMgrConn, err := grpcclient.NewConn(app, bootstrap, "filemanager", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create filemanager client: %w", err)
	}
	fileMgr := filemanager.NewFileManagerServiceGRPC2HTTP(fileMgrConn)

	// Get JWT secret from config or use default
	jwtSecret := "default-secret-change-in-production" // TODO: get from config properly

	p := &ObjectStoreProxy{
		app:       app,
		discovery: discovery,
		logger:    log.NewHelper(log.With(app.Logger(), "module", "gateway.proxy")),
		client:    client,
		fileMgr:   fileMgr,
		jwtSecret: jwtSecret,
	}
	// Initial endpoint update
	go p.updateEndpoints()

	return p, nil
}

// ServeHTTP implements http.Handler. It finds a healthy objectstore instance and proxies the request.
func (p *ObjectStoreProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extract object ID from path: /api/v1/objects/{id}
	path := r.URL.Path
	parts := strings.Split(path, "/")
	// Expected: ["", "api", "v1", "objects", "{id}"]
	if len(parts) < 5 || parts[3] != "objects" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	objectID := parts[4]

	// 1. Check if token parameter exists for temporary access
	token := r.URL.Query().Get("token")
	if token != "" {
		if !p.validateAccessToken(token, objectID) {
			http.Error(w, "Invalid or expired access token", http.StatusForbidden)
			return
		}
	} else {
		// 2. Get file metadata from FileManager to check access control
		fileMeta, err := p.getFileMetadataByObjectID(r.Context(), objectID)
		if err != nil {
			p.logger.Errorf("Failed to get file metadata: %v", err)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// 3. Check access time limit (expires_at)
		if !fileMeta.IsPermanent && fileMeta.ExpiresAt != nil {
			if time.Now().After(fileMeta.ExpiresAt.AsTime()) {
				http.Error(w, "File access expired", http.StatusGone)
				return
			}
		}

		// 4. Check download count limit
		if fileMeta.MaxDownloads > 0 && fileMeta.DownloadCount >= int32(fileMeta.MaxDownloads) {
			http.Error(w, "Download limit reached", http.StatusGone)
			return
		}

		// 5. Check visibility and permissions
		if fileMeta.Visibility == "public" {
			// public files: allow access directly
			p.logger.Debugf("Public file access: %s", objectID)
		} else {
			// private files: verify authentication and authorization
			userID, err := p.getUserIDFromRequest(r)
			if err != nil {
				p.logger.Errorf("Failed to get user ID: %v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Check ownership or admin permissions
			if !p.hasAccessPermission(userID, fileMeta) {
				p.logger.Warnf("Access denied: user %d cannot access file %s", userID, fileMeta.Id)
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			p.logger.Debugf("Private file access granted: user %d, file %s", userID, objectID)
		}

		// 6. Update download count (async)
		go p.incrementDownloadCount(fileMeta.Id)
	}

	// 7. Get Metadata (Size) first to know how to chunk
	metaReq := &objectstore.GetObjectRequest{Id: objectID}
	metaResp, err := p.client.GetObject(r.Context(), metaReq)
	if err != nil {
		p.logger.Errorf("Failed to get object metadata: %v", err)
		http.Error(w, "Object Not Found", http.StatusNotFound)
		return
	}
	totalSize := metaResp.Object.Size
	contentType := metaResp.Object.ContentType

	// Set Response Headers
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Cache-Control", "private, max-age=300")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// 8. Handle Client's Range Request (if any)
	clientRange := r.Header.Get("Range")
	var start, end int64
	end = totalSize - 1

	if clientRange != "" {
		// Parse standard Range header: bytes=0-100
		if strings.HasPrefix(clientRange, "bytes=") {
			rangeStr := strings.TrimPrefix(clientRange, "bytes=")
			rangeParts := strings.Split(rangeStr, "-")
			if len(rangeParts) >= 1 {
				fmt.Sscanf(rangeParts[0], "%d", &start)
				if len(rangeParts) > 1 && rangeParts[1] != "" {
					fmt.Sscanf(rangeParts[1], "%d", &end)
				}
			}
		}
		w.WriteHeader(http.StatusPartialContent)
		w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, totalSize))
	} else {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", totalSize))
	}

	// 9. Chunking Loop
	const chunkSize = 10 * 1024 * 1024 // 10MB
	current := start

	for current <= end {
		// Calculate chunk range
		chunkEnd := current + chunkSize - 1
		if chunkEnd > end {
			chunkEnd = end
		}

		// Construct Range Header for Backend Request
		backendRange := fmt.Sprintf("bytes=%d-%d", current, chunkEnd)

		// Call Backend
		req := &objectstore.DownloadObjectRequest{
			Id:    objectID,
			Range: backendRange,
		}

		resp, err := p.client.DownloadObject(r.Context(), req)
		if err != nil {
			p.logger.Errorf("Failed to download chunk %s: %v", backendRange, err)
			return
		}

		// Write Chunk to Client
		_, err = w.Write(resp.Data)
		if err != nil {
			p.logger.Errorf("Failed to write chunk to client: %v", err)
			return
		}

		// Flush if possible to keep stream moving
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		// Move to next chunk
		current += chunkSize
	}
}

// getEndpoint gets a healthy endpoint for objectstore service from cache or discovery.
func (p *ObjectStoreProxy) getEndpoint(ctx context.Context) (string, error) {
	p.mu.RLock()
	// Use cache if it's recent
	if len(p.endpoints) > 0 && time.Since(p.lastUpdate) < 10*time.Second {
		ep := p.endpoints[rand.Intn(len(p.endpoints))]
		p.mu.RUnlock()
		return ep, nil
	}
	p.mu.RUnlock()

	// If cache is stale or empty, update it
	return p.updateEndpoints()
}

// updateEndpoints fetches service instances from the registry and updates the local cache.
func (p *ObjectStoreProxy) updateEndpoints() (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Double-check lock pattern
	if len(p.endpoints) > 0 && time.Since(p.lastUpdate) < 10*time.Second {
		return p.endpoints[rand.Intn(len(p.endpoints))], nil
	}

	p.logger.Info("Updating objectstore endpoints from service discovery...")
	serviceName := "origadmin.service.objectstore"
	instances, err := p.discovery.GetService(context.Background(), serviceName)
	if err != nil {
		return "", err
	}

	if len(instances) == 0 {
		return "", fmt.Errorf("no instances found for service %s", serviceName)
	}

	var newEndpoints []string
	for _, ins := range instances {
		for _, ep := range ins.Endpoints {
			if strings.HasPrefix(ep, "http://") {
				newEndpoints = append(newEndpoints, ep)
			}
		}
	}

	if len(newEndpoints) == 0 {
		return "", fmt.Errorf("no healthy HTTP instances found for %s", serviceName)
	}

	p.endpoints = newEndpoints
	p.lastUpdate = time.Now()
	p.logger.Infof("Updated objectstore endpoints: %v", p.endpoints)

	return p.endpoints[rand.Intn(len(p.endpoints))], nil
}

// getFileMetadataByObjectID retrieves file metadata by object ID from FileManager
func (p *ObjectStoreProxy) getFileMetadataByObjectID(ctx context.Context, objectID string) (*types.FileMetadata, error) {
	// List all files and filter by object_id (inefficient, but FileManager doesn't have query by object_id API yet)
	listResp, err := p.fileMgr.ListFiles(ctx, &filemanager.ListFilesRequest{
		Page:     1,
		PageSize: 1000, // TODO: implement pagination
	})
	if err != nil {
		return nil, err
	}

	for _, file := range listResp.Files {
		if file.ObjectId == objectID {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file not found")
}

// getUserIDFromRequest extracts user ID from JWT token in request header
func (p *ObjectStoreProxy) getUserIDFromRequest(r *http.Request) (int64, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("missing authorization header")
	}

	// Extract Bearer token
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return 0, fmt.Errorf("invalid authorization header format")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse and validate JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.jwtSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, fmt.Errorf("user_id not found in token")
		}
		return int64(userID), nil
	}

	return 0, fmt.Errorf("invalid token")
}

// hasAccessPermission checks if user has permission to access the file
func (p *ObjectStoreProxy) hasAccessPermission(userID int64, fileMeta *types.FileMetadata) bool {
	// Owner can always access
	if fileMeta.OwnerId == userID {
		return true
	}

	// TODO: Check if user is admin
	// TODO: Check if file is shared with user

	return false
}

// incrementDownloadCount increments the download count for a file (async)
func (p *ObjectStoreProxy) incrementDownloadCount(fileID int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get current file metadata
	getResp, err := p.fileMgr.GetFile(ctx, &filemanager.GetFileRequest{Id: fileID})
	if err != nil {
		p.logger.Errorf("Failed to get file metadata for count increment: %v", err)
		return
	}

	// Update download count
	// Note: UpdateFileRequest doesn't have download_count field in proto yet
	// TODO: extend UpdateFileRequest to support download_count update
	_, err = p.fileMgr.UpdateFile(ctx, &filemanager.UpdateFileRequest{
		Id:         fileID,
		Name:       getResp.FileMetadata.Name,
		Visibility: getResp.FileMetadata.Visibility,
	})
	if err != nil {
		p.logger.Errorf("Failed to update download count: %v", err)
	}
}

// validateAccessToken validates a temporary access token generated by the proxy
func (p *ObjectStoreProxy) validateAccessToken(tokenString, objectID string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.jwtSecret), nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token is for the correct object
		if objID, ok := claims["object_id"].(string); ok && objID == objectID {
			// Check expiration
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() < int64(exp) {
					return true
				}
			}
		}
	}

	return false
}

// GenerateAccessToken generates a temporary access token for file download
func (p *ObjectStoreProxy) GenerateAccessToken(objectID string, expiresIn time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"object_id": objectID,
		"exp":       time.Now().Add(expiresIn).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(p.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
