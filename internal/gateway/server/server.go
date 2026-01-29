/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"
	"fmt"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/maps"
	"github.com/google/wire"

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
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the gateway service servers (HTTP).
func NewServers(
	app *runtime.App,
	serversCfg *transportv1.Servers,
	svc *service.GatewayService,
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
			srv, err := NewHTTPServer(app, serverCfg.GetHttp(), svc, middlewareProvider)
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
	cfg *httpv1.Server,
	svc *service.GatewayService,
	middlewareProvider container.ServerMiddlewareProvider,
) (transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

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
	// Register all services using the GatewayService method.
	svc.RegisterHTTPHandlers(srv)
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
