/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/origadmin/toolkits/errors/httperr"

	pb "origadmin/application/admin/api/v1/services/basis"
)

const (
	InvalidTokenID = "com.invalid.token"
	// InvalidCaptchaID is the error code for invalid captcha ID
	InvalidCaptchaID = "com.invalid.captcha"
	// InvalidUsernameOrPasswordID is the error code for invalid username or password
	InvalidUsernameOrPasswordID = "com.invalid.username.or.password"
)

var (
	// ErrCaptchaIDNotFound is user not found.
	ErrCaptchaIDNotFound = pb.ErrorBasisErrorReasonCaptchaIdNotFound("captcha id not found")
	// ErrInvalidToken is the error code for invalid token ID
	ErrInvalidToken = pb.ErrorBasisErrorReasonInvalidToken("invalid token")
	// ErrInvalidCaptchaID is the error code for invalid captcha ID
	ErrInvalidCaptchaID = pb.ErrorBasisErrorReasonInvalidCaptchaId("invalid captcha id")
	// ErrInvalidUsername is the error code for invalid username
	ErrInvalidUsername = pb.ErrorBasisErrorReasonInvalidUsername("invalid username")
	// ErrInvalidPassword is the error code for invalid password
	ErrInvalidPassword = pb.ErrorBasisErrorReasonInvalidPassword("invalid password")
)

func ErrorRPC2HTTP(err *errors.Error) *httperr.Error {
	return &httperr.Error{
		ID:     "http.response.status." + err.Reason,
		Code:   err.Code,
		Detail: err.Message,
	}
}
