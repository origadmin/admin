/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/biz" // Import biz package
)

const (
	DefaultReloadDelay = 2 * time.Second
)

// PolicySyncService is a business logic handler that updates authorization policies based on events.
// It does NOT manage its own subscription or Watermill lifecycle.
type PolicySyncService struct {
	policyModifier authz.PolicyModifier
	reloader       authz.Reloader
	syncer         *biz.PolicySyncer
	log            *log.Helper
	reloadDelay    time.Duration
	reloadMu       sync.Mutex
	reloadTimer    *time.Timer
}

// NewPolicySyncService creates a new PolicySyncService.
func NewPolicySyncService(policyModifier authz.PolicyModifier, authorizer *casbin.Authorizer,
	syncer *biz.PolicySyncer, logger log.Logger) *PolicySyncService {
	// Use a longer delay for UI operations (user may take time to complete multiple actions)
	return &PolicySyncService{
		policyModifier: policyModifier,
		reloader:       authorizer,
		syncer:         syncer,
		log:            log.NewHelper(log.With(logger, "module", "auth.service.policy_sync")),
		reloadDelay:    DefaultReloadDelay,
	}
}

// scheduleReload schedules a delayed Reload call with debounce logic.
// Multiple rapid calls will be coalesced into a single Reload.
func (s *PolicySyncService) scheduleReload() {
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()

	// Stop existing timer if any
	if s.reloadTimer != nil {
		s.reloadTimer.Stop()
	}

	// Schedule new Reload after delay
	s.reloadTimer = time.AfterFunc(s.reloadDelay, func() {
		s.doReload()
	})

	s.log.Debugf("Reload scheduled in %v", s.reloadDelay)
}

// doReload actually executes the Reload and cleans up.
func (s *PolicySyncService) doReload() {
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()

	s.reloadTimer = nil // Reset timer

	s.log.Info("Executing Reload (full policy sync)")

	// Use a background context as the original request context might be cancelled
	ctx := context.Background()

	// Perform full synchronization from the source of truth
	if err := s.syncer.Sync(ctx); err != nil {
		s.log.Errorf("Failed to synchronize policies: %v", err)
		return
	}

	// Reload the enforcer to apply changes
	if err := s.reloader.Reload(); err != nil {
		s.log.Errorf("Failed to execute Reload: %v", err)
	}
}

// HandleUserRoleAssigned processes a UserRoleAssignedEvent.
func (s *PolicySyncService) HandleUserRoleAssigned(msg *message.Message) error {
	var event types.UserRoleAssignedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal UserRoleAssignedEvent: %v", err)
		return err
	}

	userID := event.GetUserId()
	roleKeywords := event.GetRoleKeywords()
	s.log.Infof("Received UserRoleAssignedEvent: UserID=%s, RoleKeywords=%v. Scheduling reload.", userID, roleKeywords)

	// Explicitly remove roles for this user immediately.
	// This handles the case where all roles are revoked (RoleKeywords is empty).
	// If we rely solely on Sync(), it might skip this user if they no longer have any roles in the DB,
	// leaving stale 'g' rules in Casbin.
	if _, err := s.policyModifier.RemoveRoles(msg.Context(), userID); err != nil {
		s.log.Errorf("Failed to remove old roles for UserID=%s: %v", userID, err)
		// We continue to schedule reload even if remove fails, hoping full sync might fix it.
	}

	// Schedule a full reload to ensure consistency and handle any other changes.
	s.scheduleReload()

	return nil
}

// HandleRolePolicyChanged processes a RolePolicyChangedEvent.
func (s *PolicySyncService) HandleRolePolicyChanged(msg *message.Message) error {
	var event types.RolePolicyChangedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal RolePolicyChangedEvent: %v", err)
		return err
	}
	s.log.WithContext(msg.Context()).Infof("CONFIRM: Received RolePolicyChangedEvent with keywords: %v", event.GetRoleKeywords())

	s.log.Infof("Received RolePolicyChangedEvent from source '%s' for roles %v. Scheduling reload.", event.GetSource(), event.GetRoleKeywords())

	// Don't remove permissions immediately - let the doReload's Sync() handle all changes.
	// Immediate removal causes issues during role updates (permissions are deleted before reload).
	// The full sync will correctly remove stale permissions and add new ones.

	// Schedule delayed Reload. The doReload function will handle the full sync.
	s.scheduleReload()

	return nil
}
