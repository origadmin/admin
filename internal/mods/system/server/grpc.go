/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"net/url"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/toolkits/helpers"

	"origadmin/application/admin/api/v1/services/common"
	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/loader"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(bs *configs.Bootstrap, menus pb.MenuAPIServer, l log.Logger) *grpc.Server {
	var opts []grpc.ServerOption
	cfg := bs.GetService().GetGrpc()
	if cfg == nil {
		cfg = loader.DefaultServiceGrpc()
	}
	if cfg.Network != "" {
		opts = append(opts, grpc.Network(cfg.Network))
	}
	if cfg.Addr != "" {
		opts = append(opts, grpc.Address(cfg.Addr))
	}
	if cfg.Timeout != nil {
		opts = append(opts, grpc.Timeout(cfg.Timeout.AsDuration()))
	}

	//middlewares, err := bootstrap.LoadMiddlewares(bs.GetServiceName(), bs, l)
	//if err == nil && len(middlewares) > 0 {
	//	opts = append(opts, grpc.Middleware(middlewares...))
	//}
	ms := middleware.NewServer(bs.GetMiddlewares().GetMiddleware())
	if len(ms) > 0 {
		opts = append(opts, grpc.Middleware(ms...))
	}

	cfg.Endpoint = helpers.ServiceDiscoveryEndpoint(cfg.Endpoint, "grpc", bs.GetService().Host, cfg.Addr)
	log.Infof("Server.GinHttp.Endpoint: %v", cfg.Endpoint)
	ep, _ := url.Parse(cfg.Endpoint)
	opts = append(opts, grpc.Endpoint(ep))
	srv := grpc.NewServer(opts...)
	pb.RegisterMenuAPIServer(srv, menus)
	return srv
}

func RegisterGRPCServer(srv *grpc.Server, menus pb.MenuAPIServer, login common.LoginAPIServer) {
	common.RegisterLoginAPIServer(srv, login)
	pb.RegisterMenuAPIServer(srv, menus)
}
