/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package idutil

import "fmt"

// Constants for ID format consistency
const (
	// UserPrefix is the prefix for user IDs in Casbin and messages
	// IMPORTANT: This prefix must be consistent across all components:
	// - Principal ID (auth/auth.go:63)
	// - Casbin g rules (authorization.go:88)
	// - UserRoleAssignedEvent messages (user.go:62)
	//
	// Current implementation uses empty prefix (plain ID string)
	UserPrefix = ""
)

// FormatUserID formats a user ID with the standard prefix
// This ensures consistency across all components that work with user IDs
func FormatUserID(userID int64) string {
	if UserPrefix == "" {
		return fmt.Sprintf("%d", userID)
	}
	return fmt.Sprintf("%s%d", UserPrefix, userID)
}
