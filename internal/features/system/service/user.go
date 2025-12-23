/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"origadmin/application/admin/api/v1/services/system"
)

func (s *SystemService) ListUserResources(ctx context.Context, req *system.ListUserResourcesRequest) (*system.ListUserResourcesResponse, error) {
	resources, err := s.user.ListUserResources(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &system.ListUserResourcesResponse{
		Resources: resources,
	}, nil
}

func (s *SystemService) UpdateUserRoles(ctx context.Context, req *system.UpdateUserRolesRequest) (*system.UpdateUserRolesResponse, error) {
	err := s.user.UpdateUserRoles(ctx, req.Id, req.RoleIds)
	if err != nil {
		return nil, err
	}
	return &system.UpdateUserRolesResponse{}, nil
}

func (s *SystemService) UpdateUserStatus(ctx context.Context, req *system.UpdateUserStatusRequest) (*system.UpdateUserStatusResponse, error) {
	err := s.user.UpdateUserStatus(ctx, req.Id, req.Status)
	if err != nil {
		return nil, err
	}
	return &system.UpdateUserStatusResponse{}, nil
}

func (s *SystemService) ResetUserPassword(ctx context.Context, req *system.ResetUserPasswordRequest) (*system.ResetUserPasswordResponse, error) {
	err := s.user.ResetUserPassword(ctx, req.Id, req.Password)
	if err != nil {
		return nil, err
	}
	return &system.ResetUserPasswordResponse{}, nil
}

func (s *SystemService) ListUsers(ctx context.Context, req *system.ListUsersRequest) (*system.ListUsersResponse, error) {
	users, total, err := s.user.ListUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListUsersResponse{
		Users: users,
		Total: total,
	}, nil
}

func (s *SystemService) GetUser(ctx context.Context, req *system.GetUserRequest) (*system.User, error) {
	return s.user.GetUser(ctx, req.Id)
}

func (s *SystemService) CreateUser(ctx context.Context, req *system.CreateUserRequest) (*system.User, error) {
	return s.user.CreateUser(ctx, req.User, req.Password)
}

func (s *SystemService) UpdateUser(ctx context.Context, req *system.UpdateUserRequest) (*system.User, error) {
	return s.user.UpdateUser(ctx, req.User)
}

func (s *SystemService) DeleteUser(ctx context.Context, req *system.DeleteUserRequest) (*system.DeleteUserResponse, error) {
	err := s.user.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &system.DeleteUserResponse{}, nil
}
