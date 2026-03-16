/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"
	stdhttp "net/http"

	"github.com/google/wire"

	"github.com/origadmin/contrib/transport/watermill"
	"github.com/origadmin/runtime"
	grpcv1 "github.com/origadmin/runtime/api/gen/go/config/transport/grpc/v1"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	watermillv1 "github.com/origadmin/runtime/api/gen/go/config/transport/watermill/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/grpc"
	"github.com/origadmin/runtime/service/transport/http"
	"origadmin/application/admin/internal/broker"
	"origadmin/application/admin/internal/features/system/service"
	"origadmin/application/admin/internal/helpers/providers"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the system service servers (gRPC, HTTP, and Watermill).
func NewServers(
	app *runtime.App,
	cfg *transportv1.Servers,
	svc *service.SystemService,
	policySyncHdl *service.PolicySyncHandler,
) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		if serverCfg.GetName() != "system" && serverCfg.GetName() != "origadmin.service.system" {
			continue
		}
		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), svc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), svc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "watermill":
			srv, err := NewWatermillServer(app, serverCfg.GetWatermill(), policySyncHdl)
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
func NewHTTPServer(app *runtime.App, cfg *httpv1.Server, svc *service.SystemService) (*transport.HTTPServer, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// Fetch middlewares from container with 'feature' tag
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithInScope(runtime.ServerScope),
		runtime.WithInTags(providers.FeatureTag))
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

	_ = svc.RegisterHTTP(app.Context(), srv)

	helper := log.NewHelper(app.Logger())
	_ = srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		helper.Infow(log.DefaultMessageKey, "Registered http handler", "method", method, "path", path)
	})
	return srv, nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(app *runtime.App, cfg *grpcv1.Server, svc *service.SystemService) (*transport.GRPCServer, error) {
	if cfg == nil {
		return nil, errors.New("grpc config is nil")
	}

	// Fetch middlewares from container with 'feature' tag
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithInScope(runtime.ServerScope),
		runtime.WithInTags(providers.FeatureTag))
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

	_ = svc.RegisterGRPC(app.Context(), srv)

	return srv, nil
}

// NewWatermillServer creates a new Watermill server and registers event handlers.
func NewWatermillServer(
	app *runtime.App,
	cfg *watermillv1.Watermill,
	policySyncHdl *service.PolicySyncHandler,
) (transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("watermill config is nil")
	}

	logger := log.NewHelper(app.Logger())
	srv, err := watermill.NewServer(cfg, log.WithLogger(app.Logger()))
	if err != nil {
		return nil, err
	}

	// Register the handler for user-role changes.
	srv.AddConsumerHandler(
		"PolicySyncUserRoleChanged",
		broker.UserRoleAssignedTopic,
		policySyncHdl.HandleUserRoleAssigned,
	)

	// Register the policy sync handler for role-permission changes.
	srv.AddConsumerHandler(
		"PolicySyncRolePolicyChanged",
		broker.RolePolicyChangedTopic,
		policySyncHdl.HandleRolePolicyChanged,
	)

	logger.Info("System Watermill server and policy sync handlers initialized.")
	return srv, nil
}
