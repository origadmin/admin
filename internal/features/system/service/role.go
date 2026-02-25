/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"time"

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
)

type RoleService struct {
	system.UnimplementedRoleServiceServer
	uc        *biz.RoleUseCase
	publisher message.Publisher
	log       *log.Helper
}

func NewRoleService(uc *biz.RoleUseCase, publisher message.Publisher, logger log.Logger) *RoleService {
	return &RoleService{
		uc:        uc,
		publisher: publisher,
		log:       log.NewHelper(log.With(logger, "module", "system.service.role")),
	}
}

// publishPolicyChangeEvent is a helper to publish a policy change event for specific roles.
func (s *RoleService) publishPolicyChangeEvent(roleKeywords ...string) {
	if len(roleKeywords) == 0 {
		s.log.Warn("publishPolicyChangeEvent called with no roleKeywords, which will trigger a full sync.")
	}
	s.log.WithContext(context.Background()).Infof("CONFIRM: Preparing to publish RolePolicyChangedEvent with keywords: %v", roleKeywords)
	event := &types.RolePolicyChangedEvent{
		Timestamp:    timestamppb.New(time.Now()),
		Source:       "system.service",
		RoleKeywords: roleKeywords,
	}
	payload, err := proto.Marshal(event)
	if err != nil {
		s.log.Errorf("Failed to marshal RolePolicyChangedEvent: %v", err)
		return
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	if err := s.publisher.Publish(broker.RolePolicyChangedTopic, msg); err != nil {
		s.log.Errorf("Failed to publish RolePolicyChangedEvent: %v", err)
	}
}

func (s *RoleService) ListRoles(ctx context.Context, req *system.ListRolesRequest) (*system.ListRolesResponse, error) {
	queryOpt := dto.ListRolesRequestToQueryOption(req)
	roles, total, err := s.uc.ListRoles(ctx, queryOpt)
	if err != nil {
		return nil, err
	}

	page, pageSize, token, err := db.CalculatePagination(roles, &queryOpt.QueryOption)
	if err != nil {
		return nil, err
	}

	return &system.ListRolesResponse{
		Roles:         roles,
		Total:         total,
		PageSize:      pageSize,
		NextPageToken: token,
		Page:          page,
	}, nil
}
func (s *RoleService) GetRole(ctx context.Context, req *system.GetRoleRequest) (*system.GetRoleResponse, error) {
	queryOpt := dto.GetRoleRequestToQueryOption(req)
	role, err := s.uc.GetRole(ctx, req.GetId(), queryOpt)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "Role not found")
		}
		return nil, err
	}
	return &system.GetRoleResponse{Role: role}, nil
}
func (s *RoleService) CreateRole(ctx context.Context, req *system.CreateRoleRequest) (*system.CreateRoleResponse, error) {
	opts := dto.CreateRoleOptionsFromRequest(req)
	role, err := s.uc.CreateRole(ctx, req.GetRole(), opts)
	if err != nil {
		return nil, err
	}
	s.publishPolicyChangeEvent(role.Keyword)
	return &system.CreateRoleResponse{Role: role}, nil
}
func (s *RoleService) UpdateRole(ctx context.Context, req *system.UpdateRoleRequest) (*system.UpdateRoleResponse, error) {
	opts := dto.UpdateRoleOptionsFromRequest(req)
	role, err := s.uc.UpdateRole(ctx, req.GetRole(), opts)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "Role not found")
		}
		return nil, err
	}
	s.publishPolicyChangeEvent(role.Keyword)
	return &system.UpdateRoleResponse{Role: role}, nil
}
func (s *RoleService) DeleteRole(ctx context.Context, req *system.DeleteRoleRequest) (*system.DeleteRoleResponse, error) {
	// We need the keyword to publish, so we must fetch the role before deleting.
	role, err := s.uc.GetRole(ctx, req.GetId(), &dto.RoleQueryOption{})
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "Role not found")
		}
		return nil, err
	}

	err = s.uc.DeleteRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	// Deleting a role requires removing its 'g' and 'p' rules.
	// We trigger a targeted sync for the deleted role's keyword.
	s.publishPolicyChangeEvent(role.Keyword)
	return &system.DeleteRoleResponse{}, nil
}

// UpdateRolePermissions updates the permissions for a specific role.
func (s *RoleService) UpdateRolePermissions(ctx context.Context, roleID int64, permissionIDs []int64) error {
	// We need the keyword to publish, so we must fetch the role.
	role, err := s.uc.GetRole(ctx, roleID, &dto.RoleQueryOption{})
	if err != nil {
		return err
	}

	err = s.uc.UpdateRolePermissions(ctx, roleID, permissionIDs)
	if err != nil {
		return err
	}
	s.publishPolicyChangeEvent(role.Keyword)
	return nil
}
