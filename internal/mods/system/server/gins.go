/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"net/url"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/toolkits/helpers"

	"origadmin/application/admin/helpers/net"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
)

// NewGINSServer new a gin server.
func NewGINSServer(bootstrap *configs.Bootstrap, l log.Logger) *gins.Server {
	ms := middleware.NewServer(bootstrap.GetMiddlewares().GetMiddleware())
	var opts = []gins.ServerOption{
		gins.Middleware(ms...),
	}
	cfg := bootstrap.GetService().GetGins()
	if cfg == nil {
		cfg = loader.DefaultServiceGins()
	}

	if cfg.Network != "" {
		opts = append(opts, gins.Network(cfg.Network))
	}
	if cfg.Addr != "" {
		opts = append(opts, gins.Address(cfg.Addr))
	}
	if cfg.Timeout != nil {
		opts = append(opts, gins.Timeout(cfg.Timeout.AsDuration()))
	}

	//middlewares, err := bootstrap.LoadMiddlewares(bootstrap.GetServiceName(), bootstrap, l)
	//if err == nil && len(middlewares) > 0 {
	//	opts = append(opts, http.Middleware(middlewares...))
	//}

	if l != nil {
		opts = append(opts, gins.WithLogger(log.With(l, "module", "gins")))
	}

	log.Infof("Register.GinHttp.Endpoint: %v, type: %v, host: %v, addr: %v", cfg.Endpoint, "http", net.GetHostAddr(Host), cfg.Addr)
	cfg.Endpoint = helpers.ServiceDiscoveryEndpoint(cfg.Endpoint, "http", net.GetHostAddr(Host), cfg.Addr)
	log.Infof("Register.GinHttp.Endpoint: %v", cfg.Endpoint)
	ep, _ := url.Parse(cfg.Endpoint)
	opts = append(opts, gins.Endpoint(ep))
	srv := gins.NewServer(opts...)
	//if register != nil {
	//	register(srv)
	//}
	return srv
}
