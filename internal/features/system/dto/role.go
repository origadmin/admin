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
	Get(context.Context, int64, ...*RoleQueryOption) (*types.Role, error)
	List(context.Context, ...*RoleQueryOption) ([]*types.Role, int32, error)
	Create(context.Context, *types.Role, ...*RoleCreateOption) (*types.Role, error)
	Update(context.Context, *types.Role, ...*RoleUpdateOption) (*types.Role, error)
	Delete(context.Context, int64) error

	// Business-specific methods
	GetPermissions(context.Context, int64) ([]*types.Permission, error)
	UpdatePermissions(context.Context, int64, []int64) error
}

// RoleQueryOption specifies options for querying roles.
type RoleQueryOption struct {
	repo.QueryOption
	WithPermissions bool
}

// RoleCreateOption specifies options for creating a role.
type RoleCreateOption struct {
}

// RoleUpdateOption specifies options for updating a role.
type RoleUpdateOption struct {
	repo.UpdateOption
}

// ListRolesRequestToQueryOption converts an API request to a query option object.
func ListRolesRequestToQueryOption(req *system.ListRolesRequest) *RoleQueryOption {
	if req == nil {
		return &RoleQueryOption{}
	}
	return &RoleQueryOption{
		QueryOption: repo.QueryOptionFromRequest(req),
	}
}

// UpdateRoleRequestToUpdateOption converts an API request to an update option object.
func UpdateRoleRequestToUpdateOption(req *system.UpdateRoleRequest) *RoleUpdateOption {
	if req == nil {
		return &RoleUpdateOption{}
	}
	return &RoleUpdateOption{
		UpdateOption: repo.UpdateOptionFromRequest(req),
	}
}
