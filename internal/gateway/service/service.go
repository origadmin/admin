/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/http"
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
// onto the provided HTTP router.
func (s *GatewayService) RegisterHTTPHandlers(router *transport.RouterHTTP) {
	// Register handlers for the 'system' service
	system.RegisterUserServiceHTTPServer(router, s.System.User)
	system.RegisterRoleServiceHTTPServer(router, s.System.Role)
	system.RegisterPermissionServiceHTTPServer(router, s.System.Permission)
	system.RegisterResourceServiceHTTPServer(router, s.System.Resource)
	system.RegisterViewServiceHTTPServer(router, s.System.View)

	// Register handlers for the 'auth' service
	auth.RegisterAuthServiceHTTPServer(router, s.Auth.Auth)
	auth.RegisterMeServiceHTTPServer(router, s.Auth.Me)
}
