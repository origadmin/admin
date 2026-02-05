/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"fmt"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

const (
	// ServiceNameSystem is the short name for the system service, which provides all needed clients.
	ServiceNameSystem = "system"
)

// ProviderSet provides all gRPC clients required by the auth feature.
var ProviderSet = wire.NewSet(
	NewAuthorizationServiceClient,
	NewUserServiceClient,
	NewViewServiceClient,
)

// NewAuthorizationServiceClient creates a gRPC client for the system's AuthorizationService.
func NewAuthorizationServiceClient(
	app *runtime.App,
	bootstrap *conf.Config,
	middlewareProvider container.ClientMiddlewareProvider,
) (systemv1.AuthorizationServiceClient, error) {
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameSystem, middlewareProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to system service for AuthorizationService: %w", err)
	}
	return systemv1.NewAuthorizationServiceClient(conn.(*grpc.ClientConn)), nil
}

// NewUserServiceClient creates a gRPC client for the system's UserService.
func NewUserServiceClient(
	app *runtime.App,
	bootstrap *conf.Config,
	middlewareProvider container.ClientMiddlewareProvider,
) (systemv1.UserServiceClient, error) {
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameSystem, middlewareProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to system service for UserService: %w", err)
	}
	return systemv1.NewUserServiceClient(conn.(*grpc.ClientConn)), nil
}

// NewViewServiceClient creates a gRPC client for the system's ViewService.
func NewViewServiceClient(
	app *runtime.App,
	bootstrap *conf.Config,
	middlewareProvider container.ClientMiddlewareProvider,
) (systemv1.ViewServiceClient, error) {
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameSystem, middlewareProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to system service for ViewService: %w", err)
	}
	return systemv1.NewViewServiceClient(conn.(*grpc.ClientConn)), nil
}
