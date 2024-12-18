/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/goexts/generic/types"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	middlewaresecurity "github.com/origadmin/runtime/middleware/security"
	"github.com/origadmin/runtime/service"
	servicehttp "github.com/origadmin/runtime/service/http"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/internal/configs"
)

type HTTPRegistrar interface {
	HTTP(server service.HTTPServer)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(bootstrap *configs.Bootstrap, authenticator security.Authenticator, authorizer security.Authorizer, l log.Logger) *service.HTTPServer {
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		panic("no serviceConfig")
	}
	paths := bootstrap.GetMiddlewares().GetSecurity().GetPublicPaths()
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
	}))
	ms = append([]middleware.Middleware{
		func(handler middleware.Handler) middleware.Handler {
			return func(ctx context.Context, req interface{}) (interface{}, error) {
				fmt.Println("Handler called")
				tr, ok := transport.FromClientContext(ctx)
				if !ok {
					log.Errorf("No transport found in context")
				} else {
					log.Debugf("Starting HTTP handler for path '%s'", tr.Operation())
				}

				return handler(ctx, req)
			}
		},
	}, ms...)
	serviceConfig.Name = types.ZeroOr(serviceConfig.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(bootstrap.GetService(), service.WithHTTP(servicehttp.WithMiddlewares(ms...)))
	if err != nil {
		panic(err)
	}

	return srv
}

//func RegisterHTTPServer(srv *http.Register, menus pb.MenuAPIServer, login common.LoginAPIServer) {
//	common.RegisterLoginAPIHTTPServer(srv, login)
//	pb.RegisterMenuAPIHTTPServer(srv, menus)
//}
