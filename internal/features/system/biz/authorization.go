/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"

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

// ListAllPolicies fetches, processes, and returns all authorization policies.
func (uc *AuthorizationUseCase) ListAllPolicies(ctx context.Context, _ *systemv1.ListAllPoliciesRequest) (*systemv1.ListAllPoliciesResponse, error) {
	uc.log.WithContext(ctx).Info("Listing all policies")

	// 1. Fetch protobuf data directly from the repository.
	rolePerms, err := uc.repo.ListRolePermissions(ctx)
	if err != nil {
		return nil, err // Error is already logged in the DAL layer
	}
	// Permissions are now fetched as part of the RolePermissions join.
	// We will iterate through them to build the final rules.

	userRoles, err := uc.repo.ListUserRoles(ctx)
	if err != nil {
		return nil, err // Error is already logged in the DAL layer
	}

	// 2. Perform the "Join" and transformation logic.
	accessRules := make([]*systemv1.AccessRule, 0, len(rolePerms))
	for _, rp := range rolePerms {
		if rp.Role == nil || rp.Permission == nil {
			uc.log.Warnf("incomplete RolePermission protobuf found, ID: %d, skipping", rp.Id)
			continue
		}

		// A permission can be linked to multiple resources (API endpoints).
		// We must create a rule for each resource.
		if len(rp.Permission.Resources) == 0 {
			uc.log.Debugf("permission '%s' has no associated resources, skipping", rp.Permission.Keyword)
			continue
		}

		for _, resource := range rp.Permission.Resources {
			if resource == nil {
				continue
			}
			accessRules = append(accessRules, &systemv1.AccessRule{
				Subject: rp.Role.Keyword,
				Object:  resource.Path,
				Action:  resource.Method,
			})
		}
	}

	groupingRules := make([]*systemv1.GroupingRule, 0, len(userRoles))
	for _, ur := range userRoles {
		if ur.User == nil || ur.Role == nil {
			uc.log.Warnf("incomplete UserRole protobuf found, ID: %d, skipping", ur.Id)
			continue
		}
		groupingRules = append(groupingRules, &systemv1.GroupingRule{
			User:  fmt.Sprintf("user:%d", ur.User.Id),
			Group: ur.Role.Keyword,
		})
	}

	return &systemv1.ListAllPoliciesResponse{
		AccessRules:   accessRules,
		GroupingRules: groupingRules,
	}, nil
}
