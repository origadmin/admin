/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mock implements the functions, types, and interfaces for the module.
package test

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/origadmin/runtime"
	configv1 "github.com/origadmin/runtime/api/gen/go/config/v1"
	"github.com/origadmin/runtime/interfaces/security"
	"github.com/origadmin/toolkits/codec/toml"

	pb "origadmin/application/admin/api/v1/services/auth"
	_ "origadmin/application/admin/contrib/consul/config"
	_ "origadmin/application/admin/contrib/consul/registry"
	_ "origadmin/application/admin/contrib/database"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/loader"
	"origadmin/application/admin/internal/mods/auth/dal"
	"origadmin/application/admin/internal/mods/auth/service"
)

type mockData struct {
}

func (d mockData) QueryRoles(ctx context.Context, subject string) ([]string, error) {
	return []string{
		"role_1",
	}, nil
}

func (d mockData) QueryPermissions(ctx context.Context, subject string) ([]string, error) {
	return []string{
		"user_1",
	}, nil
}

func init() {
	encoding.RegisterCodec(toml.Codec)
}

func TestGenerateToken(t *testing.T) {
	sourceConfig := &configv1.SourceConfig{
		Types: []string{"file"},
		File: &configv1.SourceConfig_File{
			Path: "..\\resources\\configs\\config_test.toml",
		},
	}
	bootstrap, err := loader.LoadBootstrap(sourceConfig)
	if err != nil {
		t.Fatalf("failed to load bootstrap: %v", err)
	}
	r := runtime.Global()
	dataData, cleanup, err := data.NewData(r, bootstrap)
	if err != nil {
		return
	}
	defer cleanup()
	authRepo := dal.NewAuthRepo(r, dataData)
	//authServiceBiz := biz.NewAuthServiceBiz(r, authRepo)
	//authServiceServer := service.NewAuthServiceServerPB(authServiceBiz)
	//casbinSourceRepo, err := dal.NewCasbinSourceRepo(dataData)
	//if err != nil {
	//	cleanup()
	//	return
	//}
	//casbinSourceServiceBiz := biz.NewCasbinSourceServiceBiz(r, casbinSourceRepo)
	//casbinSourceServiceServer := service.NewCasbinSourceServiceServerPB(casbinSourceServiceBiz)
	//tokenizer, err := data.NewTokenizer(bootstrap)
	//if err != nil {
	//	cleanup()
	//	return
	//}
	//refreshTokenizer := dal.RefreshTokenizer(tokenizer)
	//loginData := data.NewLoginData(bootstrap, refreshTokenizer)
	//loginRepo := dal.NewLoginRepo(dataData, loginData)
	//loginServiceBiz := biz.NewLoginServiceBiz(r, loginRepo)
	//loginServiceServer := service.NewLoginServiceServerPB(loginServiceBiz)
	//personalRepo := dal.NewPersonalRepo(r, dataData)
	//personalServiceBiz := biz.NewPersonalServiceBiz(r, personalRepo)
	//personalServiceServer := service.NewPersonalServiceServerPB(r, personalServiceBiz)
	//registerServer := service.NewRegisterServer(authServiceServer, casbinSourceServiceServer, loginServiceServer, personalServiceServer)
	//v := server.NewAuthServer(r, bootstrap, registerServer)
	authenticator, err := securityx.NewAuthenticator(bootstrap)
	if err != nil {
		panic(err)
	}
	clients := loader.NewProxyGRPCClients(r, bootstrap)
	ruleSource := service.NewCasbinSourceClient(r, clients)
	opts := []casbin.AuthorizerOption{
		casbin.WithSource(ruleSource),
	}

	authorizer, err := securityx.NewAuthorizer(bootstrap, opts...)
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
		}, nil
	}
	bridge.Authenticator = authenticator
	bridge.Authorizer = authorizer
	ctx := context.Background()
	token, err := authRepo.CreateToken(ctx, &pb.CreateTokenRequest{
		Data: &pb.CreateTokenRequest_Data{
			UserId: "user_1",
		},
	})
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}
	claims, err := bridge.Authenticator.Authenticate(ctx, token.GetToken())
	if err != nil {
		t.Fatalf("failed to authenticate: %v", err)
	}
	policy, err := bridge.PolicyParser(ctx, claims)
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
