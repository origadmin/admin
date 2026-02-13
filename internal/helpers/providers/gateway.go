package providers

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/wire"

	contribsecurity "github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/request"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/extensions/configutil"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"origadmin/application/admin/internal/conf"
)

// ProviderGatewaySet provides gateway-specific dependencies.
var ProviderGatewaySet = wire.NewSet(
	ProviderCommonSet,
	ProvideAuthenticator,
	ProvideGatewayMiddlewares,
	ProvideClientMiddlewares,
	ProvideGatewaySkipper,
)

// ProvideAuthenticator creates the JWT authenticator, which also serves as a credential.Creator.
func ProvideAuthenticator(app *runtime.App, c *conf.Config) (*jwt.Authenticator, error) {
	securityConfig := c.GetBootstrap().GetSecurity()
	if securityConfig == nil {
		return nil, errors.New("security configuration not found")
	}
	authnConfig := securityConfig.GetAuthn()
	if authnConfig == nil {
		return nil, errors.New("authn configuration not found")
	}

	jwtConfig, _, err := configutil.Normalize(authnConfig.GetActive(), authnConfig.GetDefault(), authnConfig.GetConfigs())
	if err != nil {
		return nil, fmt.Errorf("failed to normalize JWT authenticator configuration: %w", err)
	}

	if jwtConfig == nil || jwtConfig.GetJwt() == nil || jwtConfig.GetType() != "jwt" {
		return nil, errors.New("JWT authenticator configuration not found")
	}
	opts, err := jwt.NewOptions(jwtConfig)
	if err != nil {
		return nil, err
	}
	return jwt.New(opts, app.Logger())
}

// ProvideGatewayMiddlewares creates gateway-specific middlewares.
func ProvideGatewayMiddlewares(app *runtime.App, authenticator *jwt.Authenticator, skip contribsecurity.Skipper) (container.ServerMiddlewareProvider, error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewAuthnGateway(authenticator, skip, log.WithLogger(app.Logger()))
	provider.RegisterServerMiddleware("authn", m)
	provider.RegisterClientMiddleware("authn", middleware.Noop())
	log.NewHelper(app.Logger()).Infof("registered %+v middlewares", provider.Names())
	return provider, nil
}

// ProvideClientMiddlewares creates client-specific middlewares for propagation.
func ProvideClientMiddlewares(app *runtime.App) (container.ClientMiddlewareProvider, error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewPropagationClient(log.WithLogger(app.Logger()))
	provider.RegisterClientMiddleware("propagation", m)
	log.NewHelper(app.Logger()).Infof("registered %+v client middlewares", provider.Names())
	return provider, nil
}

// ProvideGatewaySkipper creates a skipper for the gateway.
func ProvideGatewaySkipper(app *runtime.App, _ *conf.Config) contribsecurity.Skipper {
	return func(ctx context.Context, req contribsecurity.Request) bool {
		helper := log.NewHelper(log.With(app.Logger(),
			"kind", req.Kind(),
			"operation", req.GetOperation(),
			"path", req.GetRouteTemplate(),
		))
		if req, err := request.NewFromServerContext(ctx); err == nil {
			helper.Infof("method: %s, path: %s, operation: %s",
				req.GetMethod(),
				req.GetRouteTemplate(),
				req.GetOperation(),
			)
		} else {
			return false
		}
		// Use the pre-filtered map for a fast lookup.
		if _, ok := gatewaySkipMap[req.GetOperation()]; ok {
			helper.Infof("skip gateway checker: %s", req.GetOperation())
			return true
		}
		helper.Infof("unskipped request: %s", req.GetOperation())
		return false
	}
}
