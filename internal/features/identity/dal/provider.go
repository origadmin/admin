/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal implements the data access layer for the module.
package dal

import (
	"github.com/google/wire"
	"github.com/origadmin/contrib/security/authz"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(
	NewAuthRepo,
	NewMeRepo,

	// Provide authz.PolicyReader with either the DB or gRPC implementation.
	// Use only one of the following blocks.
	//
	// For direct database access:
	wire.Bind(new(authz.PolicyReader), new(*policyDBProvider)),
	NewPolicyDBProvider,
	//
	// For gRPC-based access:
	// wire.Bind(new(authz.PolicyReader), new(*policyGRPCProvider)),
	// NewPolicyGRPCProvider,
)
