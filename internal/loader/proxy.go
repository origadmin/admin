/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"strings"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/origadmin/runtime"
	msecurity "github.com/origadmin/runtime/agent/middleware/security"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/interfaces/security"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicehttp "github.com/origadmin/runtime/service/http"

	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/configs"
)

type data struct {
}

func (d data) QueryRoles(ctx context.Context, subject string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (d data) QueryPermissions(ctx context.Context, subject string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

// NewProxyServer creates a new proxy server.
func NewProxyServer(r runtime.Runtime, bootstrap *configs.Bootstrap, registrars []service.ServerRegistrar,
	client auth.CasbinSourceServiceClient) []transport.Server {
	clients := bootstrap.GetClients()
	if clients == nil {
		panic("no service config")
	}
	paths := bootstrap.GetSecurity().GetSecurity().GetPublicPaths()
	paths = append(DefaultPaths(), paths...)
	ms := []middleware.KMiddleware{
		recovery.Recovery(),
	}
	authenticator, err := securityx.NewAuthenticator(bootstrap)
	if err != nil {
		panic(err)
	}

	opts := []casbin.AuthorizerOption{casbin.WithServiceClient(client)}

	authorizer, err := securityx.NewAuthorizer(bootstrap, opts...)
	if err != nil {
		panic(err)
	}
	bridge := securityx.SecurityBridge{
		TokenSource:          security.TokenSourceHeader,
		Scheme:               security.SchemeBearer,
		AuthenticationHeader: security.HeaderAuthorize,
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
		Provider:    &data{},
		TokenParser: nil,
	}
	serv := selector.Server(bridge.Middleware()).Match(func(ctx context.Context, operation string) bool {
		for _, p := range paths {
			if strings.HasPrefix(operation, p) {
				log.Debugf("Operation '%s' matches public path '%s', returning true", operation, p)
				return false
			}
		}
		log.Debugf("Operation '%s' no matches public path '%s'", operation, "*")
		return true
	})
	ms = append(ms, serv.Build(), CallLoggerMiddleware())
	//clients.Get
	for i := range clients {
		clients[i].GetCore().GetName()

	}
	//clients.Name = types.ZeroOr(clients.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetEntry().GetServer(),
		servicehttp.WithServerOptions(
			http.ErrorEncoder(resp.ResponseErrorEncoder),
			http.Filter(handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			))),
		servicehttp.WithMiddlewares(ms...),
		servicehttp.WithPrefix(runtime.DefaultEnvPrefix),
	)
	if err != nil {
		panic(err)
	}
	for _, registrar := range registrars {
		registrar.Register(r.Context(), srv)
	}
	srv.WalkRoute(func(info http.RouteInfo) error {
		log.Infof("Registered HTTP route: %s %s", info.Method, info.Path)
		return nil
	})

	return []transport.Server{srv}
}

func DefaultPaths() []string {
	return []string{
		auth.OperationLoginServiceCaptchaId,
		auth.OperationLoginServiceCaptcha,
		auth.OperationLoginServiceCaptchaImage,
		auth.OperationLoginServiceCaptchaAudio,
		auth.OperationLoginServiceLogin,
		auth.OperationLoginServiceRegister,
		auth.OperationLoginServiceTokenRefresh,
	}
}

func CallLoggerMiddleware() middleware.KMiddleware {
	return func(handler middleware.KHandler) middleware.KHandler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Infof("CallLoggerMiddleware: %+v", ctx)
			tr, ok := transport.FromServerContext(ctx)
			log.Infof("Caller Server: %+v, ok: %+v", tr, ok)
			tr, ok = transport.FromClientContext(ctx)
			log.Infof("Caller ServiceClient: %+v, ok: %+v", tr, ok)
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
