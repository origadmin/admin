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
	NewLoginAPIService,
	NewLoginAPIServer,
	NewMenuAPIServer,
	NewMenuAPIService,
	NewRoleAPIServer,
	NewRoleAPIService,
	NewUserAPIServer,
	NewUserAPIService,
	NewCurrentAPIServer,
	NewCurrentAPIService,
	NewAuthAPIServer,
	NewAuthAPIService,
)

type RegisterServer struct {
	Menu    pb.MenuAPIServer
	Role    pb.RoleAPIServer
	User    pb.UserAPIServer
	Auth    pb.AuthAPIServer
	Login   pb.LoginAPIServer
	Current pb.CurrentAPIServer
}

func (s RegisterServer) GRPCServer(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
	pb.RegisterMenuAPIServer(server, s.Menu)
	pb.RegisterRoleAPIServer(server, s.Role)
	pb.RegisterUserAPIServer(server, s.User)
	pb.RegisterAuthAPIServer(server, s.Auth)
	pb.RegisterLoginAPIServer(server, s.Login)
	pb.RegisterCurrentAPIServer(server, s.Current)
}

func (s RegisterServer) HTTPServer(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	pb.RegisterMenuAPIHTTPServer(server, s.Menu)
	pb.RegisterRoleAPIHTTPServer(server, s.Role)
	pb.RegisterUserAPIHTTPServer(server, s.User)
	pb.RegisterAuthAPIHTTPServer(server, s.Auth)
	pb.RegisterLoginAPIHTTPServer(server, s.Login)
	pb.RegisterCurrentAPIHTTPServer(server, s.Current)
}

func (s RegisterServer) Server(ctx context.Context, grpcServer *service.GRPCServer, httpServer *service.HTTPServer) {
	s.HTTPServer(ctx, httpServer)
	s.GRPCServer(ctx, grpcServer)
}

var _ service.ServerRegister = (*RegisterServer)(nil)
