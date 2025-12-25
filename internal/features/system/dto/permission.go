/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"
	"time"

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
	DataScopes []string
}

// PermissionCreateOption specifies options for creating a permission.
type PermissionCreateOption struct{}

// PermissionUpdateOption specifies options for updating a permission.
type PermissionUpdateOption struct{}

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
