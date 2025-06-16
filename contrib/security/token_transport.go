/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package security

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/interfaces/security"
)

func injectHeaderTransportContext(ctx context.Context, scheme string, token string) context.Context {
	if header, ok := transport.FromServerContext(ctx); ok {
		header.RequestHeader().Set(security.HeaderAuthorize, formatToken(scheme, token))
		return transport.NewServerContext(ctx, header)
	}
	if header, ok := transport.FromClientContext(ctx); ok {
		header.RequestHeader().Set(security.HeaderAuthorize, formatToken(scheme, token))
		return transport.NewClientContext(ctx, header)
	}
	return ctx
}
func extractHeaderTransportContext(ctx context.Context) string {
	if header, ok := transport.FromServerContext(ctx); ok {
		return header.RequestHeader().Get(security.HeaderAuthorize)
	}
	if header, ok := transport.FromClientContext(ctx); ok {
		return header.RequestHeader().Get(security.HeaderAuthorize)
	}
	return ""
}

func injectServerTransportContext(ctx context.Context, scheme string, token string) context.Context {
	if header, ok := transport.FromServerContext(ctx); ok {
		header.RequestHeader().Set(security.HeaderAuthorize, formatToken(scheme, token))
		return transport.NewServerContext(ctx, header)
	}
	return ctx
}
func extractServerTransportContext(ctx context.Context) string {
	if header, ok := transport.FromServerContext(ctx); ok {
		return header.RequestHeader().Get(security.HeaderAuthorize)
	}
	return ""
}
func injectClientTransportContext(ctx context.Context, scheme string, token string) context.Context {
	if header, ok := transport.FromClientContext(ctx); ok {
		header.RequestHeader().Set(security.HeaderAuthorize, formatToken(scheme, token))
		return transport.NewClientContext(ctx, header)
	}
	return ctx
}

func extractClientTransportContext(ctx context.Context) string {
	if header, ok := transport.FromClientContext(ctx); ok {
		return header.RequestHeader().Get(security.HeaderAuthorize)
	}
	return ""
}
