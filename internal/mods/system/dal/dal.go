/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	//NewAuthRepo,
	//NewLoginRepo,
	//NewPersonalRepo,
	NewMenuRepo,
	NewResourceRepo,
	NewRoleRepo,
	NewUserRepo,
	NewPermissionRepo,
	//NewCasbinSourceRepo,
	//RefreshTokenizer,
)
