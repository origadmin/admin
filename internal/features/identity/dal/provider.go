/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal implements the data access layer for the module.
package dal

import (
	"github.com/google/wire"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(
	NewAuthnRepo,
	NewAuthzRepo,
	NewMeRepo,

	// Provide authz.PolicyReader with either the DB or gRPC implementation.
	// Use only one of the following blocks.
	//
	// For direct database access:
	NewPolicyProvider,
	//
	// For gRPC-based access:
	// wire.Bind(new(authz.PolicyReader), new(*policyGRPCProvider)),
	// NewPolicyGRPCProvider,
)
