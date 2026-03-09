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

	fmPb "origadmin/application/admin/api/v1/services/filemanager"
	fmSvc "origadmin/application/admin/internal/features/filemanager/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the filemanager service servers.
func NewServers(
	app *runtime.App,
	cfg *transportv1.Servers,
	fileManagerSvc *fmSvc.FileManagerService,
) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		if serverCfg.GetName() != "filemanager" && serverCfg.GetName() != "origadmin.service.filemanager" {
			continue
		}
		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), fileManagerSvc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		case "grpc":
			srv, err := NewGRPCServer(app, serverCfg.GetGrpc(), fileManagerSvc)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		default:
			log.Warnf("protocol '%s' is not supported by the filemanager service, skipping", serverCfg.GetProtocol())
		}
	}
	if len(transportServers) == 0 {
		return nil, errors.New("no servers named 'filemanager' or 'origadmin.service.filemanager' were created")
	}
	return transportServers, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	app *runtime.App,
	cfg *httpv1.Server,
	fileManagerSvc *fmSvc.FileManagerService,
) (*transport.HTTPServer, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// Fetch middlewares from container with 'feature' tag
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithScope(runtime.ServerScope),
		runtime.WithInTags("feature"))
	mwMap, err := middleware.GetMiddlewares(app.Context(), h)
	if err != nil {
		return nil, err
	}

	srv, err := http.NewServer(cfg, &http.ServerOptions{
		ServerMiddlewares: mwMap,
	})
	if err != nil {
		return nil, err
	}

	fmPb.RegisterFileManagerServiceHTTPServer(srv, fileManagerSvc)

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
	fileManagerSvc *fmSvc.FileManagerService,
) (*transport.GRPCServer, error) {
	if cfg == nil {
		return nil, errors.New("grpc config is nil")
	}

	// Fetch middlewares from container with 'feature' tag
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithScope(runtime.ServerScope),
		runtime.WithInTags("feature"))
	mwMap, err := middleware.GetMiddlewares(app.Context(), h)
	if err != nil {
		return nil, err
	}

	srv, err := grpc.NewServer(cfg, &grpc.ServerOptions{
		ServerMiddlewares: mwMap,
	})
	if err != nil {
		return nil, err
	}

	fmPb.RegisterFileManagerServiceServer(srv, fileManagerSvc)

	return srv, nil
}
