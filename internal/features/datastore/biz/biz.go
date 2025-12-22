/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"github.com/origadmin/runtime/errors"
	typespb "origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/helpers/pagination"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.New(50001, typespb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), "user not found")
)

var (
	defaultLimiter = pagination.PageLimiter{}
)

type UpdateHooker interface {
	UpdateRules()
}
