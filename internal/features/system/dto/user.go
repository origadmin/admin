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

// UserRepo is a User repository interface.
type UserRepo interface {
	Get(context.Context, int64, ...*UserQueryOptions) (*types.User, error)
	List(context.Context, *system.ListUsersRequest, ...*UserQueryOptions) ([]*types.User, int32, error)
	Create(context.Context, *types.User, ...*UserCreateOptions) (*types.User, error)
	Update(context.Context, *types.User, ...*UserUpdateOptions) (*types.User, error)
	Delete(context.Context, int64) error

	// Business-specific methods
	AddRoleIDs(context.Context, int64, []int64) error
	GetByUsername(context.Context, string) (*types.User, error)
	GetRoleIDs(context.Context, int64) ([]int64, error)
	ListResourceByUserID(context.Context, int64) ([]*types.Resource, error)
	UpdateUserStatus(ctx context.Context, id int64, status int32) error
}

// UserQueryOptions specifies options for listing users.
type UserQueryOptions struct {
	repo.QueryOption
	WithRoles bool
}

// UserCreateOptions specifies options for creating a user.
type UserCreateOptions struct {
	// Example: Immediately load roles after creation
	LoadRoles bool
}

// UserUpdateOptions specifies options for updating a user.
type UserUpdateOptions struct {
	// Example: For partial updates (PATCH)
	UpdateFields []string
}
