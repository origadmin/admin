/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package authSecurity implements the functions, types, and interfaces for the module.
package securityx

import (
	"context"

	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/contrib/security/authn/jwt"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/internal/configs"
)

func NewAuthenticator(bootstrap *configs.Bootstrap, ss ...jwt.Setting) (security.Authenticator, error) {
	tokenizer, err := jwt.NewTokenizer(bootstrap.GetSecurity(), ss...)
	if err != nil {
		return nil, err
	}

	authenticator := jwt.NewAuthenticator(tokenizer)
	return authenticator, nil
}

func NewAuthorizer(bootstrap *configs.Bootstrap, ss ...casbin.Setting) (security.Authorizer, error) {
	authorizer, err := casbin.NewAuthorizer(bootstrap.GetSecurity(), ss...)
	if err != nil {
		return nil, err
	}
	return authorizer, nil
}

type securityMiddleware struct {
	Security security.Security
}

func (s securityMiddleware) Authenticate(ctx context.Context, s2 string) (security.Claims, error) {
	return s.Security.Authenticate(ctx, s2)
}

func (s securityMiddleware) AuthenticateContext(ctx context.Context, tokenType security.TokenType) (security.Claims, error) {
	return s.Security.AuthenticateContext(ctx, tokenType)
}

func (s securityMiddleware) Authorized(ctx context.Context, policy security.Policy, object string, action string) (bool, error) {
	return s.Security.Authorized(ctx, policy, object, action)
}

func (s securityMiddleware) AuthorizedWithDomain(ctx context.Context, policy security.Policy, domain string, object string, action string) (bool, error) {
	return s.Security.AuthorizedWithDomain(ctx, policy, domain, object, action)
}

func (s securityMiddleware) AuthorizedWithExtra(ctx context.Context, data security.ExtraData) (bool, error) {
	return s.Security.AuthorizedWithExtra(ctx, data)
}
