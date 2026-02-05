/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/grpcclient"
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
	Auth  auth.AuthServiceHTTPServer
	Me    auth.MeServiceHTTPServer
	Admin auth.AdminServiceHTTPServer
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
// This function now uses the shared helper from helpers/grpcclient.
func NewGRPCConn(app *runtime.App, bootstrap *conf.Config, name string, middlewareProvider container.ClientMiddlewareProvider) (*grpc.ClientConn, error) {
	conn, err := grpcclient.NewConn(app, bootstrap, name, middlewareProvider)
	if err != nil {
		return nil, err
	}
	return conn.(*grpc.ClientConn), nil
}

// NewAuthBridgeSet creates a set of clients for the auth service.
func NewAuthBridgeSet(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*AuthBridgeSet, error) {
	// Use the application's root context. This ensures that the client's lifecycle
	// is tied to the application's lifecycle.
	conn, err := NewGRPCConn(app, bootstrap, ServiceNameAuth, middlewareProvider)
	if err != nil {
		return nil, err
	}
	return &AuthBridgeSet{
		Auth:  auth.NewAuthServiceGRPC2HTTP(conn),
		Me:    auth.NewMeServiceGRPC2HTTP(conn),
		Admin: auth.NewAdminServiceGRPC2HTTP(conn),
	}, nil
}

// NewSystemBridgeSet creates a set of clients for the system service.
func NewSystemBridgeSet(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*SystemBridgeSet, error) {
	// Use the application's root context.
	conn, err := NewGRPCConn(app, bootstrap, ServiceNameSystem, middlewareProvider)
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
