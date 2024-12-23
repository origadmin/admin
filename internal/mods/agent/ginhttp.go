/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"

	"origadmin/application/admin/internal/configs"
)

type GINRegistrar interface {
	GIN(server gins.IRouter)
}

func NewGinServer(bootstrap *configs.Bootstrap, l log.KLogger) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	return gin.New()
}

func NewHTTPServerAgentX(bootstrap *configs.Bootstrap, l log.KLogger) *http.Server {
	ms := middleware.NewClient(bootstrap.GetMiddleware())
	var opts = []http.ServerOption{
		http.Middleware(ms...),
	}
	//service := bootstrap.GetService()
	//cfg := service.GetGins()
	//if cfg == nil {
	//	return nil
	//}
	//
	//if cfg.Network != "" {
	//	opts = append(opts, http.Network(cfg.Network))
	//}
	//if cfg.Addr != "" {
	//	opts = append(opts, http.Address(cfg.Addr))
	//}
	//if cfg.Timeout != nil {
	//	opts = append(opts, http.Timeout(cfg.Timeout.AsDuration()))
	//}

	//var endpoint string
	//if cfg.Endpoint == "" {
	//	endpoint, _ = helpers.ServiceEndpoint("http", net.HostAddr(service.Host), cfg.Addr)
	//} else {
	//	endpoint = cfg.Endpoint
	//}
	//
	//log.Infof("Register.GinHttp.Endpoint: %v", endpoint)
	//ep, _ := url.Parse(endpoint)
	//opts = append(opts, http.Endpoint(ep))
	srv := http.NewServer(opts...)
	//srv.Server = &stdhttp.Server{
	//	Addr:         cfg.Addr,
	//	ReadTimeout:  cfg.ReadTimeout.AsDuration(),
	//	WriteTimeout: cfg.WriteTimeout.AsDuration(),
	//	IdleTimeout:  cfg.IdleTimeout.AsDuration(),
	//}
	return srv
}
