/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/internal/configs"
)

func NewAuthenticator(bootstrap *configs.Bootstrap) (security.Authenticator, error) {
	authenticator, err := jwt.NewAuthenticator(bootstrap.GetMiddlewares().GetSecurity())
	if err != nil {
		return nil, err
	}
	return authenticator, nil
}

func NewAuthorizer(bootstrap *configs.Bootstrap) (security.Authorizer, error) {
	authorizer, err := casbin.NewAuthorizer(bootstrap.GetMiddlewares().GetSecurity())
	if err != nil {
		return nil, err
	}
	return authorizer, nil
}
