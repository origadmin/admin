/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

// PermissionRepo is a Permission repository interface.
type PermissionRepo interface {
	Get(context.Context, int64, ...PermissionQueryOption) (*types.Permission, error)
	Create(context.Context, *types.Permission, ...PermissionQueryOption) (*types.Permission, error)
	Delete(context.Context, int64) error
	Update(context.Context, *types.Permission, ...PermissionQueryOption) (*types.Permission, error)
	List(context.Context, *system.ListPermissionsRequest, ...PermissionQueryOption) ([]*types.Permission, int32, error)
}

type PermissionQueryOption struct {
	OrderFields      []string
	Fields           []string
	IncludeResources bool
	IncludeRoles     bool
}
