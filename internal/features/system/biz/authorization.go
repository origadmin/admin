/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/idutil"
)

// AuthorizationUseCase is the use case for authorization policy management.
// It implements the auth.biz.PolicyProvider interface.
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
func (uc *AuthorizationUseCase) ListAllPolicies(ctx context.Context) (*systemv1.ListAllPoliciesResponse, error) {
	uc.log.WithContext(ctx).Info("Listing all policies")

	// 1. Fetch protobuf data directly from the repository in a single atomic transaction.
	// This ensures that we get a consistent snapshot of role permissions and user roles.
	rolePerms, userRoles, err := uc.repo.ListAllPolicies(ctx)
	if err != nil {
		return nil, err // Error is already logged in the DAL layer
	}

	// 2. Perform the "Join" and transformation logic.
	accessRules := make([]*systemv1.AccessRule, 0, len(rolePerms))
	for _, rp := range rolePerms {
		// If Role or Permission edges are not loaded, this indicates inconsistent data.
		// This can happen if a role or permission is deleted but the join table entry remains.
		// We log a warning and skip this record to prevent a crash.
		if rp.Role == nil || rp.Permission == nil {
			uc.log.Warnf("Inconsistent data: RolePermission record (ID: %d) has a nil Role (ID: %d) or Permission (ID: %d). Skipping.",
				rp.Id, rp.RoleId, rp.PermissionId)
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
			// Use Operation (gRPC method name) and "ANY" action for gRPC authorization compatibility.
			// This matches the logic in ListPoliciesForRoles.
			accessRules = append(accessRules, &systemv1.AccessRule{
				Subject: rp.Role.Keyword,
				Object:  resource.Operation,
				Action:  "ANY",
				Domain:  "*", // Explicitly set the global domain
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
			User:   idutil.FormatUserID(ur.User.Id),
			Group:  ur.Role.Keyword,
			Domain: "*", // Explicitly set the global domain
		})
	}

	return &systemv1.ListAllPoliciesResponse{
		AccessRules:   accessRules,
		GroupingRules: groupingRules,
	}, nil
}

// ListPoliciesForRoles fetches all access rules for a specific set of roles.
func (uc *AuthorizationUseCase) ListPoliciesForRoles(ctx context.Context, roleKeywords ...string) ([]*systemv1.AccessRule, error) {
	uc.log.WithContext(ctx).Infof("BIZ: Listing policies for roles: %v", roleKeywords)

	rolePerms, err := uc.repo.ListRolePermissionsByRoleKeywords(ctx, roleKeywords...)
	if err != nil {
		return nil, err
	}
	uc.log.WithContext(ctx).Infof("BIZ: Received %d role_permission records from DAL.", len(rolePerms))

	accessRules := make([]*systemv1.AccessRule, 0, len(rolePerms))
	for _, rp := range rolePerms {
		// If Role or Permission edges are not loaded, this indicates inconsistent data.
		// This can happen if a role or permission is deleted but the join table entry remains.
		// We log a warning and skip this record to prevent a crash.
		if rp.Role == nil || rp.Permission == nil {
			uc.log.Warnf("Inconsistent data: RolePermission record (ID: %d) has a nil Role (ID: %d) or Permission (ID: %d). Skipping.",
				rp.Id, rp.RoleId, rp.PermissionId)
			continue
		}

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
				Object:  resource.Operation,
				Action:  "ANY",
				Domain:  "*", // Explicitly set the global domain
			})
		}
	}
	uc.log.WithContext(ctx).Infof("BIZ: Transformed into %d access rules.", len(accessRules))

	return accessRules, nil
}

//var _ authbiz.PolicyProvider = (*AuthorizationUseCase)(nil)
