/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	authv1 "origadmin/application/admin/api/v1/services/auth"
	systempb "origadmin/application/admin/api/v1/services/system"
)

const (
	DefaultReloadDelay = 2 * time.Second
)

// PolicyProvider defines the interface for fetching the source-of-truth policies
// in a pre-processed, implementation-agnostic format.
type PolicyProvider interface {
	ListAllPolicies(ctx context.Context) (*systempb.ListAllPoliciesResponse, error)
	// ListPoliciesForRoles fetches all access rules for a specific set of roles, identified by their keywords.
	ListPoliciesForRoles(ctx context.Context, roleKeywords ...string) ([]*systempb.AccessRule, error)
}

// PolicySyncer is responsible for synchronizing policies from a PolicyProvider
// to a PolicyModifier. It also manages the synchronization schedule and metrics.
type PolicySyncer struct {
	provider    PolicyProvider
	modifier    authz.PolicyModifier
	reloader    authz.Reloader
	log         *log.Helper
	reloadDelay time.Duration
	reloadMu    sync.Mutex
	reloadTimer *time.Timer

	// Metrics
	lastSyncTime     atomic.Value // time.Time
	lastSyncDuration atomic.Int64 // milliseconds
	totalEvents      atomic.Int64
	successfulSyncs  atomic.Int64
	failedSyncs      atomic.Int64
}

// NewPolicySyncer creates a new PolicySyncer.
func NewPolicySyncer(provider PolicyProvider, modifier authz.PolicyModifier, reloader authz.Reloader, logger log.Logger) *PolicySyncer {
	return &PolicySyncer{
		provider:    provider,
		modifier:    modifier,
		reloader:    reloader,
		log:         log.NewHelper(log.With(logger, "module", "auth.biz.policy_syncer")),
		reloadDelay: DefaultReloadDelay,
	}
}

// ScheduleFullSync schedules a delayed full synchronization with debounce logic.
func (s *PolicySyncer) ScheduleFullSync() {
	s.totalEvents.Add(1)
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()

	if s.reloadTimer != nil {
		s.reloadTimer.Stop()
	}

	s.reloadTimer = time.AfterFunc(s.reloadDelay, func() {
		s.DoFullSyncAndReload()
	})

	s.log.Debugf("Full policy sync scheduled in %v", s.reloadDelay)
}

// ForceSync triggers an immediate full policy synchronization, bypassing the debounce timer.
func (s *PolicySyncer) ForceSync(ctx context.Context) error {
	s.log.Info("Force triggering policy sync (manual request)")

	// Cancel any pending scheduled sync
	s.reloadMu.Lock()
	if s.reloadTimer != nil {
		s.reloadTimer.Stop()
		s.reloadTimer = nil
	}
	s.reloadMu.Unlock()

	// Execute sync immediately
	s.DoFullSyncAndReload()
	return nil
}

// DoFullSyncAndReload performs a full synchronization of all policies and reloads the enforcer.
func (s *PolicySyncer) DoFullSyncAndReload() {
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()

	s.reloadTimer = nil

	s.log.Info("Executing full policy synchronization...")
	start := time.Now()
	ctx := context.Background()

	if err := s.Sync(ctx); err != nil {
		s.failedSyncs.Add(1)
		s.log.Errorf("Failed to perform full policy synchronization: %v", err)
		return
	}

	if err := s.reloader.Reload(); err != nil {
		s.failedSyncs.Add(1)
		s.log.Errorf("Failed to reload policies after full sync: %v", err)
		return
	}

	duration := time.Since(start).Milliseconds()
	s.lastSyncDuration.Store(duration)
	s.lastSyncTime.Store(time.Now())
	s.successfulSyncs.Add(1)
	s.log.Infof("Full policy synchronization completed successfully in %dms", duration)
}

// GetMetrics returns the current status and metrics of the syncer.
func (s *PolicySyncer) GetMetrics() *authv1.PolicySyncStatusResponse {
	s.reloadMu.Lock()
	pending := s.reloadTimer != nil
	s.reloadMu.Unlock()

	lastTime, _ := s.lastSyncTime.Load().(time.Time)
	var lastTimeProto *timestamppb.Timestamp
	if !lastTime.IsZero() {
		lastTimeProto = timestamppb.New(lastTime)
	}

	total := s.totalEvents.Load()
	failed := s.failedSyncs.Load()
	var failureRate float64
	if total > 0 {
		totalSyncs := s.successfulSyncs.Load() + failed
		if totalSyncs > 0 {
			failureRate = float64(failed) / float64(totalSyncs) * 100
		}
	}

	return &authv1.PolicySyncStatusResponse{
		SyncPending:      pending,
		LastSyncTime:     lastTimeProto,
		LastSyncDuration: s.lastSyncDuration.Load(),
		TotalEvents:      total,
		SuccessfulSyncs:  s.successfulSyncs.Load(),
		FailedSyncs:      failed,
		FailureRate:      failureRate,
	}
}

// Sync performs a full, destructive synchronization of policies.
// It fetches all generic rules and applies them, clearing all previous rules.
func (s *PolicySyncer) Sync(ctx context.Context) error {
	s.log.WithContext(ctx).Info("Starting full policy synchronization...")

	// 1. Fetch generic, pre-processed rules from the source of truth.
	resp, err := s.provider.ListAllPolicies(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to fetch policies from provider: %v", err)
		return fmt.Errorf("failed to fetch policies: %w", err)
	}

	s.log.WithContext(ctx).Infof("DIAGNOSIS: Fetched %d access rules ('p' rules) and %d grouping rules ('g' rules) from provider.",
		len(resp.GetAccessRules()), len(resp.GetGroupingRules()))

	// 2. Atomically clear all existing policies by calling ClearPolicies without arguments.
	// This is the correct, definitive way to ensure a clean slate for the new policies.
	if _, err := s.modifier.ClearPolicies(ctx); err != nil {
		s.log.WithContext(ctx).Errorf("Failed to clear all policies during sync: %v", err)
		return err
	}

	// 3. Add all new access rules (translating to 'p' rules for Casbin).
	if len(resp.GetAccessRules()) > 0 {
		s.log.WithContext(ctx).Info("Applying new access policies...")
		for _, rule := range resp.GetAccessRules() {
			spec := authz.RuleSpec{
				Resource: rule.GetObject(),
				Action:   rule.GetAction(),
				Domain:   rule.GetDomain(),
			}
			subject := rule.GetSubject()
			if _, err := s.modifier.AddPermissions(ctx, subject, spec); err != nil {
				s.log.WithContext(ctx).Errorf("Failed to add permission for subject '%s': %v", subject, err)
			}
		}
	}

	// 4. Add all new grouping rules (translating to 'g' rules for Casbin).
	if len(resp.GetGroupingRules()) > 0 {
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
