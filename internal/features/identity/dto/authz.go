/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto

import (
	"context"
)

// AuthzRepo defines the interface for authorization-related data access.
// It focuses on querying the permissions and roles of the current user.
type AuthzRepo interface {
	// GetPermissionKeywordsByUserID retrieves all permission keywords for the current user.
	GetPermissionKeywordsByUserID(ctx context.Context, userID int64) ([]string, error)

	// HasSystemRole checks if the current user has a role with the 'system' type.
	HasSystemRole(ctx context.Context, userID int64) (bool, error)
}
