/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"

	systempb "origadmin/application/admin/api/v1/services/system"
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

	// 1. Fetch the new, correct set of access rules for these roles.
	rules, err := s.provider.ListPoliciesForRoles(ctx, roleKeywords...)
	if err != nil {
		s.log.WithContext(ctx).Errorf("SYNCER: Failed to fetch policies for roles %v: %v", roleKeywords, err)
		return fmt.Errorf("failed to fetch policies for roles: %w", err)
	}

	s.log.WithContext(ctx).Infof("SYNCER: Received %d access rules from provider.", len(rules))

	// 2. Now that we have the new rules, clear the old ones.
	for _, roleKeyword := range roleKeywords {
		s.log.WithContext(ctx).Infof("SYNCER: Removing old permissions for role '%s'", roleKeyword)
		if _, err := s.modifier.RemovePermissions(ctx, roleKeyword); err != nil {
			s.log.WithContext(ctx).Warnf("Failed to clear permissions for role '%s': %v", roleKeyword, err)
		}
	}

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

	// 4. Sync user-role associations for these roles.
	// This is crucial because role permissions changed, and we need to ensure
	// users assigned to these roles maintain their group relationships in Casbin.
	s.log.WithContext(ctx).Infof("SYNCER: Syncing user-role associations for roles %v.", roleKeywords)

	// Fetch all policies to get user-role mappings
	allPolicies, err := s.provider.ListAllPolicies(ctx)
	if err != nil {
		s.log.WithContext(ctx).Warnf("SYNCER: Failed to fetch all policies for user-role sync: %v", err)
		// Continue anyway - we've at least updated the role permissions
	} else {
		// Filter grouping rules for the target roles
		usersToResync := make(map[string]authz.RoleSpec)
		for _, gr := range allPolicies.GetGroupingRules() {
			// Check if this grouping rule is for one of our target roles
			for _, roleKeyword := range roleKeywords {
				if gr.GetGroup() == roleKeyword {
					// Need to re-sync this user-role relationship
					usersToResync[gr.GetUser()] = authz.RoleSpec{
						Role:   gr.GetGroup(),
						Domain: gr.GetDomain(),
					}
					break
				}
			}
		}

		// Clear and re-add user-role assignments for these roles
		for user, roleSpec := range usersToResync {
			s.log.WithContext(ctx).Infof("SYNCER: Re-syncing user-role assignment: %s -> %s", user, roleSpec.Role)
			// Remove old assignment
			if _, err := s.modifier.RemoveRoles(ctx, user, roleSpec); err != nil {
				s.log.WithContext(ctx).Warnf("Failed to remove role assignment for user '%s': %v", user, err)
			}
			// Re-add assignment
			if _, err := s.modifier.AddRoles(ctx, user, roleSpec); err != nil {
				s.log.WithContext(ctx).Errorf("Failed to add role assignment for user '%s': %v", user, err)
			}
		}
		if len(usersToResync) > 0 {
			s.log.WithContext(ctx).Infof("SYNCER: Re-synced %d user-role assignments", len(usersToResync))
		}
	}

	s.log.WithContext(ctx).Infof("SYNCER: Targeted policy synchronization for roles %v finished successfully.", roleKeywords)
	return nil
}
