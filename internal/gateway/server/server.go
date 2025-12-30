/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"

	gatewayAPI "origadmin/application/admin/api/v1/services/gateway"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/gateway/service"
	httpv1 "github.com/origadmin/runtime/api/gen/go/config/transport/http/v1"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *confpb.Bootstrap, gw *service.GatewayService, logger log.Logger) (transport.Server,
	error) {
	var opts = []http.ServerOption{
		http.Middleware(
			// Add any HTTP middleware here if needed
		),
	}

	var httpServerConfig *httpv1.Server
	if bootstrap.Servers != nil {
		for _, srv := range bootstrap.Servers.Configs {
			if srv.Protocol == "http" {
				httpServerConfig = srv.GetHttp()
				break
			}
		}
	}

	if httpServerConfig == nil {
		return nil, errors.New("http server config is not found")
	}

	if httpServerConfig.Addr != "" {
		opts = append(opts, http.Address(httpServerConfig.Addr))
	}
	if httpServerConfig.Timeout != nil {
		opts = append(opts, http.Timeout(httpServerConfig.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	gatewayAPI.RegisterGatewayServiceHTTPServer(srv, gw)
	return srv, nil
}
