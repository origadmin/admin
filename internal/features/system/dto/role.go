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

// RoleRepo is a Role repository interface.
type RoleRepo interface {
	Get(context.Context, int64, ...RoleQueryOption) (*types.Role, error)
	List(context.Context, *system.ListRolesRequest, ...RoleQueryOption) ([]*types.Role, int32, error)
	Create(context.Context, *types.Role, ...RoleUpdateOption) (*types.Role, error)
	Update(context.Context, *types.Role, ...RoleUpdateOption) (*types.Role, error)
	Delete(context.Context, int64) error
}

type RoleQueryOption struct {
	IncludePermissions bool
	OrderFields        []string
}

type RoleUpdateOption struct {
	Fields []string
}
