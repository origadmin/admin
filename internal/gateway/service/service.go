/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"

	"origadmin/application/admin/internal/gateway/client"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

// GatewayService is a container for downstream service clients.
// It does not implement any gRPC service interfaces itself. The gateway is
// transparent, and HTTP handlers are registered directly with client connections.
type GatewayService struct {
	Auth   *client.AuthBridgeSet
	System *client.SystemBridgeSet
}

// NewGatewayService new a gateway service.
func NewGatewayService(authClient *client.AuthBridgeSet, systemClient *client.SystemBridgeSet) (*GatewayService, error) {
	return &GatewayService{
		Auth:   authClient,
		System: systemClient,
	}, nil
}
