/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"github.com/google/wire"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(
	NewAuthRepo,
	NewMeRepo,

	// For runtime, incremental policy updates.
	NewCasbinModifier,
)
