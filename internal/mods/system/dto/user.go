/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/protobuf/proto"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
)

type (
	User            = ent.User
	UserPB          = pb.User
	UserEdges       = ent.UserEdges
	UserEdgesPB     = pb.UserEdges
	UserRole        = ent.UserRole
	UserRolePB      = pb.UserRole
	UserRoleEdges   = ent.UserRoleEdges
	UserRoleEdgesPB = pb.UserRoleEdges

	ListUsersRequest  = pb.ListUsersRequest
	ListUsersResponse = pb.ListUsersResponse
)

func UserObject(user *UserPB) *User {
	if user == nil {
		return nil
	}
	return &User{
		ID: user.Id,
	}
}

// UserRepo is a UserPB repository interface.
type UserRepo interface {
	Get(context.Context, string, ...UserQueryOption) (*UserPB, error)
	Create(context.Context, *UserPB, ...UserQueryOption) (*UserPB, error)
	Delete(context.Context, string) error
	Update(context.Context, *UserPB, ...UserQueryOption) (*UserPB, error)
	List(context.Context, *ListUsersRequest, ...UserQueryOption) ([]*UserPB, int64, error)
}

type UserQueryOption struct {
	Name         string `form:"name" json:"name,omitempty"`
	Username     string `form:"username" json:"username,omitempty"`
	Status       int8   `form:"status" json:"status,omitempty"`
	SelectFields []string
	OmitFields   []string
	OrderFields  []string
	Fields       []string
}

func (o UserQueryOption) FromListRequest(in *ListUsersRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o UserQueryOption) FromGetRequest(in *pb.GetUserRequest, limiter pagination.PageLimiter) error {
	return nil
}

func (o UserQueryOption) FromCreateRequest(in *pb.CreateUserRequest, limiter pagination.PageLimiter) error {
	return nil
}

func ListUserResponse(result []*UserPB, in *ListUsersRequest, total int64, args ...any) (*ListUsersResponse, error) {
	response := &ListUsersResponse{
		TotalSize: int32(total),
		Current:   in.Current,
		PageSize:  in.PageSize,
		Users:     result,
		Extra:     resp.Any(args...),
	}

	return response, nil
}

func ConvertUsers(users []*User) []*UserPB {
	var result []*UserPB
	for _, user := range users {
		result = append(result, ConvertUser2PB(user))
	}
	return result
}

func ConvertUserRoles(roles []*UserRole) []*UserRolePB {
	var result []*UserRolePB
	for _, role := range roles {
		result = append(result, ConvertUserRole2PB(role))
	}
	return result
}

type UserQueryResult struct {
	Current  int                      `json:"current"`
	PageSize int                      `json:"page_size"`
	Data     []*UserPB                `json:"data"`
	Total    int64                    `json:"total"`
	Args     map[string]proto.Message `json:"args"`
}
