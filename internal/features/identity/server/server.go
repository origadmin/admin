/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"
	stdhttp "net/http"

	"github.com/google/wire"

	"github.com/origadmin/runtime"
	grpcv1 "github.com/origadmin/runtime/api/gen/go/config/transport/grpc/v1"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/grpc"
	"github.com/origadmin/runtime/service/transport/http"
	identityv1 "origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/internal/features/identity/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the system service servers (gRPC, HTTP, and Watermill).
func NewServers(
	app *runtime.App,
	cfg *transportv1.Servers,
	identitySvc *service.AuthService,
	meSvc *service.MeService,
	adminSvc *service.AdminService,
) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		if serverCfg.GetName() != "identity" && serverCfg.GetName() != "origadmin.service.identity" {
			continue
		}
		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), identitySvc, meSvc, adminSvc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), identitySvc, meSvc, adminSvc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		default:
			// Gracefully ignore unsupported protocols for this service
			log.Warnf("protocol '%s' is not supported by the system service, skipping", serverCfg.GetProtocol())
		}
	}
	if len(transportServers) == 0 {
		return nil, errors.New("no servers named 'system' or 'origadmin.service.system' were created")
	}
	return transportServers, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	app *runtime.App,
	cfg *httpv1.Server,
	identitySvc *service.AuthService,
	meSvc *service.MeService,
	adminSvc *service.AdminService,
) (*transport.HTTPServer, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// Fetch middlewares from container with 'feature' tag
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithInScope(runtime.ServerScope),
		runtime.WithInTags("feature"))
	mwMap, err := middleware.GetMiddlewares(app.Context(), h)
	if err != nil {
		return nil, err
	}

	opts := &http.ServerOptions{
		ServerMiddlewares: mwMap,
	}

	srv, err := http.NewServer(cfg, opts)
	if err != nil {
		return nil, err
	}

	identityv1.RegisterAuthServiceHTTPServer(srv, identitySvc)
	identityv1.RegisterMeServiceHTTPServer(srv, meSvc)
	identityv1.RegisterAdminServiceHTTPServer(srv, adminSvc)
	helper := log.NewHelper(app.Logger())
	srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		helper.Infow(log.DefaultMessageKey, "Registered http handler", "method", method, "path", path)
	})
	return srv, nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	app *runtime.App,
	cfg *grpcv1.Server,
	identitySvc *service.AuthService,
	meSvc *service.MeService,
	adminSvc *service.AdminService,
) (*transport.GRPCServer, error) {
	if cfg == nil {
		return nil, errors.New("grpc config is nil")
	}

	// Fetch middlewares from container with 'feature' tag
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithInScope(runtime.ServerScope),
		runtime.WithInTags("feature"))
	mwMap, err := middleware.GetMiddlewares(app.Context(), h)
	if err != nil {
		return nil, err
	}

	opts := &grpc.ServerOptions{
		ServerMiddlewares: mwMap,
	}
	srv, err := grpc.NewServer(cfg, opts)
	if err != nil {
		return nil, err
	}

	identityv1.RegisterAuthServiceServer(srv, identitySvc)
	identityv1.RegisterMeServiceServer(srv, meSvc)
	identityv1.RegisterAdminServiceServer(srv, adminSvc)

	return srv, nil
}
