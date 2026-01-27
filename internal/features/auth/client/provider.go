/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"fmt"

	"github.com/google/wire"

	"github.com/origadmin/runtime"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/container"
	runtimegrpc "github.com/origadmin/runtime/service/transport/grpc"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
)

const (
	// ServiceNameSystem is the short name for the system service.
	ServiceNameSystem = "system"
)

// ProviderSet is client providers for the auth module.
var ProviderSet = wire.NewSet(
	NewAuthorizationServiceClient,
	NewSystemPolicyProvider,
	//wire.Bind(new(biz.PolicyProvider), new(*systemPolicyProvider)),
)

// NewAuthorizationServiceClient creates a gRPC client for the system's AuthorizationService.
func NewAuthorizationServiceClient(
	app *runtime.App,
	bootstrap *conf.Config,
	middlewareProvider container.ClientMiddlewareProvider,
) (systemv1.AuthorizationServiceClient, error) {
	var clientConfig *transportv1.Client

	// The conventional name for the system service's gRPC client
	convention := fmt.Sprintf("origadmin.service.%s.client.grpc", ServiceNameSystem)

	clients := bootstrap.Clients()
	if clients != nil {
		for _, cli := range clients.Configs {
			if cli.GetGrpc() == nil {
				continue
			}
			if cli.Name == ServiceNameSystem || cli.Name == convention {
				clientConfig = cli
				break
			}
		}
	}

	if clientConfig == nil {
		return nil, fmt.Errorf("gRPC client config not found for service: %s (checked names: '%s', '%s')", ServiceNameSystem, ServiceNameSystem, convention)
	}

	registryProvider, err := app.RegistryProvider()
	if err != nil {
		return nil, err
	}

	discoveries, err := registryProvider.Discoveries()
	if err != nil {
		return nil, err
	}

	middlewares, err := middlewareProvider.ClientMiddlewares()
	if err != nil {
		return nil, err
	}

	conn, err := runtimegrpc.NewClient(app.Context(), clientConfig.GetGrpc(), &runtimegrpc.ClientOptions{
		Discoveries:       discoveries,
		ClientMiddlewares: middlewares,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to system service: %w", err)
	}

	return systemv1.NewAuthorizationServiceClient(conn), nil
}
