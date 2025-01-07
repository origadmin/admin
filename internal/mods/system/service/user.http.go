/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type UserServiceHTTPService struct {
	pb.UnimplementedUserServiceServer

	client pb.UserServiceHTTPClient
}

func (s UserServiceHTTPService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.client.ListUsers(ctx, req)
}

func (s UserServiceHTTPService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.client.GetUser(ctx, req)
}

func (s UserServiceHTTPService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.client.CreateUser(ctx, req)
}

func (s UserServiceHTTPService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.client.UpdateUser(ctx, req)
}

func (s UserServiceHTTPService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.client.DeleteUser(ctx, req)
}

// NewUserServiceHTTPServer new a user service.
func NewUserServiceHTTPServer(client pb.UserServiceHTTPClient) *UserServiceHTTPService {
	return &UserServiceHTTPService{
		client: client,
	}
}

// NewUserServiceHTTPServerPB new a user service.
func NewUserServiceHTTPServerPB(client pb.UserServiceHTTPClient) pb.UserServiceHTTPServer {
	return &UserServiceHTTPService{
		client: client,
	}
}

var _ pb.UserServiceServer = (*UserServiceHTTPService)(nil)
