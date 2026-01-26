/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"context"
	"errors"
	stdhttp "net/http"

	"github.com/google/wire"

	"github.com/origadmin/contrib/transport/watermill"
	"github.com/origadmin/runtime"
	grpcv1 "github.com/origadmin/runtime/api/gen/go/config/transport/grpc/v1"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	watermillv1 "github.com/origadmin/runtime/api/gen/go/config/transport/watermill/v1"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/grpc"
	"github.com/origadmin/runtime/service/transport/http"
	authv1 "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/internal/broker"
	"origadmin/application/admin/internal/features/auth/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewServers,
)

// NewServers creates and configures the auth service servers (gRPC, HTTP, and Watermill).
func NewServers(
	app *runtime.App,
	cfg *transportv1.Servers,
	authSvc *service.AuthService,
	meSvc *service.MeService,
	casbinSvc *service.CasbinService,
	policySyncSvc *service.PolicySyncService,
	bootstrap *service.CasbinBootstrap,
	middlewareProvider container.ServerMiddlewareProvider,
) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	// Register BeforeStart hook to bootstrap Casbin policies
	app.AddHookBeforeStart(func(ctx context.Context) error {
		log.Info("Executing Casbin bootstrap before server starts...")
		if err := bootstrap.Bootstrap(ctx); err != nil {
			log.Errorf("Casbin bootstrap failed: %v", err)
			// Don't fail startup, allow admin to manually sync later
		} else {
			log.Info("Casbin bootstrap completed successfully")
		}
		return nil
	})

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		// Check if the server configuration is for the 'auth' service.
		if serverCfg.GetName() != "auth" && serverCfg.GetName() != "origadmin.service.auth" {
			continue
		}

		// Create server based on the specified protocol.
		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), authSvc, meSvc, casbinSvc, middlewareProvider)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), authSvc, meSvc, casbinSvc, middlewareProvider)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "watermill":
			srv, err := NewWatermillServer(app, serverCfg.GetWatermill(), policySyncSvc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		}
	}

	if len(transportServers) == 0 {
		return nil, errors.New("no servers named 'auth' or 'origadmin.service.auth' were created")
	}
	return transportServers, nil
}

// NewWatermillServer creates a new Watermill server and registers event handlers.
func NewWatermillServer(
	_ *runtime.App,
	cfg *watermillv1.Watermill,
	policySyncSvc *service.PolicySyncService,
) (transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("watermill config is nil")
	}

	srv, err := watermill.NewServer(cfg)
	if err != nil {
		return nil, err
	}

	// Register the policy sync handler using AddConsumerHandler
	srv.AddConsumerHandler(
		"PolicySyncUserRoleAssigned",
		broker.UserRoleAssignedTopic,
		policySyncSvc.HandleUserRoleAssigned,
	)

	log.Info("Watermill server and policy sync handler initialized.")
	return srv, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	_ *runtime.App,
	cfg *httpv1.Server,
	authSvc *service.AuthService,
	meSvc *service.MeService,
	casbinSvc *service.CasbinService,
	provider container.ServerMiddlewareProvider,
) (*transport.HTTPServer, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	mws, err := provider.ServerMiddlewares()
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

	// Register HTTP handlers
	authv1.RegisterAuthServiceHTTPServer(srv, authSvc)
	authv1.RegisterMeServiceHTTPServer(srv, meSvc)
	authv1.RegisterCasbinServiceHTTPServer(srv, casbinSvc)

	srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		log.Infof("HTTP %s %s", method, path)
	})
	return srv, nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	_ *runtime.App,
	cfg *grpcv1.Server,
	authSvc *service.AuthService,
	meSvc *service.MeService,
	casbinSvc *service.CasbinService,
	provider container.ServerMiddlewareProvider,
) (*transport.GRPCServer, error) {
	if cfg == nil {
		return nil, errors.New("grpc config is nil")
	}

	mws, err := provider.ServerMiddlewares()
	if err != nil {
		return nil, err
	}
	opts := &grpc.ServerOptions{
		ServerMiddlewares: mws,
	}
	srv, err := grpc.NewServer(cfg, opts)
	if err != nil {
		return nil, err
	}

	// Register gRPC handlers
	authv1.RegisterAuthServiceServer(srv, authSvc)
	authv1.RegisterMeServiceServer(srv, meSvc)
	authv1.RegisterCasbinServiceServer(srv, casbinSvc)

	return srv, nil
}
