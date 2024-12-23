/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
)

type (
	Resource              = ent.MenuResource
	ResourcePB            = pb.MenuResource
	ListResourcesRequest  = pb.ListResourcesRequest
	ListResourcesResponse = pb.ListResourcesResponse
)

func AuthObject(resource *ResourcePB) *Resource {
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
		//Edges:      ent.MenuResourceEdges{},
	}
}

// AuthRepo is a Auth repository interface.
type AuthRepo interface {
	ListResources(context.Context, *ListResourcesRequest, ...ResourceQueryOption) ([]*ResourcePB, int32, error)
}

type ResourceQueryOption struct {
	SelectFields []string
	OmitFields   []string
	OrderFields  []string
	Fields       []string
}

func (o ResourceQueryOption) FromListRequest(in *ListResourcesRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func ToListResourcesResponse(result []*ResourcePB, in *ListResourcesRequest, total int32, args ...any) (*ListResourcesResponse, error) {
	response := &ListResourcesResponse{
		Resources: result,
		TotalSize: total,
	}
	return response, nil
}

func ConvertResources(menus []*Resource) []*ResourcePB {
	var result []*ResourcePB
	for _, menu := range menus {
		result = append(result, ConvertMenuResource2PB(menu))
	}
	return result
}
