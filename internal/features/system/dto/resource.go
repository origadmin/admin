/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/helpers/pagination"
	"origadmin/application/admin/internal/helpers/resp"
)

type (
	ListResourcesRequest  = pb.ListResourcesRequest
	ListResourcesResponse = pb.ListResourcesResponse
)

type ResourceNode struct {
	ResourcePB
	Children []*ResourceNode `json:"children"`
}

// ResourceRepo is a Resource repository interface.
type ResourceRepo interface {
	Get(context.Context, int64, ...ResourceQueryOption) (*ResourcePB, error)
	Create(context.Context, *ResourcePB, ...ResourceQueryOption) (*ResourcePB, error)
	Delete(context.Context, int64) error
	Update(context.Context, *ResourcePB, ...ResourceQueryOption) (*ResourcePB, error)
	List(context.Context, *ListResourcesRequest, ...ResourceQueryOption) ([]*ResourcePB, int32, error)
}

type ResourceQueryOption struct {
	Name               string  `form:"name" json:"name,omitempty"`
	Status             int8    `form:"status" json:"status,omitempty"`
	InIDs              []int64 `form:"-" json:"-"`
	UserID             string  `form:"-" json:"-"` // UserPB ID
	RoleID             string  `form:"-" json:"-"` // RolePB ID
	ParentID           int64   `form:"-" json:"-"` // Parent ID
	ParentPathPrefix   string  `form:"-" json:"-"`
	IncludeResources   bool    `form:"-" json:"-"` //　Include resources
	IncludePermissions bool    `form:"-" json:"-"`
	SelectFields       []string
	OmitFields         []string
	OrderFields        []string
	Fields             []string
}

func (o ResourceQueryOption) FromListRequest(in *ListResourcesRequest, limiter pagination.PageLimiter) error { // Updated usage
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o ResourceQueryOption) FromGetRequest(in *pb.GetResourceRequest, limiter pagination.PageLimiter) error { // Updated usage
	return nil
}

func (o ResourceQueryOption) FromCreateRequest(in *pb.CreateResourceRequest, limiter pagination.PageLimiter) error { // Updated usage
	return nil
}

func ToListResourcesResponse(result []*ResourcePB, in *ListResourcesRequest, total int32, args ...any) (*ListResourcesResponse, error) {
	response := &ListResourcesResponse{
		TotalSize: total,
		Current:   in.Current,
		PageSize:  in.PageSize,
		Resources: result,
		Extra:     resp.Any(args...),
	}
	return response, nil
}
