/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/helpers/grpcclient"
)

// ProviderSet is client providers.
var ProviderSet = wire.NewSet(
	NewIdentityBridgeSet,
	NewSystemBridgeSet,
	NewFileManagerBridgeSet,
	NewObjectStoreBridgeSet,
)

const (
	// ServiceNameIdentity is the short name for the auth service.
	ServiceNameIdentity = "identity"
	// ServiceNameSystem is the short name for the system service.
	ServiceNameSystem = "system"
	// ServiceNameFileManager is the short name for the filemanager service.
	ServiceNameFileManager = "filemanager"
	// ServiceNameObjectStore is the short name for the objectstore service.
	ServiceNameObjectStore = "objectstore"
)

// IdentityBridgeSet holds all the clients for the 'auth' service.
type IdentityBridgeSet struct {
	Auth  identity.AuthServiceHTTPServer
	Me    identity.MeServiceHTTPServer
	Admin identity.AdminServiceHTTPServer
}

// SystemBridgeSet holds all the clients for the 'system' service.
type SystemBridgeSet struct {
	User       system.UserServiceHTTPServer
	Role       system.RoleServiceHTTPServer
	Permission system.PermissionServiceHTTPServer
	Resource   system.ResourceServiceHTTPServer
	View       system.ViewServiceHTTPServer
}

// FileManagerBridgeSet holds all the clients for the 'filemanager' service.
type FileManagerBridgeSet struct {
	FileManager filemanager.FileManagerServiceHTTPServer
}

// ObjectStoreBridgeSet holds all the clients for the 'objectstore' service.
type ObjectStoreBridgeSet struct {
	ObjectStore objectstore.ObjectStoreServiceHTTPServer
}

// NewIdentityBridgeSet creates a set of clients for the auth service.
func NewIdentityBridgeSet(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*IdentityBridgeSet, error) {
	// Use the application's root context. This ensures that the client's lifecycle
	// is tied to the application's lifecycle.
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameIdentity, middlewareProvider)
	if err != nil {
		return nil, err
	}
	return &IdentityBridgeSet{
		Auth:  identity.NewAuthServiceGRPC2HTTP(conn),
		Me:    identity.NewMeServiceGRPC2HTTP(conn),
		Admin: identity.NewAdminServiceGRPC2HTTP(conn),
	}, nil
}

// NewSystemBridgeSet creates a set of clients for the system service.
func NewSystemBridgeSet(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*SystemBridgeSet, error) {
	// Use the application's root context.
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameSystem, middlewareProvider)
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

// NewFileManagerBridgeSet creates a set of clients for the filemanager service.
func NewFileManagerBridgeSet(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*FileManagerBridgeSet, error) {
	// Use the application's root context.
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameFileManager, middlewareProvider)
	if err != nil {
		return nil, err
	}
	return &FileManagerBridgeSet{
		FileManager: filemanager.NewFileManagerServiceGRPC2HTTP(conn),
	}, nil
}

// NewObjectStoreBridgeSet creates a set of clients for the objectstore service.
func NewObjectStoreBridgeSet(app *runtime.App, bootstrap *conf.Config, middlewareProvider container.ClientMiddlewareProvider) (*ObjectStoreBridgeSet, error) {
	// Use the application's root context.
	conn, err := grpcclient.NewConn(app, bootstrap, ServiceNameObjectStore, middlewareProvider)
	if err != nil {
		return nil, err
	}
	return &ObjectStoreBridgeSet{
		ObjectStore: objectstore.NewObjectStoreServiceGRPC2HTTP(conn),
	}, nil
}
