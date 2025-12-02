/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewRegisterServer,
	NewAuthServiceServerPB,
	NewAuthServiceHTTPServerPB,
	NewCasbinSourceServiceServerPB,
	NewCasbinSourceServiceHTTPServerPB,
	NewLoginServiceServerPB,
	NewLoginServiceHTTPServerPB,
	NewPersonalServiceServerPB,
	NewPersonalServiceHTTPServerPB,
	NewCasbinSourceBiz,
)

// LocalProviderSet is service providers.
var LocalProviderSet = wire.NewSet(
	NewRegisterBridgeServer,
	NewAuthServiceServerPB,
	NewAuthServiceHTTPServerPB,
	NewCasbinSourceServiceServerPB,
	NewCasbinSourceServiceHTTPServerPB,
	NewLoginServiceServerPB,
	NewLoginServiceHTTPServerPB,
	NewPersonalServiceServerPB,
	NewPersonalServiceHTTPServerPB,
	NewCasbinSourceBiz,
)

var RemoteProviderSet = wire.NewSet(
	NewRegisterBridgeServer,
	NewAuthServiceBridgeClient,
	NewCasbinServiceBridgeClient,
	NewLoginServiceBridgeClient,
	NewPersonalServiceBridgeClient,
	NewCasbinSourceClient,
)
