/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto

// PolicyRepo defines the interface for retrieving authorization policies.
// This abstraction allows switching between different implementations:
// - Direct database access (when databases are shared)
// - gRPC calls to system service (when databases are separated)
//
// PolicyRepo is used by PolicySyncer to fetch source data
// for synchronizing the casbin_rule database with URPR data.
type PolicyRepo interface {
}
