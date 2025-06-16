/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"strings"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/interfaces/security"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicegrpc "github.com/origadmin/runtime/service/grpc"
	servicehttp "github.com/origadmin/runtime/service/http"

	"origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/configs"
)

type ProxyOptions struct {
	Authenticator security.Authenticator
	Authorizer    security.Authorizer
}

func NewProxyOptions(r runtime.Runtime, bootstrap *configs.Bootstrap,
	source casbin.RuleSource) (*ProxyOptions, error) {
	authenticator, err := securityx.NewAuthenticator(bootstrap)
	if err != nil {
		return nil, err
	}
	opts := []casbin.AuthorizerOption{
		casbin.WithSource(source),
	}

	authorizer, err := securityx.NewAuthorizer(bootstrap, opts...)
	if err != nil {
		return nil, err
	}
	return &ProxyOptions{
		Authenticator: authenticator,
		Authorizer:    authorizer,
	}, nil
}

// NewProxyServer creates a new proxy server.
func NewProxyServer(
	r runtime.Runtime,
	bootstrap *configs.Bootstrap,
	registrars []service.ServerRegistrar,
	opts *ProxyOptions) []transport.Server {
	paths := bootstrap.GetSecurity().GetSecurity().GetPublicPaths()
	paths = append(DefaultPaths(), paths...)
	ms := []middleware.KMiddleware{
		recovery.Recovery(),
	}

	bridge := securityx.DefaultBridge()
	bridge.Authenticator = opts.Authenticator
	bridge.Authorizer = opts.Authorizer
	bridge.IsRoot = func(ctx context.Context, claims security.Claims) bool {
		return claims.GetSubject() == "root" || claims.GetSubject() == "admin"
	}
	serv := selector.Server(bridge.Middleware()).Match(func(ctx context.Context, operation string) bool {
		for _, p := range paths {
			if strings.HasPrefix(operation, p) {
				log.Debugf("Operation '%s' matches public path '%s', returning true", operation, p)
				return false
			}
		}
		log.Infof("Operation '%s' no matches public path '%s'", operation, "*")
		return true
	})
	ms = append(ms, serv.Build(), CallLoggerMiddleware())
	//clients.Get
	//for i := range clients {
	//	clients[i].GetCore().GetName()
	//
	//}
	//clients.Name = types.ZeroOr(clients.Name, "ORIGADMIN_SERVICE")
	var servers []transport.Server
	services := bootstrap.GetEntry().GetServices()
	for i := range services {
		if services[i].GetType() != "http" {
			continue
		}
		srv, err := r.Builder().NewHTTPServer(services[i],
			servicehttp.WithServerOptions(
				http.PathPrefix("/api/v1"),
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
		servers = append(servers, srv)
	}
	return servers
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
			log.Infof("Caller Client: %+v, ok: %+v", tr, ok)
			return handler(ctx, req)
		}
	}
}

func BridgeMiddleware() middleware.KMiddleware {
	return func(handler middleware.KHandler) middleware.KHandler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			meta, _ := metadata.FromClientContext(ctx)
			log.Infof("Caller Client Metadata: %+v", meta)
			smd, ok := metadata.FromServerContext(ctx)
			if !ok {
				smd = metadata.New(nil)
			}
			log.Infof("Caller Server Metadata: %+v", smd)
			for k, v := range meta {
				smd[k] = v
			}
			ctx = metadata.NewServerContext(ctx, smd)
			return handler(ctx, req)
		}
	}
}

func NewProxyGRPCClients(r runtime.Runtime, bootstrap *configs.Bootstrap) map[string]*service.GRPCClient {
	ll := log.NewHelper(r.WithLogger("module", "proxy"))
	ll.Infof("NewProxyGRPCClients bootstrap: %+v", bootstrap)
	clients := bootstrap.GetClients()
	clientServices := make(map[string]*service.GRPCClient, len(clients))
	for i := range clients {
		services := clients[i].GetServices()
		if len(services) == 0 {
			continue
		}
		ll.Infof("NewProxyGRPCClients: %+v", clients[i].GetCore().GetName())
		var options []service.GRPCOption
		discovery, err := r.Builder().NewDiscovery(clients[i].GetCore().GetDiscovery())
		if err == nil {
			options = append(options, servicegrpc.WithDiscovery(clients[i].GetCore().GetDiscovery().GetServiceName(),
				discovery))
		}
		for idx := range services {
			if services[idx].GetType() == "grpc" {
				ll.Infof("NewProxyGRPCClient Middleware: %+v", clients[i].GetMiddleware())
				ms := r.Builder().NewMiddlewaresClient(clients[i].GetMiddleware())
				if len(ms) > 0 {
					options = append(options, servicegrpc.WithMiddlewares(ms...))
				}
				client, err := r.Builder().NewGRPCClient(r.Context(), services[idx], options...)
				if err != nil {
					ll.Warnf("NewGRPCClient failed: %v", err)
					continue
				}
				ll.Infof("NewProxyGRPCClients: %+v", clients[i].GetCore().GetName())
				clientServices[clients[i].GetCore().GetName()] = client
			}
		}
	}
	return clientServices
}

func NewProxyHTTPClients(r runtime.Runtime, bootstrap *configs.Bootstrap) map[string]*service.HTTPClient {
	ll := log.NewHelper(r.WithLogger("module", "proxy"))
	clients := bootstrap.GetClients()
	clientServices := make(map[string]*service.HTTPClient, len(clients))
	for i := range clients {
		services := clients[i].GetServices()
		if len(services) == 0 {
			continue
		}
		var options []service.HTTPOption
		discovery, err := r.Builder().NewDiscovery(clients[i].GetCore().GetDiscovery())
		if err == nil {
			options = append(options, servicehttp.WithDiscovery(clients[i].GetCore().GetDiscovery().GetServiceName(), discovery))
		}
		for idx := range services {
			if services[idx].GetType() == "http" {
				ll.Infof("NewProxyHTTPClient Middleware: %+v", clients[i].GetMiddleware())
				ms := r.Builder().NewMiddlewaresClient(clients[i].GetMiddleware())
				if len(ms) > 0 {
					options = append(options, servicehttp.WithMiddlewares(ms...))
				}
				client, err := r.Builder().NewHTTPClient(r.Context(), services[idx], options...)
				if err != nil {
					ll.Warnf("NewHTTPClient failed: %v", err)
					continue
				}
				clientServices[clients[i].GetCore().GetName()] = client
			}
		}
	}
	return clientServices
}
