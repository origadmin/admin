/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/service"
)

var (
	// ProviderSet is server providers.
	//ProviderSet = wire.NewSet(NewHTTPServerAgent)
	ProviderSet = wire.NewSet(NewHTTPServer)
)

func init() {
	runtime.RegisterService("agent", service.DefaultServiceBuilder)
}
