/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"net/http"

	"github.com/origadmin/toolkits/errors/httperr"

	pb "origadmin/application/admin/api/v1/services/system"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = httperr.New("http.response.status."+pb.SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND.String(), http.StatusNotFound, "user not found")
)
