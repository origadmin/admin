/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz implements the data access layer for the module.
package biz

import (
	"github.com/google/wire"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(
	NewResourceUseCase,
	NewRoleUseCase,
	NewUserUseCase,
	NewPermissionUseCase,
	NewViewUseCase,
	NewAuthorizationUseCase,
	NewPolicyUseCase,
)
