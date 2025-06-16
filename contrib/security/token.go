/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package security

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/interfaces/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TokenToContext .
func TokenToContext(ctx context.Context, tokenType security.TokenSource, scheme string, token string) context.Context {
	switch tokenType {
	case security.TokenSourceMetadata:
		return injectTokenMetadataContext(ctx, scheme, token)
	case security.TokenSourceMetadataClient:
		return injectTokenMetadataClientContext(ctx, scheme, token)
	case security.TokenSourceMetadataServer:
		return injectTokenMetadataServerContext(ctx, scheme, token)
	case security.TokenSourceHeader:
		return injectHeaderTransportContext(ctx, scheme, token)
	case security.TokenSourceServerHeader:
		return injectServerTransportContext(ctx, scheme, token)
	case security.TokenSourceClientHeader:
		return injectClientTransportContext(ctx, scheme, token)
	case security.TokenSourceContext:
		return security.NewTokenContext(ctx, formatToken(scheme, token))
	default:
		return injectTokenMetadataContext(ctx, scheme, token)
	}
}

func extractTokenFromContext(ctx context.Context, tokenType security.TokenSource) string {
	switch tokenType {
	case security.TokenSourceMetadata:
		return extractTokenMetadataContext(ctx)
	case security.TokenSourceMetadataClient:
		return extractTokenMetadataClientContext(ctx)
	case security.TokenSourceMetadataServer:
		return extractTokenMetadataServerContext(ctx)
	case security.TokenSourceHeader:
		return extractHeaderTransportContext(ctx)
	case security.TokenSourceServerHeader:
		return extractServerTransportContext(ctx)
	case security.TokenSourceClientHeader:
		return extractClientTransportContext(ctx)
	case security.TokenSourceContext:
		return security.TokenFromContext(ctx)
	default:
		return extractTokenMetadataContext(ctx)
	}
}

// TokenFromContext .
func TokenFromContext(ctx context.Context, tokenType security.TokenSource, scheme string) (string, error) {
	val := extractTokenFromContext(ctx, tokenType)
	if val == "" {
		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+scheme)
	}

	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return "", status.Errorf(codes.Unauthenticated, "Bad authorization string")
	}

	if !strings.EqualFold(splits[0], scheme) {
		return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+scheme)
	}

	return splits[1], nil
}

func formatToken(scheme string, tokenStr string) string {
	return fmt.Sprintf("%s %s", scheme, tokenStr)
}

func TokenFromTransportClient(authorize string, scheme string) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		if tr, ok := transport.FromClientContext(ctx); ok {
			token := tr.RequestHeader().Get(authorize)
			splits := strings.SplitN(token, " ", 2)
			if len(splits) > 1 && strings.EqualFold(splits[0], scheme) {
				return splits[1]
			}
		}
		return ""
	}
}

func TokenFromTransportServer(authorize string, scheme string) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		if tr, ok := transport.FromServerContext(ctx); ok {
			token := tr.RequestHeader().Get(authorize)
			splits := strings.SplitN(token, " ", 2)
			if len(splits) > 1 && strings.EqualFold(splits[0], scheme) {
				return splits[1]
			}
		}
		return ""
	}
}
