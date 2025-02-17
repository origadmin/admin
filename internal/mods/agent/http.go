/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"strings"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/types"
	"github.com/gorilla/handlers"
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
		panic("no service config")
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
		SkipKey:              msecurity.MetadataSecuritySkipKey,
		PublicPaths:          nil,
		Skipper: func(path string) bool {
			return false
		},
		IsRoot: func(ctx context.Context, claims security.Claims) bool {
			return claims.GetSubject() == "root" || claims.GetSubject() == "admin"
		},
		Data:        &data{},
		TokenParser: nil,
	}
	serv := selector.Server(bridge.Build()).Match(func(ctx context.Context, operation string) bool {
		for _, p := range paths {
			if strings.HasPrefix(operation, p) {
				log.Debugf("Operation '%s' matches public path '%s', returning true", operation, p)
				return false
			}
		}
		log.Debugf("Operation '%s' no matches public path '%s'", operation)
		return true
	})
	ms = append(ms, serv.Build(), CallerMiddleware())

	serviceConfig.Name = types.ZeroOr(serviceConfig.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService(), service.WithHTTP(
		servicehttp.WithServerOptions(
			http.ErrorEncoder(resp.ResponseErrorEncoder),
			http.Filter(handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			))),
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
		system.OperationLoginServiceCaptchaId,
		system.OperationLoginServiceCaptcha,
		system.OperationLoginServiceCaptchaImage,
		system.OperationLoginServiceCaptchaAudio,
		system.OperationLoginServiceLogin,
		system.OperationLoginServiceRegister,
	}
}

func CallerMiddleware() middleware.KMiddleware {
	return func(handler middleware.KHandler) middleware.KHandler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Infof("CallerMiddleware: %+v", ctx)
			tr, ok := transport.FromServerContext(ctx)
			log.Infof("Caller Server: %+v, ok: %+v", tr, ok)
			tr, ok = transport.FromClientContext(ctx)
			log.Infof("Caller Client: %+v, ok: %+v", tr, ok)
			return handler(ctx, req)
		}
	}
}

func CorsMiddleware() middleware.KMiddleware {
	return func(handler middleware.KHandler) middleware.KHandler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Infof("CorsMiddleware: %+v", ctx)
			return handler(ctx, req)
		}
	}
}
