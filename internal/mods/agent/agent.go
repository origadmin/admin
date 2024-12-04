/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"github.com/google/wire"
)

var (
	// ProviderSet is server providers.
	ProviderSet = wire.NewSet(NewHTTPServerAgent)
	//ProviderSet = wire.NewSet(NewHTTPServer)
)
