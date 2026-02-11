/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"github.com/google/wire"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(
	NewUserService,
	NewRoleService,
	NewResourceService,
	NewPermissionService,
	NewViewService,
	NewSystemService,
	NewPolicyBootstrap,
	NewPolicySyncHandler,
	NewPolicyQueryService,

	// Bind the concrete implementation to the proto-defined interface.
	//wire.Bind(new(systemv1.PolicyQueryServiceServer), new(*PolicyQueryService)),
)
