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
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/grpc"
	"github.com/origadmin/runtime/service/transport/http"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/broker"
	"origadmin/application/admin/internal/features/system/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the system service servers (gRPC, HTTP, and Watermill).
func NewServers(
	app *runtime.App,
	cfg *transportv1.Servers,
	svc *service.SystemService,
	policySyncSvc *service.PolicyService,
	middlewareProvider container.ServerMiddlewareProvider,
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
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), svc, middlewareProvider)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), svc, middlewareProvider)
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
func NewHTTPServer(_ *runtime.App, cfg *httpv1.Server, svc *service.SystemService, provider container.ServerMiddlewareProvider) (*transport.HTTPServer, error) {
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
	systemv1.RegisterUserServiceHTTPServer(srv, svc.User)
	systemv1.RegisterRoleServiceHTTPServer(srv, svc.Role)
	systemv1.RegisterPermissionServiceHTTPServer(srv, svc.Permission)
	systemv1.RegisterResourceServiceHTTPServer(srv, svc.Resource)
	systemv1.RegisterViewServiceHTTPServer(srv, svc.View)
	srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		log.Infof("HTTP %s %s", method, path)
	})
	return srv, nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(_ *runtime.App, cfg *grpcv1.Server, svc *service.SystemService, provider container.ServerMiddlewareProvider) (*transport.GRPCServer, error) {
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
	systemv1.RegisterUserServiceServer(srv, svc.User)
	systemv1.RegisterRoleServiceServer(srv, svc.Role)
	systemv1.RegisterPermissionServiceServer(srv, svc.Permission)
	systemv1.RegisterResourceServiceServer(srv, svc.Resource)
	systemv1.RegisterViewServiceServer(srv, svc.View)
	systemv1.RegisterAuthorizationServiceServer(srv, svc.Authorization)

	return srv, nil
}

// NewWatermillServer creates a new Watermill server and registers event handlers.
func NewWatermillServer(
	app *runtime.App,
	cfg *watermillv1.Watermill,
	policySyncSvc *service.PolicyService,
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
		policySyncSvc.HandleUserRoleAssigned,
	)

	// Register the policy sync handler for role-permission changes.
	srv.AddConsumerHandler(
		"PolicySyncRolePolicyChanged",
		broker.RolePolicyChangedTopic,
		policySyncSvc.HandleRolePolicyChanged,
	)

	logger.Info("System Watermill server and policy sync handlers initialized.")
	return srv, nil
}
