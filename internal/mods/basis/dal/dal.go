/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(LoginData), "*"),
	NewLoginRepo,
)
