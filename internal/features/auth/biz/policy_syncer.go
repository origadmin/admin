/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"

	systempb "origadmin/application/admin/api/v1/services/system"
)

const (
	// retryInterval is the duration to wait before retrying a policy fetch
	// if the initial attempt returns no rules, to account for transactional delays.
	retryInterval = 200 * time.Millisecond
)

// PolicyProvider defines the interface for fetching the source-of-truth policies
// in a pre-processed, implementation-agnostic format.
type PolicyProvider interface {
	ListAllPolicies(ctx context.Context) (*systempb.ListAllPoliciesResponse, error)
	// ListPoliciesForRoles fetches all access rules for a specific set of roles, identified by their keywords.
	ListPoliciesForRoles(ctx context.Context, roleKeywords ...string) ([]*systempb.AccessRule, error)
}

// PolicySyncer is responsible for synchronizing policies from a PolicyProvider
// to a PolicyModifier. Its logic is greatly simplified because the provider
// now returns ready-to-use generic policy rules.
type PolicySyncer struct {
	provider PolicyProvider
	modifier authz.PolicyModifier
	log      *log.Helper
}

// NewPolicySyncer creates a new PolicySyncer.
func NewPolicySyncer(provider PolicyProvider, modifier authz.PolicyModifier, logger log.Logger) *PolicySyncer {
	return &PolicySyncer{
		provider: provider,
		modifier: modifier,
		log:      log.NewHelper(log.With(logger, "module", "auth.biz.policy_syncer")),
	}
}

// Sync performs a full, destructive synchronization of policies.
// It fetches all generic rules and applies them, clearing previous rules for managed subjects.
func (s *PolicySyncer) Sync(ctx context.Context) error {
	s.log.WithContext(ctx).Info("Starting policy synchronization...")

	// 1. Fetch generic, pre-processed rules from the source of truth.
	resp, err := s.provider.ListAllPolicies(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to fetch policies from provider: %v", err)
		return fmt.Errorf("failed to fetch policies: %w", err)
	}

	s.log.WithContext(ctx).Infof("Fetched %d access rules and %d grouping rules.",
		len(resp.GetAccessRules()), len(resp.GetGroupingRules()))

	// 2. Collect all unique subjects and users to be managed.
	// This allows us to clear only the policies for the entities we are about to sync.
	subjectsToClear := make(map[string]bool) // For 'p' rules (roles and direct users)
	usersToClear := make(map[string]bool)    // For 'g' rules (users in groups)

	for _, rule := range resp.GetAccessRules() {
		subjectsToClear[rule.GetSubject()] = true
	}
	for _, rule := range resp.GetGroupingRules() {
		usersToClear[rule.GetUser()] = true
	}

	// 3. Clear existing policies for all managed entities before applying new ones.
	// This ensures a clean, consistent state aligned with the source of truth.
	s.log.WithContext(ctx).Info("Clearing existing policies for managed subjects and users...")
	for subject := range subjectsToClear {
		// This handles both role permissions and direct user permissions ('p' rules)
		if _, err := s.modifier.RemovePermissions(ctx, subject); err != nil {
			s.log.WithContext(ctx).Warnf("Failed to clear permissions for subject '%s': %v", subject, err)
		}
	}
	for user := range usersToClear {
		// This handles user-role assignments ('g' rules)
		if _, err := s.modifier.RemoveRoles(ctx, user); err != nil {
			s.log.WithContext(ctx).Warnf("Failed to clear roles for user '%s': %v", user, err)
		}
	}

	// 4. Add all new access rules (translating to 'p' rules for Casbin).
	s.log.WithContext(ctx).Info("Applying new access policies...")
	for _, rule := range resp.GetAccessRules() {
		spec := authz.RuleSpec{
			Resource: rule.GetObject(),
			Action:   rule.GetAction(),
			Domain:   rule.GetDomain(),
		}
		subject := rule.GetSubject()

		// The PolicyModifier interface abstracts away the difference between role and user permissions.
		// We can use a single method for both.
		if _, err := s.modifier.AddPermissions(ctx, subject, spec); err != nil {
			s.log.WithContext(ctx).Errorf("Failed to add permission for subject '%s': %v", subject, err)
		}
	}

	// 5. Add all new grouping rules (translating to 'g' rules for Casbin).
	s.log.WithContext(ctx).Info("Applying new grouping policies...")
	for _, rule := range resp.GetGroupingRules() {
		roleSpec := authz.RoleSpec{
			Role:   rule.GetGroup(),
			Domain: rule.GetDomain(),
		}
		if _, err := s.modifier.AddRoles(ctx, rule.GetUser(), roleSpec); err != nil {
			s.log.WithContext(ctx).Errorf("Failed to add user '%s' to group '%s': %v", rule.GetUser(), rule.GetGroup(), err)
		}
	}

	s.log.WithContext(ctx).Info("Policy synchronization finished successfully.")
	return nil
}

// SyncRoles performs a targeted synchronization for a specific set of roles.
// If roleKeywords is empty, it falls back to a full synchronization.
func (s *PolicySyncer) SyncRoles(ctx context.Context, roleKeywords ...string) error {
	if len(roleKeywords) == 0 {
		s.log.WithContext(ctx).Info("No role keywords provided, performing a full policy synchronization.")
		return s.Sync(ctx)
	}

	s.log.WithContext(ctx).Infof("SYNCER: Starting targeted policy synchronization for roles: %v", roleKeywords)

	// 1. Clear existing permissions for the specified roles.
	// The subject in 'p' rules for roles is the role keyword itself.
	for _, roleKeyword := range roleKeywords {
		s.log.WithContext(ctx).Infof("SYNCER: Removing permissions for role '%s'", roleKeyword)
		if _, err := s.modifier.RemovePermissions(ctx, roleKeyword); err != nil {
			s.log.WithContext(ctx).Warnf("Failed to clear permissions for role '%s': %v", roleKeyword, err)
			// Continue even if clearing fails, as adding new rules might fix inconsistencies.
		}
	}

	// 2. Fetch the new, correct set of access rules for these roles.
	rules, err := s.provider.ListPoliciesForRoles(ctx, roleKeywords...)
	if err != nil {
		s.log.WithContext(ctx).Errorf("SYNCER: Failed to fetch policies for roles %v: %v", roleKeywords, err)
		return fmt.Errorf("failed to fetch policies for roles: %w", err)
	}

	// HACK: Retry once if no rules are found, to mitigate race conditions where the message
	// is consumed before the writing transaction is fully committed and visible.
	if len(rules) == 0 {
		s.log.WithContext(ctx).Warnf("SYNCER: Received 0 access rules on first attempt. Retrying after %v...", retryInterval)
		time.Sleep(retryInterval)
		rules, err = s.provider.ListPoliciesForRoles(ctx, roleKeywords...)
		if err != nil {
			s.log.WithContext(ctx).Errorf("SYNCER: Failed to fetch policies on retry for roles %v: %v", roleKeywords, err)
			return fmt.Errorf("failed to fetch policies on retry: %w", err)
		}
	}

	s.log.WithContext(ctx).Infof("SYNCER: Received %d access rules from provider.", len(rules))

	// 3. Add the new access rules.
	if len(rules) > 0 {
		s.log.WithContext(ctx).Infof("SYNCER: Applying %d new access rules for roles %v.", len(rules), roleKeywords)
		for _, rule := range rules {
			spec := authz.RuleSpec{
				Resource: rule.GetObject(),
				Action:   rule.GetAction(),
				Domain:   rule.GetDomain(),
			}
			if _, err := s.modifier.AddPermissions(ctx, rule.GetSubject(), spec); err != nil {
				s.log.WithContext(ctx).Errorf("Failed to add permission for subject '%s': %v", rule.GetSubject(), err)
			}
		}
	}

	s.log.WithContext(ctx).Infof("SYNCER: Targeted policy synchronization for roles %v finished successfully.", roleKeywords)
	return nil
}
