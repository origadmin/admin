/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	identityv1 "origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/internal/features/system/biz"
)

// AdminService implements the identity.AdminServiceServer interface.
type AdminService struct {
	identityv1.UnimplementedAdminServiceServer

	syncer     *biz.PolicyUseCase
	authorizer authz.Authorizer
	log        *log.Helper
}

// NewAdminService creates a new AdminService.
func NewAdminService(syncer *biz.PolicyUseCase, authorizer authz.Authorizer, logger log.Logger) *AdminService {
	return &AdminService{
		syncer:     syncer,
		authorizer: authorizer,
		log:        log.NewHelper(log.With(logger, "module", "system.service.admin")),
	}
}

// ForcePolicySync triggers an immediate, full synchronization of all identityorization policies.
func (s *AdminService) ForcePolicySync(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	s.log.WithContext(ctx).Info("Received ForcePolicySync request")
	if err := s.syncer.ForceSync(ctx); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// GetPolicySyncStatus retrieves the current status and metrics of the policy synchronization service.
func (s *AdminService) GetPolicySyncStatus(ctx context.Context, _ *emptypb.Empty) (*identityv1.PolicySyncStatusResponse, error) {
	return s.syncer.GetMetrics(), nil
}

// GetEnforcerPolicies retrieves all policy rules currently loaded into the Casbin Enforcer's memory.
func (s *AdminService) GetEnforcerPolicies(ctx context.Context, _ *emptypb.Empty) (*identityv1.EnforcerPoliciesResponse, error) {
	// TODO: Implement this once casbin.Authorizer exposes the underlying Enforcer or a method to list policies.
	// For now, return an empty list to satisfy the interface.
	s.log.WithContext(ctx).Warn("GetEnforcerPolicies is not yet implemented")
	return &identityv1.EnforcerPoliciesResponse{}, nil
}
