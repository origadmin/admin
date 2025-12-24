/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

func (s *SystemService) ListRoles(ctx context.Context, req *system.ListRolesRequest) (*system.ListRolesResponse, error) {
	roles, total, err := s.Role.ListRoles(ctx, req)
	if err != nil {
		return nil, err
	}

	return &system.ListRolesResponse{
		Roles: roles,
		Total: total,
	}, nil
}
func (s *SystemService) GetRole(ctx context.Context, req *system.GetRoleRequest) (*types.Role, error) {
	return s.Role.GetRole(ctx, req.Id)
}
func (s *SystemService) CreateRole(ctx context.Context, req *system.CreateRoleRequest) (*types.Role, error) {
	return s.Role.CreateRole(ctx, req.Role)
}
func (s *SystemService) UpdateRole(ctx context.Context, req *system.UpdateRoleRequest) (*types.Role, error) {
	return s.Role.UpdateRole(ctx, req.Role)
}
func (s *SystemService) DeleteRole(ctx context.Context, req *system.DeleteRoleRequest) (*system.DeleteRoleResponse, error) {
	err := s.Role.DeleteRole(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &system.DeleteRoleResponse{}, nil
}
