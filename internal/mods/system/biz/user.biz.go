/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// UsersBiz is a UserPB use case.
type UsersBiz struct {
	dao     dto.UserRepo
	limiter pagination.PageLimiter
	log     *log.Helper
}

func (m UsersBiz) ListUsers(ctx context.Context, in *pb.ListUsersRequest, opts ...grpc.CallOption) (*pb.ListUsersResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromListRequest(in, m.limiter); err != nil {
		return nil, err
	}

	result, total, err := m.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ListUserResponse(result, in, total)
}

func (m UsersBiz) GetUser(ctx context.Context, in *pb.GetUserRequest, opts ...grpc.CallOption) (*pb.GetUserResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromGetRequest(in, m.limiter); err != nil {
		return nil, err
	}
	result, err := m.dao.Get(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		User: result,
	}, nil
}

func (m UsersBiz) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.CreateUserResponse, error) {
	var option dto.UserQueryOption
	if err := option.FromCreateRequest(in, m.limiter); err != nil {
		return nil, err
	}
	result, err := m.dao.Create(ctx, in.User, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: result,
	}, nil
}

func (m UsersBiz) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest, opts ...grpc.CallOption) (*pb.UpdateUserResponse, error) {
	//var option dto.UpdateUserOption
	//if err := option.FromListRequest(in, m.limiter); err != nil {
	//	return nil, err
	//}
	result, err := m.dao.Update(ctx, in.User)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		User: result,
	}, nil
}

func (m UsersBiz) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.DeleteUserResponse, error) {
	//var option dto.DeleteUserOption
	//if err := option.FromListRequest(in, m.limiter); err != nil {
	//	return nil, err
	//}
	//_, err := m.dao.Get(ctx, in.GetId())
	//if err != nil {
	//	return nil, err
	//}
	if err := m.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}

// NewUsersBiz new a UserPB use case.
func NewUsersBiz(repo dto.UserRepo, logger log.Logger) *UsersBiz {
	return &UsersBiz{dao: repo, log: log.NewHelper(logger)}
}

// NewUsersClient new a UserPB use case.
func NewUsersClient(repo dto.UserRepo, logger log.Logger) pb.UserAPIClient {
	return &UsersBiz{dao: repo, log: log.NewHelper(logger)}
}

var _ pb.UserAPIClient = (*UsersBiz)(nil)
