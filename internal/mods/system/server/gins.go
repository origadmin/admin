/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"net/url"

	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	"github.com/origadmin/toolkits/env"
	"github.com/origadmin/toolkits/net"

	"origadmin/application/admin/internal/configs"
)

// NewGINSServer new a gin server.
func NewGINSServer(bootstrap *configs.Bootstrap, l log.KLogger, ss ...service.OptionSetting) *gins.Server {
	ms := middleware.NewServer(bootstrap.GetMiddleware())
	//option := settings.ApplyOrZero(ss...)
	var opts = []gins.ServerOption{
		gins.Middleware(ms...),
	}
	//serviceConfig := bootstrap.GetService()
	//cfg := serviceConfig.GetGins()
	//if cfg == nil {
	//	return nil
	//}
	//
	//if cfg.Network != "" {
	//	opts = append(opts, gins.Network(cfg.Network))
	//}
	//if cfg.Addr != "" {
	//	opts = append(opts, gins.Address(cfg.Addr))
	//}
	//if cfg.Timeout != nil {
	//	opts = append(opts, gins.Timeout(cfg.Timeout.AsDuration()))
	//}

	//middlewares, err := bootstrap.LoadMiddlewares(bootstrap.GetServiceName(), bootstrap, l)
	//if err == nil && len(middlewares) > 0 {
	//	opts = append(opts, http.BuildMiddleware(middlewares...))
	//}

	if l != nil {
		opts = append(opts, gins.WithLogger(log.With(l, "module", "gins")))
	}
	log.Infof("GetHostName: %s", env.Var(runtime.DefaultEnvPrefix, "host"))
	host := env.Var(runtime.DefaultEnvPrefix, "host")
	hostIP := env.GetEnv(env.Var(runtime.DefaultEnvPrefix, "host_ip"))
	if hostIP == "" {
		log.Debugf("HostIP is empty, replacing with HostAddr: %s", host)
		hostIP = net.HostAddr(host)
		log.Debugf("HostIP after replacement: %s", hostIP)
	}

	var endpoint string
	//if cfg.Endpoint == "" {
	//	endpoint, _ = helpers.ServiceEndpoint("http", hostIP, cfg.Addr)
	//} else {
	//	endpoint = cfg.Endpoint
	//}
	log.Debugf("GINS.Endpoint: %v", endpoint)
	ep, _ := url.Parse(endpoint)
	opts = append(opts, gins.Endpoint(ep))
	srv := gins.NewServer(opts...)
	//if register != nil {
	//	register(srv)
	//}
	return srv
}
