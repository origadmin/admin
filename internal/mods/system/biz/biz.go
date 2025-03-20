/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"net/http"

	"github.com/google/wire"
	"github.com/origadmin/toolkits/errors/httperr"
	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAuthServiceBiz,
	NewLoginServiceBiz,
	NewPersonalServiceBiz,
	NewResourceServiceBiz,
	NewRoleServiceBiz,
	NewUserServiceBiz,
	NewPermissionServiceBiz,
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), http.StatusNotFound, "user not found")
)

var (
	defaultLimiter = pagination.DefaultLimiter()
)

type UpdateHooker interface {
	UpdateRules()
}
