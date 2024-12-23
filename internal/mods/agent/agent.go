/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"github.com/goexts/generic/types"
	"github.com/google/wire"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/service"
	servicehttp "github.com/origadmin/runtime/service/http"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/internal/configs"
)

var ProviderSet = wire.NewSet(
	NewHTTPServerAgent,
)

// NewGINServer new an HTTP server.
func NewGINServer(bootstrap *configs.Bootstrap, registrars []HTTPRegistrar, authenticator security.Authenticator, authorizer security.Authorizer, l log.KLogger) *service.HTTPServer {
	serviceConfig := bootstrap.GetService()
	if serviceConfig == nil {
		panic("no service config")
	}

	serviceConfig.Name = types.ZeroOr(serviceConfig.Name, "ORIGADMIN_SERVICE")
	srv, err := runtime.NewHTTPServiceServer(serviceConfig, service.WithHTTP(
		servicehttp.WithMiddlewares(middleware.NewClient(bootstrap.GetMiddleware())...),
	))
	if err != nil {
		panic(err)
	}

	//engine := gin.New()
	//engine.ContextWithFallback = true
	//config := bootstrap.GetMiddlewares()
	//paths := config.GetSecurity().GetPublicPaths()
	//ms := middleware.New(config,
	//	middleware.WithAuthorizer(authorizer),
	//	middleware.WithAuthenticator(authenticator),
	//	middleware.WithSkipper(paths),
	//)
	//engine.Use(gin.Recovery())
	for _, registrar := range registrars {
		registrar.HTTP(srv)
	}
	//srv.HandlePrefix("/", engine)
	return srv
}
