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
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
)

type RoleService struct {
	system.UnimplementedRoleServiceServer
	uc *biz.RoleUseCase
}

func NewRoleService(uc *biz.RoleUseCase) *RoleService {
	return &RoleService{uc: uc}
}

func (s *RoleService) ListRoles(ctx context.Context, req *system.ListRolesRequest) (*system.ListRolesResponse, error) {
	queryOpt := dto.ListRolesRequestToQueryOption(req)
	roles, total, err := s.uc.ListRoles(ctx, queryOpt)
	if err != nil {
		return nil, err
	}

	page, pageSize, token, err := db.CalculatePagination(roles, &queryOpt.QueryOption)
	if err != nil {
		return nil, err
	}

	return &system.ListRolesResponse{
		Roles:         roles,
		Total:         total,
		PageSize:      pageSize,
		NextPageToken: token,
		Page:          page,
	}, nil
}
func (s *RoleService) GetRole(ctx context.Context, req *system.GetRoleRequest) (*system.GetRoleResponse, error) {
	role, err := s.uc.GetRole(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "Role not found")
		}
		return nil, err
	}
	return &system.GetRoleResponse{Role: role}, nil
}
func (s *RoleService) CreateRole(ctx context.Context, req *system.CreateRoleRequest) (*system.CreateRoleResponse, error) {
	opts := dto.CreateRoleOptionsFromRequest(req)
	role, err := s.uc.CreateRole(ctx, req.GetRole(), opts)
	if err != nil {
		return nil, err
	}
	return &system.CreateRoleResponse{Role: role}, nil
}
func (s *RoleService) UpdateRole(ctx context.Context, req *system.UpdateRoleRequest) (*system.UpdateRoleResponse, error) {
	opts := dto.UpdateRoleOptionsFromRequest(req)
	role, err := s.uc.UpdateRole(ctx, req.GetRole(), opts)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "Role not found")
		}
		return nil, err
	}
	return &system.UpdateRoleResponse{Role: role}, nil
}
func (s *RoleService) DeleteRole(ctx context.Context, req *system.DeleteRoleRequest) (*system.DeleteRoleResponse, error) {
	err := s.uc.DeleteRole(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "Role not found")
		}
		return nil, err
	}
	return &system.DeleteRoleResponse{}, nil
}
