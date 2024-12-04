/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	commonpb "origadmin/application/admin/api/v1/services/common"
	pb "origadmin/application/admin/api/v1/services/system"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(RegisterServer), "*"),
	NewMenuAPIService,
	NewMenuAPIServer,
	NewRoleAPIServer,
	NewRoleAPIService,
	NewUserAPIServer,
	NewUserAPIService,
)

type RegisterServer struct {
	Menu  pb.MenuAPIServer
	Role  pb.RoleAPIServer
	User  pb.UserAPIServer
	Login commonpb.LoginAPIServer
}

func (s RegisterServer) GRPC(server *service.GRPCServer) *service.GRPCServer {
	log.Info("grpc server init")
	pb.RegisterMenuAPIServer(server, s.Menu)
	pb.RegisterRoleAPIServer(server, s.Role)
	pb.RegisterUserAPIServer(server, s.User)
	commonpb.RegisterLoginAPIServer(server, s.Login)
	return server
}

func (s RegisterServer) GINS(server *gins.Server) *gins.Server {
	log.Info("gins server init")
	pb.RegisterMenuAPIGINSServer(server, s.Menu)
	pb.RegisterRoleAPIGINSServer(server, s.Role)
	pb.RegisterUserAPIGINSServer(server, s.User)
	commonpb.RegisterLoginAPIGINSServer(server, s.Login)
	return server
}

func (s RegisterServer) HTTP(server *service.HTTPServer) *service.HTTPServer {
	log.Info("http server init")
	pb.RegisterMenuAPIHTTPServer(server, s.Menu)
	pb.RegisterRoleAPIHTTPServer(server, s.Role)
	pb.RegisterUserAPIHTTPServer(server, s.User)
	commonpb.RegisterLoginAPIHTTPServer(server, s.Login)
	return server
}
