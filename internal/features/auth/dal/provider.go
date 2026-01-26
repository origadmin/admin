/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"github.com/google/wire"
	"github.com/origadmin/contrib/security/authz"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(
	NewAuthRepo,
	NewMeRepo,

	// For runtime, incremental policy updates.
	NewCasbinModifier,
	wire.Bind(new(authz.PolicyModifier), new(*casbinModifier)),

	// For initial, full policy synchronization.
	NewCasbinStorageManager,
	wire.Bind(new(CasbinStorageManager), new(*casbinStorageManager)),
)
