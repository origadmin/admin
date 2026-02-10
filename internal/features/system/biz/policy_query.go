/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/system/dto"
)

// PolicyQueryUseCase is the use case for querying authorization policies.
type PolicyQueryUseCase struct {
	repo dto.PolicyRepo
	log  *log.Helper
}

// NewPolicyQueryUseCase creates a new PolicyQueryUseCase.
func NewPolicyQueryUseCase(repo dto.PolicyRepo, logger log.Logger) *PolicyQueryUseCase {
	return &PolicyQueryUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz.policy_query")),
	}
}

// ListPolicies fetches, processes, and returns authorization policies from casbin_rule table.
func (uc *PolicyQueryUseCase) ListPolicies(ctx context.Context, req *systemv1.ListPoliciesRequest) (*systemv1.ListPoliciesResponse, error) {
	uc.log.WithContext(ctx).Info("Querying policies from casbin_rule")

	policies, total, err := uc.repo.ListPolicies(ctx, req)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Failed to list policies: %v", err)
		return nil, err
	}

	return &systemv1.ListPoliciesResponse{
		Policies:      policies,
		Total:         total,
		Page:          req.GetPage(),
		PageSize:      req.GetPageSize(),
		NextPageToken: "",
	}, nil
}
