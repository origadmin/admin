/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package grpcclient

import (
	"fmt"

	"github.com/origadmin/runtime"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/container"
	runtimegrpc "github.com/origadmin/runtime/service/transport/grpc"
	"origadmin/application/admin/internal/conf"
)

// NewConn finds a client configuration by service name or convention
// and establishes a gRPC connection.
//
// This is a shared helper for creating gRPC connections to backend services.
// Both gateway and backend services use this to create their inter-service clients.
//
// Parameters:
//   - app: The runtime application instance
//   - bootstrap: Configuration containing client settings
//   - name: Service name (e.g., "auth", "system")
//   - middlewareProvider: Provider for client middlewares
//
// Returns a gRPC connection or an error if configuration is not found.
func NewConn(
	app *runtime.App,
	bootstrap *conf.Config,
	name string,
	middlewareProvider container.ClientMiddlewareProvider,
) (interface{}, error) {
	var clientConfig *transportv1.Client

	// The conventional name for gRPC clients
	convention := fmt.Sprintf("origadmin.service.%s.client.grpc", name)

	// Find the client configuration
	clients := bootstrap.Clients()
	if clients != nil {
		for _, cli := range clients.Configs {
			// Capability Check: Must have gRPC config
			if cli.GetGrpc() == nil {
				continue
			}

			// Smart Matching: Match exact name OR convention name
			if cli.Name == name || cli.Name == convention {
				clientConfig = cli
				break
			}
		}
	}

	if clientConfig == nil {
		return nil, fmt.Errorf("gRPC client config not found for service: %s (checked name: '%s' and '%s')",
			name, name, convention)
	}

	// Get registry provider for service discovery
	registryProvider, err := app.RegistryProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to get registry provider: %w", err)
	}

	// Get discoveries
	discoveries, err := registryProvider.Discoveries()
	if err != nil {
		return nil, fmt.Errorf("failed to get discoveries: %w", err)
	}

	// Get client middlewares
	middlewares, err := middlewareProvider.ClientMiddlewares()
	if err != nil {
		return nil, fmt.Errorf("failed to get client middlewares: %w", err)
	}

	// Create and return the gRPC connection
	return runtimegrpc.NewClient(app.Context(), clientConfig.GetGrpc(), &runtimegrpc.ClientOptions{
		Discoveries:       discoveries,
		ClientMiddlewares: middlewares,
	})
}
