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

type RoleService struct {
	system.UnimplementedRoleServiceServer
	uc *biz.RoleUseCase
}

func NewRoleService(uc *biz.RoleUseCase) *RoleService {
	return &RoleService{uc: uc}
}

func (s *RoleService) ListRoles(ctx context.Context, req *system.ListRolesRequest) (*system.ListRolesResponse, error) {
	roles, total, err := s.uc.ListRoles(ctx, req)
	if err != nil {
		return nil, err
	}

	pageSize := db.GetPageSize(req)
	resp := &system.ListRolesResponse{
		Roles:    roles,
		Total:    total,
		PageSize: int32(pageSize),
	}

	if req.GetPagingMode() == db.PagingModeCursor {
		// Only generate a next page token if the number of results equals the page size,
		// which implies there might be more data.
		if len(roles) > 0 && len(roles) == pageSize {
			nextToken, err := db.GenerateNextPageToken(roles, req)
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
	role, err := s.uc.CreateRole(ctx, req.GetRole())
	if err != nil {
		return nil, err
	}
	return &system.CreateRoleResponse{Role: role}, nil
}
func (s *RoleService) UpdateRole(ctx context.Context, req *system.UpdateRoleRequest) (*system.UpdateRoleResponse, error) {
	role, err := s.uc.UpdateRole(ctx, req.GetRole())
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
