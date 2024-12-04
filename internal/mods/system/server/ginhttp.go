/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	stdhttp "net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/toolkits/helpers"
	"github.com/origadmin/toolkits/net"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
)

func NewGinHTTPServer(bootstrap *configs.Bootstrap, engine *gin.Engine, l log.Logger) *http.Server {
	ms := middleware.NewServer(bootstrap.GetMiddlewares().GetMiddleware())
	var opts = []http.ServerOption{
		http.Middleware(ms...),
	}
	cfg := bootstrap.GetService().GetGins()
	if cfg == nil {
		cfg = loader.DefaultServiceGins()
	}

	if cfg.Network != "" {
		opts = append(opts, http.Network(cfg.Network))
	}
	if cfg.Addr != "" {
		opts = append(opts, http.Address(cfg.Addr))
	}
	if cfg.Timeout != nil {
		opts = append(opts, http.Timeout(cfg.Timeout.AsDuration()))
	}

	var endpoint string
	if cfg.Endpoint == "" {
		endpoint, _ = helpers.ServiceEndpoint("http", net.HostAddr(Host), cfg.Addr)
	} else {
		endpoint = cfg.Endpoint
	}

	log.Infof("Register.GinHttp.Endpoint: %v", endpoint)
	ep, _ := url.Parse(endpoint)
	opts = append(opts, http.Endpoint(ep))
	srv := http.NewServer(opts...)
	srv.Server = &stdhttp.Server{
		Addr:         cfg.Addr,
		Handler:      engine.Handler(),
		ReadTimeout:  cfg.ReadTimeout.AsDuration(),
		WriteTimeout: cfg.WriteTimeout.AsDuration(),
		IdleTimeout:  cfg.IdleTimeout.AsDuration(),
	}
	return srv
}
