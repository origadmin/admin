/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"
	"time"

	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/protobuf/proto"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema"
)

type (
	RoleEdges       = ent.RoleEdges
	RoleEdgesPB     = pb.RoleEdges
	RoleMenu        = ent.RoleMenu
	RoleMenuPB      = pb.RoleMenu
	RoleMenuEdges   = ent.RoleMenuEdges
	RoleMenuEdgesPB = pb.RoleMenuEdges

	ListRolesRequest  = pb.ListRolesRequest
	ListRolesResponse = pb.ListRolesResponse
)

func RoleObject(rolePB *RolePB) *Role {
	if rolePB == nil {
		return nil
	}
	return &Role{
		ID:          rolePB.Id,
		CreateTime:  rolePB.CreateTime.AsTime(),
		UpdateTime:  rolePB.UpdateTime.AsTime(),
		Keyword:     rolePB.Keyword,
		Name:        rolePB.Name,
		Description: rolePB.Description,
		Type:        schema.RoleTypeUser,
		Sequence:    int(rolePB.Sequence),
		Status:      int8(rolePB.Status),
		IsSystem:    false,
		Edges:       ent.RoleEdges{},
	}
}

// RoleRepo is a RolePB repository interface.
type RoleRepo interface {
	Get(context.Context, int64, ...RoleQueryOption) (*RolePB, error)
	Create(context.Context, *RolePB, ...RoleQueryOption) (*RolePB, error)
	Delete(context.Context, int64) error
	Update(context.Context, *RolePB, ...RoleQueryOption) (*RolePB, error)
	List(context.Context, *ListRolesRequest, ...RoleQueryOption) ([]*RolePB, int32, error)
}

type RoleQueryOption struct {
	Name         string  `form:"name" json:"name,omitempty"`
	Status       int8    `form:"status" json:"status,omitempty"`
	InIDs        []int64 `form:"-" json:"-"`
	UpdateTimeGT *time.Time
	SelectFields []string
	OmitFields   []string
	OrderFields  []string
	Fields       []string
}

func (o RoleQueryOption) FromListRequest(in *ListRolesRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o RoleQueryOption) FromGetRequest(in *pb.GetRoleRequest, limiter pagination.PageLimiter) error {
	return nil
}

func (o RoleQueryOption) FromCreateRequest(in *pb.CreateRoleRequest, limiter pagination.PageLimiter) error {
	return nil
}

func ToListRolesResponse(result []*RolePB, in *ListRolesRequest, total int32, args ...any) (*ListRolesResponse, error) {
	response := &ListRolesResponse{
		TotalSize: total,
		Current:   in.Current,
		PageSize:  in.PageSize,
		Roles:     result,
		Extra:     resp.Any(args...),
	}
	return response, nil
}

func ConvertRoles(roles []*Role) []*RolePB {
	var result []*RolePB
	for _, role := range roles {
		result = append(result, ConvertRole2PB(role))
	}
	return result
}

func ConvertRoleMenus(menus []*RoleMenu) []*RoleMenuPB {
	var result []*RoleMenuPB
	for _, menu := range menus {
		result = append(result, ConvertRoleMenu2PB(menu))
	}
	return result
}

type RoleQueryResult struct {
	Current  int                      `json:"current"`
	PageSize int                      `json:"page_size"`
	Data     []*RolePB                `json:"data"`
	Total    int64                    `json:"total"`
	Args     map[string]proto.Message `json:"args"`
}
