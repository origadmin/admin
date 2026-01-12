/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/errors"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/helpers/db"
)

type UserService struct {
	system.UnimplementedUserServiceServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) ListUserResources(ctx context.Context, req *system.ListUserResourcesRequest) (*system.ListUserResourcesResponse, error) {
	resources, err := s.uc.ListUserResources(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.ListUserResourcesResponse{
		Resources: resources,
	}, nil
}

func (s *UserService) UpdateUserRoles(ctx context.Context, req *system.UpdateUserRolesRequest) (*system.UpdateUserRolesResponse, error) {
	err := s.uc.UpdateUserRoles(ctx, req.GetId(), req.GetRoleIds())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.UpdateUserRolesResponse{}, nil
}

func (s *UserService) UpdateUserStatus(ctx context.Context, req *system.UpdateUserStatusRequest) (*system.UpdateUserStatusResponse, error) {
	err := s.uc.UpdateUserStatus(ctx, req.GetId(), int8(req.GetStatus()))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.UpdateUserStatusResponse{}, nil
}

func (s *UserService) ResetUserPassword(ctx context.Context, req *system.ResetUserPasswordRequest) (*system.ResetUserPasswordResponse, error) {
	err := s.uc.ResetUserPassword(ctx, req.GetId(), req.GetPassword())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.ResetUserPasswordResponse{}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *system.ListUsersRequest) (*system.ListUsersResponse, error) {
	users, total, err := s.uc.ListUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &system.ListUsersResponse{
		Users:    users,
		Total:    total,
		PageSize: req.GetPageSize(),
	}

	if req.GetPagingMode() == db.PagingModeCursor {
		nextToken, err := db.GenerateNextPageToken(users, req)
		if err != nil {
			return nil, errors.InternalServer("TOKEN_GENERATION_FAILED", err.Error())
		}
		resp.NextPageToken = nextToken
	} else {
		resp.Page = req.GetPage()
	}

	return resp, nil
}

func (s *UserService) GetUser(ctx context.Context, req *system.GetUserRequest) (*system.GetUserResponse, error) {
	user, err := s.uc.GetUser(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.GetUserResponse{User: user}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *system.CreateUserRequest) (*system.CreateUserResponse, error) {
	user, err := s.uc.CreateUser(ctx, req.GetUser(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &system.CreateUserResponse{User: user}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *system.UpdateUserRequest) (*system.UpdateUserResponse, error) {
	user, err := s.uc.UpdateUser(ctx, req.GetUser())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.UpdateUserResponse{User: user}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *system.DeleteUserRequest) (*system.DeleteUserResponse, error) {
	err := s.uc.DeleteUser(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.DeleteUserResponse{}, nil
}
