/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"context"
	"errors"

	"github.com/google/wire"

	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	"github.com/origadmin/runtime/service/transport/grpc"
	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/system"
	confpb "origadmin/application/admin/internal/conf/pb"
)

// ProviderSet is client providers.
var ProviderSet = wire.NewSet(NewAuthClient, NewSystemClient)

// NewAuthClient creates a new AuthAPI client.
func NewAuthClient(bootstrap *confpb.Bootstrap) (auth.AuthServiceClient, error) {
	var clientConfig *transportv1.Client
	if bootstrap.Clients != nil {
		for _, cli := range bootstrap.Clients.Configs {
			if cli.Name == "client.auth" {
				clientConfig = cli
				break
			}
		}
	}

	if clientConfig == nil {
		return nil, errors.New("client config not found: client.auth")
	}

	grpcConfig := clientConfig.GetGrpc()
	if grpcConfig == nil {
		return nil, errors.New("grpc client config not found: client.auth")
	}

	conn, err := grpc.NewClient(context.Background(), grpcConfig, &grpc.ClientOptions{})
	if err != nil {
		return nil, err
	}
	return auth.NewAuthServiceClient(conn), nil
}

// NewSystemClient creates a new SystemAPI client.
func NewSystemClient(bootstrap *confpb.Bootstrap) (system.UserServiceClient, error) {
	var clientConfig *transportv1.Client
	if bootstrap.Clients != nil {
		for _, cli := range bootstrap.Clients.Configs {
			if cli.Name == "client.system" {
				clientConfig = cli
				break
			}
		}
	}

	if clientConfig == nil {
		return nil, errors.New("client config not found: client.system")
	}

	grpcConfig := clientConfig.GetGrpc()
	if grpcConfig == nil {
		return nil, errors.New("grpc client config not found: client.system")
	}

	conn, err := grpc.NewClient(context.Background(), grpcConfig, &grpc.ClientOptions{})
	if err != nil {
		return nil, err
	}
	return system.NewUserServiceClient(conn), nil
}
