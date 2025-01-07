/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package securityx

import (
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/origadmin/runtime/context"
)

const (
	GlobalSecurityUserID = "x-md-global-security-user-id"
)

func AppendUserID(ctx context.Context, userID string) context.Context {
	return metadata.AppendToClientContext(ctx, GlobalSecurityUserID, userID)
}

func GetUserID(ctx context.Context) string {
	if md, ok := metadata.FromServerContext(ctx); ok {
		return md.Get(GlobalSecurityUserID)
	}
	if md, ok := metadata.FromClientContext(ctx); ok {
		return md.Get(GlobalSecurityUserID)
	}
	return ""
}
