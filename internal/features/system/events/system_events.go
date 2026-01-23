/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package events

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"origadmin/application/admin/api/v1/services/types"
)

// Define topic constants for system-related events.
const (
	UserRoleAssignedTopic   = "system.user.role_assigned"
	UserRoleUnassignedTopic = "system.user.role_unassigned"
)

// NewUserRoleAssignedMessage creates a new Watermill message for a UserRoleAssignedEvent using Protobuf serialization.
func NewUserRoleAssignedMessage(userID string, roleIDs []string, sourceService string) (*message.Message, error) {
	// Create the Protobuf event message.
	event := &types.UserRoleAssignedEvent{
		Timestamp: timestamppb.Now(),
		UserId:    userID,
		RoleIds:   roleIDs,
		Source:    sourceService,
	}

	// Marshal the Protobuf message into a binary payload.
	payload, err := proto.Marshal(event)
	if err != nil {
		return nil, err
	}

	// Create a new Watermill message with the binary payload.
	msg := message.NewMessage(watermill.NewUUID(), payload)
	return msg, nil
}
