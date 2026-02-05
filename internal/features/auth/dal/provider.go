/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"github.com/google/wire"
)

// ProviderSet is dal providers for shared database mode.
// Use this when auth and system services share the same database.
var ProviderSet = wire.NewSet(
	NewAuthRepo,
	NewMeRepo,
	NewPolicyDBProvider, // DB implementation for PolicyProvider
	NewCasbinModifier,
)

// ProviderSetWithGRPC is dal providers for separate database mode.
// It provides gRPC-based implementations for the data access interfaces.
// The required gRPC clients are expected to be provided by the client package.
var ProviderSetWithGRPC = wire.NewSet(
	NewAuthRepo,           // AuthRepo might still use the DB for local auth state.
	NewMeGRPCRepo,         // gRPC implementation for MeRepo
	NewPolicyGRPCProvider, // gRPC implementation for PolicyProvider
)
