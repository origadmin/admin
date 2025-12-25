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
	Get(context.Context, int64, ...*UserQueryOption) (*types.User, error)
	List(context.Context, ...*UserQueryOption) ([]*types.User, int32, error)
	Create(context.Context, *types.User, string, ...*UserCreateOption) (*types.User, error)
	Update(context.Context, *types.User, ...*UserUpdateOption) (*types.User, error)
	Delete(context.Context, int64) error

	// Business-specific methods
	AddRoleIDs(context.Context, int64, []int64) error
	GetByUsername(context.Context, string) (*types.User, error)
	GetRoleIDs(context.Context, int64) ([]int64, error)
	ListResourceByUserID(context.Context, int64) ([]*types.Resource, error)
	UpdateUserStatus(ctx context.Context, id int64, status int32) error
}

// UserQueryOption specifies options for querying users.
type UserQueryOption struct {
	repo.QueryOption
	WithRoles bool
}

// UserCreateOption specifies options for creating a user.
type UserCreateOption struct {
	// Example: Immediately load roles after creation
	LoadRoles bool
}

// UserUpdateOption specifies options for updating a user.
type UserUpdateOption struct {
	// Example: For partial updates (PATCH)
	UpdateFields []string
}

// ListUsersRequestToQueryOption converts an API request to a query option object.
func ListUsersRequestToQueryOption(req *system.ListUsersRequest) *UserQueryOption {
	if req == nil {
		return &UserQueryOption{}
	}
	return &UserQueryOption{
		QueryOption: repo.OptionFromRequest(req),
	}
}
