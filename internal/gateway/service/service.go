/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime/service/transport"
	"origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/gateway/client"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

// GatewayService acts as a dependency injection container for the various
// generated bridge sets. Each bridge set contains the client-side logic
// to forward requests to a specific downstream gRPC service.
// This approach avoids implementing downstream service interfaces directly in the gateway.
type GatewayService struct {
	Identity    *client.IdentityBridgeSet
	System      *client.SystemBridgeSet
	FileManager *client.FileManagerBridgeSet
	ObjectStore *client.ObjectStoreBridgeSet
}

// NewGatewayService creates a new GatewayService, aggregating the generated
// bridge clients for all downstream services.
func NewGatewayService(
	identityClient *client.IdentityBridgeSet,
	systemClient *client.SystemBridgeSet,
	fileManagerClient *client.FileManagerBridgeSet,
	objectStoreClient *client.ObjectStoreBridgeSet,
) (*GatewayService, error) {
	return &GatewayService{
		Identity:    identityClient,
		System:      systemClient,
		FileManager: fileManagerClient,
		ObjectStore: objectStoreClient,
	}, nil
}

// RegisterHTTPHandlers registers all the HTTP handlers for the downstream services
// onto the provided HTTP server. It also logs the registered routes.
func (s *GatewayService) RegisterHTTPHandlers(srv *transport.HTTPServer) {
	// Register handlers for the 'system' service
	system.RegisterUserServiceHTTPServer(srv, s.System.User)
	system.RegisterRoleServiceHTTPServer(srv, s.System.Role)
	system.RegisterPermissionServiceHTTPServer(srv, s.System.Permission)
	system.RegisterResourceServiceHTTPServer(srv, s.System.Resource)
	system.RegisterViewServiceHTTPServer(srv, s.System.View)

	// Register handlers for the 'identity' service
	identity.RegisterAuthServiceHTTPServer(srv, s.Identity.Auth)
	identity.RegisterMeServiceHTTPServer(srv, s.Identity.Me)
	identity.RegisterAdminServiceHTTPServer(srv, s.Identity.Admin)

	// Register handlers for the 'filemanager' service
	filemanager.RegisterFileManagerServiceHTTPServer(srv, s.FileManager.FileManager)

	// Register handlers for the 'objectstore' service
	objectstore.RegisterObjectStoreServiceHTTPServer(srv, s.ObjectStore.ObjectStore)
}
