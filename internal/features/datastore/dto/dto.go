/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"net/http"

	"github.com/origadmin/runtime/errors"
	"origadmin/application/admin/internal/data/enums"

	typespb "origadmin/application/admin/api/v1/services/types"
)

//go:generate abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/data/entity/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="ent,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound      = errors.New(http.StatusNotFound, "http.response.status."+typespb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), "user not found")
	ErrInvalidCaptchaID  = errors.New(http.StatusBadRequest, "http.response.status."+typespb.SystemErrorReason_SYSTEM_ERROR_REASON_INVALID_CAPTCHA_ID.String(), "invalid captcha id")
	ErrInvalidPassword   = errors.New(http.StatusBadRequest, "http.response.status."+typespb.SystemErrorReason_SYSTEM_ERROR_REASON_INVALID_PASSWORD.String(), "invalid password")
	ErrInvalidUsername   = errors.New(http.StatusBadRequest, "http.response.status."+typespb.SystemErrorReason_SYSTEM_ERROR_REASON_INVALID_USERNAME.String(), "invalid username")
	ErrCaptchaIDNotFound = errors.New(http.StatusBadRequest, "http.response.status."+typespb.SystemErrorReason_SYSTEM_ERROR_REASON_CAPTCHA_ID_NOT_FOUND.String(), "captcha id not found")
)

const (
	UserStatusActive = enums.StatusActive
	UserStatusFrozen = enums.StatusFrozen
)

const (
	ResourceStatusEnabled  = enums.StatusEnabled
	ResourceStatusDisabled = enums.StatusDisabled
)
