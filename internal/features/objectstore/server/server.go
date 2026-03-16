/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"
	stdhttp "net/http"

	kratosgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/origadmin/runtime"
	grpcv1 "github.com/origadmin/runtime/api/gen/go/config/transport/grpc/v1"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service/transport"
	runtimegrpc "github.com/origadmin/runtime/service/transport/grpc"
	runtimehttp "github.com/origadmin/runtime/service/transport/http"
	"origadmin/application/admin/internal/helpers/providers"

	objPb "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/features/objectstore/dal"
	objSvc "origadmin/application/admin/internal/features/objectstore/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the objectstore service servers.
func NewServers(
	app *runtime.App,
	cfg *transportv1.Servers,
	objectStoreSvc *objSvc.ObjectStoreService,
	storageCfg *dal.LocalStorageConfig,
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
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), objectStoreSvc, storageCfg)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), objectStoreSvc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		}
	}
	return transportServers, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	app *runtime.App,
	cfg *httpv1.Server,
	objectStoreSvc *objSvc.ObjectStoreService,
	storageCfg *dal.LocalStorageConfig,
) (*transport.HTTPServer, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// Fetch middlewares from container with 'feature' tag
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithInScope(runtime.ServerScope),
		runtime.WithInTags(providers.FeatureTag, providers.GatewayTag))
	mwMap, err := middleware.GetMiddlewares(app.Context(), h)
	if err != nil {
		return nil, err
	}

	srv, err := runtimehttp.NewServer(cfg, &runtimehttp.ServerOptions{
		ServerMiddlewares: mwMap,
	})
	if err != nil {
		return nil, err
	}

	objPb.RegisterObjectStoreServiceHTTPServer(srv, objectStoreSvc)
	// Manually register handlers for paths not covered by proto (e.g., PUT for multipart)
	objectStoreSvc.RegisterHandlers(srv)

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
	objectStoreSvc *objSvc.ObjectStoreService,
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

	srv, err := runtimegrpc.NewServer(cfg, &runtimegrpc.ServerOptions{
		ServerOptions:     []kratosgrpc.ServerOption{kratosgrpc.Options(grpc.MaxRecvMsgSize(512 << 20))},
		ServerMiddlewares: mwMap,
	})
	if err != nil {
		return nil, err
	}
	objPb.RegisterObjectStoreServiceServer(srv, objectStoreSvc)
	return srv, nil
}
