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

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	identityv1 "origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/debounce"
)

// PolicyType defines the type of policy.
type PolicyType int

const (
	PolicyTypeUnspecified PolicyType = iota
	PolicyTypeAccess
	PolicyTypeGrouping
)

// PolicyEffect defines the effect of the policy.
type PolicyEffect int

const (
	PolicyEffectAllow PolicyEffect = iota
	PolicyEffectDeny
)

// Policy is the core entity for authorization policy.
type Policy struct {
	ID         string
	Type       PolicyType
	Effect     PolicyEffect
	Subject    string
	Domain     string
	Resource   string
	Action     string
	Attributes map[string]string
	Condition  string
}

// PolicyOption defines the functional option for querying policies.
type PolicyOption func(*PolicyFilter)

// PolicyFilter is the value object for querying policies.
type PolicyFilter struct {
	Type       PolicyType
	Subjects   []string
	Domain     string
	Resource   string
	Action     string
	Attributes map[string]string
	Page       int
	PageSize   int
}

// WithSubjects sets the subjects filter.
func WithSubjects(subjects ...string) PolicyOption {
	return func(f *PolicyFilter) {
		f.Subjects = subjects
	}
}

// WithDomain sets the domain filter.
func WithDomain(domain string) PolicyOption {
	return func(f *PolicyFilter) {
		f.Domain = domain
	}
}

// WithResource sets the resource filter.
func WithResource(resource string) PolicyOption {
	return func(f *PolicyFilter) {
		f.Resource = resource
	}
}

// WithAction sets the action filter.
func WithAction(action string) PolicyOption {
	return func(f *PolicyFilter) {
		f.Action = action
	}
}

// WithAttributes sets the attributes filter.
func WithAttributes(attributes map[string]string) PolicyOption {
	return func(f *PolicyFilter) {
		f.Attributes = attributes
	}
}

// WithPagination sets the pagination options.
func WithPagination(page, pageSize int) PolicyOption {
	return func(f *PolicyFilter) {
		f.Page = page
		f.PageSize = pageSize
	}
}

// PolicySyncUseCase is responsible for synchronizing policies from a PolicyProvider to a PolicyModifier.
type PolicySyncUseCase struct {
	repo      dto.PolicyRepo
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

// NewPolicySyncUseCase creates a new PolicySyncUseCase.
func NewPolicySyncUseCase(repo dto.PolicyRepo, modifier authz.PolicyModifier, authorizer authz.Authorizer, watcher persist.Watcher, debouncer debounce.Executor, logger log.Logger) *PolicySyncUseCase {
	logHelper := log.NewHelper(log.With(logger, "module", "system.biz.policy_sync"))

	var reloader authz.Reloader

	directReloader, ok := authorizer.(authz.Reloader)
	if !ok {
		logHelper.Warn("The provided authorizer does not implement authz.Reloader. Local policy reloading will not be possible.")
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

	return &PolicySyncUseCase{
		repo:      repo,
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
func (s *PolicySyncUseCase) ScheduleSync() {
	s.totalEvents.Add(1)
	s.log.Debugf("Policy sync event received, scheduling execution.")
	s.debouncer.Schedule(s.syncAndReload)
}

// SyncNow triggers an immediate policy synchronization without debounce delay.
// Use this for:
// - Application startup initialization
// - Manual trigger from admin panel
func (s *PolicySyncUseCase) SyncNow(ctx context.Context) error {
	s.log.Info("Executing immediate policy sync")
	s.debouncer.Cancel()
	s.syncAndReload()
	return nil
}

// syncAndReload performs a full synchronization and then triggers a reload.
func (s *PolicySyncUseCase) syncAndReload() {
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
func (s *PolicySyncUseCase) GetMetrics() *identityv1.PolicySyncStatusResponse {
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
func (s *PolicySyncUseCase) sync(ctx context.Context) error {
	s.log.WithContext(ctx).Info("Starting full policy database synchronization...")

	// Fetch all policies from the repository (URPR tables)
	rolePerms, userRoles, err := s.repo.ListUserRolePermissions(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to fetch policies from provider: %v", err)
		return fmt.Errorf("failed to fetch policies: %w", err)
	}

	s.log.WithContext(ctx).Infof("DIAGNOSIS: Fetched %d access rules ('p' rules) and %d grouping rules ('g' rules) from provider.",
		len(rolePerms), len(userRoles))

	// Clear existing policies in Casbin storage
	if _, err := s.modifier.ClearPolicies(ctx); err != nil {
		s.log.WithContext(ctx).Errorf("Failed to clear all policies during sync: %v", err)
		return err
	}

	// Apply new Access Policies (p-rules)
	if len(rolePerms) > 0 {
		s.log.WithContext(ctx).Info("Applying new access policies...")
		var policies []*authzv1.PolicySpec
		for _, rp := range rolePerms {
			if rp.Role == nil || rp.Permission == nil {
				continue
			}
			for _, resource := range rp.Permission.Resources {
				if resource == nil {
					continue
				}
				// Construct PolicySpec for access rule
				spec := &authzv1.PolicySpec{
					Type:      "p",
					Subject:   rp.Role.Keyword,
					Domain:    dto.StrPtr("*"),
					Resources: []string{resource.Operation},
					Actions:   []string{"ANY"},
				}
				policies = append(policies, spec)
			}
		}
		if _, err := s.modifier.AddPolicies(ctx, policies...); err != nil {
			s.log.WithContext(ctx).Errorf("Failed to add access policies: %v", err)
		}
	}

	// Apply new Grouping Policies (g-rules)
	if len(userRoles) > 0 {
		s.log.WithContext(ctx).Info("Applying new grouping policies...")
		var policies []*authzv1.PolicySpec
		for _, ur := range userRoles {
			if ur.User == nil || ur.Role == nil {
				continue
			}
			// Construct PolicySpec for grouping rule
			user := fmt.Sprintf("%d", ur.User.Id)
			spec := &authzv1.PolicySpec{
				Type:    "g",
				Subject: user,
				Domain:  dto.StrPtr("*"),
				Roles:   []string{ur.Role.Keyword},
			}
			policies = append(policies, spec)
		}
		if _, err := s.modifier.AddPolicies(ctx, policies...); err != nil {
			s.log.WithContext(ctx).Errorf("Failed to add grouping policies: %v", err)
		}
	}

	s.log.WithContext(ctx).Info("Policy database synchronization finished successfully.")
	return nil
}
