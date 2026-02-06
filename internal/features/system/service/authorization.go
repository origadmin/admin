/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/system/biz"
)

// AuthorizationService implements the authorization service.
type AuthorizationService struct {
	systemv1.UnimplementedAuthorizationServiceServer

	uc  *biz.AuthorizationUseCase
	log *log.Helper
}

// NewAuthorizationService creates a new authorization service.
func NewAuthorizationService(uc *biz.AuthorizationUseCase, logger log.Logger) *AuthorizationService {
	return &AuthorizationService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service.authorization")),
	}
}

// ListPolicies retrieves authorization policies from casbin_rule table.
// This is the ONLY data source interface for providing policy data to other services.
func (s *AuthorizationService) ListPolicies(ctx context.Context, req *systemv1.ListPoliciesRequest) (*systemv1.ListPoliciesResponse, error) {
	s.log.WithContext(ctx).Info("SERVICE: Retrieving policies from casbin_rule")

	return s.uc.ListPolicies(ctx, req)
}
