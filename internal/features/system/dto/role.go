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
	WithPermissionIDs []int64
	WithResourceIDs   []int64
	WithViewIDs       []int64
}

// RoleUpdateOption specifies options for updating a role.
type RoleUpdateOption struct {
	repo.UpdateOption
	WithPermissionIDs []int64
	WithResourceIDs   []int64
	WithViewIDs       []int64
}

// ListRolesRequestToQueryOption converts a ListRolesRequest to a query option object.
func ListRolesRequestToQueryOption(req *system.ListRolesRequest) *RoleQueryOption {
	if req == nil {
		return &RoleQueryOption{}
	}
	return &RoleQueryOption{
		QueryOption: repo.QueryOptionFromRequest(req),
	}
}

// CreateRoleOptionsFromRequest converts a CreateRoleRequest to a create option object.
// It intelligently extracts all association IDs from the request.
func CreateRoleOptionsFromRequest(req *system.CreateRoleRequest) *RoleCreateOption {
	if req == nil {
		return &RoleCreateOption{}
	}
	opts := &RoleCreateOption{
		WithPermissionIDs: req.GetPermissionIds(),
		WithResourceIDs:   req.GetResourceIds(),
		WithViewIDs:       req.GetViewIds(),
	}
	return opts
}

// UpdateRoleOptionsFromRequest converts an UpdateRoleRequest to an update option object.
// It intelligently extracts all association IDs from the request.
func UpdateRoleOptionsFromRequest(req *system.UpdateRoleRequest) *RoleUpdateOption {
	if req == nil {
		return &RoleUpdateOption{}
	}
	opts := &RoleUpdateOption{
		UpdateOption: repo.UpdateOptionFromRequest(req),
		WithPermissionIDs: req.GetPermissionIds(),
		WithResourceIDs:   req.GetResourceIds(),
		WithViewIDs:       req.GetViewIds(),
	}
	return opts
}
