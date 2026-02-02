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
type PolicySyncService struct {
	reloader    authz.Reloader
	syncer      *biz.PolicySyncer
	log         *log.Helper
	reloadDelay time.Duration
	reloadMu    sync.Mutex
	reloadTimer *time.Timer
}

// NewPolicySyncService creates a new PolicySyncService.
func NewPolicySyncService(authorizer *casbin.Authorizer,
	syncer *biz.PolicySyncer, logger log.Logger) *PolicySyncService {
	return &PolicySyncService{
		reloader:    authorizer,
		syncer:      syncer,
		log:         log.NewHelper(log.With(logger, "module", "auth.service.policy_sync")),
		reloadDelay: DefaultReloadDelay,
	}
}

// scheduleFullSync schedules a delayed full synchronization with debounce logic.
func (s *PolicySyncService) scheduleFullSync() {
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()

	if s.reloadTimer != nil {
		s.reloadTimer.Stop()
	}

	s.reloadTimer = time.AfterFunc(s.reloadDelay, func() {
		s.doFullSync()
	})

	s.log.Debugf("Full policy sync scheduled in %v", s.reloadDelay)
}

// doFullSync performs a full synchronization of all policies and reloads the enforcer.
func (s *PolicySyncService) doFullSync() {
	s.reloadMu.Lock()
	defer s.reloadMu.Unlock()

	s.reloadTimer = nil

	s.log.Info("Executing full policy synchronization...")
	ctx := context.Background()

	if err := s.syncer.Sync(ctx); err != nil {
		s.log.Errorf("Failed to perform full policy synchronization: %v", err)
		return
	}

	if err := s.reloader.Reload(); err != nil {
		s.log.Errorf("Failed to reload policies after full sync: %v", err)
	}
}

// HandleUserRoleAssigned receives a notification that a user's roles have changed
// and triggers a full policy synchronization.
func (s *PolicySyncService) HandleUserRoleAssigned(msg *message.Message) error {
	var event types.UserRoleAssignedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal UserRoleAssignedEvent: %v", err)
		return err
	}

	s.log.Infof("Received UserRoleAssignedEvent for UserID=%s. Scheduling a full policy sync.", event.GetUserId())
	s.scheduleFullSync()

	return nil
}

// HandleRolePolicyChanged receives a notification that a role's permissions have changed
// and triggers a full policy synchronization.
func (s *PolicySyncService) HandleRolePolicyChanged(msg *message.Message) error {
	var event types.RolePolicyChangedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal RolePolicyChangedEvent: %v", err)
		return err
	}

	s.log.Infof("Received RolePolicyChangedEvent for roles %v. Scheduling a full policy sync.", event.GetRoleKeywords())
	s.scheduleFullSync()

	return nil
}
