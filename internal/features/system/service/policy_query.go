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

// PolicyQueryService implements the policy query service.
type PolicyQueryService struct {
	systemv1.UnimplementedPolicyQueryServiceServer

	uc  *biz.PolicyQueryUseCase
	log *log.Helper
}

// NewPolicyQueryService creates a new policy query service.
func NewPolicyQueryService(uc *biz.PolicyQueryUseCase, logger log.Logger) *PolicyQueryService {
	return &PolicyQueryService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service.policy_query")),
	}
}

// ListPolicies retrieves authorization policies from casbin_rule table.
func (s *PolicyQueryService) ListPolicies(ctx context.Context, req *systemv1.ListPoliciesRequest) (*systemv1.ListPoliciesResponse, error) {
	s.log.WithContext(ctx).Info("SERVICE: Retrieving policies from casbin_rule")

	return s.uc.ListPolicies(ctx, req)
}
