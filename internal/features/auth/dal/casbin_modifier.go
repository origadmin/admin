/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strings"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data"
)

// casbinModifier implements the authz.PolicyModifier interface using the low-level Casbin adapter.
// Its purpose is to translate business-level requests (e.g., "add role to user")
// into Casbin-specific operations, hiding the implementation details from the caller.
type casbinModifier struct {
	adapter *data.CasbinAdapter
	log     *log.Helper
}

// NewCasbinModifier creates a new PolicyModifier implementation.
func NewCasbinModifier(adapter *data.CasbinAdapter, logger log.Logger) (authz.PolicyModifier, error) {
	return &casbinModifier{
		adapter: adapter,
		log:     log.NewHelper(log.With(logger, "module", "auth.dal.casbin_modifier")),
	}, nil
}

// --- User-Role Management ---

func (m *casbinModifier) AddUserRole(ctx context.Context, userID string, roleID string) (bool, error) {
	m.log.WithContext(ctx).Debugf("Adding user-role link: User=%s, Role=%s", userID, roleID)
	// In Casbin, a user-role assignment is a 'g' (grouping) policy.
	// The format is g, user_subject, role_name
	err := m.adapter.AddPolicy("g", "g", []string{userID, roleID})
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemoveUserRole(ctx context.Context, userID string, roleID string) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing user-role link: User=%s, Role=%s", userID, roleID)
	err := m.adapter.RemovePolicy("g", "g", []string{userID, roleID})
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemoveAllUserRoles(ctx context.Context, userID string) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing all roles for user: %s", userID)
	// This removes all 'g' policies where the first field (V0) matches the userID.
	err := m.adapter.RemoveFilteredPolicy("g", "g", 0, userID)
	return m.handleAdapterResult(ctx, err)
}

// --- Role-Permission Management ---

func (m *casbinModifier) AddRolePermission(ctx context.Context, roleID string, spec authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Adding role-permission link: Role=%s, Resource=%s, Action=%s", roleID, spec.Resource, spec.Action)
	// In Casbin, a role-permission assignment is a 'p' (policy) rule.
	// The format is p, role_name, resource, action
	err := m.adapter.AddPolicy("p", "p", []string{roleID, spec.Resource, spec.Action})
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemoveRolePermission(ctx context.Context, roleID string, spec authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing role-permission link: Role=%s, Resource=%s, Action=%s", roleID, spec.Resource, spec.Action)
	err := m.adapter.RemovePolicy("p", "p", []string{roleID, spec.Resource, spec.Action})
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemoveAllRolePermissions(ctx context.Context, roleID string) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing all permissions for role: %s", roleID)
	// This removes all 'p' policies where the first field (V0) matches the roleID.
	err := m.adapter.RemoveFilteredPolicy("p", "p", 0, roleID)
	return m.handleAdapterResult(ctx, err)
}

// --- Direct User-Permission Management ---

func (m *casbinModifier) AddUserPermission(ctx context.Context, userID string, spec authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Adding direct user-permission: User=%s, Resource=%s, Action=%s", userID, spec.Resource, spec.Action)
	// This is also a 'p' rule, but the subject is a user instead of a role.
	err := m.adapter.AddPolicy("p", "p", []string{userID, spec.Resource, spec.Action})
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemoveUserPermission(ctx context.Context, userID string, spec authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing direct user-permission: User=%s, Resource=%s, Action=%s", userID, spec.Resource, spec.Action)
	err := m.adapter.RemovePolicy("p", "p", []string{userID, spec.Resource, spec.Action})
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemoveAllUserPermissions(ctx context.Context, userID string) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing all direct permissions for user: %s", userID)
	// This is tricky. We only want to remove 'p' rules for this user, but not 'g' rules.
	// We assume direct user permissions are 'p' rules where the subject is the user ID.
	err := m.adapter.RemoveFilteredPolicy("p", "p", 0, userID)
	return m.handleAdapterResult(ctx, err)
}

// handleAdapterResult is a helper to interpret the result from the adapter.
// The adapter might return an error for duplicates, which we treat as "not added/removed".
func (m *casbinModifier) handleAdapterResult(ctx context.Context, err error) (bool, error) {
	if err == nil {
		return true, nil // Operation was successful and changed state.
	}
	// Check for unique constraint violation, which means the rule already exists/doesn't exist.
	if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "duplicate") {
		m.log.WithContext(ctx).Debugf("Adapter returned a duplicate/unique constraint error, treating as no-op: %v", err)
		return false, nil // No change was made.
	}
	// For RemoveFilteredPolicy, if no rows are affected, it might not be an error.
	// The current adapter implementation doesn't distinguish this, so we rely on error text.
	// A better adapter might return a specific error type or affected rows count.
	if strings.Contains(err.Error(), "not found") {
		return false, nil
	}

	m.log.WithContext(ctx).Errorf("Adapter operation failed with an unexpected error: %v", err)
	return false, err
}
