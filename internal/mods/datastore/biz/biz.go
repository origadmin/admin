/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"net/http"

	"github.com/origadmin/runtime/interfaces/pagination"
	"github.com/origadmin/toolkits/errors/httperr"

	typespb "origadmin/application/admin/api/v1/services/types"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = httperr.New("http.response.status."+typespb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), http.StatusNotFound, "user not found")
)

var (
	defaultLimiter = pagination.DefaultLimiter()
)

type UpdateHooker interface {
	UpdateRules()
}
