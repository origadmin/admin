/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
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
func (s *SystemService) GetRole(ctx context.Context, req *system.GetRoleRequest) (*system.GetRoleResponse, error) {
	role, err := s.Role.GetRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &system.GetRoleResponse{Role: role}, nil
}
func (s *SystemService) CreateRole(ctx context.Context, req *system.CreateRoleRequest) (*system.CreateRoleResponse, error) {
	role, err := s.Role.CreateRole(ctx, req.GetRole())
	if err != nil {
		return nil, err
	}
	return &system.CreateRoleResponse{Role: role}, nil
}
func (s *SystemService) UpdateRole(ctx context.Context, req *system.UpdateRoleRequest) (*system.UpdateRoleResponse, error) {
	role, err := s.Role.UpdateRole(ctx, req.GetRole())
	if err != nil {
		return nil, err
	}
	return &system.UpdateRoleResponse{Role: role}, nil
}
func (s *SystemService) DeleteRole(ctx context.Context, req *system.DeleteRoleRequest) (*system.DeleteRoleResponse, error) {
	err := s.Role.DeleteRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &system.DeleteRoleResponse{}, nil
}
