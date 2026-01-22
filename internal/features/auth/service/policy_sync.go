/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/events"
)

// PolicySyncService is a business logic handler that updates authorization policies based on events.
// It does NOT manage its own subscription or Watermill lifecycle.
type PolicySyncService struct {
	policyManager authz.PolicyManager
	log           *log.Helper
}

// NewPolicySyncService creates a new PolicySyncService.
// It only needs a PolicyManager and a Logger.
func NewPolicySyncService(pm authz.PolicyManager, logger log.Logger) *PolicySyncService {
	return &PolicySyncService{
		policyManager: pm,
		log:           log.NewHelper(log.With(logger, "module", "auth.service.policy_sync")),
	}
}

// HandleUserRoleAssigned processes a UserRoleAssignedEvent.
// This method is designed to be registered directly with the watermill.Server.
func (s *PolicySyncService) HandleUserRoleAssigned(msg *message.Message) error {
	var event events.UserRoleAssignedEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal UserRoleAssignedEvent: %v", err)
		return err // Return error to Watermill for potential retry/DLQ
	}

	s.log.Infof("Processing UserRoleAssignedEvent: UserID=%s, RoleIDs=%v", event.UserID, event.RoleIDs)

	// Here you might want to clear existing roles for the user before adding new ones,
	// depending on whether the event represents a full sync or an incremental update.
	// For this implementation, we assume an incremental add.

	for _, roleID := range event.RoleIDs {
		added, err := s.policyManager.AddGroupingPolicy(context.Background(), event.UserID, roleID)
		if err != nil {
			s.log.Errorf("Failed to add grouping policy for UserID=%s, RoleID=%s: %v", event.UserID, roleID, err)
			return err // Return error to Watermill
		}

		if added {
			s.log.Infof("Successfully added grouping policy for UserID=%s, RoleID=%s", event.UserID, roleID)
		} else {
			s.log.Infof("Grouping policy for UserID=%s, RoleID=%s already exists.", event.UserID, roleID)
		}
	}

	return nil // Return nil for successful processing
}
