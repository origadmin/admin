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
	Restore(context.Context, int64) error

	AddRoleIDs(context.Context, int64, []int64) ([]*types.Role, error)
	GetByUsername(context.Context, string) (*types.User, error)
	GetRoleIDs(context.Context, int64) ([]int64, error)
	ListResourceByUserID(context.Context, int64) ([]*types.Resource, error)
	UpdateUserStatus(ctx context.Context, id int64, status int8) error
}

// UserQueryOption specifies options for querying users.
type UserQueryOption struct {
	repo.QueryOption
	WithRoles bool
}

// UserCreateOption specifies options for creating a user.
type UserCreateOption struct {
	WithRoleIDs []int64
}

// UserUpdateOption specifies options for updating a user.
type UserUpdateOption struct {
	repo.UpdateOption
	WithRoleIDs []int64
}

// GetUserRequestToQueryOption converts a GetUserRequest to a query option object,
// always enabling the loading of associated roles for detail views.
func GetUserRequestToQueryOption(req *system.GetUserRequest) *UserQueryOption {
	if req == nil {
		return &UserQueryOption{WithRoles: true}
	}
	return &UserQueryOption{
		QueryOption: repo.QueryOptionFromRequest(req),
		WithRoles:   true, // Always load roles for a single user
	}
}

// ListUsersRequestToQueryOption converts an API request to a query option object.
func ListUsersRequestToQueryOption(req *system.ListUsersRequest) *UserQueryOption {
	if req == nil {
		return &UserQueryOption{}
	}
	return &UserQueryOption{
		QueryOption: repo.QueryOptionFromRequest(req),
		WithRoles:   req.GetWithRoles(),
	}
}

// CreateUserOptionsFromRequest converts a CreateUserRequest to a create option object.
func CreateUserOptionsFromRequest(req *system.CreateUserRequest) *UserCreateOption {
	if req == nil {
		return &UserCreateOption{}
	}
	opts := &UserCreateOption{
		WithRoleIDs: req.GetRoleIds(),
	}
	return opts
}

// UpdateUserOptionsFromRequest converts an UpdateUserRequest to an update option object.
func UpdateUserOptionsFromRequest(req *system.UpdateUserRequest) *UserUpdateOption {
	if req == nil {
		return &UserUpdateOption{}
	}
	opts := &UserUpdateOption{
		UpdateOption: repo.UpdateOptionFromRequest(req),
		WithRoleIDs:  req.GetRoleIds(),
	}
	return opts
}
