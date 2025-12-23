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

// ResourceRepo is a Resource repository interface.
type ResourceRepo interface {
	Get(context.Context, int64, ...ResourceQueryOption) (*types.Resource, error)
	Create(context.Context, *types.Resource, ...ResourceQueryOption) (*types.Resource, error)
	Delete(context.Context, int64) error
	Update(context.Context, *types.Resource, ...ResourceQueryOption) (*types.Resource, error)
	List(context.Context, *system.ListResourcesRequest, ...ResourceQueryOption) ([]*types.Resource, int32, error)
}

type ResourceQueryOption struct {
	OrderFields []string
	Fields      []string
}
