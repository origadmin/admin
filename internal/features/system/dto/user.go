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

// UserRepo is a User repository interface.
type UserRepo interface {
	Get(context.Context, int64, ...UserQueryOption) (*types.User, error)
	Create(context.Context, *types.User, ...UserMutationOption) (*types.User, error)
	Delete(context.Context, int64) error
	Update(context.Context, *types.User, ...UserMutationOption) (*types.User, error)
	List(context.Context, *system.ListUsersRequest, ...UserQueryOption) ([]*types.User, int32, error)
	AddRoleIDs(context.Context, int64, []int64, ...UserMutationOption) error
	GetByUsername(context.Context, string, ...string) (*types.User, error)
	GetRoleIDs(context.Context, int64) ([]int64, error)
	ListResourceByUserID(context.Context, int64, ...UserQueryOption) ([]*types.Resource, error)
	Current(context.Context, int64) (*types.User, error)
	UpdateUserStatus(ctx context.Context, id int64, status int32, options ...UserQueryOption) error
}

type UserMutationOption struct {
	Fields []string
}

type UserQueryOption struct {
	IncludeRoles bool
	OrderFields  []string
	Fields       []string
}
