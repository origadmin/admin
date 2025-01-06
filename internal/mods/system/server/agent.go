/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package server implements the functions, types, and interfaces for the module.
package server

import (
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/origadmin/runtime/context"
)

type agentCtx struct{}

func NewAgentContext(ctx context.Context, metadata metadata.Metadata) context.Context {
	return context.WithValue(ctx, agentCtx{}, metadata)
}

func FromAgentContext(ctx context.Context) (metadata.Metadata, bool) {
	if v, ok := ctx.Value(agentCtx{}).(metadata.Metadata); ok {
		return v, true
	}
	return nil, false
}
