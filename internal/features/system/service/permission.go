/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"origadmin/application/admin/api/v1/services/system"
)

func (s *SystemService) ListPermissions(ctx context.Context, req *system.ListPermissionsRequest) (*system.ListPermissionsResponse, error) {
	permissions, total, err := s.permission.ListPermissions(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListPermissionsResponse{
		Permissions: permissions,
		Total:       total,
	}, nil
}

func (s *SystemService) GetPermission(ctx context.Context, req *system.GetPermissionRequest) (*system.Permission, error) {
	return s.permission.GetPermission(ctx, req.Id)
}

func (s *SystemService) CreatePermission(ctx context.Context, req *system.CreatePermissionRequest) (*system.Permission, error) {
	return s.permission.CreatePermission(ctx, req.Permission)
}

func (s *SystemService) UpdatePermission(ctx context.Context, req *system.UpdatePermissionRequest) (*system.Permission, error) {
	return s.permission.UpdatePermission(ctx, req.Permission)
}

func (s *SystemService) DeletePermission(ctx context.Context, req *system.DeletePermissionRequest) (*system.DeletePermissionResponse, error) {
	err := s.permission.DeletePermission(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &system.DeletePermissionResponse{}, nil
}
