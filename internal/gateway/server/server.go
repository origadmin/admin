/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"

	"github.com/google/wire"

	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service/transport"
	"github.com/origadmin/runtime/service/transport/http"
	gatewayAPI "origadmin/application/admin/api/v1/services/gateway"
	"origadmin/application/admin/internal/gateway/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers)

// NewServers creates and configures the gateway service servers (HTTP).
func NewServers(cfg *transportv1.Servers, svc *service.GatewayService, logger log.Logger) ([]transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("servers config is nil")
	}

	var transportServers []transport.Server
	for _, serverCfg := range cfg.GetConfigs() {
		// Filter server configurations by name.
		if serverCfg.GetName() != "gateway" {
			continue
		}

		switch serverCfg.GetProtocol() {
		case "http":
			srv, err := NewHTTPServer(serverCfg.GetHttp(), svc, logger)
			if err != nil {
				return nil, err
			}
			transportServers = append(transportServers, srv)
		default:
			// Log a warning for unsupported protocols but don't return an error
			// to allow other servers to start.
			log.NewHelper(logger).Warnf("protocol is not supported: %s", serverCfg.GetProtocol())
		}
	}

	if len(transportServers) == 0 {
		return nil, errors.New("no servers named 'gateway' were created")
	}

	return transportServers, nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(cfg *httpv1.Server, svc *service.GatewayService, logger log.Logger) (transport.Server, error) {
	if cfg == nil {
		return nil, errors.New("http config is nil")
	}

	// Create server options. The runtime's NewServer will handle middleware
	// and other configurations like CORS based on the provided cfg.
	opts := &http.ServerOptions{}

	// Create the HTTP server.
	srv, err := http.NewServer(cfg, opts)
	if err != nil {
		return nil, err
	}

	// Register the gateway service.
	gatewayAPI.RegisterGatewayServiceHTTPServer(srv, svc)
	return srv, nil
}
