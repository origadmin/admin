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
	NewRegisterServer,
	NewAuthServiceServerPB,
	NewCasbinSourceServiceServerPB,
	NewLoginServiceServerPB,
	NewPersonalServiceServerPB,
	NewPersonalServiceHTTPServerPB,
	NewCasbinSourceBiz,
)

var RemoteProviderSet = wire.NewSet(
	NewRegisterServer,
	NewAuthServiceBridgeClient,
	NewCasbinServiceBridgeClient,
	NewLoginServiceBridgeClient,
	NewPersonalServiceBridgeClient,
	NewCasbinSourceClient,

)

type RegisterServer struct {
	Auth     pb.AuthServiceServer
	Casbin   pb.CasbinSourceServiceServer
	Login    pb.LoginServiceServer
	Personal pb.PersonalServiceServer
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
	log.Info("grpc server auth init")
	pb.RegisterAuthServiceServer(server, s.Auth)
	pb.RegisterCasbinSourceServiceServer(server, s.Casbin)
	pb.RegisterLoginServiceServer(server, s.Login)
	pb.RegisterPersonalServiceServer(server, s.Personal)
}

func (s RegisterServer) RegisterHTTP(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server auth init")
	pb.RegisterAuthServiceHTTPServer(server, s.Auth)
	pb.RegisterCasbinSourceServiceHTTPServer(server, s.Casbin)
	pb.RegisterLoginServiceHTTPServer(server, s.Login)
	pb.RegisterPersonalServiceHTTPServer(server, s.Personal)
}

func NewRegisterServer(
	Auth pb.AuthServiceServer,
	Casbin pb.CasbinSourceServiceServer,
	Login pb.LoginServiceServer,
	Personal pb.PersonalServiceServer,
) *RegisterServer {
	return &RegisterServer{
		Auth:     Auth,
		Casbin:   Casbin,
		Login:    Login,
		Personal: Personal,
	}
}

var _ service.ServerRegistrar = (*RegisterServer)(nil)
