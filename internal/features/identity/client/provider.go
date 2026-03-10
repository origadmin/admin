/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"fmt"

	"github.com/google/wire"

	"github.com/origadmin/runtime"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

const (
	// ServiceNameSystem is the short name for the system service, which provides all needed clients.
	ServiceNameSystem = "system"
)

// ProviderSet provides all gRPC clients required by the identity feature.
var ProviderSet = wire.NewSet(
	NewAuthorizationServiceClient,
	NewUserServiceClient,
	NewViewServiceClient,
)

// NewAuthorizationServiceClient creates a gRPC client for the system's AuthorizationService.
func NewAuthorizationServiceClient(
	app *runtime.App,
	bootstrap *confpb.Bootstrap,
) (systemv1.PolicyQueryServiceClient, error) {
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameSystem)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to system service for AuthorizationService: %w", err)
	}
	return systemv1.NewPolicyQueryServiceClient(conn), nil
}

// NewUserServiceClient creates a gRPC client for the system's UserService.
func NewUserServiceClient(
	app *runtime.App,
	bootstrap *confpb.Bootstrap,
) (systemv1.UserServiceClient, error) {
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameSystem)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to system service for UserService: %w", err)
	}
	return systemv1.NewUserServiceClient(conn), nil
}

// NewViewServiceClient creates a gRPC client for the system's ViewService.
func NewViewServiceClient(
	app *runtime.App,
	bootstrap *confpb.Bootstrap,
) (systemv1.ViewServiceClient, error) {
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameSystem)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to system service for ViewService: %w", err)
	}
	return systemv1.NewViewServiceClient(conn), nil
}
