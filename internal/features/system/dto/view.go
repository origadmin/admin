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
	Scope         string
	WithResources bool
}

// ViewCreateOption specifies options for creating a view.
type ViewCreateOption struct {
	WithResourceIDs []int64
	WithRoleIDs     []int64
}

// ViewUpdateOption specifies options for updating a view.
type ViewUpdateOption struct {
	repo.UpdateOption
	WithResourceIDs []int64
	WithRoleIDs     []int64
}

// ListViewsRequestToQueryOption converts an API request to a query option object.
func ListViewsRequestToQueryOption(req *system.ListViewsRequest) *ViewQueryOption {
	if req == nil {
		return &ViewQueryOption{}
	}
	return &ViewQueryOption{
		QueryOption:   repo.QueryOptionFromRequest(req),
		Scope:         req.GetScope(),
		WithResources: req.GetWithResources(),
	}
}

// CreateViewOptionsFromRequest converts a CreateViewRequest to a create option object.
func CreateViewOptionsFromRequest(req *system.CreateViewRequest) *ViewCreateOption {
	if req == nil {
		return &ViewCreateOption{}
	}
	opts := &ViewCreateOption{
		WithResourceIDs: req.GetResourceIds(),
		WithRoleIDs:     req.GetRoleIds(),
	}
	return opts
}

// UpdateViewOptionsFromRequest converts an UpdateViewRequest to an update option object.
func UpdateViewOptionsFromRequest(req *system.UpdateViewRequest) *ViewUpdateOption {
	if req == nil {
		return &ViewUpdateOption{}
	}
	opts := &ViewUpdateOption{
		UpdateOption:    repo.UpdateOptionFromRequest(req),
		WithResourceIDs: req.GetResourceIds(),
		WithRoleIDs:     req.GetRoleIds(),
	}
	return opts
}
