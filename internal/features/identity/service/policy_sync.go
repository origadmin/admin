/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/identity/biz"
)

// PolicySyncService is an event handler that triggers policy synchronization.
// It acts as a simple adapter between the messaging system and the business logic syncer.
type PolicySyncService struct {
	syncer *biz.PolicySyncer
	log    *log.Helper
}

// NewPolicySyncService creates a new PolicySyncService.
func NewPolicySyncService(syncer *biz.PolicySyncer, logger log.Logger) *PolicySyncService {
	return &PolicySyncService{
		syncer: syncer,
		log:    log.NewHelper(log.With(logger, "module", "identity.service.policy_sync")),
	}
}

// HandleUserRoleAssigned receives a notification and schedules a policy synchronization.
func (s *PolicySyncService) HandleUserRoleAssigned(msg *message.Message) error {
	var event types.UserRoleAssignedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal UserRoleAssignedEvent: %v", err)
		return err // Return error to prevent acking the message
	}

	s.log.Infof("Received UserRoleAssignedEvent for UserID=%s. Scheduling a policy sync.", event.GetUserId())
	s.syncer.ScheduleSync()
	return nil
}

// HandleRolePolicyChanged receives a notification and schedules a policy synchronization.
func (s *PolicySyncService) HandleRolePolicyChanged(msg *message.Message) error {
	var event types.RolePolicyChangedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal RolePolicyChangedEvent: %v", err)
		return err // Return error to prevent acking the message
	}

	s.log.Infof("Received RolePolicyChangedEvent for roles %v. Scheduling a policy sync.", event.GetRoleKeywords())
	s.syncer.ScheduleSync()
	return nil
}
