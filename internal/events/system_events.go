/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package events

import (
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Define topic constants for system-related events.
const (
	UserRoleAssignedTopic   = "system.user.role_assigned"
	UserRoleUnassignedTopic = "system.user.role_unassigned"
	// Add other system event topics here as needed, e.g.,
	// UserCreatedTopic = "system.user.created"
	// RoleCreatedTopic = "system.role.created"
)

// UserRoleAssignedEvent represents an event when a user is assigned to one or more roles.
type UserRoleAssignedEvent struct {
	Timestamp int64    `json:"timestamp"`
	UserID    string   `json:"userID"`
	RoleIDs   []string `json:"roleIDs"` // Changed to slice to reflect UpdateUserRolesRequest
	Source    string   `json:"source"`  // Service that triggered the event
}

// NewUserRoleAssignedMessage creates a new Watermill message for a UserRoleAssignedEvent.
func NewUserRoleAssignedMessage(userID string, roleIDs []string, sourceService string) (*message.Message, error) {
	event := UserRoleAssignedEvent{
		Timestamp: time.Now().Unix(),
		UserID:    userID,
		RoleIDs:   roleIDs,
		Source:    sourceService,
	}
	payload, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	return msg, nil
}

// PolicyUpdateEvent (from permission.go) is still relevant for generic policy updates,
// but specific events like UserRoleAssignedEvent provide more context.
// We might deprecate PolicyUpdateEvent in favor of more granular events later.
