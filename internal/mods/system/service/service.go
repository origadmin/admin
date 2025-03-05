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
	NewResourceServiceServerPB,
	NewResourceServiceHTTPServerPB,
	NewRoleServiceServerPB,
	NewRoleServiceHTTPServerPB,
	NewUserServiceServerPB,
	NewUserServiceHTTPServerPB,
	NewPersonalServiceServerPB,
	NewPersonalServiceHTTPServerPB,
	NewAuthServiceServerPB,
	NewAuthServiceHTTPServerPB,
	NewPermissionServiceServerPB,
	NewPermissionServiceHTTPServerPB,
)

type RegisterServer struct {
	Resource   pb.ResourceServiceServer
	Role       pb.RoleServiceServer
	User       pb.UserServiceServer
	Auth       pb.AuthServiceServer
	Login      pb.LoginServiceServer
	Personal   pb.PersonalServiceServer
	Permission pb.PermissionServiceServer
}

func (s RegisterServer) GRPCServer(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
	pb.RegisterResourceServiceServer(server, s.Resource)
	pb.RegisterRoleServiceServer(server, s.Role)
	pb.RegisterUserServiceServer(server, s.User)
	pb.RegisterAuthServiceServer(server, s.Auth)
	pb.RegisterLoginServiceServer(server, s.Login)
	pb.RegisterPersonalServiceServer(server, s.Personal)
	pb.RegisterPermissionServiceServer(server, s.Permission)
}

func (s RegisterServer) HTTPServer(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	pb.RegisterResourceServiceHTTPServer(server, s.Resource)
	pb.RegisterRoleServiceHTTPServer(server, s.Role)
	pb.RegisterUserServiceHTTPServer(server, s.User)
	pb.RegisterAuthServiceHTTPServer(server, s.Auth)
	pb.RegisterLoginServiceHTTPServer(server, s.Login)
	pb.RegisterPersonalServiceHTTPServer(server, s.Personal)
	pb.RegisterPermissionServiceHTTPServer(server, s.Permission)
}

func (s RegisterServer) Server(ctx context.Context, grpcServer *service.GRPCServer, httpServer *service.HTTPServer) {
	s.HTTPServer(ctx, httpServer)
	s.GRPCServer(ctx, grpcServer)
}

var _ service.ServerRegister = (*RegisterServer)(nil)
