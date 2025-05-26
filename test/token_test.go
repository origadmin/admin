/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mock implements the functions, types, and interfaces for the module.
package test

import (
	"context"
	"testing"

	_ "origadmin/application/admin/contrib/consul/config"
	_ "origadmin/application/admin/contrib/consul/registry"
	_ "github.com/origadmin/contrib/database"
	"github.com/origadmin/runtime/bootstrap"
	"github.com/origadmin/runtime/interfaces/security"

	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/loader"
	"origadmin/application/admin/internal/mods/system/dal"
	"origadmin/application/admin/internal/mods/system/server"
)

type data struct {
}

func (d data) QueryRoles(ctx context.Context, subject string) ([]string, error) {
	return []string{
		"role_1",
	}, nil
}

func (d data) QueryPermissions(ctx context.Context, subject string) ([]string, error) {
	return []string{
		"user_1",
	}, nil
}

func TestGenerateToken(t *testing.T) {
	bs, err := loader.LoadBootstrap(&loader.Bootstrap{
		Flags:      bootstrap.Flags{},
		WorkDir:    "",
		ConfigPath: "D:\\workspace\\project\\golang\\origadmin\\backend\\resources\\configs\\config_test.toml",
		Env:        "",
		Daemon:     false,
	})
	if err != nil {
		t.Fatalf("failed to load bootstrap: %v", err)
	}
	dd, cleanup, err := dal.NewData(bs, nil)
	if err != nil {
		t.Fatalf("failed to new dd: %v", err)
	}
	defer cleanup()
	//casbinRepo, err := dal.NewCasbinSourceRepo(dd)
	//if err != nil {
	//	t.Fatalf("failed to new casbin source repo: %v", err)
	//}
	resourceRepo := dal.NewResourceRepo(dd, nil)
	roleRepo := dal.NewRoleRepo(dd, nil)
	userRepo := dal.NewUserRepo(dd, nil)
	basisConfig := loader.NewBasisConfig(bs)
	//v, err := server.NewSystemClient(bs, nil)
	//if err != nil {
	//	t.Fatalf("failed to new system client: %v", err)
	//}
	//auth := system.NewAuthServiceClient(v)
	tokenizer, err := loader.NewTokenizer(bs)
	if err != nil {
		t.Fatalf("failed to new tokenizer: %v", err)
	}
	refreshTokenizer := dal.RefreshTokenizer(tokenizer)
	loginData := &dal.LoginData{
		Tokenizer: refreshTokenizer,
		Resource:  resourceRepo,
		Role:      roleRepo,
		User:      userRepo,
	}
	ctx := context.Background()
	claims, err := loginData.Tokenizer.CreateClaims(ctx, "user_1")
	if err != nil {
		return
	}
	token, err := loginData.Tokenizer.CreateToken(ctx, claims)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}
	t.Logf("token: %s", token)

	v, err := server.NewSystemClient(bs, nil)
	if err != nil {
		t.Fatalf("failed to new system client: %v", err)
	}
	//registerAgent, err := server.NewSystemServiceAgentClient(v, nil)
	//if err != nil {
	//	t.Fatalf("failed to new system service agent client: %v", err)
	//}
	//_ := agent.NewRegisterAgent(registerAgent)
	casbinSourceServiceClient := server.NewCasbinServiceClient(v, nil)
	_ = casbinSourceServiceClient
	//casbinBiz := biz.NewCasbinSourceServiceBiz(casbinRepo, nil)
	//client := service.NewCasbinSourceServiceServerPB(casbinBiz)
	authenticator, err := securityx.NewAuthenticator(bs)
	if err != nil {
		panic(err)
	}
	//adapter := casbin.NewAdapter()
	authorizer, err := securityx.NewAuthorizer(bs, casbin.WithServiceClient(casbinSourceServiceClient))
	if err != nil {
		panic(err)
	}
	bridge := securityx.DefaultBridge()
	bridge.PolicyParser = func(ctx context.Context, claims security.Claims) (security.Policy, error) {
		return security.RegisteredPolicy{
			Subject: "user_1",
			Object:  "/api/v1/sys/users",
			Action:  "GET",
			Domain:  "*",
			//Roles:       roles,
			//Permissions: permissions,
		}, nil
	}
	bridge.Authenticator = authenticator
	bridge.Authorizer = authorizer
	claims2, err := bridge.Authenticator.Authenticate(ctx, token)
	if err != nil {
		t.Fatalf("failed to authenticate: %v", err)
	}
	policy, err := bridge.PolicyParser(ctx, claims2)
	if err != nil {
		t.Fatalf("failed to parse policy: %v", err)
	}

	authorized, err := bridge.Authorizer.AuthorizedWithExtra(ctx, security.DataWithExtra(claims,
		policy, nil))
	if err != nil {
		t.Errorf("failed to authorize: %v", err)
	}
	if !authorized {
		t.Errorf("failed to authorize: %v", err)
	}
	t.Logf("authorized: %v", authorized)
}
