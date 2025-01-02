/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package authSecurity implements the functions, types, and interfaces for the module.
package securityx

import (
	"github.com/origadmin/toolkits/security"
)

type CasbinPolicy struct {
	Subject     string
	Method      string
	Path        string
	Scopes      string
	Roles       []string
	Permissions []string
}

func (obj CasbinPolicy) GetRoles() []string {
	return obj.Roles
}

func (obj CasbinPolicy) GetPermissions() []string {
	return obj.Permissions
}

func (obj CasbinPolicy) GetSubject() string {
	return obj.Subject
}

func (obj CasbinPolicy) GetObject() string {
	return obj.Path
}

func (obj CasbinPolicy) GetAction() string {
	return obj.Method
}

func (obj CasbinPolicy) GetDomain() string {
	return obj.Scopes
}

var _ security.Policy = (*CasbinPolicy)(nil)
