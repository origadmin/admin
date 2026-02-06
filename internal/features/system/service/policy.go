/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/biz"
)

// PolicyService is an event handler that triggers policy synchronization.
// It acts as a simple adapter between the messaging system and the business logic syncer.
//
// Policy Update Process:
// 1. User permission tables are updated (users, roles, permissions, resources)
// 2. Update notification is sent via message queue (NATS)
// 3. PolicyService receives the notification
// 4. PolicyService schedules a debounced policy sync
// 5. PolicySyncer fetches data from URPR and writes to casbin_rule
// 6. Watcher notifies other instances of policy changes
type PolicyService struct {
	syncer *biz.PolicyUseCase
	log    *log.Helper
}

// NewPolicyService creates a new PolicyService.
func NewPolicyService(syncer *biz.PolicyUseCase, logger log.Logger) *PolicyService {
	return &PolicyService{
		syncer: syncer,
		log:    log.NewHelper(log.With(logger, "module", "system.service.policy_sync")),
	}
}

// HandleUserRoleAssigned receives a notification and schedules a policy synchronization.
func (s *PolicyService) HandleUserRoleAssigned(msg *message.Message) error {
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
func (s *PolicyService) HandleRolePolicyChanged(msg *message.Message) error {
	var event types.RolePolicyChangedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal RolePolicyChangedEvent: %v", err)
		return err // Return error to prevent acking the message
	}

	s.log.Infof("Received RolePolicyChangedEvent for roles %v. Scheduling a policy sync.", event.GetRoleKeywords())
	s.syncer.ScheduleSync()
	return nil
}
