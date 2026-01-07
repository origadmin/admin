/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"github.com/origadmin/contrib/security"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/helpers/repo"
)

// ResourceRepo is a Resource repository interface.
type ResourceRepo interface {
	Get(context.Context, int64, ...*ResourceQueryOption) (*types.Resource, error)
	List(context.Context, ...*ResourceQueryOption) ([]*types.Resource, int32, error)
	Create(context.Context, *types.Resource, ...*ResourceCreateOption) (*types.Resource, error)
	CreateFromPolicy(ctx context.Context, policy *security.Policy) (*types.Resource, error)
	Update(context.Context, *types.Resource, ...*ResourceUpdateOption) (*types.Resource, error)
	Delete(context.Context, int64) error
}

// ResourceQueryOption specifies options for querying resources.
type ResourceQueryOption struct {
	repo.QueryOption
	WithPermissions bool
}

// ResourceCreateOption specifies options for creating a resource.
type ResourceCreateOption struct {
}

// ResourceUpdateOption specifies options for updating a resource.
type ResourceUpdateOption struct {
	repo.UpdateOption
}

// ListResourcesRequestToQueryOption converts an API request to a query option object.
func ListResourcesRequestToQueryOption(req *system.ListResourcesRequest) *ResourceQueryOption {
	if req == nil {
		return &ResourceQueryOption{}
	}
	return &ResourceQueryOption{
		QueryOption: repo.QueryOptionFromRequest(req),
	}
}

// UpdateResourceRequestToUpdateOption converts an API request to an update option object.
func UpdateResourceRequestToUpdateOption(req *system.UpdateResourceRequest) *ResourceUpdateOption {
	if req == nil {
		return &ResourceUpdateOption{}
	}
	return &ResourceUpdateOption{
		UpdateOption: repo.UpdateOptionFromRequest(req),
	}
}
