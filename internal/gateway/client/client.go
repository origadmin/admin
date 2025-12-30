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
var ProviderSet = wire.NewSet(NewAuthClient, NewSystemClient, NewSystemClientSet)

// SystemClientSet holds all the clients for the 'system' service.
// This avoids creating multiple connections to the same downstream service.
type SystemClientSet struct {
	UserClient       system.UserServiceClient
	RoleClient       system.RoleServiceClient
	PermissionClient system.PermissionServiceClient
	ResourceClient   system.ResourceServiceClient
	ViewClient       system.ViewServiceClient
}

// newGRPCConn is a private helper to create a gRPC connection from config.
func newGRPCConn(bootstrap *conf.Config, clientName string) (*grpc.ClientConn, error) {
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

// NewAuthClient creates a new AuthAPI client.
func NewAuthClient(bootstrap *conf.Config) (auth.AuthServiceClient, error) {
	conn, err := newGRPCConn(bootstrap, "client.auth")
	if err != nil {
		return nil, err
	}
	return auth.NewAuthServiceClient(conn), nil
}

// NewSystemClientSet creates a set of clients for the system service.
// It establishes a single gRPC connection and initializes all related clients.
func NewSystemClientSet(bootstrap *conf.Config) (*SystemClientSet, error) {
	conn, err := newGRPCConn(bootstrap, "client.system")
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

// NewSystemClient creates a new SystemAPI client.
// Deprecated: Use NewSystemClientSet instead to access all clients for the system service.
func NewSystemClient(bootstrap *conf.Config) (system.UserServiceClient, error) {
	conn, err := newGRPCConn(bootstrap, "client.system")
	if err != nil {
		return nil, err
	}
	return system.NewUserServiceClient(conn), nil
}
