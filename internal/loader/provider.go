/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewServiceServerRegistrars,
)
