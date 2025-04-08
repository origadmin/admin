/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"errors"
	"fmt"

	"github.com/origadmin/runtime/context"
	"github.com/origadmin/toolkits/security"
)

type mockAuthenticator struct {
}

func (m mockAuthenticator) Authenticate(ctx context.Context, s string) (security.Claims, error) {
	fmt.Println("token:", s)
	if s == "valid_token" {
		return &security.RegisteredClaims{
			Subject: "valid_token",
		}, nil
	}
	return nil, errors.New("invalid token")
}

func (m mockAuthenticator) AuthenticateContext(ctx context.Context, source security.TokenSource) (security.Claims, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockAuthenticator) DestroyToken(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func (m mockAuthenticator) DestroyRefreshToken(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

type mockAuthorizer struct {
}

func (m mockAuthorizer) Authorized(ctx context.Context, policy security.Policy, object string, action string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockAuthorizer) AuthorizedWithDomain(ctx context.Context, policy security.Policy, object string, action string, domain string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockAuthorizer) AuthorizedWithExtra(ctx context.Context, data security.ExtraData) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type mockData struct {
}

func (m mockData) QueryRoles(ctx context.Context, subject string) ([]string, error) {
	return []string{"admin"}, nil
}

func (m mockData) QueryPermissions(ctx context.Context, subject string) ([]string, error) {
	return []string{"admin"}, nil
}
