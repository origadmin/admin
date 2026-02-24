/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"
	"fmt"
	stdhttp "net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/maps"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/origadmin/runtime"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/http"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/gateway/service"
	"origadmin/application/admin/internal/gateway/web"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewServers,
	NewObjectStoreProxy, // Register the Proxy provider
)

// NewServers creates and configures the gateway service servers (HTTP).
func NewServers(
	app *runtime.App,
	bootstrap *conf.Config, // Pass bootstrap config
	serversCfg *transportv1.Servers,
	svc *service.GatewayService,
	proxy *ObjectStoreProxy, // Inject the Proxy
	middlewareProvider container.ServerMiddlewareProvider,
) ([]transport.Server, error) {
	if serversCfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range serversCfg.GetConfigs() {
		// Filter server configurations by name.
		if serverCfg.GetName() != "gateway" && serverCfg.GetName() != "origadmin.server.gateway" {
			continue
		}

		switch serverCfg.GetProtocol() {
		case "http":
			// Pass proxy to NewHTTPServer
			srv, err := NewHTTPServer(app, bootstrap, serverCfg.GetHttp(), svc, proxy, middlewareProvider)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		default:
			log.NewHelper(app.Logger()).Warn("protocol", serverCfg.GetProtocol(), "msg", "protocol is not supported")
		}
	}

	if len(transportServers) == 0 {
		return nil, errors.New("no servers named 'gateway' or 'origadmin.server.gateway' were created")
	}

	return transportServers, nil
}

// NewHTTPServer creates a new HTTP server and registers all downstream service handlers.
func NewHTTPServer(
	app *runtime.App,
	bootstrap *conf.Config, // Pass bootstrap config
	cfg *httpv1.Server,
	svc *service.GatewayService,
	proxy *ObjectStoreProxy, // Inject the Proxy
	middlewareProvider container.ServerMiddlewareProvider,
) (transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// Set a longer timeout for the HTTP server to accommodate file operations.
	cfg.Timeout = durationpb.New(30 * time.Second)

	mws, err := middlewareProvider.ServerMiddlewares()
	if err != nil {
		return nil, err
	}

	serverOpts := []kratoshttp.ServerOption{
		kratoshttp.PathPrefix(conf.APIPrefix),
	}

	logger := log.NewHelper(app.Logger())
	logger.Infow("msg", "Registering middleware", "middlewares", maps.Keys(mws))
	opts := &http.ServerOptions{
		ServerOptions:     serverOpts,
		ServerMiddlewares: mws,
	}

	srv, err := http.NewServer(cfg, opts)
	if err != nil {
		return nil, err
	}

	// 1. Register generated gRPC-Gateway handlers (standard API)
	svc.RegisterHTTPHandlers(srv)

	// 2. Register Custom Proxy Handlers (Frontend-friendly API)
	// Register each handler separately with proper middleware support
	proxy.RegisterHandlers(srv)

	// Try to get the handler for the embedded Web UI.
	webUIHandler, err := web.GetHandler()
	if err == nil {
		logger.Infow("msg", "Embedded Web UI is enabled and will be served.")
		srv.HandlePrefix("/", webUIHandler)
	} else {
		logger.Warnw("msg", "Embedded Web UI is disabled. To enable, build with '-tags embed_ui'.")
	}

	// Log all registered HTTP routes for debugging and verification
	srv.WalkHandle(func(method, path string, handler stdhttp.HandlerFunc) {
		fmt.Printf("HTTP %s %s\n", method, path)
	})
	return srv, nil
}
