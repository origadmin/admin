/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal implements the functions, types, and interfaces for the module.
package dal

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	//NewData,
	NewAuthRepo,
	NewLoginRepo,
	NewCasbinSourceRepo,
	NewPersonalRepo,
	RefreshTokenizer,
)
