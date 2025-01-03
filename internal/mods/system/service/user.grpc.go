/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type UserAPIService struct {
	pb.UnimplementedUserAPIServer

	client pb.UserAPIClient
}

func (s UserAPIService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.client.ListUsers(ctx, req)
}

func (s UserAPIService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.client.GetUser(ctx, req)
}

func (s UserAPIService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.client.CreateUser(ctx, req)
}

func (s UserAPIService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.client.UpdateUser(ctx, req)
}

func (s UserAPIService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.client.DeleteUser(ctx, req)
}

// NewUserAPIService new a user service.
func NewUserAPIService(client pb.UserAPIClient) *UserAPIService {
	return &UserAPIService{}
}

// NewUserAPIServer new a user service.
func NewUserAPIServer(client pb.UserAPIClient) pb.UserAPIServer {
	return &UserAPIService{
		client: client,
	}
}

var _ pb.UserAPIServer = (*UserAPIService)(nil)
