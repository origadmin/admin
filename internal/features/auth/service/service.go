/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/auth"
)

type AuthServerRegistrar service.ServerRegistrar

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
) AuthServerRegistrar {
	return &RegisterServer{
		Auth:     Auth,
		Casbin:   Casbin,
		Login:    Login,
		Personal: Personal,
	}
}

type RegisterBridgeServer struct {
	Auth     pb.AuthServiceHookedBridger
	Casbin   pb.CasbinSourceServiceHookedBridger
	Login    pb.LoginServiceHookedBridger
	Personal pb.PersonalServiceHookedBridger
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
	log.Info("http server auth init")
	pb.RegisterAuthServiceBridgeServer(server, s.Auth)
	pb.RegisterCasbinSourceServiceBridgeServer(server, s.Casbin)
	pb.RegisterLoginServiceBridgeServer(server, s.Login)
	pb.RegisterPersonalServiceBridgeServer(server, s.Personal)
}

func (s RegisterBridgeServer) RegisterGRPC(ctx context.Context, server *service.GRPCServer) {
	log.Info("http server system init")
	//pb.RegisterResourceServiceBridgeServer(server, s.Resource)
	//pb.RegisterRoleServiceBridgeServer(server, s.Role)
	//pb.RegisterUserServiceBridgeServer(server, s.User)
	//pb.RegisterPermissionServiceBridgeServer(server, s.Permission)
}

func NewRegisterBridgeServer(r runtime.Runtime,
	Auth pb.AuthServiceServer,
	Casbin pb.CasbinSourceServiceServer,
	Login pb.LoginServiceServer,
	Personal pb.PersonalServiceServer,
) AuthServerRegistrar {
	return &RegisterBridgeServer{
		Auth:     NewAuthServiceHookedBridge(r, Auth),
		Casbin:   NewCasbinServiceHookedBridge(r, Casbin),
		Login:    NewLoginServiceHookedBridge(r, Login),
		Personal: NewPersonalServiceHookedBridge(r, Personal),
	}
}

var _ service.ServerRegistrar = (*RegisterServer)(nil)
