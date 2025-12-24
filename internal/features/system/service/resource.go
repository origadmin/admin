/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

func (s *SystemService) ListResources(ctx context.Context, req *system.ListResourcesRequest) (*system.ListResourcesResponse, error) {
	resources, total, err := s.Resource.ListResources(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListResourcesResponse{
		Resources: resources,
		Total:     total,
	}, nil
}

func (s *SystemService) GetResource(ctx context.Context, req *system.GetResourceRequest) (*types.Resource, error) {
	return s.Resource.GetResource(ctx, req.Id)
}

func (s *SystemService) CreateResource(ctx context.Context, req *system.CreateResourceRequest) (*types.Resource, error) {
	return s.Resource.CreateResource(ctx, req.Resource)
}

func (s *SystemService) UpdateResource(ctx context.Context, req *system.UpdateResourceRequest) (*types.Resource, error) {
	return s.Resource.UpdateResource(ctx, req.Resource)
}

func (s *SystemService) DeleteResource(ctx context.Context, req *system.DeleteResourceRequest) (*system.DeleteResourceResponse, error) {
	err := s.Resource.DeleteResource(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &system.DeleteResourceResponse{}, nil
}
