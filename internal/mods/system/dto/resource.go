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
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
)

type (
	ResourceEdges   = ent.ResourceEdges
	ResourceEdgesPB = pb.ResourceEdges

	ListResourcesRequest  = pb.ListResourcesRequest
	ListResourcesResponse = pb.ListResourcesResponse
)

func ResourceObject(resource *ResourcePB) *Resource {
	if resource == nil {
		return nil
	}
	return &Resource{
		ID:         resource.Id,
		CreateTime: resource.CreateTime.AsTime(),
		UpdateTime: resource.UpdateTime.AsTime(),
		MenuID:     resource.MenuId,
		Method:     resource.Method,
		Path:       resource.Path,
	}
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
	Name             string   `form:"name" json:"name,omitempty"`
	Status           int8     `form:"status" json:"status,omitempty"`
	InIDs            []string `form:"-" json:"-"`
	UserID           string   `form:"-" json:"-"` // UserPB ID
	RoleID           string   `form:"-" json:"-"` // RolePB ID
	ParentID         string   `form:"-" json:"-"` // Parent ID
	ParentPathPrefix string   `form:"-" json:"-"`
	IncludeResources bool     `form:"-" json:"-"` //　Include resources
	SelectFields     []string
	OmitFields       []string
	OrderFields      []string
	Fields           []string
}

func (o ResourceQueryOption) FromListRequest(in *ListResourcesRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o ResourceQueryOption) FromGetRequest(in *pb.GetResourceRequest, limiter pagination.PageLimiter) error {
	return nil
}

func (o ResourceQueryOption) FromCreateRequest(in *pb.CreateResourceRequest, limiter pagination.PageLimiter) error {
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

func ConvertResources(resources []*Resource) []*ResourcePB {
	var result []*ResourcePB
	for _, resource := range resources {
		result = append(result, ConvertResource2PB(resource))
	}
	return result
}
