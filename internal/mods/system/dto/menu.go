/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
)

type (
	ListMenusRequest  = pb.ListMenusRequest
	ListMenusResponse = pb.ListMenusResponse
)

// MenuRepo is a Menu repository interface.
type MenuRepo interface {
	//Get(context.Context, int64, ...MenuQueryOption) (*MenuPB, error)
	//Create(context.Context, *MenuPB, ...MenuQueryOption) (*MenuPB, error)
	//Delete(context.Context, int64) error
	//Update(context.Context, *MenuPB, ...MenuQueryOption) (*MenuPB, error)
	//List(context.Context, *ListMenusRequest, ...MenuQueryOption) ([]*MenuPB, int32, error)
}

type MenuQueryOption struct {
	Name             string  `form:"name" json:"name,omitempty"`
	Status           int8    `form:"status" json:"status,omitempty"`
	InIDs            []int64 `form:"-" json:"-"`
	UserID           int64   `form:"-" json:"-"` // UserPB ID
	RoleID           int64   `form:"-" json:"-"` // RolePB ID
	ParentID         int64   `form:"-" json:"-"` // Parent ID
	ParentPathPrefix string  `form:"-" json:"-"`
	IncludeResources bool    `form:"-" json:"-"` //　Include resources
	SelectFields     []string
	OmitFields       []string
	OrderFields      []string
	Fields           []string
}

func (o MenuQueryOption) FromListRequest(in *ListMenusRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o MenuQueryOption) FromGetRequest(in *pb.GetMenuRequest, limiter pagination.PageLimiter) error {
	return nil
}

func (o MenuQueryOption) FromCreateRequest(in *pb.CreateMenuRequest, limiter pagination.PageLimiter) error {
	return nil
}

//
//func ToListMenusResponse(result []*MenuPB, in *ListMenusRequest, total int32, args ...any) (*ListMenusResponse, error) {
//	response := &ListMenusResponse{
//		TotalSize: total,
//		Current:   in.Current,
//		PageSize:  in.PageSize,
//		Menus:     result,
//		Extra:     resp.Any(args...),
//	}
//	return response, nil
//}
//
//func ConvertMenus(menus []*Menu) []*MenuPB {
//	var result []*MenuPB
//	for _, menu := range menus {
//		result = append(result, ConvertMenu2PB(menu))
//	}
//	return result
//}
