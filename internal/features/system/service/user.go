/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
)

func (s *SystemService) ListUserResources(ctx context.Context, req *system.ListUserResourcesRequest) (*system.ListUserResourcesResponse, error) {
	resources, err := s.User.ListUserResources(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &system.ListUserResourcesResponse{
		Resources: resources,
	}, nil
}

func (s *SystemService) UpdateUserRoles(ctx context.Context, req *system.UpdateUserRolesRequest) (*system.UpdateUserRolesResponse, error) {
	err := s.User.UpdateUserRoles(ctx, req.GetId(), req.GetRoleIds())
	if err != nil {
		return nil, err
	}
	return &system.UpdateUserRolesResponse{}, nil
}

func (s *SystemService) UpdateUserStatus(ctx context.Context, req *system.UpdateUserStatusRequest) (*system.UpdateUserStatusResponse, error) {
	err := s.User.UpdateUserStatus(ctx, req.GetId(), req.GetStatus())
	if err != nil {
		return nil, err
	}
	return &system.UpdateUserStatusResponse{}, nil
}

func (s *SystemService) ResetUserPassword(ctx context.Context, req *system.ResetUserPasswordRequest) (*system.ResetUserPasswordResponse, error) {
	err := s.User.ResetUserPassword(ctx, req.GetId(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &system.ResetUserPasswordResponse{}, nil
}

func (s *SystemService) ListUsers(ctx context.Context, req *system.ListUsersRequest) (*system.ListUsersResponse, error) {
	users, total, err := s.User.ListUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListUsersResponse{
		Users: users,
		Total: total,
	}, nil
}

func (s *SystemService) GetUser(ctx context.Context, req *system.GetUserRequest) (*types.User, error) {
	return s.User.GetUser(ctx, req.GetId())
}

func (s *SystemService) CreateUser(ctx context.Context, req *system.CreateUserRequest) (*types.User, error) {
	return s.User.CreateUser(ctx, req.GetUser(), req.GetPassword())
}

func (s *SystemService) UpdateUser(ctx context.Context, req *system.UpdateUserRequest) (*types.User, error) {
	return s.User.UpdateUser(ctx, req.GetUser())
}

func (s *SystemService) DeleteUser(ctx context.Context, req *system.DeleteUserRequest) (*system.DeleteUserResponse, error) {
	err := s.User.DeleteUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &system.DeleteUserResponse{}, nil
}
