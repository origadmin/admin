/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"

	gatewayAPI "origadmin/application/admin/api/v1/services/gateway"
	"origadmin/application/admin/internal/gateway/client"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGatewayService)

// GatewayService is a gateway service.
type GatewayService struct {
	gatewayAPI.UnimplementedGatewayServiceServer
	authClient   *client.AuthClientSet
	systemClient *client.SystemClientSet
}

// NewGatewayService new a gateway service.
func NewGatewayService(authClient *client.AuthClientSet, systemClient *client.SystemClientSet) (*GatewayService, error) {
	return &GatewayService{
		authClient:   authClient,
		systemClient: systemClient,
	}, nil
}
