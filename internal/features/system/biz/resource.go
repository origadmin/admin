/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/dto"
)

// ResourceUseCase is a Resource use case.
type ResourceUseCase struct {
	repo dto.ResourceRepo
}

func (uc *ResourceUseCase) ListResources(ctx context.Context, in *system.ListResourcesRequest) ([]*types.Resource, int32, error) {
	result, total, err := uc.repo.List(ctx, in)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (uc *ResourceUseCase) GetResource(ctx context.Context, id int64) (*types.Resource, error) {
	result, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *ResourceUseCase) CreateResource(ctx context.Context, in *types.Resource) (*types.Resource, error) {
	result, err := uc.repo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *ResourceUseCase) UpdateResource(ctx context.Context, in *types.Resource) (*types.Resource, error) {
	result, err := uc.repo.Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *ResourceUseCase) DeleteResource(ctx context.Context, id int64) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

// NewResourceUseCase new a Resource use case.
func NewResourceUseCase(repo dto.ResourceRepo) (*ResourceUseCase, error) {
	return &ResourceUseCase{repo: repo}, nil
}
