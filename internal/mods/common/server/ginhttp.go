/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	stdhttp "net/http"
	"net/url"

	"github.com/origadmin/toolkits/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/middleware"

	"origadmin/application/admin/internal/configs"
)

func NewGinHTTPServer(bootstrap *configs.Bootstrap, engine *gin.Engine, l log.Logger) *http.Server {
	ms := middleware.NewServer(bootstrap.GetMiddlewares().GetMiddleware())
	var opts = []http.ServerOption{
		http.Middleware(ms...),
	}
	cfg := bootstrap.GetService().GetGins()
	if cfg == nil {
		return nil
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

	cfg.Endpoint = helpers.ServiceDiscoveryEndpoint(cfg.Endpoint, "http", bootstrap.GetService().Host, cfg.Addr)
	log.Infof("Register.GinHttp.Endpoint: %v", cfg.Endpoint)
	ep, _ := url.Parse(cfg.Endpoint)
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
