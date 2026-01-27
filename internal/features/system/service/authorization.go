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

// ListAllPolicies retrieves the entire set of policies, pre-processed into a generic,
// implementation-agnostic format.
func (s *AuthorizationService) ListAllPolicies(ctx context.Context, req *systemv1.ListAllPoliciesRequest) (*systemv1.ListAllPoliciesResponse, error) {
	// The business logic is delegated to the use case.
	return s.uc.ListAllPolicies(ctx, req)
}
