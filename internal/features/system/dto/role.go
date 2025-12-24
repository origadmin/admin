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

// RoleRepo is a Role repository interface.
type RoleRepo interface {
	Get(context.Context, int64, ...*RoleQueryOptions) (*types.Role, error)
	List(context.Context, *system.ListRolesRequest, ...*RoleQueryOptions) ([]*types.Role, int32, error)
	Create(context.Context, *types.Role, ...*RoleCreateOptions) (*types.Role, error)
	Update(context.Context, *types.Role, ...*RoleUpdateOptions) (*types.Role, error)
	Delete(context.Context, int64) error

	// Business-specific methods
	GetPermissions(context.Context, int64) ([]*types.Permission, error)
	UpdatePermissions(context.Context, int64, []int64) error
}

// RoleQueryOptions specifies options for listing roles.
type RoleQueryOptions struct {
	repo.QueryOption
	WithPermissions bool
}

// RoleCreateOptions specifies options for creating a role.
type RoleCreateOptions struct {
}

// RoleUpdateOptions specifies options for updating a role.
type RoleUpdateOptions struct {
}
