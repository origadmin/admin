/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/errors"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/data/entity/ent"
)

func (s *SystemService) ListResources(ctx context.Context, req *system.ListResourcesRequest) (*system.ListResourcesResponse, error) {
	resources, total, err := s.Resource.ListResources(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListResourcesResponse{
		Resources: resources,
		Total:     total,
		Page:      req.GetPage(),
		PageSize:  req.GetPageSize(),
	}, nil
}

func (s *SystemService) GetResource(ctx context.Context, req *system.GetResourceRequest) (*system.GetResourceResponse, error) {
	resource, err := s.Resource.GetResource(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("RESOURCE_NOT_FOUND", "Resource not found")
		}
		return nil, err
	}
	return &system.GetResourceResponse{Resource: resource}, nil
}

func (s *SystemService) CreateResource(ctx context.Context, req *system.CreateResourceRequest) (*system.CreateResourceResponse, error) {
	resource, err := s.Resource.CreateResource(ctx, req.GetResource())
	if err != nil {
		return nil, err
	}
	return &system.CreateResourceResponse{Resource: resource}, nil
}

func (s *SystemService) UpdateResource(ctx context.Context, req *system.UpdateResourceRequest) (*system.UpdateResourceResponse, error) {
	resource, err := s.Resource.UpdateResource(ctx, req.GetResource())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("RESOURCE_NOT_FOUND", "Resource not found")
		}
		return nil, err
	}
	return &system.UpdateResourceResponse{Resource: resource}, nil
}

func (s *SystemService) DeleteResource(ctx context.Context, req *system.DeleteResourceRequest) (*system.DeleteResourceResponse, error) {
	err := s.Resource.DeleteResource(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("RESOURCE_NOT_FOUND", "Resource not found")
		}
		return nil, err
	}
	return &system.DeleteResourceResponse{}, nil
}
