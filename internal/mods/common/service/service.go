/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/common"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(RegisterServer), "*"),
	NewLoginAPIService,
	NewLoginAPIServer,
)

type RegisterServer struct {
	Login pb.LoginAPIServer
}

func (s RegisterServer) GRPC(server *service.GRPCServer) *service.GRPCServer {
	log.Info("grpc server common init")
	pb.RegisterLoginAPIServer(server, s.Login)
	return server
}

func (s RegisterServer) GINS(server *gins.Server) *gins.Server {
	log.Info("gins server common init")
	pb.RegisterLoginAPIGINSServer(server, s.Login)
	return server
}

func (s RegisterServer) HTTP(server *service.HTTPServer) *service.HTTPServer {
	log.Info("http server common init")
	pb.RegisterLoginAPIHTTPServer(server, s.Login)
	return server
}
