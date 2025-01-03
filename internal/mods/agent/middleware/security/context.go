/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
)

type skipCtx struct{}

func NewSkipContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, skipCtx{}, true)
}

func IsSkipped(ctx context.Context) bool {
	return ctx.Value(skipCtx{}) != nil
}
