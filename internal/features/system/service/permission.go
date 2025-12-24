/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

func (s *SystemService) ListPermissions(ctx context.Context, req *system.ListPermissionsRequest) (*system.ListPermissionsResponse, error) {
	permissions, total, err := s.Permission.ListPermissions(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListPermissionsResponse{
		Permissions: permissions,
		Total:       total,
	}, nil
}

func (s *SystemService) GetPermission(ctx context.Context, req *system.GetPermissionRequest) (*types.Permission, error) {
	return s.Permission.GetPermission(ctx, req.Id)
}

func (s *SystemService) CreatePermission(ctx context.Context, req *system.CreatePermissionRequest) (*types.Permission, error) {
	return s.Permission.CreatePermission(ctx, req.Permission)
}

func (s *SystemService) UpdatePermission(ctx context.Context, req *system.UpdatePermissionRequest) (*types.Permission, error) {
	return s.Permission.UpdatePermission(ctx, req.Permission)
}

func (s *SystemService) DeletePermission(ctx context.Context, req *system.DeletePermissionRequest) (*system.DeletePermissionResponse, error) {
	err := s.Permission.DeletePermission(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &system.DeletePermissionResponse{}, nil
}
