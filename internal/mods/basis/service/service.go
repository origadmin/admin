/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
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

func (s RegisterServer) GRPCServer(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server basis init")
	pb.RegisterLoginAPIServer(server, s.Login)
}

func (s RegisterServer) HTTPServer(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server basis init")
	pb.RegisterLoginAPIHTTPServer(server, s.Login)
}

func (s RegisterServer) Server(ctx context.Context, grpcServer *service.GRPCServer, httpServer *service.HTTPServer) {
	s.HTTPServer(ctx, httpServer)
	s.GRPCServer(ctx, grpcServer)
}

var _ service.ServerRegister = (*RegisterServer)(nil)
