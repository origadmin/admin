/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package securityx

import (
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/contrib/security/authn/jwt"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/internal/configs"
)

func NewAuthenticator(bootstrap *configs.Bootstrap) (security.Authenticator, error) {
	authenticator, err := jwt.NewAuthenticator(bootstrap.GetSecurity())
	if err != nil {
		return nil, err
	}
	return authenticator, nil
}

func NewAuthorizer(bootstrap *configs.Bootstrap, roleQuery security.RoleQuery) (security.Authorizer, error) {
	authorizer, err := casbin.NewAuthorizer(bootstrap.GetSecurity(), casbin.WithRoleQuery(roleQuery))
	if err != nil {
		return nil, err
	}
	return authorizer, nil
}
