/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dto"
)

// RoleUseCase is a Role use case.
type RoleUseCase struct {
	repo dto.RoleRepo
}

// NewRoleUseCase new a Role use case.
func NewRoleUseCase(repo dto.RoleRepo) *RoleUseCase {
	return &RoleUseCase{repo: repo}
}

func (uc *RoleUseCase) ListRoles(ctx context.Context, opts ...*dto.RoleQueryOption) ([]*types.Role, int32, error) {
	return uc.repo.List(ctx, opts...)
}

func (uc *RoleUseCase) GetRole(ctx context.Context, id int64, opts ...*dto.RoleQueryOption) (*types.Role, error) {
	return uc.repo.Get(ctx, id, opts...)
}

// CreateRole creates a new role, ensuring essential fields have valid default values.
func (uc *RoleUseCase) CreateRole(ctx context.Context, in *types.Role, opts ...*dto.RoleCreateOption) (*types.Role, error) {
	// The backend must always enforce data integrity, regardless of frontend behavior.
	if in.Status == 0 {
		in.Status = int32(enums.StatusEnabled)
	}

	return uc.repo.Create(ctx, in, opts...)
}

func (uc *RoleUseCase) UpdateRole(ctx context.Context, in *types.Role, opts ...*dto.RoleUpdateOption) (*types.Role, error) {
	return uc.repo.Update(ctx, in, opts...)
}

func (uc *RoleUseCase) DeleteRole(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *RoleUseCase) UpdateRolePermissions(ctx context.Context, id int64, permissionIDs []int64) error {
	return uc.repo.UpdatePermissions(ctx, id, permissionIDs)
}
