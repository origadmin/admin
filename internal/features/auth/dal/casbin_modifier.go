/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strings"

	"github.com/casbin/casbin/v3/persist"

	watcher "github.com/origadmin/casbin-watcher/v3"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data"
)

// casbinModifier implements the authz.PolicyModifier interface by writing directly
// to the Casbin adapter and then triggering a watcher notification. This decouples
// the policy writing logic from the policy reading (enforcer) logic.
type casbinModifier struct {
	adapter persist.Adapter
	watcher persist.Watcher
	log     *log.Helper
}

// NewCasbinModifier creates a new PolicyModifier that requires both the data adapter
// and a watcher to broadcast updates. It accepts interfaces for maximum flexibility.
func NewCasbinModifier(adapter *data.CasbinAdapter, watcher *watcher.Watcher, logger log.Logger) (authz.PolicyModifier,
	error) {
	return &casbinModifier{
		adapter: adapter,
		watcher: watcher,
		log:     log.NewHelper(log.With(logger, "module", "auth.dal.casbin_modifier")),
	}, nil
}

// --- User-Role Management ---

func (m *casbinModifier) AddUserRole(ctx context.Context, userID string, roleID string) (bool, error) {
	m.log.WithContext(ctx).Debugf("Adding user-role link: User=%s, Role=%s", userID, roleID)
	// In Casbin, a user-role assignment is a 'g' (grouping) policy.
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
	err := m.adapter.RemoveFilteredPolicy("g", "g", 0, userID)
	return m.handleAdapterResult(ctx, err)
}

// --- Role-Permission Management ---

func (m *casbinModifier) AddRolePermission(ctx context.Context, roleID string, spec authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Adding role-permission link: Role=%s, Resource=%s, Action=%s", roleID, spec.Resource, spec.Action)
	// In Casbin, a role-permission assignment is a 'p' (policy) rule.
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
	// This removes 'p' rules where the subject is the user ID.
	err := m.adapter.RemoveFilteredPolicy("p", "p", 0, userID)
	return m.handleAdapterResult(ctx, err)
}

// handleAdapterResult interprets the adapter's result and triggers the watcher on success.
func (m *casbinModifier) handleAdapterResult(ctx context.Context, err error) (bool, error) {
	if err == nil {
		// The operation changed the state in the adapter.
		// Now, notify other instances via the watcher.
		if m.watcher != nil {
			if err := m.watcher.Update(); err != nil {
				m.log.WithContext(ctx).Errorf("Failed to broadcast policy update via watcher: %v", err)
				// Return the watcher error, as the update notification is critical.
				return false, err
			}
			m.log.WithContext(ctx).Info("Policy update broadcasted via watcher.")
		} else {
			m.log.WithContext(ctx).Warn("Watcher is not configured; policy changes will not be broadcasted.")
		}
		return true, nil
	}

	// Handle cases where no change was made (e.g., duplicate or non-existent rule).
	if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "not found") {
		m.log.WithContext(ctx).Debugf("Adapter returned a no-op error, not triggering watcher: %v", err)
		return false, nil // No change was made, so no error and no notification.
	}

	m.log.WithContext(ctx).Errorf("Adapter operation failed with an unexpected error: %v", err)
	return false, err
}
