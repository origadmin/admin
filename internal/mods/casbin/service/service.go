/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/google/wire"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/casbin"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(RegisterServer), "*"),
	NewCasbinSourceServiceServer,
	NewCasbinSourceServiceServerPB,
)

type RegisterServer struct {
	CasbinSource pb.CasbinSourceServiceServer
}

func (s RegisterServer) GRPCServer(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
	pb.RegisterCasbinSourceServiceServer(server, s.CasbinSource)
}

func (s RegisterServer) HTTPServer(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
}

func (s RegisterServer) Server(ctx context.Context, grpcServer *service.GRPCServer, httpServer *service.HTTPServer) {
	s.HTTPServer(ctx, httpServer)
	s.GRPCServer(ctx, grpcServer)
}

var _ service.ServerRegister = (*RegisterServer)(nil)
