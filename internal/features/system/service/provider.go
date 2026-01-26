/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"github.com/google/wire"
	systempb "origadmin/application/admin/api/v1/services/system" // Import the new system proto
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewUserService,
	NewRoleService,
	NewResourceService,
	NewPermissionService,
	NewViewService,
	NewSystemService,
	NewCasbinService,
	// Bind the concrete implementation to the new proto-defined interface.
	wire.Bind(new(systempb.PolicyServiceServer), new(*CasbinService)),
)
