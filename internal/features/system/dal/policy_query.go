/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/permission"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/data/entity/ent/rolepermission"
	"origadmin/application/admin/internal/features/system/dto"
)

// policyQueryRepo implements the dto.PolicyQueryRepo interface.
type policyQueryRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewPolicyQueryRepo creates a new policyQuery repository.
func NewPolicyQueryRepo(database *ent.Database, logger log.Logger) dto.PolicyRepo {
	return &policyQueryRepo{
		db:  database,
		log: log.NewHelper(log.With(logger, "module", "dal.policy_query")),
	}
}

// ListRolePermissions queries the role_permissions join table, converts entities to PB types, and returns them.
func (r *policyQueryRepo) ListRolePermissions(ctx context.Context) ([]*types.RolePermission, error) {
	// Query the explicit join table entity `RolePermission`.
	// We need to eager-load Role, and Permission with its Resources.
	rolePerms, err := r.db.RolePermission(ctx).Query().
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query role_permissions: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("DAL: Loaded %d role_permissions with edges", len(rolePerms))

	// Manually convert the slice using the generated item converter.
	dtos := make([]*types.RolePermission, len(rolePerms))
	for i, rp := range rolePerms {
		dtos[i] = dto.ConvertRolePermissionToRolePermissionPB(rp)
		// Debug: check if edges are loaded
		if rp.Edges.Role == nil {
			r.log.WithContext(ctx).Warnf("DAL: RolePermission #%d has nil Role edge", rp.ID)
		}
		if rp.Edges.Permission == nil {
			r.log.WithContext(ctx).Warnf("DAL: RolePermission #%d has nil Permission edge", rp.ID)
		}
	}

	return dtos, nil
}

// ListRolePermissionsByRoleKeywords queries role-permission relations for a specific set of role keywords.
func (r *policyQueryRepo) ListRolePermissionsByRoleKeywords(ctx context.Context, roleKeywords ...string) ([]*types.RolePermission, error) {
	r.log.WithContext(ctx).Infof("DAL: Querying role_permissions for keywords: %v", roleKeywords)
	rolePerms, err := r.db.RolePermission(ctx).Query().
		Where(rolepermission.HasRoleWith(role.KeywordIn(roleKeywords...))).
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query role_permissions by role keywords: %v", err)
		return nil, err
	}

	r.log.WithContext(ctx).Infof("DAL: Found %d role_permission records for keywords: %v", len(rolePerms), roleKeywords)
	for i, rp := range rolePerms {
		r.log.WithContext(ctx).Infof("CONFIRM: DAL Result #%d: RoleID=%d, PermissionID=%d", i, rp.Edges.Role.ID, rp.Edges.Permission.ID)
	}

	dtos := make([]*types.RolePermission, len(rolePerms))
	for i, rp := range rolePerms {
		dtos[i] = dto.ConvertRolePermissionToRolePermissionPB(rp)
	}

	return dtos, nil
}

// ListPermissions queries all permissions, converts entities to PB types, and returns them.
func (r *policyQueryRepo) ListPermissions(ctx context.Context) ([]*types.Permission, error) {
	// Eager-load resources as they contain the actual service/method info.
	permissions, err := r.db.Permission(ctx).Query().WithResources().All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query permissions: %v", err)
		return nil, err
	}

	// Use the generated slice converter.
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

// ListRolesByIDs queries roles by their IDs.
func (r *policyQueryRepo) ListRolesByIDs(ctx context.Context, ids ...int64) ([]*types.Role, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	r.log.WithContext(ctx).Infof("DAL: Querying roles by IDs: %v", ids)
	roles, err := r.db.Role(ctx).Query().
		Where(role.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query roles by IDs %v: %v", ids, err)
		return nil, err
	}
	r.log.WithContext(ctx).Infof("DAL: Found %d roles by IDs %v", len(roles), ids)
	return dto.ConvertRolesToRolesPB(roles), nil
}

// ListPermissionsByIDs queries permissions by their IDs.
func (r *policyQueryRepo) ListPermissionsByIDs(ctx context.Context, ids ...int64) ([]*types.Permission, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	permissions, err := r.db.Permission(ctx).Query().
		Where(permission.IDIn(ids...)).
		WithResources().
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query permissions by IDs: %v", err)
		return nil, err
	}
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

// ListUserRoles queries the user_roles join table, converts entities to PB types, and returns them.
func (r *policyQueryRepo) ListUserRoles(ctx context.Context) ([]*types.UserRole, error) {
	// Query the explicit join table entity `UserRole`.
	userRoles, err := r.db.UserRole(ctx).Query().
		WithUser(). // Eager-load the associated User.
		WithRole(). // Eager-load the associated Role.
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("failed to query user_roles: %v", err)
		return nil, err
	}

	// Manually convert the slice using the generated item converter.
	dtos := make([]*types.UserRole, len(userRoles))
	for i, ur := range userRoles {
		dtos[i] = dto.ConvertUserRoleToUserRolePB(ur)
	}

	return dtos, nil
}

// ListAllPolicies retrieves all role permissions and user roles in a single, atomic operation.
func (r *policyQueryRepo) ListAllPolicies(ctx context.Context) ([]*types.RolePermission, []*types.UserRole, error) {
	var rolePerms []*types.RolePermission
	var userRoles []*types.UserRole

	// Execute both queries within a single read-only transaction to ensure data consistency.
	// The 'ent' database wrapper handles transaction context propagation.
	err := r.db.Tx(ctx, func(txCtx context.Context) error {
		var err error

		// Reuse existing methods, passing the transaction context.
		rolePerms, err = r.ListRolePermissions(txCtx)
		if err != nil {
			return err
		}

		userRoles, err = r.ListUserRoles(txCtx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return rolePerms, userRoles, nil
}

// ListPolicies queries policies from casbin_rule table and converts them to authzv1.PolicySpec format.
// This implements the dto.PolicyQueryRepo interface for providing policy data to other services.
func (r *policyQueryRepo) ListPolicies(ctx context.Context, req *system.ListPoliciesRequest) ([]*authzv1.PolicySpec, int32, error) {
	r.log.WithContext(ctx).Info("DAL: Listing policies from casbin_rule table")

	// Build filter predicates
	preds := make([]predicate.CasbinRule, 0)

	// Filter by type (ptype)
	if req.GetType() != "" {
		preds = append(preds, casbinrule.PtypeEQ(req.GetType()))
	}

	// Filter by subject (v0)
	if req.GetSubject() != "" {
		preds = append(preds, casbinrule.V0EQ(req.GetSubject()))
	}

	// Filter by domain (v1)
	if req.GetDomain() != "" {
		preds = append(preds, casbinrule.V1EQ(req.GetDomain()))
	}

	// Build query with pagination
	query := r.db.CasbinRule(ctx).Query().Where(preds...)

	// Apply pagination
	if req.GetPageSize() > 0 {
		if req.GetPage() > 0 {
			query.Offset(int((req.GetPage() - 1) * req.GetPageSize()))
		}
		query.Limit(int(req.GetPageSize()))
	}

	// Query from database
	rules, err := query.All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("Failed to query casbin rules: %v", err)
		return nil, 0, err
	}

	// Convert Casbin rules to PolicySpec format
	policies := make([]*authzv1.PolicySpec, 0, len(rules))
	for _, rule := range rules {
		policy := r.casbinRuleToPolicySpec(rule)
		policies = append(policies, policy)
	}

	// Get total count
	total, err := r.db.CasbinRule(ctx).Query().Where(preds...).Count(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorf("Failed to count casbin rules: %v", err)
		return policies, int32(len(policies)), nil // Return current count as fallback
	}

	r.log.WithContext(ctx).Infof("DAL: Retrieved %d policies (total: %d) from casbin_rule", len(policies), total)
	return policies, int32(total), nil
}

// casbinRuleToPolicySpec converts a Casbin rule to a PolicySpec.
func (r *policyQueryRepo) casbinRuleToPolicySpec(rule *ent.CasbinRule) *authzv1.PolicySpec {
	policy := &authzv1.PolicySpec{
		Type:    rule.Ptype, // Directly use ptype: "p" or "g"
		Subject: rule.V0,
	}

	// Map based on ptype
	switch rule.Ptype {
	case "p":
		// Permission rule: subject, domain, resource, action
		if rule.V1 != "" {
			policy.Domain = &rule.V1
		}
		if rule.V2 != "" {
			policy.Resources = []string{rule.V2}
		}
		if rule.V3 != "" {
			policy.Actions = []string{rule.V3}
		}
		//effect := "allow"
		//policy.Effect = &effect
	case "g":
		// Grouping/Role rule: subject, role, domain
		if rule.V1 != "" {
			policy.Roles = []string{rule.V1} // V1 is role
		}
		if rule.V2 != "" {
			policy.Domain = &rule.V2
		}
		//effect := "allow"
		//policy.Effect = &effect
	}

	return policy
}

// ListUserRolePermissions retrieves all role permissions and user roles in a single, atomic operation.
func (r *policyQueryRepo) ListUserRolePermissions(ctx context.Context) ([]*types.RolePermission, []*types.UserRole, error) {
	var rolePerms []*types.RolePermission
	var userRoles []*types.UserRole

	// Execute both queries within a single read-only transaction to ensure data consistency.
	// The 'ent' database wrapper handles transaction context propagation.
	err := r.db.Tx(ctx, func(txCtx context.Context) error {
		var err error

		// Reuse existing methods, passing the transaction context.
		rolePerms, err = r.ListRolePermissions(txCtx)
		if err != nil {
			return err
		}

		userRoles, err = r.ListUserRoles(txCtx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return rolePerms, userRoles, nil
}
