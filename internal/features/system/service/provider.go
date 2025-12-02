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
	NewResourceServiceServerPB,
	NewResourceServiceHTTPServerPB,
	NewRoleServiceServerPB,
	NewRoleServiceHTTPServerPB,
	NewUserServiceServerPB,
	NewUserServiceHTTPServerPB,
	NewPermissionServiceServerPB,
	NewPermissionServiceHTTPServerPB,
)

// LocalProviderSet is service providers.
var LocalProviderSet = wire.NewSet(
	NewRegisterBridgeServer,
	NewResourceServiceServerPB,
	NewResourceServiceHTTPServerPB,
	NewRoleServiceServerPB,
	NewRoleServiceHTTPServerPB,
	NewUserServiceServerPB,
	NewUserServiceHTTPServerPB,
	NewPermissionServiceServerPB,
	NewPermissionServiceHTTPServerPB,
)

var RemoteProviderSet = wire.NewSet(
	NewRegisterBridgeServer,
	NewResourceServiceBridgeClient,
	//NewResourceServiceBridge,
	NewRoleServiceBridgeClient,
	//NewRoleServiceBridge,
	NewUserServiceBridgeClient,
	//NewUserServiceBridge,
	NewPermissionServiceBridgeClient,
	//NewPermissionServiceBridge,
)
