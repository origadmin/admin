/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	context2 "context"
	"strings"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/types"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	middlewaresecurity "github.com/origadmin/runtime/middleware/security"
	"github.com/origadmin/runtime/service"
	servicehttp "github.com/origadmin/runtime/service/http"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/api/v1/services/basis"
	"origadmin/application/admin/internal/configs"
)

type HTTPRegistrar interface {
	HTTP(server *service.HTTPServer)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *configs.Bootstrap, registrars []HTTPRegistrar, authenticator security.Authenticator, authorizer security.Authorizer, l log.Logger) *service.HTTPServer {
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		panic("no serviceConfig")
	}
	paths := bootstrap.GetMiddlewares().GetSecurity().GetPublicPaths()
	paths = append(DefaultPaths(), paths...)
	ms := middleware.NewClient(bootstrap.GetMiddlewares(), middleware.WithSecurityOptions(func(option *middlewaresecurity.Option) {
		option.Authenticator = authenticator
		option.Authorizer = authorizer
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
		option.Parser = func(ctx context2.Context, id string) (security.UserClaims, error) {
			//todo: 从token中获取用户信息
			return nil, nil
		}
	}))
	ms = append([]middleware.Middleware{MiddlewareAdapter()}, ms...)
	serviceConfig.Name = types.ZeroOr(serviceConfig.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService(), service.WithHTTP(servicehttp.WithMiddlewares(ms...)))
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

func MiddlewareAdapter() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
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
