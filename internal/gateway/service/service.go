/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"

	"origadmin/application/admin/api/v1/services/auth"
	gatewayAPI "origadmin/application/admin/api/v1/services/gateway"
	"origadmin/application/admin/api/v1/services/system"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

// GatewayService is a gateway service.
type GatewayService struct {
	gatewayAPI.UnimplementedGatewayServiceServer
	authClient   auth.AuthServiceClient
	systemClient system.UserServiceClient
}

// NewGatewayService new a gateway service.
func NewGatewayService(authClient auth.AuthServiceClient, systemClient system.UserServiceClient) (*GatewayService, error) {
	return &GatewayService{
		authClient:   authClient,
		systemClient: systemClient,
	}, nil
}
