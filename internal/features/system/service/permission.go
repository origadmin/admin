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

type PermissionService struct {
	system.UnimplementedPermissionServiceServer
	uc *biz.PermissionUseCase
}

func NewPermissionService(uc *biz.PermissionUseCase) *PermissionService {
	return &PermissionService{uc: uc}
}

func (s *PermissionService) ListPermissions(ctx context.Context, req *system.ListPermissionsRequest) (*system.ListPermissionsResponse, error) {
	queryOpt := dto.ListPermissionsRequestToQueryOption(req)
	permissions, total, err := s.uc.ListPermissions(ctx, queryOpt)
	if err != nil {
		return nil, err
	}

	pageSize := db.GetPageSize(req)
	resp := &system.ListPermissionsResponse{
		Permissions: permissions,
		Total:       total,
		PageSize:    int32(pageSize),
	}

	if req.GetPagingMode() == db.PagingModeCursor {
		// Only generate a next page token if the number of results equals the page size,
		// which implies there might be more data.
		if len(permissions) > 0 && len(permissions) == pageSize {
			nextToken, err := db.GenerateNextPageToken(permissions, req)
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

func (s *PermissionService) GetPermission(ctx context.Context, req *system.GetPermissionRequest) (*system.GetPermissionResponse, error) {
	permission, err := s.uc.GetPermission(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("PERMISSION_NOT_FOUND", "Permission not found")
		}
		return nil, err
	}
	return &system.GetPermissionResponse{Permission: permission}, nil
}

func (s *PermissionService) CreatePermission(ctx context.Context, req *system.CreatePermissionRequest) (*system.CreatePermissionResponse, error) {
	opts := dto.CreatePermissionOptionsFromRequest(req)
	permission, err := s.uc.CreatePermission(ctx, req.GetPermission(), opts)
	if err != nil {
		return nil, err
	}
	return &system.CreatePermissionResponse{Permission: permission}, nil
}

func (s *PermissionService) UpdatePermission(ctx context.Context, req *system.UpdatePermissionRequest) (*system.UpdatePermissionResponse, error) {
	opts := dto.UpdatePermissionOptionsFromRequest(req)
	permission, err := s.uc.UpdatePermission(ctx, req.GetPermission(), opts)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("PERMISSION_NOT_FOUND", "Permission not found")
		}
		return nil, err
	}
	return &system.UpdatePermissionResponse{Permission: permission}, nil
}

func (s *PermissionService) DeletePermission(ctx context.Context, req *system.DeletePermissionRequest) (*system.DeletePermissionResponse, error) {
	err := s.uc.DeletePermission(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("PERMISSION_NOT_FOUND", "Permission not found")
		}
		return nil, err
	}
	return &system.DeletePermissionResponse{}, nil
}
