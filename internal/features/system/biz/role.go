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

// RoleUseCase is a Role use case.
type RoleUseCase struct {
	repo dto.RoleRepo
}

func (uc *RoleUseCase) ListRoles(ctx context.Context, in *system.ListRolesRequest) ([]*types.Role, int32, error) {
	result, total, err := uc.repo.List(ctx, in)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (uc *RoleUseCase) GetRole(ctx context.Context, id int64) (*types.Role, error) {
	result, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *RoleUseCase) CreateRole(ctx context.Context, in *types.Role) (*types.Role, error) {
	result, err := uc.repo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *RoleUseCase) UpdateRole(ctx context.Context, in *types.Role) (*types.Role, error) {
	result, err := uc.repo.Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *RoleUseCase) DeleteRole(ctx context.Context, id int64) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

// NewRoleUseCase new a Role use case.
func NewRoleUseCase(repo dto.RoleRepo) (*RoleUseCase, error) {
	return &RoleUseCase{repo: repo}, nil
}
