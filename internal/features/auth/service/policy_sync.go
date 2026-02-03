/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/runtime/log"
	watcher "github.com/origadmin/casbin-watcher/v3"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/biz"
)

// PolicySyncService is an event handler that triggers policy synchronization.
// It acts as an adapter between the messaging system and the business logic.
type PolicySyncService struct {
	syncer  *biz.PolicySyncer
	watcher *watcher.Watcher
	log     *log.Helper
}

// NewPolicySyncService creates a new PolicySyncService.
func NewPolicySyncService(syncer *biz.PolicySyncer, w *watcher.Watcher, logger log.Logger) *PolicySyncService {
	return &PolicySyncService{
		syncer:  syncer,
		watcher: w,
		log:     log.NewHelper(log.With(logger, "module", "auth.service.policy_sync")),
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

	s.log.Infof("Received UserRoleAssignedEvent for UserID=%s. Syncing policies to database.", event.GetUserId())

	// Sync policies from business data to casbin_rule table
	ctx := context.Background()
	if err := s.syncer.Sync(ctx); err != nil {
		s.log.Errorf("Failed to sync policies to database: %v", err)
		return err
	}

	// Broadcast global update notification to all auth service instances
	if s.watcher != nil {
		if err := s.watcher.Update(); err != nil {
			s.log.Errorf("Failed to broadcast casbin update: %v", err)
			return err
		}
		s.log.Info("Broadcasted casbin policy update notification to all instances.")
	}

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

	s.log.Infof("Received RolePolicyChangedEvent for roles %v. Syncing policies to database.", event.GetRoleKeywords())

	// Sync policies from business data to casbin_rule table
	ctx := context.Background()
	if err := s.syncer.Sync(ctx); err != nil {
		s.log.Errorf("Failed to sync policies to database: %v", err)
		return err
	}

	// Broadcast global update notification to all auth service instances
	if s.watcher != nil {
		if err := s.watcher.Update(); err != nil {
			s.log.Errorf("Failed to broadcast casbin update: %v", err)
			return err
		}
		s.log.Info("Broadcasted casbin policy update notification to all instances.")
	}

	return nil
}
