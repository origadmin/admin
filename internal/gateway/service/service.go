/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime/service/transport"
	"origadmin/application/admin/api/v1/services/auth"
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
	Auth   *client.AuthBridgeSet
	System *client.SystemBridgeSet
}

// NewGatewayService creates a new GatewayService, aggregating the generated
// bridge clients for all downstream services.
func NewGatewayService(authClient *client.AuthBridgeSet, systemClient *client.SystemBridgeSet) (*GatewayService, error) {
	return &GatewayService{
		Auth:   authClient,
		System: systemClient,
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

	// Register handlers for the 'auth' service
	auth.RegisterAuthServiceHTTPServer(srv, s.Auth.Auth)
	auth.RegisterMeServiceHTTPServer(srv, s.Auth.Me)
	auth.RegisterAdminServiceHTTPServer(srv, s.Auth.Admin)
}
