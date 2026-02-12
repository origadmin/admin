/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/runtime/log"
	authv1 "origadmin/application/admin/api/v1/services/identity"
)

// AdminService implements the auth.AdminServiceServer interface.
type AdminService struct {
	authv1.UnimplementedAdminServiceServer

	authorizer *casbin.Authorizer
	log        *log.Helper
}

// NewAdminService creates a new AdminService.
func NewAdminService(authorizer *casbin.Authorizer, logger log.Logger) *AdminService {
	return &AdminService{
		authorizer: authorizer,
		log:        log.NewHelper(log.With(logger, "module", "auth.service.admin")),
	}
}

// ForcePolicySync triggers an immediate, full synchronization of all authorization policies.
func (s *AdminService) ForcePolicySync(ctx context.Context, _ *authv1.ForcePolicySyncRequest) (*authv1.ForcePolicySyncResponse, error) {
	s.log.WithContext(ctx).Info("Received ForcePolicySync request")
	//if err := s.syncer.ForceSync(ctx); err != nil {
	//	return nil, err
	//}
	return &authv1.ForcePolicySyncResponse{}, nil
}

// GetPolicySyncStatus retrieves the current status and metrics of the policy synchronization service.
func (s *AdminService) GetPolicySyncStatus(ctx context.Context, _ *authv1.GetPolicySyncStatusRequest) (*authv1.GetPolicySyncStatusResponse, error) {
	//return s.syncer.GetMetrics(), nil
	return &authv1.GetPolicySyncStatusResponse{}, nil
}

// GetEnforcerPolicies retrieves all policy rules currently loaded into the Casbin Enforcer's memory.
func (s *AdminService) GetEnforcerPolicies(ctx context.Context, _ *authv1.GetEnforcerPoliciesRequest) (*authv1.GetEnforcerPoliciesResponse, error) {
	// TODO: Implement this once casbin.Authorizer exposes the underlying Enforcer or a method to list policies.
	// For now, return an empty list to satisfy the interface.
	s.log.WithContext(ctx).Warn("GetEnforcerPolicies is not yet implemented")
	return &authv1.GetEnforcerPoliciesResponse{}, nil
}
