/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewUserService,
	NewRoleService,
	NewResourceService,
	NewPermissionService,
	NewViewService,
	NewSystemService,
)
