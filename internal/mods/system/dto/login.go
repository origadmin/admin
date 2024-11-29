/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"
)

type LoginRepo interface {
	Login(ctx context.Context, username, password string) (any, error)
}
