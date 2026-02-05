/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/casbin/casbin/v3/persist"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	identityv1 "origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/internal/features/identity/dto"
	"origadmin/application/admin/internal/helpers/debounce"
)

// PolicySyncer is responsible for synchronizing policies from a PolicyProvider to a PolicyModifier.
type PolicySyncer struct {
	provider  dto.PolicyProvider
	modifier  authz.PolicyModifier
	reloader  authz.Reloader
	log       *log.Helper
	debouncer debounce.Executor

	// Metrics
	lastSyncTime     atomic.Value // time.Time
	lastSyncDuration atomic.Int64 // milliseconds
	totalEvents      atomic.Int64
	successfulSyncs  atomic.Int64
	failedSyncs      atomic.Int64
}

// NewPolicySyncer creates a new PolicySyncer.
func NewPolicySyncer(provider dto.PolicyProvider, modifier authz.PolicyModifier, identityorizer authz.Authorizer, watcher persist.Watcher, debouncer debounce.Executor, logger log.Logger) *PolicySyncer {
	logHelper := log.NewHelper(log.With(logger, "module", "identity.biz.policy_syncer"))

	var reloader authz.Reloader

	directReloader, ok := identityorizer.(authz.Reloader)
	if !ok {
		logHelper.Warn("The provided identityorizer does not implement authz.Reloader. Local policy reloading will not be possible.")
		directReloader = noopReloader{}
	}

	if watcher != nil {
		logHelper.Info("Watcher detected. Using watcher-based reloading strategy.")
		reloader = &watcherReloader{
			watcher:  watcher,
			reloader: directReloader,
			log:      logHelper,
		}
	} else {
		logHelper.Info("No watcher detected. Using direct reloading strategy.")
		reloader = directReloader
	}

	return &PolicySyncer{
		provider:  provider,
		modifier:  modifier,
		reloader:  reloader,
		log:       logHelper,
		debouncer: debouncer,
	}
}

// watcherReloader adapts a watcher to the authz.Reloader interface.
type watcherReloader struct {
	watcher  persist.Watcher
	reloader authz.Reloader
	log      *log.Helper
}

// Reload broadcasts an update via the watcher.
func (r *watcherReloader) Reload(force bool) error {
	if !force {
		return nil
	}
	r.log.Info("Propagating policy update via watcher...")
	if err := r.watcher.Update(); err != nil {
		r.log.Errorf("Failed to broadcast casbin update: %v. Falling back to local reload.", err)
		return r.reloader.Reload(true)
	}
	r.log.Info("Successfully broadcasted policy update.")
	return nil
}

// noopReloader is a safe, no-operation implementation of authz.Reloader.
type noopReloader struct{}

func (r noopReloader) Reload(force bool) error { return nil }

// ScheduleSync schedules a delayed full synchronization with debounce logic.
func (s *PolicySyncer) ScheduleSync() {
	s.totalEvents.Add(1)
	s.log.Debugf("Policy sync event received, scheduling execution.")
	s.debouncer.Schedule(s.syncAndReload)
}

// ForceSync triggers an immediate full policy synchronization.
func (s *PolicySyncer) ForceSync(ctx context.Context) error {
	s.log.Info("Force triggering policy sync (manual request)")
	s.debouncer.Cancel()
	s.syncAndReload()
	return nil
}

// syncAndReload performs a full synchronization and then triggers a reload.
func (s *PolicySyncer) syncAndReload() {
	s.log.Info("Executing full policy synchronization and reload...")
	start := time.Now()
	ctx := context.Background()

	if err := s.sync(ctx); err != nil {
		s.failedSyncs.Add(1)
		s.log.Errorf("Failed to perform full policy synchronization: %v", err)
		return
	}

	if err := s.reloader.Reload(true); err != nil {
		s.failedSyncs.Add(1)
		s.log.Errorf("Policy reload failed: %v", err)
		return
	}

	duration := time.Since(start).Milliseconds()
	s.lastSyncDuration.Store(duration)
	s.lastSyncTime.Store(time.Now())
	s.successfulSyncs.Add(1)
	s.log.Infof("Full policy synchronization and reload completed in %dms", duration)
}

// GetMetrics returns the current status and metrics of the syncer.
func (s *PolicySyncer) GetMetrics() *identityv1.PolicySyncStatusResponse {
	pending := s.debouncer.IsPending()
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

	return &identityv1.PolicySyncStatusResponse{
		SyncPending:      pending,
		LastSyncTime:     lastTimeProto,
		LastSyncDuration: s.lastSyncDuration.Load(),
		TotalEvents:      total,
		SuccessfulSyncs:  s.successfulSyncs.Load(),
		FailedSyncs:      failed,
		FailureRate:      failureRate,
	}
}

// sync performs a full, destructive synchronization of policies.
func (s *PolicySyncer) sync(ctx context.Context) error {
	s.log.WithContext(ctx).Info("Starting full policy database synchronization...")

	resp, err := s.provider.ListAllPolicies(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to fetch policies from provider: %v", err)
		return fmt.Errorf("failed to fetch policies: %w", err)
	}

	s.log.WithContext(ctx).Infof("DIAGNOSIS: Fetched %d access rules ('p' rules) and %d grouping rules ('g' rules) from provider.",
		len(resp.GetAccessRules()), len(resp.GetGroupingRules()))

	if _, err := s.modifier.ClearPolicies(ctx); err != nil {
		s.log.WithContext(ctx).Errorf("Failed to clear all policies during sync: %v", err)
		return err
	}

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

	s.log.WithContext(ctx).Info("Policy database synchronization finished successfully.")
	return nil
}
