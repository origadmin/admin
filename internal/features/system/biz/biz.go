/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"net/http"

	"github.com/origadmin/runtime/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.New(http.StatusNotFound, "USER_NOT_FOUND", "user not found")
)
