/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"errors"
	"fmt"
	"net/http"

	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/origadmin/runtime/log"
)

func decodeError(alwaysSucceed bool, code int, err error) (int, *Error) {
	var ierr *Error
	var status int
	if ok := errors.As(err, &ierr); ok {
		status = int(ierr.Code)
	} else if ke := kerr.FromError(err); ke != nil {
		status = int(ke.Code)
		if ke.Reason == "" {
			ke.Reason = "UNKNOWN_REASON"
		}
		ierr = &Error{
			Id:     HTTPErrorPrefix + ke.Reason,
			Code:   ke.Code,
			Detail: ke.Message,
		}
	} else {
		status = http.StatusInternalServerError
		ierr = &Error{
			Id:     HTTPErrorPrefix + "UNKNOWN",
			Code:   http.StatusInternalServerError,
			Detail: fmt.Sprintf("unhandled error: %v", err),
		}
	}
	if code >= 500 {
		log.Debugf("Error: %v, 5xx status: %d", ierr, code)
	}
	if alwaysSucceed {
		status = http.StatusOK
	}
	if !alwaysSucceed && code != status {
		ierr.Code = int32(status)
	}
	if code == 0 {
		code = status
	}
	return code, ierr
}
