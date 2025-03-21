/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"fmt"

	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// UserServiceBiz is a UserPB use case.
type UserServiceBiz struct {
	dao     dto.UserRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz UserServiceBiz) ListUserResources(ctx context.Context, in *pb.ListUserResourcesRequest, opts ...grpc.CallOption) (*pb.ListUserResourcesResponse, error) {
	var option dto.UserQueryOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("ListUserResources")
	//option.IncludeRoles = true
	result, err := biz.dao.ListResourceByUserID(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	log.Info("ListUserResources result:", result)
	//return dto.ToListResourcesResponse(result, in, total)
	return &pb.ListUserResourcesResponse{
		TotalSize: int32(len(result)),
		//Current:   in.Current,
		//PageSize:  in.PageSize,
		Resources: result,
		//Extra:     resp.Any(args...),
	}, nil
}

func (biz UserServiceBiz) UpdateUserRoles(ctx context.Context, in *pb.UpdateUserRolesRequest, opts ...grpc.CallOption) (*pb.UpdateUserRolesResponse, error) {
	var option dto.UserQueryOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("UpdateUserRoles")
	//option.IncludeRoles = true
	err := biz.dao.AddRoleIDs(ctx, in.GetUser().GetId(), in.GetRoleIds(), option)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserRolesResponse{
		//User: result,
	}, nil
}

func (biz UserServiceBiz) UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest, opts ...grpc.CallOption) (*pb.UpdateUserStatusResponse, error) {
	var option dto.UserQueryOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("UpdateUserStatus")
	option.Fields = []string{"status"}
	//option.IncludeRoles = true
	err := biz.dao.UpdateUserStatus(ctx, in.GetUser().GetId(), int8(in.GetUser().GetStatus()), option)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserStatusResponse{
		//User: result,
	}, nil
}

func (biz UserServiceBiz) ResetUserPassword(ctx context.Context, in *pb.ResetUserPasswordRequest, opts ...grpc.CallOption) (*pb.ResetUserPasswordResponse, error) {
	var option dto.UserQueryOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("ResetUserPassword")
	option.IncludeRoles = true
	//result, total, err := biz.dao.ResetUserPassword(ctx, in, option)
	//if err != nil {
	//	return nil, err
	//}
	//return dto.ToListUsersResponse(result, in, total)
	return &pb.ResetUserPasswordResponse{}, nil
}

func (biz UserServiceBiz) ListUsers(ctx context.Context, in *pb.ListUsersRequest, opts ...grpc.CallOption) (*pb.ListUsersResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("ListUsers")
	option.IncludeRoles = true
	result, total, err := biz.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ToListUsersResponse(result, in, total)
}

func (biz UserServiceBiz) GetUser(ctx context.Context, in *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.GetUserResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromGetRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("GetUser")
	result, err := biz.dao.Get(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		User: result,
	}, nil
}

func (biz UserServiceBiz) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.CreateUserResponse, error) {
	var option dto.UserMutationOption
	if err := option.FromCreateRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("MakeCreateUser")
	username := in.GetUser().GetUsername()
	password := in.GetUser().GetPassword()
	createUser, ps, err := dto.MakeCreateUser(in.User, username, password, option)
	if err != nil {
		return nil, err
	}
	// TODO: Send email or sms to user
	_ = ps
	fmt.Println("Create new user username:", username, "password:", ps)

	result, err := biz.dao.Create(ctx, createUser, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: result,
	}, nil
}

func (biz UserServiceBiz) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.UpdateUserResponse, error) {
	//var option dto.UpdateUserOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("UpdateUser")
	result, err := biz.dao.Update(ctx, in.User)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		User: result,
	}, nil
}

func (biz UserServiceBiz) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error) {
	log.Info("DeleteUser")
	if err := biz.dao.Delete(ctx, in.GetUser().GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}

// NewUserServiceBiz new a UserPB use case.
func NewUserServiceBiz(repo dto.UserRepo, logger log.KLogger) *UserServiceBiz {
	return &UserServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}
