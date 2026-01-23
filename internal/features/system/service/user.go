/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"fmt"

	"github.com/origadmin/runtime/errors"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/broker"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/features/system/events"
	"origadmin/application/admin/internal/helpers/db"
)

type UserService struct {
	system.UnimplementedUserServiceServer
	uc        *biz.UserUseCase
	publisher broker.Publisher
	log       *log.Helper
}

// NewUserService creates a new UserService.
func NewUserService(uc *biz.UserUseCase, publisher broker.Publisher, logger log.Logger) *UserService {
	return &UserService{
		uc:        uc,
		publisher: publisher,
		log:       log.NewHelper(log.With(logger, "module", "system.service.user")),
	}
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

	userIDStr := fmt.Sprintf("%d", req.GetId())
	roleIDsStr := make([]string, len(req.GetRoleIds()))
	for i, roleID := range req.GetRoleIds() {
		roleIDsStr[i] = fmt.Sprintf("%d", roleID)
	}

	msg, err := events.NewUserRoleAssignedMessage(userIDStr, roleIDsStr, "system-service")
	if err != nil {
		s.log.Errorf("failed to create UserRoleAssignedEvent message for user %s: %v", userIDStr, err)
	} else {
		// Corrected call to s.publisher.Publish, removing the ctx parameter.
		err = s.publisher.Publish(events.UserRoleAssignedTopic, msg)
		if err != nil {
			s.log.Errorf("failed to publish UserRoleAssignedEvent for user %s: %v", userIDStr, err)
		}
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
	queryOpt := dto.ListUsersRequestToQueryOption(req)
	users, total, err := s.uc.ListUsers(ctx, queryOpt)
	if err != nil {
		return nil, err
	}

	page, pageSize, token, err := db.CalculatePagination(users, &queryOpt.QueryOption)
	if err != nil {
		return nil, err
	}

	return &system.ListUsersResponse{
		Users:         users,
		Total:         total,
		PageSize:      pageSize,
		NextPageToken: token,
		Page:          page,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *system.GetUserRequest) (*system.GetUserResponse, error) {
	queryOpt := dto.GetUserRequestToQueryOption(req)
	user, err := s.uc.GetUser(ctx, req.GetId(), queryOpt)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.GetUserResponse{User: user}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *system.CreateUserRequest) (*system.CreateUserResponse, error) {
	opts := dto.CreateUserOptionsFromRequest(req)
	user, err := s.uc.CreateUser(ctx, req.GetUser(), req.GetPassword(), opts)
	if err != nil {
		return nil, err
	}
	return &system.CreateUserResponse{User: user}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *system.UpdateUserRequest) (*system.UpdateUserResponse, error) {
	opts := dto.UpdateUserOptionsFromRequest(req)
	user, err := s.uc.UpdateUser(ctx, req.GetUser(), opts)
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
