/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz implements the functions, types, and interfaces for the module.
package biz

import (
	"github.com/google/wire"

	identitydto "origadmin/application/admin/internal/features/identity/dto"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewResourceUseCase,
	NewRoleUseCase,
	NewUserUseCase,
	NewPermissionUseCase,
	NewViewUseCase,
	NewAuthorizationUseCase,
	wire.Bind(new(identitydto.PolicyProvider), new(*AuthorizationUseCase)),
)
