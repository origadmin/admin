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
	Get(context.Context, int64, ...*PermissionQueryOption) (*types.Permission, error)
	List(context.Context, ...*PermissionQueryOption) ([]*types.Permission, int32, error)
	Create(context.Context, *types.Permission, ...*PermissionCreateOption) (*types.Permission, error)
	Update(context.Context, *types.Permission, ...*PermissionUpdateOption) (*types.Permission, error)
	Delete(context.Context, int64) error
}

// PermissionQueryOption specifies options for querying permissions.
type PermissionQueryOption struct {
	repo.QueryOption
	DataScopes    []string
	WithResources bool
	WithRoles     bool
}

// PermissionCreateOption specifies options for creating a permission.
type PermissionCreateOption struct {
}

// PermissionUpdateOption specifies options for updating a permission.
type PermissionUpdateOption struct {
}

// ListPermissionsRequestToQueryOption converts an API request to a query option object.
func ListPermissionsRequestToQueryOption(req *system.ListPermissionsRequest) *PermissionQueryOption {
	if req == nil {
		return &PermissionQueryOption{}
	}
	return &PermissionQueryOption{
		QueryOption: repo.OptionFromRequest(req),
		DataScopes:  req.GetDataScopes(),
	}
}
