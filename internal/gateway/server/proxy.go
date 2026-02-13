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
	"github.com/origadmin/runtime"
	"origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

// ObjectStoreProxy handles reverse proxying to the objectstore service using service discovery.
type ObjectStoreProxy struct {
	app       *runtime.App
	discovery registry.Discovery
	logger    *log.Helper
	client    objectstore.ObjectStoreServiceHTTPClient

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

	// Create HTTP client for ObjectStore
	conn, err := grpcclient.NewConn(app, bootstrap, "objectstore", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create objectstore client: %w", err)
	}
	client := objectstore.NewObjectStoreServiceHTTPClient(conn)

	p := &ObjectStoreProxy{
		app:       app,
		discovery: discovery, // Use the correct, validated discovery instance
		logger:    log.NewHelper(log.With(app.Logger(), "module", "gateway.proxy")),
		client:    client,
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

	// 1. Get Metadata (Size) first to know how to chunk
	// We can use GetObject for this.
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

	// 2. Handle Client's Range Request (if any)
	clientRange := r.Header.Get("Range")
	var start, end int64
	end = totalSize - 1

	if clientRange != "" {
		// Parse standard Range header: bytes=0-100
		// Simplified parsing
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

	// 3. Chunking Loop
	// Define chunk size (e.g., 10MB to be safe and responsive)
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

		// p.logger.Infof("Requesting chunk: %s", backendRange)
		resp, err := p.client.DownloadObject(r.Context(), req)
		if err != nil {
			p.logger.Errorf("Failed to download chunk %s: %v", backendRange, err)
			// If we already wrote headers, we can't really send an error status code now.
			// We just stop.
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
