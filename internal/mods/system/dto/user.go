/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/rand"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/protobuf/proto"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/id"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
)

type (
	UserRole        = ent.UserRole
	UserRolePB      = pb.UserRole
	UserRoleEdges   = ent.UserRoleEdges
	UserRoleEdgesPB = pb.UserRoleEdges

	ListUsersRequest  = pb.ListUsersRequest
	ListUsersResponse = pb.ListUsersResponse
)

func UserObject(userPB *UserPB) *User {
	if userPB == nil {
		return nil
	}
	return &User{
		ID:   userPB.Id,
		UUID: userPB.Uuid,
		//CreateAuthor:  userPB.CreateAuthor,
		//UpdateAuthor:  userPB.UpdateAuthor,
		//CreateTime: userPB.CreateTime.AsTime(),
		//UpdateTime: userPB.UpdateTime.AsTime(),
		//Index:         userPB.Index,
		//Department:    userPB.Department,
		//AllowedIP:     userPB.AllowedIP,
		Username: userPB.Username,
		Name:     userPB.Name,
		Gender:   user.Gender(userPB.Gender),
		Avatar:   userPB.Avatar,
		Password: userPB.Password,
		Salt:     userPB.Salt,
		Phone:    userPB.Phone,
		Email:    userPB.Email,
		Remark:   userPB.Remark,
		//LastLoginTime: userPB.LastLoginTime.AsTime(),
		//SanctionDate:  userPB.SanctionDate.AsTime(),
		Status: int8(userPB.Status),
		//ManagerID:     userPB.ManagerID,
		//Manager:       userPB.Manager,
	}
}

// UserRepo is a UserPB repository interface.
type UserRepo interface {
	Get(context.Context, int64, ...UserQueryOption) (*UserPB, error)
	Create(context.Context, *UserPB, ...UserQueryOption) (*UserPB, error)
	Delete(context.Context, int64) error
	Update(context.Context, *UserPB, ...UserQueryOption) (*UserPB, error)
	List(context.Context, *ListUsersRequest, ...UserQueryOption) ([]*UserPB, int32, error)
	AddRoleIDs(context.Context, int64, []int64) error
	GetByUserName(context.Context, string, ...string) (*UserPB, error)
	GetRoleIDs(context.Context, int64) ([]int64, error)
	ListMenuByUserID(context.Context, int64) ([]*MenuPB, error)
	Current(context.Context, int64) (*UserPB, error)
	UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest, option UserQueryOption) (*pb.UpdateUserStatusResponse, error)
}

type UserQueryOption struct {
	IsAdmin      bool
	NoPasswd     bool
	RandomPasswd bool
	Status       int8 `form:"status" json:"status,omitempty"`
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

func (o *UserQueryOption) FromListRequest(in *ListUsersRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func (o *UserQueryOption) FromGetRequest(in *pb.GetUserRequest, limiter pagination.PageLimiter) error {
	return nil
}

func (o *UserQueryOption) FromCreateRequest(in *pb.CreateUserRequest, limiter pagination.PageLimiter) error {
	o.RandomPasswd = in.RandomPassword
	//o.NoPasswd = in.NoPassword
	o.IsAdmin = in.IsAdmin
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

// CreateUser functions are used to create new users
func CreateUser(user *UserPB, username, password string, option UserQueryOption) (*UserPB, string, error) {
	var err error

	if !option.NoPasswd {
		if option.RandomPasswd && (user.Email != "" || user.Phone != "") {
			password = rand.GenerateRandom(8)
		} else {
			password = ""
		}
	} else {
		password = ""
	}
	var passwd, salt string
	if password != "" {
		salt = rand.GenerateSalt()
		passwd, err = hash.Generate(password, salt)
		if err != nil {
			return nil, "", err
		}
	}
	registerID := id.Gen()
	user.Id = registerID
	user.Uuid = uuid.Must(uuid.NewRandom()).String()
	user.Username = username
	user.Name = "user_" + strconv.Itoa(int(registerID))
	user.Password = passwd
	user.Salt = salt
	user.Status = 1
	return user, password, nil
}
