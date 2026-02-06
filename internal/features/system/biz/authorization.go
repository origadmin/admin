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

// AuthorizationUseCase is the use case for authorization policy management.
type AuthorizationUseCase struct {
	repo dto.AuthorizationRepo
	log  *log.Helper
}

// NewAuthorizationUseCase creates a new AuthorizationUseCase.
func NewAuthorizationUseCase(repo dto.AuthorizationRepo, logger log.Logger) *AuthorizationUseCase {
	return &AuthorizationUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz.authorization")),
	}
}

// ListPolicies fetches, processes, and returns authorization policies from casbin_rule table.
func (uc *AuthorizationUseCase) ListPolicies(ctx context.Context, req *systemv1.ListPoliciesRequest) (*systemv1.ListPoliciesResponse, error) {
	uc.log.WithContext(ctx).Info("Listing policies from casbin_rule")

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
