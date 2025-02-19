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

// UserServiceClientBiz is a UserPB use case.
type UserServiceClientBiz struct {
	dao     dto.UserRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz UserServiceClientBiz) UpdateUserRoles(ctx context.Context, in *pb.UpdateUserRolesRequest, opts ...grpc.CallOption) (*pb.UpdateUserRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz UserServiceClientBiz) UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest, opts ...grpc.CallOption) (*pb.UpdateUserStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz UserServiceClientBiz) ResetUserPassword(ctx context.Context, in *pb.ResetUserPasswordRequest, opts ...grpc.CallOption) (*pb.ResetUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz UserServiceClientBiz) ListUsers(ctx context.Context, in *pb.ListUsersRequest, opts ...grpc.CallOption) (*pb.ListUsersResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("ListUsers")
	result, total, err := biz.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ToListUsersResponse(result, in, total)
}

func (biz UserServiceClientBiz) GetUser(ctx context.Context, in *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.GetUserResponse, error) {
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

func (biz UserServiceClientBiz) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.CreateUserResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromCreateRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("CreateUser")
	username := in.GetUser().GetUsername()
	password := in.GetUser().GetPassword()
	createUser, ps, err := dto.CreateUser(in.User, username, password, option)
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

func (biz UserServiceClientBiz) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.UpdateUserResponse, error) {
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

func (biz UserServiceClientBiz) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error) {
	log.Info("DeleteUser")
	if err := biz.dao.Delete(ctx, in.GetUser().GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}

// NewUserServiceClientBiz new a UserPB use case.
func NewUserServiceClientBiz(repo dto.UserRepo, logger log.KLogger) *UserServiceClientBiz {
	return &UserServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewUserServiceClient new a UserPB use case.
func NewUserServiceClient(repo dto.UserRepo, logger log.KLogger) pb.UserServiceClient {
	return &UserServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.UserServiceClient = (*UserServiceClientBiz)(nil)
