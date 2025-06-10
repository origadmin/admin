/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/google/wire"
	"github.com/origadmin/runtime"
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

// LocalProviderSet is service providers.
var LocalProviderSet = wire.NewSet(
	NewRegisterBridgeServer,
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
	NewRegisterBridgeServer,
	NewResourceServiceBridgeClient,
	//NewResourceServiceBridge,
	NewRoleServiceBridgeClient,
	//NewRoleServiceBridge,
	NewUserServiceBridgeClient,
	//NewUserServiceBridge,
	NewPermissionServiceBridgeClient,
	//NewPermissionServiceBridge,
)

type SystemServerRegistrar service.ServerRegistrar

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
) SystemServerRegistrar {
	return &RegisterServer{
		Resource:   Resource,
		Role:       Role,
		User:       User,
		Permission: Permission,
	}
}

type RegisterBridgeServer struct {
	Resource   pb.ResourceServiceHookedBridger
	Role       pb.RoleServiceHookedBridger
	User       pb.UserServiceHookedBridger
	Permission pb.PermissionServiceHookedBridger
}

func (s RegisterBridgeServer) Register(ctx context.Context, svc any) {
	switch v := svc.(type) {
	case *service.GRPCServer:
		s.RegisterGRPC(ctx, v)
	case *service.HTTPServer:
		s.RegisterHTTP(ctx, v)
	}
}

func (s RegisterBridgeServer) RegisterHTTP(ctx context.Context, server *service.HTTPServer) {
	log.Info("http server system init")
	pb.RegisterResourceServiceBridgeServer(server, s.Resource)
	pb.RegisterRoleServiceBridgeServer(server, s.Role)
	pb.RegisterUserServiceBridgeServer(server, s.User)
	pb.RegisterPermissionServiceBridgeServer(server, s.Permission)
}

func (s RegisterBridgeServer) RegisterGRPC(ctx context.Context, server *service.GRPCServer) {
	log.Info("grpc server system init")
	//pb.RegisterResourceServiceBridgeServer(server, s.Resource)
	//pb.RegisterRoleServiceBridgeServer(server, s.Role)
	//pb.RegisterUserServiceBridgeServer(server, s.User)
	//pb.RegisterPermissionServiceBridgeServer(server, s.Permission)
}

func NewRegisterBridgeServer(r runtime.Runtime,
	Resource pb.ResourceServiceServer,
	Role pb.RoleServiceServer,
	User pb.UserServiceServer,
	Permission pb.PermissionServiceServer,
) SystemServerRegistrar {
	return &RegisterBridgeServer{
		Resource:   NewResourceServiceHookedBridge(r, Resource),
		Role:       NewRoleServiceHookedBridge(r, Role),
		User:       NewUserServiceHookedBridge(r, User),
		Permission: NewPermissionServiceHookedBridge(r, Permission),
	}
}

var _ service.ServerRegistrar = (*RegisterServer)(nil)
