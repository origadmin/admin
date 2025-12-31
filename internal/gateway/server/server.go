/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/http"
	"origadmin/application/admin/api/v1/services/auth"
	gatewayAPI "origadmin/application/admin/api/v1/services/gateway"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/gateway/client"
	"origadmin/application/admin/internal/gateway/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the gateway service servers (HTTP).
func NewServers(
	app *runtime.App,
	serversCfg *transportv1.Servers,
	appCfg *conf.Config,
	svc *service.GatewayService,
) ([]transport.Server, error) {
	if serversCfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range serversCfg.GetConfigs() {
		// Filter server configurations by name.
		if serverCfg.GetName() != "gateway" {
			continue
		}

		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), appCfg, svc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		default:
			log.NewHelper(app.Logger()).Warn("protocol", serverCfg.GetProtocol(), "msg", "protocol is not supported")
		}
	}

	if len(transportServers) == 0 {
		return nil, errors.New("no servers named 'gateway' were created")
	}

	return transportServers, nil
}

// NewHTTPServer creates a new HTTP server and registers all downstream service handlers.
func NewHTTPServer(
	app *runtime.App,
	cfg *httpv1.Server,
	appCfg *conf.Config,
	svc *service.GatewayService,
) (transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// 1. Create a new ServeMux for registering routes.
	mux := http.NewServeMux()

	// 2. Register all services to the mux.
	// Register the gateway's own service.
	if err := gatewayAPI.RegisterGatewayServiceHandlerServer(context.Background(), mux, svc); err != nil {
		return nil, err
	}
	// Dynamically register all downstream services.
	if err := registerDownstreamServices(context.Background(), mux, appCfg); err != nil {
		return nil, err
	}

	// 3. Create the server, passing the configured mux.
	middlewareProvider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	mws, err := middlewareProvider.ServerMiddlewares()
	if err != nil {
		return nil, err
	}
	opts := &http.ServerOptions{
		ServerMiddlewares: mws,
		Mux:               mux, // Pass the configured mux to the server.
	}

	return http.NewServer(cfg, opts)
}

// registerDownstreamServices creates connections and registers handlers to the provided ServeMux.
func registerDownstreamServices(ctx context.Context, mux *http.ServeMux, cfg *conf.Config) error {
	// --- Register System Service ---
	systemConn, err := client.NewGRPCConn(cfg, "client.system")
	if err != nil {
		return err
	}
	if err := system.RegisterUserServiceHandler(ctx, mux, systemConn); err != nil {
		return err
	}
	if err := system.RegisterRoleServiceHandler(ctx, mux, systemConn); err != nil {
		return err
	}
	if err := system.RegisterPermissionServiceHandler(ctx, mux, systemConn); err != nil {
		return err
	}
	if err := system.RegisterResourceServiceHandler(ctx, mux, systemConn); err != nil {
		return err
	}
	if err := system.RegisterViewServiceHandler(ctx, mux, systemConn); err != nil {
		return err
	}

	// --- Register Auth Service ---
	authConn, err := client.NewGRPCConn(cfg, "client.auth")
	if err != nil {
		return err
	}
	if err := auth.RegisterAuthServiceHandler(ctx, mux, authConn); err != nil {
		return err
	}
	if err := auth.RegisterMeServiceHandler(ctx, mux, authConn); err != nil {
		return err
	}

	return nil
}
