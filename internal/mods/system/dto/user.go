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
		//CreateAuthor:  user.CreateAuthor,
		//UpdateAuthor:  user.UpdateAuthor,
		CreateTime: user.CreateTime.AsTime(),
		UpdateTime: user.UpdateTime.AsTime(),
		//Index:         user.Index,
		//Department:    user.Department,
		//AllowedIP:     user.AllowedIP,
		Username: user.Username,
		Name:     user.Name,
		//UserID:        user.UserID,
		Avatar:   user.Avatar,
		Password: user.Password,
		Salt:     user.Salt,
		Phone:    user.Phone,
		Email:    user.Email,
		Remark:   user.Remark,
		//LastLoginTime: user.LastLoginTime.AsTime(),
		//SanctionDate:  user.SanctionDate.AsTime(),
		Status: int8(user.Status),
		//ManagerID:     user.ManagerID,
		//Manager:       user.Manager,
	}
}

// UserRepo is a UserPB repository interface.
type UserRepo interface {
	Get(context.Context, string, ...UserQueryOption) (*UserPB, error)
	Create(context.Context, *UserPB, ...UserQueryOption) (*UserPB, error)
	Delete(context.Context, string) error
	Update(context.Context, *UserPB, ...UserQueryOption) (*UserPB, error)
	List(context.Context, *ListUsersRequest, ...UserQueryOption) ([]*UserPB, int32, error)
	GetByUserName(context.Context, string, ...string) (*UserPB, error)
	GetRoleIDs(context.Context, string) ([]string, error)
	ListMenuByUserID(context.Context, string) ([]*MenuPB, error)
	Current(context.Context, string) (*UserPB, error)
	UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest, option UserQueryOption) (*pb.UpdateUserStatusResponse, error)
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

type UserQueryResult struct {
	Current  int                      `json:"current"`
	PageSize int                      `json:"page_size"`
	Data     []*UserPB                `json:"data"`
	Total    int64                    `json:"total"`
	Args     map[string]proto.Message `json:"args"`
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

func ToListUsersResponse(result []*UserPB, in *ListUsersRequest, total int32, args ...any) (*ListUsersResponse, error) {
	response := &ListUsersResponse{
		TotalSize: total,
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
