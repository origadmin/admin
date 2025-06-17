/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz implements the functions, types, and interfaces for the module.
package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAuthServiceBiz,
	NewLoginServiceBiz,
	NewPersonalServiceBiz,
	NewCasbinSourceServiceBiz,
)
