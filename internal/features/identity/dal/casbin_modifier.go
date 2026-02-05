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

// casbinModifier implements the authz.PolicyModifier interface.
type casbinModifier struct {
	adapter *data.CasbinAdapter // Use concrete type to access custom methods
	log     *log.Helper
}

// NewCasbinModifier creates a new PolicyModifier.
func NewCasbinModifier(adapter *data.CasbinAdapter, logger log.Logger) (authz.PolicyModifier, error) {
	return &casbinModifier{
		adapter: adapter,
		log:     log.NewHelper(log.With(logger, "module", "identity.dal.casbin_modifier")),
	}, nil
}

func (m *casbinModifier) AddRoles(ctx context.Context, subject string, roles ...authz.RoleSpec) (bool, error) {
	if len(roles) == 0 {
		return false, nil
	}
	m.log.WithContext(ctx).Debugf("Adding roles for subject: %s", subject)

	rules := make([][]string, len(roles))
	for i, r := range roles {
		rules[i] = []string{subject, r.Role, r.Domain}
	}

	err := m.adapter.AddPolicies("g", "g", rules)
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemoveRoles(ctx context.Context, subject string, roles ...authz.RoleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing roles for subject: %s", subject)

	// If no specs are provided, remove all roles for the subject across all domains.
	if len(roles) == 0 {
		err := m.adapter.RemoveFilteredPolicy("", "g", 0, subject)
		return m.handleAdapterResult(ctx, err)
	}

	// Process each spec as a separate filter.
	var firstErr error
	for _, r := range roles {
		args := []string{subject}
		if r.Role != "" {
			args = append(args, r.Role)
			if r.Domain != "" {
				args = append(args, r.Domain)
			}
		}
		err := m.adapter.RemoveFilteredPolicy("", "g", 0, args...)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return m.handleAdapterResult(ctx, firstErr)
}

func (m *casbinModifier) UpdateRole(ctx context.Context, subject string, oldRole authz.RoleSpec, newRole authz.RoleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Updating role for subject: %s", subject)

	oldRule := []string{subject, oldRole.Role, oldRole.Domain}
	newRule := []string{subject, newRole.Role, newRole.Domain}

	err := m.adapter.UpdatePolicy("g", "g", oldRule, newRule)
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) AddPermissions(ctx context.Context, subject string, permissions ...authz.RuleSpec) (bool, error) {
	if len(permissions) == 0 {
		return false, nil
	}
	m.log.WithContext(ctx).Debugf("Adding permissions for subject: %s", subject)

	rules := make([][]string, len(permissions))
	for i, p := range permissions {
		rules[i] = []string{subject, p.Domain, p.Resource, p.Action}
	}

	err := m.adapter.AddPolicies("p", "p", rules)
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemovePermissions(ctx context.Context, subject string, permissions ...authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing permissions for subject: %s", subject)

	if len(permissions) == 0 {
		err := m.adapter.RemoveFilteredPolicy("", "p", 0, subject)
		return m.handleAdapterResult(ctx, err)
	}

	var firstErr error
	for _, p := range permissions {
		args := []string{subject, p.Domain, p.Resource, p.Action}
		err := m.adapter.RemoveFilteredPolicy("", "p", 0, args...)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return m.handleAdapterResult(ctx, firstErr)
}

func (m *casbinModifier) UpdatePermission(ctx context.Context, subject string, oldPerm authz.RuleSpec, newPerm authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Updating permission for subject: %s", subject)

	oldRule := []string{subject, oldPerm.Domain, oldPerm.Resource, oldPerm.Action}
	newRule := []string{subject, newPerm.Domain, newPerm.Resource, newPerm.Action}

	err := m.adapter.UpdatePolicy("p", "p", oldRule, newRule)
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) ClearPolicies(ctx context.Context, subjects ...string) (bool, error) {
	// Case 1: No subjects provided, which means clear ALL policies.
	if len(subjects) == 0 {
		m.log.WithContext(ctx).Info("Clearing ALL 'p' and 'g' policies from storage.")
		// Clear all 'p' (permission) rules.
		_, errP := m.handleAdapterResult(ctx, m.adapter.RemoveFilteredPolicy("", "p", 0))
		if errP != nil {
			m.log.WithContext(ctx).Errorf("Failed to clear all 'p' policies: %v", errP)
			return false, errP
		}
		// Clear all 'g' (grouping/role) rules.
		_, errG := m.handleAdapterResult(ctx, m.adapter.RemoveFilteredPolicy("", "g", 0))
		if errG != nil {
			m.log.WithContext(ctx).Errorf("Failed to clear all 'g' policies: %v", errG)
			return false, errG
		}
		m.log.WithContext(ctx).Info("Successfully cleared all 'p' and 'g' policies.")
		return true, nil
	}

	// Case 2: One or more subjects are provided, clear policies for each one.
	m.log.WithContext(ctx).Debugf("Clearing all policies for subjects: %v", subjects)
	var firstErr error
	for _, subject := range subjects {
		// Remove all g-rules for subject
		errG := m.adapter.RemoveFilteredPolicy("", "g", 0, subject)
		// Remove all p-rules for subject
		errP := m.adapter.RemoveFilteredPolicy("", "p", 0, subject)

		if errG != nil && firstErr == nil {
			firstErr = errG
		}
		if errP != nil && firstErr == nil {
			firstErr = errP
		}
	}
	return m.handleAdapterResult(ctx, firstErr)
}

// handleAdapterResult interprets the adapter's result.
func (m *casbinModifier) handleAdapterResult(ctx context.Context, err error) (bool, error) {
	if err == nil {
		return true, nil
	}

	// "not found" is a valid case for removal operations, indicating the policy was already gone.
	// We can treat it as a success from the modifier's perspective (the state is what we want it to be).
	if strings.Contains(err.Error(), "not found") {
		m.log.WithContext(ctx).Debugf("Adapter returned a 'not found' error, which is acceptable for removals: %v", err)
		return false, nil // Return false for "changed", but no error.
	}

	if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "duplicate") {
		m.log.WithContext(ctx).Debugf("Adapter returned a no-op error: %v", err)
		return false, nil
	}

	m.log.WithContext(ctx).Errorf("Adapter operation failed with an unexpected error: %v", err)
	return false, err
}
