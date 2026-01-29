/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"github.com/google/wire"
	systemv1 "origadmin/application/admin/api/v1/services/system"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewUserService,
	NewRoleService,
	NewResourceService,
	NewPermissionService,
	NewViewService,
	NewSystemService,
	NewAuthorizationService,

	// Bind the concrete implementation to the proto-defined interface.
	wire.Bind(new(systemv1.AuthorizationServiceServer), new(*AuthorizationService)),
)
