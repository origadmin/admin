/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/origadmin/runtime"
	transportv1 "github.com/origadmin/runtime/api/gen/go/config/transport/v1"
	runtimegrpc "github.com/origadmin/runtime/service/transport/grpc"
	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
)

// ProviderSet is client providers.
var ProviderSet = wire.NewSet(
	NewAuthBridgeSet,
	NewSystemBridgeSet,
)

const (
	// ServiceNameAuth is the short name for the auth service.
	ServiceNameAuth = "auth"
	// ServiceNameSystem is the short name for the system service.
	ServiceNameSystem = "system"
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

// NewGRPCConn finds a client configuration by service name or convention
// and establishes a gRPC connection.
//
// The provided context is used for the client lifecycle.
func NewGRPCConn(ctx context.Context, bootstrap *conf.Config, name string) (*grpc.ClientConn, error) {
	var clientConfig *transportv1.Client

	// The conventional name for gRPC clients
	convention := fmt.Sprintf("origadmin.service.%s.client.grpc", name)

	if bootstrap.Bootstrap.Clients != nil {
		for _, cli := range bootstrap.Bootstrap.Clients.Configs {
			// Capability Check: Must have gRPC config
			if cli.GetGrpc() == nil {
				continue
			}

			// Smart Matching: Match exact name OR convention name
			if cli.Name == name || cli.Name == convention {
				clientConfig = cli
				break
			}
		}
	}

	if clientConfig == nil {
		return nil, fmt.Errorf("gRPC client config not found for service: %s (checked name: '%s' and '%s')", name, name, convention)
	}

	return runtimegrpc.NewClient(ctx, clientConfig.GetGrpc(), &runtimegrpc.ClientOptions{})
}

// NewAuthBridgeSet creates a set of clients for the auth service.
func NewAuthBridgeSet(app *runtime.App, bootstrap *conf.Config) (*AuthBridgeSet, error) {
	// Use the application's root context. This ensures that the client's lifecycle
	// is tied to the application's lifecycle.
	conn, err := NewGRPCConn(app.Context(), bootstrap, ServiceNameAuth)
	if err != nil {
		return nil, err
	}
	return &AuthBridgeSet{
		Auth: auth.NewAuthServiceGRPC2HTTP(conn),
		Me:   auth.NewMeServiceGRPC2HTTP(conn),
	}, nil
}

// NewSystemBridgeSet creates a set of clients for the system service.
func NewSystemBridgeSet(app *runtime.App, bootstrap *conf.Config) (*SystemBridgeSet, error) {
	// Use the application's root context.
	conn, err := NewGRPCConn(app.Context(), bootstrap, ServiceNameSystem)
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
