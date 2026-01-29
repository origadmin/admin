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
		log:     log.NewHelper(log.With(logger, "module", "auth.dal.casbin_modifier")),
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
		filters := map[string]string{"v0": subject}
		err := m.adapter.RemovePoliciesByFields("g", filters)
		return m.handleAdapterResult(ctx, err)
	}

	// Process each spec as a separate filter.
	var firstErr error
	for _, r := range roles {
		filters := map[string]string{"v0": subject}
		if r.Role != "" {
			filters["v1"] = r.Role
		}
		// Only filter by domain if it is explicitly provided.
		if r.Domain != "" {
			filters["v2"] = r.Domain
		}

		err := m.adapter.RemovePoliciesByFields("g", filters)
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
		rules[i] = []string{subject, p.Resource, p.Action, p.Domain}
	}

	err := m.adapter.AddPolicies("p", "p", rules)
	return m.handleAdapterResult(ctx, err)
}

func (m *casbinModifier) RemovePermissions(ctx context.Context, subject string, permissions ...authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Removing permissions for subject: %s", subject)

	// If no specs are provided, remove all permissions for the subject across all domains.
	if len(permissions) == 0 {
		filters := map[string]string{"v0": subject}
		err := m.adapter.RemovePoliciesByFields("p", filters)
		return m.handleAdapterResult(ctx, err)
	}

	// Process each spec as a separate filter.
	var firstErr error
	for _, p := range permissions {
		filters := map[string]string{"v0": subject}
		if p.Resource != "" {
			filters["v1"] = p.Resource
		}
		if p.Action != "" {
			filters["v2"] = p.Action
		}
		// Only filter by domain if it is explicitly provided.
		if p.Domain != "" {
			filters["v3"] = p.Domain
		}

		err := m.adapter.RemovePoliciesByFields("p", filters)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return m.handleAdapterResult(ctx, firstErr)
}

func (m *casbinModifier) UpdatePermission(ctx context.Context, subject string, oldPerm authz.RuleSpec, newPerm authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Debugf("Updating permission for subject: %s", subject)

	oldRule := []string{subject, oldPerm.Resource, oldPerm.Action, oldPerm.Domain}
	newRule := []string{subject, newPerm.Resource, newPerm.Action, newPerm.Domain}

	err := m.adapter.UpdatePolicy("p", "p", oldRule, newRule)
	return m.handleAdapterResult(ctx, err)
}

// handleAdapterResult interprets the adapter's result.
func (m *casbinModifier) handleAdapterResult(ctx context.Context, err error) (bool, error) {
	if err == nil {
		return true, nil
	}

	if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "not found") {
		m.log.WithContext(ctx).Debugf("Adapter returned a no-op error: %v", err)
		return false, nil
	}

	m.log.WithContext(ctx).Errorf("Adapter operation failed with an unexpected error: %v", err)
	return false, err
}
