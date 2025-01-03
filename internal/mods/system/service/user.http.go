/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type UserAPIHTTPService struct {
	pb.UnimplementedUserAPIServer

	client pb.UserAPIHTTPClient
}

func (s UserAPIHTTPService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.client.ListUsers(ctx, req)
}

func (s UserAPIHTTPService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.client.GetUser(ctx, req)
}

func (s UserAPIHTTPService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.client.CreateUser(ctx, req)
}

func (s UserAPIHTTPService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.client.UpdateUser(ctx, req)
}

func (s UserAPIHTTPService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.client.DeleteUser(ctx, req)
}

// NewUserAPIHTTPService new a user service.
func NewUserAPIHTTPService(client pb.UserAPIHTTPClient) *UserAPIHTTPService {
	return &UserAPIHTTPService{
		client: client,
	}
}

// NewUserAPIHTTPServer new a user service.
func NewUserAPIHTTPServer(client pb.UserAPIHTTPClient) pb.UserAPIHTTPServer {
	return &UserAPIHTTPService{
		client: client,
	}
}

var _ pb.UserAPIServer = (*UserAPIHTTPService)(nil)
