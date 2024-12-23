/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"strings"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/types"
	"github.com/origadmin/runtime"
	msecurity "github.com/origadmin/runtime/agent/middleware/security"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicehttp "github.com/origadmin/runtime/service/http"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/api/v1/services/basis"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/configs"
)

type HTTPRegistrar interface {
	HTTP(server *service.HTTPServer)
}

// NewHTTPServerAgent new an HTTP server.
func NewHTTPServerAgent(bootstrap *configs.Bootstrap, registrars []HTTPRegistrar, l log.KLogger) *service.HTTPServer {
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
	authorizer, err := securityx.NewAuthorizer(bootstrap)
	if err != nil {
		panic(err)
	}

	options := []msecurity.OptionSetting{
		msecurity.WithAuthorizer(authorizer),
		msecurity.WithAuthenticator(authenticator),
		func(option *msecurity.Option) {
			option.Skipper = func(path string) bool {
				log.Debugf("Checking if path '%s' should be skipped", path)
				for _, p := range paths {
					if strings.HasPrefix(path, p) {
						log.Debugf("Path '%s' starts with '%s', skipping", path, p)
						return true
					}
				}
				log.Debugf("Path '%s' does not match any public path, not skipping", path)
				return false
			}
			option.Parser = func(ctx context.Context, id string) (security.UserClaims, error) {
				req, ok := http.RequestFromServerContext(ctx)
				if !ok {
					return nil, errors.New("no request in context")
				}
				c := msecurity.ClaimsFromContext(ctx)
				if c == nil {
					return nil, errors.New("no token in context")
				}
				return &securityx.CasbinUserClaims{
					Claims:  c,
					Subject: c.GetSubject(),
					Method:  req.Method,
					Path:    req.URL.Path,
					Scopes:  c.GetIssuer(),
					Root:    c.GetSubject() == "admin",
				}, nil
			}
		},
	}
	s, ok := msecurity.SkipperServer(bootstrap.GetSecurity(), options...)
	log.Debugf("Skipper: %v", ok)
	if ok {
		ms = append(ms, s)
	}

	n, err := msecurity.NewAuthN(bootstrap.GetSecurity(), options...)
	if err != nil {
		panic(err)
	}
	z, err := msecurity.NewAuthZ(bootstrap.GetSecurity(), options...)
	if err != nil {
		panic(err)
	}
	ms = append(ms, n, z)
	serviceConfig.Name = types.ZeroOr(serviceConfig.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService(), service.WithHTTP(
		servicehttp.WithServerOptions(http.ErrorEncoder(resp.ResponseErrorEncoder)),
		servicehttp.WithMiddlewares(ms...),
		servicehttp.WithPrefix(runtime.DefaultEnvPrefix),
	))
	if err != nil {
		panic(err)
	}
	for _, registrar := range registrars {
		registrar.HTTP(srv)
	}
	srv.WalkRoute(func(info http.RouteInfo) error {
		log.Infof("Registered HTTP route: %s %s", info.Method, info.Path)
		return nil
	})

	return srv
}

func MiddlewareAdapter() middleware.KMiddleware {
	return func(handler middleware.KHandler) middleware.KHandler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				return handler(transport.NewClientContext(ctx, tr), req)
			}
			return handler(ctx, req)
		}
	}
}

func DefaultPaths() []string {
	return []string{
		basis.OperationLoginAPICaptchaID,
		basis.OperationLoginAPICaptchaImage,
		basis.OperationLoginAPICaptchaResource,
		basis.OperationLoginAPICaptchaResources,
		//basis.OperationLoginAPICurrentMenus,
		//basis.OperationLoginAPICurrentUser,
		basis.OperationLoginAPILogin,
		//basis.OperationLoginAPILogout,
		basis.OperationLoginAPIRefresh,
		basis.OperationLoginAPIRegister,
	}
}
