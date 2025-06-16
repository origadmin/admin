/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package security

import (
	kmetadata "github.com/go-kratos/kratos/v2/metadata"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/interfaces/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func injectTokenMetadataContext(ctx context.Context, scheme string, token string) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		// Use pairs to create a new one.
		md = metadata.Pairs()
	}
	md.Set(security.HeaderAuthorize, formatToken(scheme, token))
	return ctx
}

func injectTokenMetadataServerContext(ctx context.Context, scheme string, token string) context.Context {
	md, ok := kmetadata.FromServerContext(ctx)
	if !ok {
		// Use make to create a new one.
		md = make(kmetadata.Metadata)
	}
	md.Set(security.HeaderAuthorize, formatToken(scheme, token))
	return ctx
}

func injectTokenMetadataClientContext(ctx context.Context, scheme string, token string) context.Context {
	md, ok := kmetadata.FromClientContext(ctx)
	if !ok {
		// Use make to create a new one.
		md = make(kmetadata.Metadata)
	}
	md.Set(security.HeaderAuthorize, formatToken(scheme, token))
	return ctx
}

func extractTokenMetadataContext(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md.Get(security.HeaderAuthorize)[0]
	}
	return ""
}

func extractTokenMetadataServerContext(ctx context.Context) string {
	if meta, ok := kmetadata.FromServerContext(ctx); ok {
		return meta.Get(security.HeaderAuthorize)
	}
	return ""
}

func extractTokenMetadataClientContext(ctx context.Context) string {
	if meta, ok := kmetadata.FromClientContext(ctx); ok {
		return meta.Get(security.HeaderAuthorize)
	}
	return ""
}

func ClaimFromTokenTypeContext(ctx context.Context, tokenType security.TokenSource) (security.Claims, error) {
	switch tokenType {
	case security.TokenSourceContext:
		return security.ClaimsFromContext(ctx), nil
	}
	return nil, status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+tokenType.String())
}
