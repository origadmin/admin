/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto

import (
	"context"

	systemv1 "origadmin/application/admin/api/v1/services/system"
)

// PolicyProvider defines the interface for retrieving identityorization policies.
// This abstraction allows switching between different implementations:
// - Direct database access (when databases are shared)
// - gRPC calls to system service (when databases are separated)
type PolicyProvider interface {
	// ListAllPolicies retrieves all identityorization policies.
	ListAllPolicies(ctx context.Context) (*systemv1.ListAllPoliciesResponse, error)

	// ListPoliciesForRoles retrieves policies for specific role keywords.
	ListPoliciesForRoles(ctx context.Context, roleKeywords ...string) ([]*systemv1.AccessRule, error)
}
