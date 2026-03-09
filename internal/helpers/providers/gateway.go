/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"github.com/google/wire"
)

// ProviderGatewaySet provides gateway-specific dependencies.
var ProviderGatewaySet = wire.NewSet(
	ProviderCommonSet,
	ProvideAuthenticator,
)
