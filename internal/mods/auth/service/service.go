/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/google/wire"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/auth"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(RegisterServer), "*"),
	NewAuthServiceServerPB,
	//NewAuthServiceHTTPServerPB,
	NewCasbinSourceServiceServerPB,
	//NewCasbinSourceServiceHTTPServerPB,
	NewLoginServiceServerPB,
	//NewLoginServiceHTTPServerPB,
	NewRegisterServer,
)

type RegisterServer struct {
	Auth   pb.AuthServiceServer
	Casbin pb.CasbinSourceServiceServer
	Login  pb.LoginServiceServer
}

func (s RegisterServer) Register(ctx context.Context, svc any) {
	switch v := svc.(type) {
	case *service.GRPCServer:
		s.RegisterGRPC(ctx, v)
	case *service.HTTPServer:
		s.RegisterHTTP(ctx, v)
	}
}

func (s RegisterServer) RegisterGRPC(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
	pb.RegisterAuthServiceServer(server, s.Auth)
	pb.RegisterCasbinSourceServiceServer(server, s.Casbin)
	pb.RegisterLoginServiceServer(server, s.Login)
}

func (s RegisterServer) RegisterHTTP(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	pb.RegisterAuthServiceHTTPServer(server, s.Auth)
	pb.RegisterCasbinSourceServiceHTTPServer(server, s.Casbin)
	pb.RegisterLoginServiceHTTPServer(server, s.Login)
}

func NewRegisterServer(
	Auth pb.AuthServiceServer,
	Casbin pb.CasbinSourceServiceServer,
	Login pb.LoginServiceServer,
) service.ServerRegistrar {
	return &RegisterServer{
		Auth:   Auth,
		Casbin: Casbin,
		Login:  Login,
	}
}

var _ service.ServerRegistrar = (*RegisterServer)(nil)
