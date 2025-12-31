/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"context"
	"fmt"

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

// AuthBridgeSet holds all the clients for the 'auth' service.
type AuthBridgeSet struct {
	Auth auth.AuthServiceHTTPServer
	Me   auth.MeServiceHTTPServer
}

// SystemBridgeSet holds all the clients for the 'system' service.
type SystemBridgeSet struct {
	User       system.UserServiceHTTPServer
	Role       system.RoleServiceHTTPServer
	Permission system.PermissionServiceHTTPServer
	Resource   system.ResourceServiceHTTPServer
	View       system.ViewServiceHTTPServer
}

// NewGRPCConn finds a client configuration by name from the bootstrap config
// and establishes a gRPC connection.
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
		return nil, fmt.Errorf("client config not found: %s", clientName)
	}

	grpcConfig := clientConfig.GetGrpc()
	if grpcConfig == nil {
		return nil, fmt.Errorf("gRPC client config not found for: %s", clientName)
	}

	return runtimegrpc.NewClient(context.Background(), grpcConfig, &runtimegrpc.ClientOptions{})
}

// NewAuthClientSet creates a set of clients for the auth service.
func NewAuthClientSet(bootstrap *conf.Config) (*AuthBridgeSet, error) {
	conn, err := NewGRPCConn(bootstrap, "client.auth")
	if err != nil {
		return nil, err
	}
	return &AuthBridgeSet{
		Auth: auth.NewAuthServiceGRPC2HTTP(conn),
		Me:   auth.NewMeServiceGRPC2HTTP(conn),
	}, nil
}

// NewSystemClientSet creates a set of clients for the system service.
func NewSystemClientSet(bootstrap *conf.Config) (*SystemBridgeSet, error) {
	conn, err := NewGRPCConn(bootstrap, "client.system")
	if err != nil {
		return nil, err
	}
	return &SystemBridgeSet{
		User:       system.NewUserServiceGRPC2HTTP(conn),
		Role:       system.NewRoleServiceGRPC2HTTP(conn),
		Permission: system.NewPermissionServiceGRPC2HTTP(conn),
		Resource:   system.NewResourceServiceGRPC2HTTP(conn),
		View:       system.NewViewServiceGRPC2HTTP(conn),
	}, nil
}
