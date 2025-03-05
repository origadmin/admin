/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

type (
	ListPermissionsRequest  = pb.ListPermissionsRequest
	ListPermissionsResponse = pb.ListPermissionsResponse
)

// PermissionRepo is a Permission repository interface.
type PermissionRepo interface {
	Get(context.Context, int64, ...PermissionQueryOption) (*PermissionPB, error)
	Create(context.Context, *PermissionPB, ...PermissionQueryOption) (*PermissionPB, error)
	Delete(context.Context, int64) error
	Update(context.Context, *PermissionPB, ...PermissionQueryOption) (*PermissionPB, error)
	List(context.Context, *ListPermissionsRequest, ...PermissionQueryOption) ([]*PermissionPB, int32, error)
}

type PermissionQueryOption struct {
	Name               string   `form:"name" json:"name,omitempty"`
	Status             int8     `form:"status" json:"status,omitempty"`
	InIDs              []string `form:"-" json:"-"`
	UserID             string   `form:"-" json:"-"` // UserPB ID
	RoleID             string   `form:"-" json:"-"` // RolePB ID
	ParentID           string   `form:"-" json:"-"` // Parent ID
	ParentPathPrefix   string   `form:"-" json:"-"`
	IncludePermissions bool     `form:"-" json:"-"` //　Include permissions
	SelectFields       []string
	OmitFields         []string
	OrderFields        []string
	Fields             []string
}

func (o PermissionQueryOption) FromListRequest(in *ListPermissionsRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o PermissionQueryOption) FromGetRequest(in *pb.GetPermissionRequest, limiter pagination.PageLimiter) error {
	return nil
}

func (o PermissionQueryOption) FromCreateRequest(in *pb.CreatePermissionRequest, limiter pagination.PageLimiter) error {
	return nil
}

func ToListPermissionsResponse(result []*PermissionPB, in *ListPermissionsRequest, total int32, args ...any) (*ListPermissionsResponse, error) {
	response := &ListPermissionsResponse{
		TotalSize:   total,
		Current:     in.Current,
		PageSize:    in.PageSize,
		Permissions: result,
		Extra:       resp.Any(args...),
	}
	return response, nil
}

func ConvertPermissions(permissions []*Permission) []*PermissionPB {
	var result []*PermissionPB
	for _, permission := range permissions {
		result = append(result, ConvertPermission2PB(permission))
	}
	return result
}
