/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/origadmin/runtime/errors"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/broker"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/idutil"
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

// publishUserRoleChangeEvent publishes a generic event indicating a user's roles have changed.
// The consumer is responsible for triggering a full policy sync.
func (s *UserService) publishUserRoleChangeEvent(ctx context.Context, userID int64) {
	userIDStr := idutil.FormatUserID(userID)
	event := &types.UserRoleAssignedEvent{
		Timestamp: timestamppb.Now(),
		UserId:    userIDStr,
		Source:    "system.service",
	}

	s.log.Debugf("Publishing UserRoleAssignedEvent for UserID=%s to trigger policy sync.", userIDStr)

	payload, err := proto.Marshal(event)
	if err != nil {
		s.log.Errorf("failed to marshal UserRoleAssignedEvent for user %s: %v", userIDStr, err)
		return
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	if err := s.publisher.Publish(broker.UserRoleAssignedTopic, msg); err != nil {
		s.log.Errorf("failed to publish UserRoleAssignedEvent for user %s: %v", userIDStr, err)
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

func (s *UserService) ListUserRoles(ctx context.Context, req *system.ListUserRolesRequest) (*system.ListUserRolesResponse, error) {
	roles, err := s.uc.ListUserRoles(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.ListUserRolesResponse{
		Roles: roles,
	}, nil
}

func (s *UserService) ListUserViews(ctx context.Context, req *system.ListUserViewsRequest) (*system.ListUserViewsResponse, error) {
	views, err := s.uc.ListUserViews(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.ListUserViewsResponse{
		Views: views,
	}, nil
}

func (s *UserService) ListUserPermissions(ctx context.Context, req *system.ListUserPermissionsRequest) (*system.ListUserPermissionsResponse, error) {
	permissions, err := s.uc.ListUserPermissions(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.ListUserPermissionsResponse{
		Permissions: permissions,
	}, nil
}

func (s *UserService) UpdateUserRoles(ctx context.Context, req *system.UpdateUserRolesRequest) (*system.UpdateUserRolesResponse, error) {
	_, err := s.uc.UpdateUserRoles(ctx, req.GetId(), req.GetRoleIds())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	s.publishUserRoleChangeEvent(ctx, req.GetId())
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

// ChangeUserPassword updates a user's password.
// The `password` field in the request MUST contain an already hashed password.
// The identity service is responsible for any hashing logic.
func (s *UserService) ChangeUserPassword(ctx context.Context, req *system.ChangeUserPasswordRequest) (*system.ChangeUserPasswordResponse, error) {
	err := s.uc.UpdateUserPassword(ctx, req.GetId(), req.GetPassword())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("USER_NOT_FOUND", "User not found")
		}
		return nil, err
	}
	return &system.ChangeUserPasswordResponse{}, nil
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

// CreateUser creates a new user.
// The `password` field in the request MUST contain an already hashed password.
// The identity service is responsible for any hashing logic.
func (s *UserService) CreateUser(ctx context.Context, req *system.CreateUserRequest) (*system.CreateUserResponse, error) {
	opts := dto.CreateUserOptionsFromRequest(req)
	// The password from the request is expected to be already hashed by the caller (e.g., identity service).
	user, err := s.uc.CreateUser(ctx, req.GetUser(), req.GetPassword(), opts)
	if err != nil {
		return nil, err
	}
	if len(req.GetRoleIds()) > 0 {
		s.publishUserRoleChangeEvent(ctx, user.GetId())
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

	s.publishUserRoleChangeEvent(ctx, user.GetId())
	return &system.UpdateUserResponse{User: user}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *system.DeleteUserRequest) (*system.DeleteUserResponse, error) {
	err := s.uc.DeleteUser(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "Role not found")
		}
		return nil, err
	}
	s.publishUserRoleChangeEvent(ctx, req.GetId())
	return &system.DeleteUserResponse{}, nil
}
