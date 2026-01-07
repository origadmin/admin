/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dto"
)

// ResourceUseCase is a Resource use case.
type ResourceUseCase struct {
	repo dto.ResourceRepo
}

// NewResourceUseCase new a Resource use case.
func NewResourceUseCase(repo dto.ResourceRepo) *ResourceUseCase {
	return &ResourceUseCase{repo: repo}
}

func (uc *ResourceUseCase) ListResources(ctx context.Context, in *system.ListResourcesRequest) ([]*types.Resource, int32, error) {
	queryOpt := dto.ListResourcesRequestToQueryOption(in)
	return uc.repo.List(ctx, queryOpt)
}

func (uc *ResourceUseCase) GetResource(ctx context.Context, id int64) (*types.Resource, error) {
	return uc.repo.Get(ctx, id)
}

// CreateResource creates a new resource, intended for use by external APIs (e.g., frontend).
// It handles manual creation logic.
func (uc *ResourceUseCase) CreateResource(ctx context.Context, in *types.Resource) (*types.Resource, error) {
	// The backend must always enforce data integrity.
	if in.Status == 0 {
		in.Status = int32(enums.StatusEnabled)
	}

	return uc.repo.Create(ctx, in)
}

// CreateResourceFromPolicy creates a new resource based on a security policy definition.
// This is intended for internal use, like database seeding.
func (uc *ResourceUseCase) CreateResourceFromPolicy(ctx context.Context, input *dto.ResourceFromPolicyInput) (*types.Resource, error) {
	return uc.repo.CreateFromPolicy(ctx, input)
}

func (uc *ResourceUseCase) UpdateResource(ctx context.Context, in *types.Resource) (*types.Resource, error) {
	return uc.repo.Update(ctx, in)
}

func (uc *ResourceUseCase) DeleteResource(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
