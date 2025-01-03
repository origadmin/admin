/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"strings"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/types"
	"github.com/origadmin/runtime"
	msecurity "github.com/origadmin/runtime/agent/middleware/security"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicehttp "github.com/origadmin/runtime/service/http"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/configs"
)

type data struct {
}

func (d data) QueryRoles(ctx context.Context, subject string) ([]string, error) {
	return []string{}, nil
}

func (d data) QueryPermissions(ctx context.Context, subject string) ([]string, error) {
	return []string{}, nil
}

// NewHTTPServerAgent new an HTTP server.
func NewHTTPServerAgent(bootstrap *configs.Bootstrap, registrars []ServerRegisterAgent, l log.KLogger) *service.HTTPServer {
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		panic("no serviceConfig")
	}
	paths := bootstrap.GetSecurity().GetPublicPaths()
	paths = append(DefaultPaths(), paths...)
	ms := []middleware.KMiddleware{
		recovery.Recovery(),
	}
	authenticator, err := securityx.NewAuthenticator(bootstrap)
	if err != nil {
		panic(err)
	}
	adapter := casbin.NewAdapter()
	authorizer, err := securityx.NewAuthorizer(bootstrap, casbin.WithPolicyAdapter(adapter))
	if err != nil {
		panic(err)
	}
	bridge := securityx.SecurityBridge{
		TokenSource:          security.TokenSourceHeader,
		Scheme:               security.SchemeBearer,
		AuthenticationHeader: "Authorization",
		Authenticator:        authenticator,
		Authorizer:           authorizer,
		SkipKey:              "",
		PublicPaths:          nil,
		Skipper: func(path string) bool {
			for _, p := range paths {
				if strings.HasPrefix(path, p) {
					return true
				}
			}
			return false
		},
		IsRoot: func(ctx context.Context, claims security.Claims) bool {
			return claims.GetSubject() == "root" || claims.GetSubject() == "admin"
		},
		Data:        &data{},
		TokenParser: nil,
	}
	bridge.SkipKey = msecurity.MetadataSecuritySkipKey
	ms = append(ms, bridge.BuildMiddleware())
	serviceConfig.Name = types.ZeroOr(serviceConfig.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService(), service.WithHTTP(
		servicehttp.WithServerOptions(http.ErrorEncoder(resp.ResponseErrorEncoder)),
		servicehttp.WithMiddlewares(ms...),
		servicehttp.WithPrefix(runtime.DefaultEnvPrefix),
	))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	for _, registrar := range registrars {
		registrar.HTTPServer(ctx, srv)
	}
	srv.WalkRoute(func(info http.RouteInfo) error {
		log.Infof("Registered HTTP route: %s %s", info.Method, info.Path)
		return nil
	})

	return srv
}

func DefaultPaths() []string {
	return []string{
		system.OperationLoginAPICaptchaID,
		system.OperationLoginAPICaptchaImage,
		system.OperationLoginAPICaptchaResource,
		system.OperationLoginAPICaptchaResources,
		system.OperationLoginAPILogin,
		system.OperationLoginAPIRegister,
	}
}
