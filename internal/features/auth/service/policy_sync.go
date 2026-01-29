/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/biz" // Import biz package
)

// PolicySyncService is a business logic handler that updates authorization policies based on events.
// It does NOT manage its own subscription or Watermill lifecycle.
type PolicySyncService struct {
	policyModifier authz.PolicyModifier
	reloader       authz.Reloader
	syncer         *biz.PolicySyncer // Add PolicySyncer
	log            *log.Helper
}

// NewPolicySyncService creates a new PolicySyncService.
func NewPolicySyncService(policyModifier authz.PolicyModifier, authorizer *casbin.Authorizer,
	syncer *biz.PolicySyncer, logger log.Logger) *PolicySyncService { // Add syncer parameter
	return &PolicySyncService{
		policyModifier: policyModifier,
		reloader:       authorizer,
		syncer:         syncer, // Initialize syncer
		log:            log.NewHelper(log.With(logger, "module", "auth.service.policy_sync")),
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
	s.log.Infof("Processing UserRoleAssignedEvent: UserID=%s, RoleKeywords=%v", userID, roleKeywords)

	// 1. Remove all existing roles for the user
	if _, err := s.policyModifier.RemoveRoles(msg.Context(), userID); err != nil {
		s.log.Errorf("Failed to remove old roles for UserID=%s: %v", userID, err)
		return err
	}

	// 2. Add the new roles
	if len(roleKeywords) > 0 {
		newRoles := make([]authz.RoleSpec, len(roleKeywords))
		for i, roleKeyword := range roleKeywords {
			// Use "*" as the standard wildcard for the global domain.
			newRoles[i] = authz.RoleSpec{Role: roleKeyword, Domain: "*"}
		}
		if _, err := s.policyModifier.AddRoles(msg.Context(), userID, newRoles...); err != nil {
			s.log.Errorf("Failed to add new roles to user '%s': %v", userID, err)
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

// HandleRolePolicyChanged processes a RolePolicyChangedEvent.
func (s *PolicySyncService) HandleRolePolicyChanged(msg *message.Message) error {
	var event types.RolePolicyChangedEvent
	if err := proto.Unmarshal(msg.Payload, &event); err != nil {
		s.log.Errorf("Failed to unmarshal RolePolicyChangedEvent: %v", err)
		return err
	}
	s.log.WithContext(msg.Context()).Infof("CONFIRM: Received RolePolicyChangedEvent with keywords: %v", event.GetRoleKeywords())

	s.log.Infof("Processing RolePolicyChangedEvent from source '%s' for roles %v", event.GetSource(), event.GetRoleKeywords())

	// 1. Trigger a targeted policy synchronization for the affected roles.
	// If the RoleKeywords list is empty, SyncRoles will fall back to a full sync.
	if err := s.syncer.SyncRoles(msg.Context(), event.GetRoleKeywords()...); err != nil {
		s.log.Errorf("Failed to synchronize policies on role policy change: %v", err)
		return err
	}

	// 2. Reload all policies from the storage adapter into the enforcer.
	if err := s.reloader.Reload(); err != nil {
		s.log.Errorf("Failed to trigger policy reload after sync: %v", err)
		return err
	}

	s.log.Info("Successfully synchronized and reloaded policies due to RolePolicyChangedEvent.")
	return nil
}
