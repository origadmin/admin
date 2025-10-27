/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package securityx

import (
	"github.com/goexts/generic/configure"
	"github.com/origadmin/runtime/interfaces/security"
)

type AuthenticatorSetting = func(tz *authSecurity)

type authSecurity struct {
	security.Authenticator
	security.Authorizer
}

func NewSecurity(authenticator security.Authenticator, authorizer security.Authorizer, ss ...AuthenticatorSetting) security.Security {
	return configure.Apply(&authSecurity{
		Authenticator: authenticator,
		Authorizer:    authorizer,
	}, ss)
}

var _ security.Security = (*authSecurity)(nil)
