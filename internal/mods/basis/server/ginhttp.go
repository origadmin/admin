/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/toolkits/env"
	"github.com/origadmin/toolkits/net"

	"origadmin/application/admin/internal/configs"
)

func NewGinHTTPServer(bootstrap *configs.Bootstrap, engine *gin.Engine, l log.KLogger) *http.Server {
	ms := middleware.NewServer(bootstrap.GetMiddleware())
	var opts = []http.ServerOption{
		http.Middleware(ms...),
	}
	//cfg := bootstrap.GetService().GetGins()
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

	log.Infof("GetHostName: %s", env.Var(runtime.DefaultEnvPrefix, "host"))
	host := env.Var(runtime.DefaultEnvPrefix, "host")
	hostIP := env.GetEnv(env.Var(runtime.DefaultEnvPrefix, "host_ip"))
	if hostIP == "" {
		log.Debugf("HostIP is empty, replacing with HostAddr: %s", host)
		hostIP = net.HostAddr(host)
		log.Debugf("HostIP after replacement: %s", hostIP)
	}

	//if endpoint, err := helpers.ServiceEndpoint("http", hostIP, cfg.Addr); err == nil {
	//	cfg.Endpoint = endpoint
	//}
	//log.Infof("Register.GinHttp.Endpoint: %v", cfg.Endpoint)
	//ep, _ := url.Parse(cfg.Endpoint)
	//opts = append(opts, http.Endpoint(ep))
	srv := http.NewServer(opts...)
	//srv.Server = &stdhttp.Server{
	//	Addr:         cfg.Addr,
	//	Handler:      engine.Handler(),
	//	ReadTimeout:  cfg.ReadTimeout.AsDuration(),
	//	WriteTimeout: cfg.WriteTimeout.AsDuration(),
	//	IdleTimeout:  cfg.IdleTimeout.AsDuration(),
	//}
	return srv
}
