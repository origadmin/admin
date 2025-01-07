/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/google/wire"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(RegisterServer), "*"),
	NewLoginServiceServerPB,
	NewLoginServiceHTTPServerPB,
	NewMenuServiceServerPB,
	NewMenuServiceHTTPServerPB,
	NewRoleServiceServerPB,
	NewRoleServiceHTTPServerPB,
	NewUserServiceServerPB,
	NewUserServiceHTTPServerPB,
	NewCurrentServiceServerPB,
	NewCurrentServiceHTTPServerPB,
	NewAuthServiceServerPB,
	NewAuthServiceHTTPServerPB,
)

type RegisterServer struct {
	Menu    pb.MenuServiceServer
	Role    pb.RoleServiceServer
	User    pb.UserServiceServer
	Auth    pb.AuthServiceServer
	Login   pb.LoginServiceServer
	Current pb.CurrentServiceServer
}

func (s RegisterServer) GRPCServer(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
	pb.RegisterMenuServiceServer(server, s.Menu)
	pb.RegisterRoleServiceServer(server, s.Role)
	pb.RegisterUserServiceServer(server, s.User)
	pb.RegisterAuthServiceServer(server, s.Auth)
	pb.RegisterLoginServiceServer(server, s.Login)
	pb.RegisterCurrentServiceServer(server, s.Current)
}

func (s RegisterServer) HTTPServer(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	pb.RegisterMenuServiceHTTPServer(server, s.Menu)
	pb.RegisterRoleServiceHTTPServer(server, s.Role)
	pb.RegisterUserServiceHTTPServer(server, s.User)
	pb.RegisterAuthServiceHTTPServer(server, s.Auth)
	pb.RegisterLoginServiceHTTPServer(server, s.Login)
	pb.RegisterCurrentServiceHTTPServer(server, s.Current)
}

func (s RegisterServer) Server(ctx context.Context, grpcServer *service.GRPCServer, httpServer *service.HTTPServer) {
	s.HTTPServer(ctx, httpServer)
	s.GRPCServer(ctx, grpcServer)
}

var _ service.ServerRegister = (*RegisterServer)(nil)
