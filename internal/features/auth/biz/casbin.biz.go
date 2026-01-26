/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v3/model"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/features/auth/dal"
	systempb "origadmin/application/admin/api/v1/services/system"
)

// CasbinSynchronizer is responsible for performing a one-time, full synchronization
// of policies from the system's policy service to the Casbin persistence layer.
type CasbinSynchronizer struct {
	storageManager dal.CasbinStorageManager
	policyClient   systempb.PolicyServiceServer
	log            *log.Helper
}

// NewCasbinSynchronizer creates a new synchronizer.
func NewCasbinSynchronizer(
	logger log.Logger,
	storageManager dal.CasbinStorageManager,
	policyClient systempb.PolicyServiceServer,
) *CasbinSynchronizer {
	return &CasbinSynchronizer{
		storageManager: storageManager,
		policyClient:   policyClient,
		log:            log.NewHelper(log.With(logger, "module", "biz/casbin-synchronizer")),
	}
}

// Sync performs the full synchronization. It should be called once at application startup.
func (s *CasbinSynchronizer) Sync(ctx context.Context) error {
	s.log.WithContext(ctx).Info("Starting full synchronization of Casbin policies from system policy service...")

	// 1. Directly fetch all business-level authorization data from the policy service.
	allPoliciesResp, err := s.policyClient.ListAllPolicies(ctx, &systempb.ListAllPoliciesRequest{})
	if err != nil {
		return fmt.Errorf("failed to list all policies from system policy service: %w", err)
	}
	s.log.WithContext(ctx).Infof("System policy service returned %d role-permissions, %d permission definitions, and %d user-roles.",
		len(allPoliciesResp.GetRolePermissions()), len(allPoliciesResp.GetPermissionDefinitions()), len(allPoliciesResp.GetUserRoles()))

	// 2. Translate the business-level data into Casbin-specific rules.
	// This logic now resides entirely within the auth module.
	m := model.NewModel()

	// Create a map for quick lookup of permission definitions by keyword.
	permDefMap := make(map[string]*systempb.PermissionDefinition)
	for _, permDef := range allPoliciesResp.GetPermissionDefinitions() {
		permDefMap[permDef.PermissionKeyword] = permDef
	}

	// Translate role-permissions to Casbin 'p' rules.
	for _, rolePerm := range allPoliciesResp.GetRolePermissions() {
		if permDef, ok := permDefMap[rolePerm.PermissionKeyword]; ok {
			// p, role_id, resource_path, action_verb, domain_id
			m.AddPolicy("p", "p", []string{
				rolePerm.RoleId,
				permDef.ResourcePath,
				permDef.ActionVerb,
				rolePerm.DomainId,
			})
		}
	}

	// Translate user-roles to Casbin 'g' rules.
	for _, userRole := range allPoliciesResp.GetUserRoles() {
		// g, user_id, role_id, domain_id
		m.AddPolicy("g", "g", []string{
			userRole.UserId,
			userRole.RoleId,
			userRole.DomainId,
		})
	}

	// 3. Pass the complete model to the DAL to be atomically persisted.
	if err := s.storageManager.OverwriteAllPolicies(ctx, m); err != nil {
		return fmt.Errorf("failed to overwrite policies via storage manager: %w", err)
	}

	s.log.WithContext(ctx).Info("Successfully synchronized all Casbin policies.")
	return nil
}
