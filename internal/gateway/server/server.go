/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/http"
	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/gateway/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the gateway service servers (HTTP).
func NewServers(
	app *runtime.App,
	serversCfg *transportv1.Servers,
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
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), svc)
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
	svc *service.GatewayService,
) (transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

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
	}

	srv, err := http.NewServer(cfg, opts)
	if err != nil {
		return nil, err
	}
	// Register all services.
	registerServices(srv, svc)
	return srv, nil
}

func registerServices(
	srv *transport.HTTPServer,
	svc *service.GatewayService,
) {
	system.RegisterUserServiceHTTPServer(srv, svc.System.User)
	system.RegisterRoleServiceHTTPServer(srv, svc.System.Role)
	system.RegisterPermissionServiceHTTPServer(srv, svc.System.Permission)
	system.RegisterResourceServiceHTTPServer(srv, svc.System.Resource)
	system.RegisterViewServiceHTTPServer(srv, svc.System.View)
	auth.RegisterAuthServiceHTTPServer(srv, svc.Auth.Auth)
	auth.RegisterMeServiceHTTPServer(srv, svc.Auth.Me)
	srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		log.Infof("HTTP %s %s", method, path)
	})
}
