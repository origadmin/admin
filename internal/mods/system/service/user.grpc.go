/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/biz"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer

	client *biz.UserServiceBiz
	log    *log.KHelper
}

func (s UserServiceServer) ListUserResources(ctx context.Context, request *pb.ListUserResourcesRequest) (*pb.ListUserResourcesResponse, error) {
	return s.client.ListUserResources(ctx, request)
}

func (s UserServiceServer) UpdateUserStatus(ctx context.Context, request *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceServer) UpdateUserRoles(ctx context.Context, request *pb.UpdateUserRolesRequest) (*pb.UpdateUserRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceServer) ResetUserPassword(ctx context.Context, request *pb.ResetUserPasswordRequest) (*pb.ResetUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.client.ListUsers(ctx, req)
}

func (s UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.client.GetUser(ctx, req)
}

func (s UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.client.CreateUser(ctx, req)
}

func (s UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.client.UpdateUser(ctx, req)
}

func (s UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.client.DeleteUser(ctx, req)
}

// NewUserServiceServer new a user service.
func NewUserServiceServer(r runtime.Runtime, client *biz.UserServiceBiz) *UserServiceServer {
	return &UserServiceServer{
		log: log.NewHelper(r.WithLogger(
			"module", "service/user",
		)),
		client: client,
	}
}

// NewUserServiceServerPB new a user service.
func NewUserServiceServerPB(r runtime.Runtime, client *biz.UserServiceBiz) pb.UserServiceServer {
	return NewUserServiceServer(r, client)
}

var _ pb.UserServiceServer = (*UserServiceServer)(nil)
