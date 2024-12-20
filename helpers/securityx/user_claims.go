/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package securityx

import (
	"github.com/origadmin/toolkits/security"
)

type CasbinUserClaims struct {
	Root    bool
	Subject string
	Method  string
	Path    string
	Scopes  string
	Extra   map[string]string
	Claims  security.Claims
}

func (u CasbinUserClaims) IsRoot() bool {
	return u.Root
}

func (u CasbinUserClaims) GetSubject() string {
	return u.Subject
}

func (u CasbinUserClaims) GetObject() string {
	return u.Path
}

func (u CasbinUserClaims) GetAction() string {
	return u.Method
}

func (u CasbinUserClaims) GetDomain() string {
	return u.Scopes
}

func (u CasbinUserClaims) GetClaims() security.Claims {
	return u.Claims
}

func (u CasbinUserClaims) GetExtra() map[string]string {
	return u.Extra
}

var _ security.UserClaims = (*CasbinUserClaims)(nil)