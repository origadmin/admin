/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/data/entity/ent/rolepermission"
	"origadmin/application/admin/internal/features/identity/dto"
	"origadmin/application/admin/internal/helpers/idutil"
)

// policyDBProvider implements dto.PolicyProvider interface using direct database access.
type policyDBProvider struct {
	db  *ent.Database
	log *log.Helper
}

// NewPolicyDBProvider creates a new PolicyProvider that uses direct database access.
func NewPolicyDBProvider(db *ent.Database, logger log.Logger) dto.PolicyProvider {
	return &policyDBProvider{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "dal.policy_db")),
	}
}

// ListAllPolicies fetches, processes, and returns all identityorization policies directly from the database.
func (p *policyDBProvider) ListAllPolicies(ctx context.Context) (*systemv1.ListAllPoliciesResponse, error) {
	p.log.WithContext(ctx).Info("Listing all policies from database")

	// 1. Fetch RolePermissions with eager-loaded Role and Permission+Resources.
	rolePerms, err := p.db.RolePermission(ctx).Query().
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	// 2. Fetch UserRoles with eager-loaded User and Role.
	userRoles, err := p.db.UserRole(ctx).Query().
		WithUser().
		WithRole().
		All(ctx)
	if err != nil {
		return nil, err
	}

	// 3. Transform RolePermission to AccessRule.
	accessRules := make([]*systemv1.AccessRule, 0, len(rolePerms))
	for _, rp := range rolePerms {
		if rp.Edges.Role == nil || rp.Edges.Permission == nil {
			p.log.Warnf("Inconsistent data: RolePermission record (ID: %d) has nil Role or Permission. Skipping.", rp.ID)
			continue
		}

		if len(rp.Edges.Permission.Edges.Resources) == 0 {
			p.log.Debugf("permission '%s' has no associated resources, skipping", rp.Edges.Permission.Keyword)
			continue
		}

		for _, resource := range rp.Edges.Permission.Edges.Resources {
			if resource == nil {
				continue
			}
			accessRules = append(accessRules, &systemv1.AccessRule{
				Subject: rp.Edges.Role.Keyword,
				Object:  resource.Operation,
				Action:  "ANY",
				Domain:  "*",
			})
		}
	}

	// 4. Transform UserRole to GroupingRule.
	groupingRules := make([]*systemv1.GroupingRule, 0, len(userRoles))
	for _, ur := range userRoles {
		if ur.Edges.User == nil || ur.Edges.Role == nil {
			p.log.Warnf("Incomplete UserRole record found, ID: %d, skipping", ur.ID)
			continue
		}
		groupingRules = append(groupingRules, &systemv1.GroupingRule{
			User:   idutil.FormatUserID(ur.Edges.User.ID),
			Group:  ur.Edges.Role.Keyword,
			Domain: "*",
		})
	}

	return &systemv1.ListAllPoliciesResponse{
		AccessRules:   accessRules,
		GroupingRules: groupingRules,
	}, nil
}

// ListPoliciesForRoles fetches all access rules for a specific set of roles directly from the database.
func (p *policyDBProvider) ListPoliciesForRoles(ctx context.Context, roleKeywords ...string) ([]*systemv1.AccessRule, error) {
	p.log.WithContext(ctx).Infof("Listing policies for roles: %v from database", roleKeywords)

	rolePerms, err := p.db.RolePermission(ctx).Query().
		Where(rolepermission.HasRoleWith(role.KeywordIn(roleKeywords...))).
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	p.log.WithContext(ctx).Infof("Received %d role_permission records from database", len(rolePerms))

	accessRules := make([]*systemv1.AccessRule, 0, len(rolePerms))
	for _, rp := range rolePerms {
		if rp.Edges.Role == nil || rp.Edges.Permission == nil {
			p.log.Warnf("Inconsistent data: RolePermission record (ID: %d) has nil Role or Permission. Skipping.", rp.ID)
			continue
		}

		if len(rp.Edges.Permission.Edges.Resources) == 0 {
			p.log.Debugf("permission '%s' has no associated resources, skipping", rp.Edges.Permission.Keyword)
			continue
		}

		for _, resource := range rp.Edges.Permission.Edges.Resources {
			if resource == nil {
				continue
			}
			accessRules = append(accessRules, &systemv1.AccessRule{
				Subject: rp.Edges.Role.Keyword,
				Object:  resource.Operation,
				Action:  "ANY",
				Domain:  "*",
			})
		}
	}

	p.log.WithContext(ctx).Infof("Transformed into %d access rules", len(accessRules))
	return accessRules, nil
}

var _ dto.PolicyProvider = (*policyDBProvider)(nil)
