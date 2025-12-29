/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/origadmin/contrib/security/authn"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/runtime"
	"origadmin/application/admin/internal/conf"

	authnv1 "github.comcom/origadmin/contrib/api/gen/go/security/authn/v1"
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/config/middleware/v1"
)

// provideAuthenticator creates a new JWT authenticator from the application configuration.
func provideAuthenticator(c *conf.Config, logger log.Logger) (authn.Authenticator, error) {
	m, err := c.DecodeMiddlewares()
	if err != nil {
		return nil, fmt.Errorf("failed to decode middlewares config: %w", err)
	}

	var jwtMiddleware *middlewarev1.Middleware
	for _, mw := range m.GetConfigs() {
		if mw.GetType() == "jwt" {
			jwtMiddleware = mw
			break
		}
	}

	if jwtMiddleware == nil || jwtMiddleware.GetJwt() == nil {
		return nil, errors.New("JWT middleware configuration not found in bootstrap config")
	}

	authnConfig := &authnv1.Authenticator{
		Jwt: jwtMiddleware.GetJwt(),
	}

	return jwt.NewAuthenticator(authnConfig, log.NewOption(logger))
}

func provideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

var infraProviderSet = wire.NewSet(provideLogger, provideAuthenticator)
