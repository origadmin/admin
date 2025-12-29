/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package main

import (
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/credential"
	"github.com/origadmin/runtime"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	hash_types "github.com/origadmin/toolkits/crypto/hash/types"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"

	authnv1 "github.com/origadmin/contrib/api/gen/go/security/authn/v1"
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/config/middleware/v1"
)

// provideCredentialCreator creates a credential creator from the application configuration.
// It finds the JWT middleware config and uses it to initialize a jwt.Authenticator,
// which implements the credential.Creator interface.
func provideCredentialCreator(c *conf.Config, logger log.Logger) (credential.Creator, error) {
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

	// jwt.NewAuthenticator returns a *jwt.Authenticator which implements credential.Creator
	return jwt.NewAuthenticator(authnConfig, log.NewOption(logger))
}

func provideCaptchaConfig(c *conf.Config) (*confpb.Captcha, error) {
	return c.GetCaptcha()
}

func provideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(hash_types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

func provideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

var infraProviderSet = wire.NewSet(provideLogger, provideCredentialCreator, provideCaptchaConfig, provideHasher)
