/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"context"
	"errors"

	"github.com/google/wire"
	"google.golang.org/grpc"

	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	runtimegrpc "github.com/origadmin/runtime/service/transport/grpc"
	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
)

// ProviderSet is client providers.
var ProviderSet = wire.NewSet(
	NewAuthClientSet,
	NewSystemClientSet,
)

// AuthClientSet holds all the clients for the 'auth' service.
type AuthClientSet struct {
	AuthClient auth.AuthServiceClient
	MeClient   auth.MeServiceClient
}

// SystemClientSet holds all the clients for the 'system' service.
type SystemClientSet struct {
	UserClient       system.UserServiceClient
	RoleClient       system.RoleServiceClient
	PermissionClient system.PermissionServiceClient
	ResourceClient   system.ResourceServiceClient
	ViewClient       system.ViewServiceClient
}

// NewGRPCConn is a helper to create a gRPC connection from config by name.
func NewGRPCConn(bootstrap *conf.Config, clientName string) (*grpc.ClientConn, error) {
	var clientConfig *transportv1.Client
	if bootstrap.Bootstrap.Clients != nil {
		for _, cli := range bootstrap.Bootstrap.Clients.Configs {
			if cli.Name == clientName {
				clientConfig = cli
				break
			}
		}
	}

	if clientConfig == nil {
		return nil, errors.New("client config not found: " + clientName)
	}

	grpcConfig := clientConfig.GetGrpc()
	if grpcConfig == nil {
		return nil, errors.New("grpc client config not found: " + clientName)
	}

	return runtimegrpc.NewClient(context.Background(), grpcConfig, &runtimegrpc.ClientOptions{})
}

// NewAuthClientSet creates a set of clients for the auth service.
func NewAuthClientSet(bootstrap *conf.Config) (*AuthClientSet, error) {
	conn, err := NewGRPCConn(bootstrap, "client.auth")
	if err != nil {
		return nil, err
	}
	return &AuthClientSet{
		AuthClient: auth.NewAuthServiceClient(conn),
		MeClient:   auth.NewMeServiceClient(conn),
	}, nil
}

// NewSystemClientSet creates a set of clients for the system service.
func NewSystemClientSet(bootstrap *conf.Config) (*SystemClientSet, error) {
	conn, err := NewGRPCConn(bootstrap, "client.system")
	if err != nil {
		return nil, err
	}
	return &SystemClientSet{
		UserClient:       system.NewUserServiceClient(conn),
		RoleClient:       system.NewRoleServiceClient(conn),
		PermissionClient: system.NewPermissionServiceClient(conn),
		ResourceClient:   system.NewResourceServiceClient(conn),
		ViewClient:       system.NewViewServiceClient(conn),
	}, nil
}
