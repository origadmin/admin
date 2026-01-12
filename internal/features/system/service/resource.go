/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/errors"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/helpers/db"
)

type ResourceService struct {
	system.UnimplementedResourceServiceServer
	uc *biz.ResourceUseCase
}

func NewResourceService(uc *biz.ResourceUseCase) *ResourceService {
	return &ResourceService{uc: uc}
}

func (s *ResourceService) ListResources(ctx context.Context, req *system.ListResourcesRequest) (*system.ListResourcesResponse, error) {
	resources, total, err := s.uc.ListResources(ctx, req)
	if err != nil {
		return nil, err
	}

	pageSize := db.GetPageSize(req)
	resp := &system.ListResourcesResponse{
		Resources: resources,
		Total:     total,
		PageSize:  int32(pageSize),
	}

	if req.GetPagingMode() == db.PagingModeCursor {
		// Only generate a next page token if the number of results equals the page size,
		// which implies there might be more data.
		if len(resources) > 0 && len(resources) == pageSize {
			nextToken, err := db.GenerateNextPageToken(resources, req)
			if err != nil {
				return nil, errors.InternalServer("TOKEN_GENERATION_FAILED", err.Error())
			}
			resp.NextPageToken = nextToken
		}
	} else {
		resp.Page = req.GetPage()
	}

	return resp, nil
}

func (s *ResourceService) GetResource(ctx context.Context, req *system.GetResourceRequest) (*system.GetResourceResponse, error) {
	resource, err := s.uc.GetResource(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("RESOURCE_NOT_FOUND", "Resource not found")
		}
		return nil, err
	}
	return &system.GetResourceResponse{Resource: resource}, nil
}

func (s *ResourceService) CreateResource(ctx context.Context, req *system.CreateResourceRequest) (*system.CreateResourceResponse, error) {
	resource, err := s.uc.CreateResource(ctx, req.GetResource())
	if err != nil {
		return nil, err
	}
	return &system.CreateResourceResponse{Resource: resource}, nil
}

func (s *ResourceService) UpdateResource(ctx context.Context, req *system.UpdateResourceRequest) (*system.UpdateResourceResponse, error) {
	resource, err := s.uc.UpdateResource(ctx, req.GetResource())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("RESOURCE_NOT_FOUND", "Resource not found")
		}
		return nil, err
	}
	return &system.UpdateResourceResponse{Resource: resource}, nil
}

func (s *ResourceService) DeleteResource(ctx context.Context, req *system.DeleteResourceRequest) (*system.DeleteResourceResponse, error) {
	err := s.uc.DeleteResource(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("RESOURCE_NOT_FOUND", "Resource not found")
		}
		return nil, err
	}
	return &system.DeleteResourceResponse{}, nil
}
