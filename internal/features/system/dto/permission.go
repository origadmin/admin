/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"
	"time"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/helpers/repo"
)

// PermissionRepo is a Permission repository interface.
type PermissionRepo interface {
	Get(ctx context.Context, id int64, opts ...*PermissionQueryOption) (*types.Permission, error)
	List(ctx context.Context, opts ...*PermissionQueryOption) ([]*types.Permission, int32, error)
	Create(ctx context.Context, in *types.Permission, opts ...*PermissionCreateOption) (*types.Permission, error)
	Update(ctx context.Context, in *types.Permission, opts ...*PermissionUpdateOption) (*types.Permission, error)
	Delete(ctx context.Context, id int64) error
}

// PermissionQueryOption specifies options for querying permissions.
type PermissionQueryOption struct {
	repo.QueryOption
	DataScopes    []string
	WithResources bool
	WithRoles     bool
	WithViews     bool
}

// PermissionCreateOption specifies options for creating a permission.
type PermissionCreateOption struct{}

// PermissionUpdateOption specifies options for updating a permission.
type PermissionUpdateOption struct {
	repo.UpdateOption
}

// PermissionCondition represents a single condition for a permission.
type PermissionCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// PermissionAccessControl defines the access control rules for a permission.
type PermissionAccessControl struct {
	Actions    []string          `json:"actions"`
	Conditions map[string]string `json:"conditions"`
	ValidFrom  *time.Time        `json:"valid_from"`
	ValidUntil *time.Time        `json:"valid_until"`
	Attributes map[string]any    `json:"attributes"`
}

// ListPermissionsRequestToQueryOption converts an API request to a query option object.
func ListPermissionsRequestToQueryOption(req *system.ListPermissionsRequest) *PermissionQueryOption {
	if req == nil {
		return &PermissionQueryOption{}
	}
	return &PermissionQueryOption{
		QueryOption:   repo.QueryOptionFromRequest(req),
		DataScopes:    req.GetDataScopes(),
		WithResources: req.GetWithResources(),
		WithRoles:     req.GetWithRoles(),
		WithViews:     req.GetWithViews(),
	}
}

// UpdatePermissionRequestToUpdateOption converts an API request to an update option object.
func UpdatePermissionRequestToUpdateOption(req *system.UpdatePermissionRequest) *PermissionUpdateOption {
	if req == nil {
		return &PermissionUpdateOption{}
	}
	return &PermissionUpdateOption{
		UpdateOption: repo.UpdateOptionFromRequest(req),
	}
}
