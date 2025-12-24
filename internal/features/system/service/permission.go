/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
)

func (s *SystemService) ListPermissions(ctx context.Context, req *system.ListPermissionsRequest) (*system.ListPermissionsResponse, error) {
	permissions, total, err := s.Permission.ListPermissions(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListPermissionsResponse{
		Permissions: permissions,
		Total:       total,
		Page:        req.GetPage(),
		PageSize:    req.GetPageSize(),
	}, nil
}

func (s *SystemService) GetPermission(ctx context.Context, req *system.GetPermissionRequest) (*system.GetPermissionResponse, error) {
	permission, err := s.Permission.GetPermission(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &system.GetPermissionResponse{Permission: permission}, nil
}

func (s *SystemService) CreatePermission(ctx context.Context, req *system.CreatePermissionRequest) (*system.CreatePermissionResponse, error) {
	permission, err := s.Permission.CreatePermission(ctx, req.GetPermission())
	if err != nil {
		return nil, err
	}
	return &system.CreatePermissionResponse{Permission: permission}, nil
}

func (s *SystemService) UpdatePermission(ctx context.Context, req *system.UpdatePermissionRequest) (*system.UpdatePermissionResponse, error) {
	permission, err := s.Permission.UpdatePermission(ctx, req.GetPermission())
	if err != nil {
		return nil, err
	}
	return &system.UpdatePermissionResponse{Permission: permission}, nil
}

func (s *SystemService) DeletePermission(ctx context.Context, req *system.DeletePermissionRequest) (*system.DeletePermissionResponse, error) {
	err := s.Permission.DeletePermission(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &system.DeletePermissionResponse{}, nil
}
