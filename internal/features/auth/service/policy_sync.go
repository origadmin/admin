/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
)

// PolicySyncService is a business logic handler that updates authorization policies based on events.
// It does NOT manage its own subscription or Watermill lifecycle.
type PolicySyncService struct {
	policyModifier authz.PolicyModifier
	reloader       authz.Reloader
	log            *log.Helper
}

// NewPolicySyncService creates a new PolicySyncService.
func NewPolicySyncService(policyModifier authz.PolicyModifier, authorizer authz.Authorizer, logger log.Logger) *PolicySyncService {
	reloader, ok := authorizer.(authz.Reloader)
	if !ok {
		// Fallback for authorizers that don't support reloading (e.g. noop)
		reloader = &noopReloader{}
	}
	return &PolicySyncService{
		policyModifier: policyModifier,
		reloader:       reloader,
		log:            log.NewHelper(log.With(logger, "module", "auth.service.policy_sync")),
	}
}

type noopReloader struct{}

func (n *noopReloader) Reload() error { return nil }

// HandleUserRoleAssigned processes a UserRoleAssignedEvent.
func (s *PolicySyncService) HandleUserRoleAssigned(msg *message.Message) error {
	var event types.UserRoleAssignedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal UserRoleAssignedEvent: %v", err)
		return err
	}

	userID := event.GetUserId()
	roleIDs := event.GetRoleIds()
	s.log.Infof("Processing UserRoleAssignedEvent: UserID=%s, RoleIDs=%v", userID, roleIDs)

	// 1. Remove all existing roles for the user
	if _, err := s.policyModifier.RemoveAllUserRoles(msg.Context(), userID); err != nil {
		s.log.Errorf("Failed to remove old roles for UserID=%s: %v", userID, err)
		return err
	}

	// 2. Add the new roles
	for _, roleID := range roleIDs {
		if _, err := s.policyModifier.AddUserRole(msg.Context(), userID, roleID); err != nil {
			s.log.Errorf("Failed to add role '%s' to user '%s': %v", roleID, userID, err)
			// Consider transaction rollback logic here if needed
			return err
		}
	}

	// 3. Trigger policy reload via the reloader
	if err := s.reloader.Reload(); err != nil {
		s.log.Errorf("Failed to trigger policy reload: %v", err)
		return err
	}

	s.log.Infof("Successfully processed UserRoleAssignedEvent for UserID=%s", userID)
	return nil
}
