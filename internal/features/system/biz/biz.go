/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"net/http"

	"github.com/origadmin/runtime/errors" // Changed from httperr

	"origadmin/application/admin/internal/helpers/pagination"

	typespb "origadmin/application/admin/api/v1/services/types"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.New("http.response.status."+typespb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), http.StatusNotFound, "user not found")
)

var (
	defaultLimiter = pagination.DefaultLimiter()
)

type UpdateHooker interface {
	UpdateRules()
}
