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

// PermissionUseCase is a Permission use case.
type PermissionUseCase struct {
	repo dto.PermissionRepo
}

func (uc *PermissionUseCase) ListPermissions(ctx context.Context, in *system.ListPermissionsRequest) ([]*types.Permission, int32, error) {
	result, total, err := uc.repo.List(ctx, in)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (uc *PermissionUseCase) GetPermission(ctx context.Context, id int64) (*types.Permission, error) {
	result, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *PermissionUseCase) CreatePermission(ctx context.Context, in *types.Permission) (*types.Permission, error) {
	result, err := uc.repo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *PermissionUseCase) UpdatePermission(ctx context.Context, in *types.Permission) (*types.Permission, error) {
	result, err := uc.repo.Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *PermissionUseCase) DeletePermission(ctx context.Context, id int64) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

// NewPermissionUseCase new a Permission use case.
func NewPermissionUseCase(repo dto.PermissionRepo) (*PermissionUseCase, error) {
	return &PermissionUseCase{repo: repo}, nil
}
