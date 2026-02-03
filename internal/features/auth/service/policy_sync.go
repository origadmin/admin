/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/biz"
)

// PolicySyncService is an event handler that triggers policy synchronization.
// It acts as an adapter between the messaging system and the business logic.
type PolicySyncService struct {
	syncer *biz.PolicySyncer
	log    *log.Helper
}

// NewPolicySyncService creates a new PolicySyncService.
func NewPolicySyncService(syncer *biz.PolicySyncer, logger log.Logger) *PolicySyncService {
	return &PolicySyncService{
		syncer: syncer,
		log:    log.NewHelper(log.With(logger, "module", "auth.service.policy_sync")),
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
	s.syncer.ScheduleFullSync()

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
	s.syncer.ScheduleFullSync()

	return nil
}
