/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/dto"
)

// PermissionUseCase is a Permission use case.
type PermissionUseCase struct {
	repo dto.PermissionRepo
}

// NewPermissionUseCase new a Permission use case.
func NewPermissionUseCase(repo dto.PermissionRepo) *PermissionUseCase {
	return &PermissionUseCase{repo: repo}
}

func (uc *PermissionUseCase) ListPermissions(ctx context.Context, opts ...*dto.PermissionQueryOption) ([]*types.Permission, int32, error) {
	return uc.repo.List(ctx, opts...)
}

func (uc *PermissionUseCase) GetPermission(ctx context.Context, id int64) (*types.Permission, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *PermissionUseCase) CreatePermission(ctx context.Context, in *types.Permission, opts ...*dto.PermissionCreateOption) (*types.Permission, error) {
	// Permission does not have a status field, so no default value is needed here.
	return uc.repo.Create(ctx, in, opts...)
}

func (uc *PermissionUseCase) UpdatePermission(ctx context.Context, in *types.Permission, opts ...*dto.PermissionUpdateOption) (*types.Permission, error) {
	return uc.repo.Update(ctx, in, opts...)
}

func (uc *PermissionUseCase) DeletePermission(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
