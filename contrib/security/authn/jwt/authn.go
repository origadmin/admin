/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package jwt implements the functions, types, and interfaces for the module.
package jwt

import (
	"context"

	"github.com/goexts/generic/settings"
	"github.com/origadmin/runtime/interfaces/security"
	"github.com/origadmin/runtime/interfaces/security/token"

	contribsecurity "origadmin/application/admin/contrib/security"
)

type Authenticator struct {
	Tokenizer security.Tokenizer
	Cache     token.CacheStorage
	Scheme    security.Scheme
}

func (obj Authenticator) Authenticate(ctx context.Context, s string) (security.Claims, error) {
	claims, err := obj.Tokenizer.ParseClaims(ctx, s)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (obj Authenticator) AuthenticateContext(ctx context.Context, tokenType security.TokenSource) (security.Claims, error) {
	token, err := contribsecurity.TokenFromContext(ctx, tokenType, obj.Scheme.String())
	if err != nil {
		return nil, err
	}
	return obj.Authenticate(ctx, token)
}

func (obj Authenticator) DestroyToken(ctx context.Context, tokenStr string) error {
	return obj.Cache.Remove(ctx, obj.key(token.CacheAccess, tokenStr))
}

func (obj Authenticator) DestroyRefreshToken(ctx context.Context, tokenStr string) error {
	return obj.Cache.Remove(ctx, obj.key(token.CacheRefresh, tokenStr))
}

func (obj Authenticator) key(ns, token string) string {
	return ns + ":" + token
}

type AuthenticatorSetting = func(*Authenticator)

func NewAuthenticator(tokenizer security.Tokenizer, ss ...AuthenticatorSetting) security.Authenticator {
	return settings.Apply(&Authenticator{
		Tokenizer: tokenizer,
		Cache:     token.New(),
		Scheme:    security.SchemeBearer,
	}, ss)
}

var _ security.Authenticator = (*Authenticator)(nil)
