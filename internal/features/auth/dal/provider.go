/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"github.com/google/wire"

	"github.com/origadmin/contrib/security/authz"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(NewAuthRepo, NewCasbinRepo, NewMeRepo, NewAuthPolicyRepository, wire.Bind(new(authz.PolicyManager), new(*AuthPolicyRepository)))
