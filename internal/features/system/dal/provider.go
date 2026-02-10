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
	NewUserRepo,
	NewRoleRepo,
	NewResourceRepo,
	NewPermissionRepo,
	NewViewRepo,
	NewPolicyQueryRepo,
	NewCasbinModifier,
)
