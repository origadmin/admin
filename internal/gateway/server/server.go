/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"

	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service/transport"
	runtimehttp "github.com/origadmin/runtime/service/transport/http"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/gateway/service"
	"origadmin/application/admin/internal/gateway/web"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewServers,
	NewObjectStoreProxy,
)

// NewServers creates and configures the gateway service servers.
func NewServers(
	app *runtime.App,
	bootstrap *confpb.Bootstrap,
	gatewayService *service.GatewayService,
	objectStoreProxy *ObjectStoreProxy,
) ([]transport.Server, error) {
	cfg := bootstrap.GetServers()
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		if serverCfg.GetName() != "gateway" {
			continue
		}
		if serverCfg.GetProtocol() == "http" {
			srv, err := NewHTTPServer(app, bootstrap, serverCfg.GetHttp(), gatewayService, objectStoreProxy)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		}
	}
	return transportServers, nil
}

func NewHTTPServer(
	app *runtime.App,
	_ *confpb.Bootstrap,
	cfg *httpv1.Server,
	gatewayService *service.GatewayService,
	objectStoreProxy *ObjectStoreProxy,
) (*transport.HTTPServer, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// Fetch middlewares from container with 'gateway' tag as a map
	h := app.Container().In(runtime.CategoryMiddleware,
		runtime.WithScope(runtime.ServerScope),
		runtime.WithInTags("gateway"))
	mwMap, err := middleware.GetMiddlewares(app.Context(), h)
	if err != nil {
		return nil, err
	}

	serverOpts := []kratoshttp.ServerOption{
		kratoshttp.PathPrefix(conf.APIPrefix),
	}

	// Use standard runtime server creation
	srv, err := runtimehttp.NewServer(cfg, &runtimehttp.ServerOptions{
		ServerMiddlewares: mwMap,
		ServerOptions:     serverOpts,
	})
	if err != nil {
		return nil, err
	}

	// 1. Register generated service handlers
	gatewayService.RegisterHTTPHandlers(srv)

	// 2. Register custom storage proxy handlers
	objectStoreProxy.RegisterHandlers(srv)

	// 3. UI Handler (Conditional)
	if webUIHandler, err := web.GetHandler(); err == nil && webUIHandler != nil {
		srv.HandlePrefix("/", webUIHandler)
	}

	return srv, nil
}
