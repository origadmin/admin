/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/helpers/repo"
)

// PermissionRepo is a Permission repository interface.
type PermissionRepo interface {
	Get(context.Context, int64, ...*PermissionQueryOptions) (*types.Permission, error)
	List(context.Context, *system.ListPermissionsRequest, ...*PermissionQueryOptions) ([]*types.Permission, int32, error)
	Create(context.Context, *types.Permission, ...*PermissionCreateOptions) (*types.Permission, error)
	Update(context.Context, *types.Permission, ...*PermissionUpdateOptions) (*types.Permission, error)
	Delete(context.Context, int64) error
}

// PermissionQueryOptions specifies options for listing permissions.
type PermissionQueryOptions struct {
	repo.QueryOption
	WithResources bool
	WithRoles     bool
}

// PermissionCreateOptions specifies options for creating a permission.
type PermissionCreateOptions struct {
}

// PermissionUpdateOptions specifies options for updating a permission.
type PermissionUpdateOptions struct {
}
