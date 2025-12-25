/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/helpers/repo"
)

// ViewRepo is a View repository interface.
type ViewRepo interface {
	Get(ctx context.Context, id int64, opts ...*ViewQueryOption) (*types.View, error)
	List(ctx context.Context, opts ...*ViewQueryOption) ([]*types.View, int32, error)
	Create(ctx context.Context, in *types.View, opts ...*ViewCreateOption) (*types.View, error)
	Update(ctx context.Context, in *types.View, opts ...*ViewUpdateOption) (*types.View, error)
	Delete(ctx context.Context, id int64) error
}

// ViewQueryOption specifies options for querying views.
type ViewQueryOption struct {
	repo.QueryOption
	Scope string
}

// ViewCreateOption specifies options for creating a view.
type ViewCreateOption struct{}

// ViewUpdateOption specifies options for updating a view.
type ViewUpdateOption struct{}
