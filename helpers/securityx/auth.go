/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package securityx

import (
	"github.com/goexts/generic/settings"
	"github.com/origadmin/toolkits/security"
)

type AuthenticatorSetting = func(tz *authSecurity)

type authSecurity struct {
	security.Authenticator
	security.Authorizer
}

func NewSecurity(authenticator security.Authenticator, authorizer security.Authorizer, ss ...AuthenticatorSetting) security.Security {
	t := settings.Apply(&authSecurity{
		Authenticator: authenticator,
		Authorizer:    authorizer,
	}, ss)
	return t
}

var _ security.Security = (*authSecurity)(nil)
