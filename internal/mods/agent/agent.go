/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"github.com/google/wire"
	"github.com/origadmin/runtime/service"

	basisserver "origadmin/application/admin/internal/mods/basis/server"
	systemserver "origadmin/application/admin/internal/mods/system/server"
)

var ProviderSet = wire.NewSet(
	NewRegisterAgent,
	NewHTTPServerAgent,
)

type ServerRegisterAgent service.ServerRegister

func NewRegisterAgent(s1 *systemserver.RegisterAgent, s2 *basisserver.RegisterAgent) []ServerRegisterAgent {
	return []ServerRegisterAgent{
		s1, s2,
	}
}
