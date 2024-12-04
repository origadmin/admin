/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/common"
)

type (
	CaptchaIDRequest  = pb.CaptchaIDRequest
	CaptchaIDResponse = pb.CaptchaIDResponse
)

type LoginRepo interface {
	Login(ctx context.Context, username, password string) (any, error)
	CaptchaID(ctx context.Context, in *CaptchaIDRequest) (*CaptchaIDResponse, error)
}
