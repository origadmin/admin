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
	NewRegisterServer,
	NewResourceServiceServerPB,
	NewResourceServiceHTTPServerPB,
	NewRoleServiceServerPB,
	NewRoleServiceHTTPServerPB,
	NewUserServiceServerPB,
	NewUserServiceHTTPServerPB,
	NewPermissionServiceServerPB,
	NewPermissionServiceHTTPServerPB,
)

var RemoteProviderSet = wire.NewSet(
	NewRegisterServer,
	NewResourceServiceBridgeClient,
	//NewResourceServiceBridge,
	NewRoleServiceBridgeClient,
	//NewRoleServiceBridge,
	NewUserServiceBridgeClient,
	//NewUserServiceBridge,
	NewPermissionServiceBridgeClient,
	//NewPermissionServiceBridge,
)

type RegisterServer struct {
	Resource   pb.ResourceServiceServer
	Role       pb.RoleServiceServer
	User       pb.UserServiceServer
	Permission pb.PermissionServiceServer
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
	pb.RegisterResourceServiceServer(server, s.Resource)
	pb.RegisterRoleServiceServer(server, s.Role)
	pb.RegisterUserServiceServer(server, s.User)
	pb.RegisterPermissionServiceServer(server, s.Permission)
}

func (s RegisterServer) RegisterHTTP(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	pb.RegisterResourceServiceHTTPServer(server, s.Resource)
	pb.RegisterRoleServiceHTTPServer(server, s.Role)
	pb.RegisterUserServiceHTTPServer(server, s.User)
	pb.RegisterPermissionServiceHTTPServer(server, s.Permission)
}

func NewRegisterServer(
	Resource pb.ResourceServiceServer,
	Role pb.RoleServiceServer,
	User pb.UserServiceServer,
	Permission pb.PermissionServiceServer,
) *RegisterServer {
	return &RegisterServer{
		Resource:   Resource,
		Role:       Role,
		User:       User,
		Permission: Permission,
	}
}

var _ service.ServerRegistrar = (*RegisterServer)(nil)
