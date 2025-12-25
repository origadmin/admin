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

// RoleUseCase is a Role use case.
type RoleUseCase struct {
	repo dto.RoleRepo
}

// NewRoleUseCase new a Role use case.
func NewRoleUseCase(repo dto.RoleRepo) *RoleUseCase {
	return &RoleUseCase{repo: repo}
}

func (uc *RoleUseCase) ListRoles(ctx context.Context, in *system.ListRolesRequest) ([]*types.Role, int32, error) {
	queryOpt := dto.ListRolesRequestToQueryOption(in)
	return uc.repo.List(ctx, queryOpt)
}

func (uc *RoleUseCase) GetRole(ctx context.Context, id int64) (*types.Role, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *RoleUseCase) CreateRole(ctx context.Context, in *types.Role) (*types.Role, error) {
	// Set business-defined default values.
	if in.Status == 0 {
		in.Status = int32(enums.StatusEnabled)
	}

	return uc.repo.Create(ctx, in)
}

func (uc *RoleUseCase) UpdateRole(ctx context.Context, in *types.Role) (*types.Role, error) {
	return uc.repo.Update(ctx, in)
}

func (uc *RoleUseCase) DeleteRole(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *RoleUseCase) UpdateRolePermissions(ctx context.Context, id int64, permissionIDs []int64) error {
	return uc.repo.UpdatePermissions(ctx, id, permissionIDs)
}
