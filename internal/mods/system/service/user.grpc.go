/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer

	client pb.UserServiceClient
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
func NewUserServiceServer(client pb.UserServiceClient) *UserServiceServer {
	return &UserServiceServer{}
}

// NewUserServiceServerPB new a user service.
func NewUserServiceServerPB(client pb.UserServiceClient) pb.UserServiceServer {
	return &UserServiceServer{
		client: client,
	}
}

var _ pb.UserServiceServer = (*UserServiceServer)(nil)
