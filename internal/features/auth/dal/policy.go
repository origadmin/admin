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

// AuthPolicyRepository is the implementation of authz.PolicyManager.
// It is responsible for directly interacting with the database to manage Casbin policies and notifying the Watcher.
type AuthPolicyRepository struct {
	adapter persist.Adapter // The persistence adapter for Casbin, e.g., *gormadapter.Adapter.
	watcher persist.Watcher // The watcher to notify other instances of policy changes.
	log     *log.Helper
}

// NewAuthPolicyRepository creates a new instance of AuthPolicyRepository.
// It requires a Casbin Adapter and a Watcher to be injected.
func NewAuthPolicyRepository(adapter *data.CasbinAdapter, watcher *watcher.Watcher,
	logger log.Logger) *AuthPolicyRepository {
	return &AuthPolicyRepository{
		adapter: adapter,
		watcher: watcher,
		log:     log.NewHelper(log.With(logger, "module", "auth.dal.policy_repo")),
	}
}

// AddGroupingPolicy implements the authz.PolicyManager interface.
// It writes the policy directly to the database via the Adapter and then notifies the Watcher.
func (r *AuthPolicyRepository) AddGroupingPolicy(ctx context.Context, subject string, group string) (bool, error) {
	r.log.WithContext(ctx).Debugf("[AuthPolicyRepo] Adding grouping policy: sub=%s, group=%s", subject, group)

	// The Casbin Adapter interface for adding a single policy is AddPolicy.
	// For a 'g' rule, the section is 'g', ptype is 'g', and the rule is a slice of strings.
	err := r.adapter.AddPolicy("g", "g", []string{subject, group})
	if err != nil {
		// Handle potential unique constraint errors, which we treat as "already exists".
		if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "duplicate") {
			r.log.WithContext(ctx).Debugf("Grouping policy already exists in DB for sub=%s, group=%s", subject, group)
			return false, nil // Indicates that the policy was not added because it already existed.
		}
		return false, err
	}

	// Notify other instances via the watcher.
	if r.watcher != nil {
		if err := r.watcher.Update(); err != nil {
			// Log the error but do not fail the operation, as the primary data change was successful.
			r.log.WithContext(ctx).Errorf("Failed to notify watcher after policy update: %v", err)
		}
	}

	return true, nil // Indicates that the policy was successfully added.
}

// RemoveGroupingPolicy implements the authz.PolicyManager interface.
func (r *AuthPolicyRepository) RemoveGroupingPolicy(ctx context.Context, subject string, group string) (bool, error) {
	r.log.WithContext(ctx).Debugf("[AuthPolicyRepo] Removing grouping policy: sub=%s, group=%s", subject, group)

	err := r.adapter.RemovePolicy("g", "g", []string{subject, group})
	if err != nil {
		return false, err
	}

	// Notify other instances.
	if r.watcher != nil {
		if err := r.watcher.Update(); err != nil {
			r.log.WithContext(ctx).Errorf("Failed to notify watcher after policy removal: %v", err)
		}
	}
	return true, nil
}

// AddPolicy implements the authz.PolicyManager interface.
func (r *AuthPolicyRepository) AddPolicy(ctx context.Context, subject string, resource string, action string) (bool, error) {
	r.log.WithContext(ctx).Debugf("[AuthPolicyRepo] Adding policy: sub=%s, obj=%s, act=%s", subject, resource, action)

	err := r.adapter.AddPolicy("p", "p", []string{subject, resource, action})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") || strings.Contains(err.Error(), "duplicate") {
			r.log.WithContext(ctx).Debugf("Policy already exists in DB for sub=%s, obj=%s, act=%s", subject, resource, action)
			return false, nil
		}
		return false, err
	}

	// Notify other instances.
	if r.watcher != nil {
		if err := r.watcher.Update(); err != nil {
			r.log.WithContext(ctx).Errorf("Failed to notify watcher after policy update: %v", err)
		}
	}
	return true, nil
}

// RemovePolicy implements the authz.PolicyManager interface.
func (r *AuthPolicyRepository) RemovePolicy(ctx context.Context, subject string, resource string, action string) (bool, error) {
	r.log.WithContext(ctx).Debugf("[AuthPolicyRepo] Removing policy: sub=%s, obj=%s, act=%s", subject, resource, action)

	err := r.adapter.RemovePolicy("p", "p", []string{subject, resource, action})
	if err != nil {
		return false, err
	}

	// Notify other instances.
	if r.watcher != nil {
		if err := r.watcher.Update(); err != nil {
			r.log.WithContext(ctx).Errorf("Failed to notify watcher after policy removal: %v", err)
		}
	}
	return true, nil
}

// Ensure AuthPolicyRepository implements the authz.PolicyManager interface.
var _ authz.PolicyManager = (*AuthPolicyRepository)(nil)
