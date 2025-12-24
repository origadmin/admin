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

// ResourceRepo is a Resource repository interface.
type ResourceRepo interface {
	Get(context.Context, int64, ...*ResourceQueryOptions) (*types.Resource, error)
	List(context.Context, *system.ListResourcesRequest, ...*ResourceQueryOptions) ([]*types.Resource, int32, error)
	Create(context.Context, *types.Resource, ...*ResourceCreateOptions) (*types.Resource, error)
	Update(context.Context, *types.Resource, ...*ResourceUpdateOptions) (*types.Resource, error)
	Delete(context.Context, int64) error
}

// ResourceQueryOptions specifies options for listing resources.
type ResourceQueryOptions struct {
	repo.QueryOption
	WithPermissions bool
}

// ResourceCreateOptions specifies options for creating a resource.
type ResourceCreateOptions struct {
}

// ResourceUpdateOptions specifies options for updating a resource.
type ResourceUpdateOptions struct {
}
