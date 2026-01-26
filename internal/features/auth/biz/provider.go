/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"github.com/google/wire"
	"origadmin/application/admin/internal/features/auth/client"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAuthUseCase,
	NewMeUseCase,
	NewCaptchaUseCase,
	NewCasbinSynchronizer,
	client.ProviderSet,
)
