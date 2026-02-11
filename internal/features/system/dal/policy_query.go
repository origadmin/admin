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
	r.log.WithContext(ctx).Debugw("msg", "ListRolePermissions")
	// Query the explicit join table entity `RolePermission`.
	// We need to eager-load Role, and Permission with its Resources.
	rolePerms, err := r.db.RolePermission(ctx).Query().
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListRolePermissions", "err", err)
		return nil, err
	}

	r.log.WithContext(ctx).Debugw("msg", "ListRolePermissions", "count", len(rolePerms))

	// Manually convert the slice using the generated item converter.
	dtos := make([]*types.RolePermission, len(rolePerms))
	for i, rp := range rolePerms {
		dtos[i] = dto.ConvertRolePermissionToRolePermissionPB(rp)
		// Debug: check if edges are loaded
		if rp.Edges.Role == nil {
			r.log.WithContext(ctx).Warnw("msg", "ListRolePermissions", "warn", "RolePermission has nil Role edge", "id", rp.ID)
		}
		if rp.Edges.Permission == nil {
			r.log.WithContext(ctx).Warnw("msg", "ListRolePermissions", "warn", "RolePermission has nil Permission edge", "id", rp.ID)
		}
	}

	return dtos, nil
}

// ListRolePermissionsByRoleKeywords queries role-permission relations for a specific set of role keywords.
func (r *policyQueryRepo) ListRolePermissionsByRoleKeywords(ctx context.Context, roleKeywords ...string) ([]*types.RolePermission, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListRolePermissionsByRoleKeywords", "keywords", roleKeywords)
	rolePerms, err := r.db.RolePermission(ctx).Query().
		Where(rolepermission.HasRoleWith(role.KeywordIn(roleKeywords...))).
		WithRole().
		WithPermission(func(q *ent.PermissionQuery) {
			q.WithResources()
		}).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListRolePermissionsByRoleKeywords", "err", err)
		return nil, err
	}

	r.log.WithContext(ctx).Debugw("msg", "ListRolePermissionsByRoleKeywords", "count", len(rolePerms), "keywords", roleKeywords)

	dtos := make([]*types.RolePermission, len(rolePerms))
	for i, rp := range rolePerms {
		dtos[i] = dto.ConvertRolePermissionToRolePermissionPB(rp)
	}

	return dtos, nil
}

// ListPermissions queries all permissions, converts entities to PB types, and returns them.
func (r *policyQueryRepo) ListPermissions(ctx context.Context) ([]*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListPermissions")
	// Eager-load resources as they contain the actual service/method info.
	permissions, err := r.db.Permission(ctx).Query().WithResources().All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListPermissions", "err", err)
		return nil, err
	}

	// Use the generated slice converter.
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

// ListRolesByIDs queries roles by their IDs.
func (r *policyQueryRepo) ListRolesByIDs(ctx context.Context, ids ...int64) ([]*types.Role, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListRolesByIDs", "ids", ids)
	if len(ids) == 0 {
		return nil, nil
	}
	roles, err := r.db.Role(ctx).Query().
		Where(role.IDIn(ids...)).
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListRolesByIDs", "err", err)
		return nil, err
	}
	r.log.WithContext(ctx).Debugw("msg", "ListRolesByIDs", "count", len(roles), "ids", ids)
	return dto.ConvertRolesToRolesPB(roles), nil
}

// ListPermissionsByIDs queries permissions by their IDs.
func (r *policyQueryRepo) ListPermissionsByIDs(ctx context.Context, ids ...int64) ([]*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListPermissionsByIDs", "ids", ids)
	if len(ids) == 0 {
		return nil, nil
	}
	permissions, err := r.db.Permission(ctx).Query().
		Where(permission.IDIn(ids...)).
		WithResources().
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListPermissionsByIDs", "err", err)
		return nil, err
	}
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

// ListUserRoles queries the user_roles join table, converts entities to PB types, and returns them.
func (r *policyQueryRepo) ListUserRoles(ctx context.Context) ([]*types.UserRole, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListUserRoles")
	// Query the explicit join table entity `UserRole`.
	userRoles, err := r.db.UserRole(ctx).Query().
		WithUser(). // Eager-load the associated User.
		WithRole(). // Eager-load the associated Role.
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListUserRoles", "err", err)
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
	r.log.WithContext(ctx).Debugw("msg", "ListAllPolicies")
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
		r.log.WithContext(ctx).Errorw("msg", "ListAllPolicies.Tx", "err", err)
		return nil, nil, err
	}

	return rolePerms, userRoles, nil
}

// ListPolicies queries policies from casbin_rule table and converts them to authzv1.PolicySpec format.
// This implements the dto.PolicyQueryRepo interface for providing policy data to other services.
func (r *policyQueryRepo) ListPolicies(ctx context.Context, req *system.ListPoliciesRequest) ([]*authzv1.PolicySpec, int32, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListPolicies", "req", req)

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
		r.log.WithContext(ctx).Errorw("msg", "ListPolicies.All", "err", err)
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
		r.log.WithContext(ctx).Errorw("msg", "ListPolicies.Count", "err", err)
		return policies, int32(len(policies)), nil // Return current count as fallback
	}

	r.log.WithContext(ctx).Debugw("msg", "ListPolicies", "count", len(policies), "total", total)
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
	case "g":
		// Grouping/Role rule: subject, role, domain
		if rule.V1 != "" {
			policy.Roles = []string{rule.V1} // V1 is role
		}
		if rule.V2 != "" {
			policy.Domain = &rule.V2
		}
	}

	return policy
}

// ListUserRolePermissions retrieves all role permissions and user roles in a single, atomic operation.
func (r *policyQueryRepo) ListUserRolePermissions(ctx context.Context) ([]*types.RolePermission, []*types.UserRole, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListUserRolePermissions")
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
		r.log.WithContext(ctx).Errorw("msg", "ListUserRolePermissions.Tx", "err", err)
		return nil, nil, err
	}

	return rolePerms, userRoles, nil
}
