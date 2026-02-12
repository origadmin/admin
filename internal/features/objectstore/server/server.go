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
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/grpc"
	"github.com/origadmin/runtime/service/transport/http"

	objPb "origadmin/application/admin/api/v1/services/objectstore"
	objSvc "origadmin/application/admin/internal/features/objectstore/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the objectstore service servers.
func NewServers(
	app *runtime.App,
	cfg *transportv1.Servers,
	objectStoreSvc *objSvc.ObjectStoreService,
	middlewareProvider container.ServerMiddlewareProvider,
) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		if serverCfg.GetName() != "objectstore" && serverCfg.GetName() != "origadmin.service.objectstore" {
			continue
		}
		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), objectStoreSvc, middlewareProvider)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), objectStoreSvc, middlewareProvider)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		default:
			log.Warnf("protocol '%s' is not supported by the objectstore service, skipping", serverCfg.GetProtocol())
		}
	}
	return transportServers, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	_ *runtime.App,
	cfg *httpv1.Server,
	objectStoreSvc *objSvc.ObjectStoreService,
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

	objPb.RegisterObjectStoreServiceHTTPServer(srv, objectStoreSvc)

	srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		log.Infof("HTTP %s %s", method, path)
	})
	return srv, nil
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	_ *runtime.App,
	cfg *grpcv1.Server,
	objectStoreSvc *objSvc.ObjectStoreService,
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

	objPb.RegisterObjectStoreServiceServer(srv, objectStoreSvc)

	return srv, nil
}
