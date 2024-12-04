/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewLoginAPIService,
	NewLoginAPIServer,
)
