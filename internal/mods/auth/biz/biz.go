/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"github.com/google/wire"
	"github.com/origadmin/runtime/interfaces/pagination"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAuthServiceBiz,
	NewLoginServiceBiz,
	NewCasbinSourceServiceBiz,
)

var (
	defaultLimiter = pagination.DefaultLimiter()
)

type UpdateHooker interface {
	UpdateRules()
}
