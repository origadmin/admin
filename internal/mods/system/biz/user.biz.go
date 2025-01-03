/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// UsersBiz is a UserPB use case.
type UsersBiz struct {
	dao     dto.UserRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz UsersBiz) UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusRequest, opts ...grpc.CallOption) (*pb.UpdateUserStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz UsersBiz) ResetUserPassword(ctx context.Context, in *pb.ResetUserPasswordRequest, opts ...grpc.CallOption) (*pb.ResetUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz UsersBiz) ListUsers(ctx context.Context, in *pb.ListUsersRequest, opts ...grpc.CallOption) (*pb.ListUsersResponse, error) {
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

func (biz UsersBiz) GetUser(ctx context.Context, in *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.GetUserResponse, error) {
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

func (biz UsersBiz) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.CreateUserResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromCreateRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("CreateUser")
	result, err := biz.dao.Create(ctx, in.User, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: result,
	}, nil
}

func (biz UsersBiz) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.UpdateUserResponse, error) {
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

func (biz UsersBiz) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error) {
	//var option dto.DeleteUserOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	//_, err := biz.dao.Get(ctx, in.GetId())
	//if err != nil {
	//	return nil, err
	//}
	log.Info("DeleteUser")
	if err := biz.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}

// NewUsersBiz new a UserPB use case.
func NewUsersBiz(repo dto.UserRepo, logger log.KLogger) *UsersBiz {
	return &UsersBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewUsersClient new a UserPB use case.
func NewUsersClient(repo dto.UserRepo, logger log.KLogger) pb.UserAPIClient {
	return &UsersBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.UserAPIClient = (*UsersBiz)(nil)
