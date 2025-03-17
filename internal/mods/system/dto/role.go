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
)

type (
	RoleEdges   = ent.RoleEdges
	RoleEdgesPB = pb.RoleEdges

	ListRolesRequest  = pb.ListRolesRequest
	ListRolesResponse = pb.ListRolesResponse
)

// RoleRepo is a RolePB repository interface.
type RoleRepo interface {
	Get(context.Context, int64, ...RoleQueryOption) (*RolePB, error)
	List(context.Context, *ListRolesRequest, ...RoleQueryOption) ([]*RolePB, int32, error)
	Create(context.Context, *RolePB, ...RoleUpdateOption) (*RolePB, error)
	Update(context.Context, *RolePB, ...RoleUpdateOption) (*RolePB, error)
	Delete(context.Context, int64) error
}

type RoleQueryOption struct {
	Name               string  `form:"name" json:"name,omitempty"`
	Status             int8    `form:"status" json:"status,omitempty"`
	InIDs              []int64 `form:"-" json:"-"`
	UpdateTimeGT       *time.Time
	SelectFields       []string
	OmitFields         []string
	OrderFields        []string
	Fields             []string
	IncludePermissions bool
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

// RoleUpdateOption is used for creating and updating roles.
type RoleUpdateOption struct {
	Name               string `form:"name" json:"name,omitempty"`
	Status             int8   `form:"status" json:"status,omitempty"`
	UpdateTimeGT       *time.Time
	SelectFields       []string
	OmitFields         []string
	OrderFields        []string
	Fields             []string
	IncludePermissions bool
}

func (o RoleUpdateOption) FromCreateRequest(in *pb.CreateRoleRequest) error {
	return nil
}

func (o RoleUpdateOption) FromUpdateRequest(in *pb.UpdateRoleRequest) error {
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

type RoleQueryResult struct {
	Current  int                      `json:"current"`
	PageSize int                      `json:"page_size"`
	Data     []*RolePB                `json:"data"`
	Total    int64                    `json:"total"`
	Args     map[string]proto.Message `json:"args"`
}
