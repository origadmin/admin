package providers

import (
	"context"
	"errors"

	"github.com/google/wire"

	authnv1 "github.com/origadmin/contrib/api/gen/go/security/authn/v1"
	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/contrib/security/credential"
	secmiddleware "github.com/origadmin/contrib/security/middleware"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/types"
	_ "origadmin/application/admin/api/v1/services/auth"
	_ "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/helpers/captcha"
)

var (
	factory  = secmiddleware.NewFactory()
	policies = map[string]security.Policy{}
)

func init() {
	ps := security.RegisteredPolicies()
	for _, p := range ps {
		policies[p.ServiceMethod] = p
	}
}

// ProvideAuthenticatorOptions creates the JWT options from the application configuration.
func ProvideAuthenticatorOptions(c *conf.Config) (*jwt.Options, error) {
	securityConfig := c.GetBootstrap().GetSecurity()
	if securityConfig == nil {
		return nil, errors.New("security configuration not found in bootstrap config")
	}
	authnConfig := securityConfig.GetAuthn()
	if authnConfig == nil {
		return nil, errors.New("authn configuration not found in security config")
	}
	configs := authnConfig.GetConfigs()

	var jwtConfig *authnv1.Authenticator
	for _, mw := range configs {
		if mw.GetType() == "jwt" {
			jwtConfig = mw
			break
		}
	}
	if jwtConfig == nil || jwtConfig.GetJwt() == nil {
		return nil, errors.New("JWT authenticator configuration not found in bootstrap config")
	}
	return jwt.NewOptions(jwtConfig)
}

// ProvideCredentialCreator creates the JWT authenticator instance.
// It returns a credential.Creator interface, which is implemented by *jwt.Authenticator.
func ProvideCredentialCreator(opts *jwt.Options, logger log.Logger) (credential.Creator, error) {
	return jwt.New(opts, logger)
}

// ProvideAuthorizer creates the Casbin authorizer.
func ProvideAuthorizer(app *runtime.App, c *conf.Config, database *ent.Database) (authz.Authorizer, error) {
	securityConfig := c.GetBootstrap().GetSecurity()
	if securityConfig == nil {
		return nil, errors.New("security configuration not found")
	}
	authzConfig := securityConfig.GetAuthz()
	if authzConfig == nil {
		return nil, errors.New("authz configuration not found")
	}

	var casbinConfig *authzv1.Authorizer
	for _, cfg := range authzConfig.GetConfigs() {
		if cfg.GetType() == "casbin" {
			casbinConfig = cfg
			break
		}
	}

	if casbinConfig == nil || casbinConfig.GetCasbin() == nil {
		return nil, errors.New("casbin authorizer configuration not found")
	}
	adapter, err := data.NewAdapter(database)
	if err != nil {
		return nil, err
	}

	return casbin.NewAuthorizer(casbinConfig, log.WithLogger(app.Logger()), casbin.WithPolicyAdapter(adapter))
}

// ProvideAuthenticator creates the Casbin authorizer.
func ProvideAuthenticator(app *runtime.App, c *conf.Config) (authn.Authenticator, error) {
	securityConfig := c.GetBootstrap().GetSecurity()
	if securityConfig == nil {
		return nil, errors.New("security configuration not found")
	}
	authnConfig := securityConfig.GetAuthn()
	if authnConfig == nil {
		return nil, errors.New("authz configuration not found")
	}

	var jwtConfig *authnv1.Authenticator
	for _, cfg := range authnConfig.GetConfigs() {
		if cfg.GetType() == "jwt" {
			jwtConfig = cfg
			break
		}
	}

	if jwtConfig == nil || jwtConfig.GetJwt() == nil {
		return nil, errors.New("casbin authorizer configuration not found")
	}
	return jwt.NewAuthenticator(jwtConfig, log.WithLogger(app.Logger()))
}

func ProvideCache(r *runtime.App) (container.CacheProvider, error) {
	cacheProvider, err := r.CacheProvider()
	if err != nil {
		return nil, err
	}
	return cacheProvider, nil
}

func ProvideCaptcha(p container.CacheProvider, cfg *confpb.Captcha) (*captcha.Captcha, error) {
	if cfg == nil {
		cfg = &confpb.Captcha{
			CacheName: "default",
			Height:    80,
			Width:     240,
			Length:    6,
			Maxskew:   0.7,
			DotCount:  80,
		}
	}
	cache, err := p.Cache(cfg.CacheName)
	if err != nil {
		return nil, err
	}

	c := &captcha.Config{
		Store: captcha.NewStore(cache),
		DriverDigit: &captcha.DriverDigit{
			Height:   int(cfg.GetHeight()),
			Width:    int(cfg.GetWidth()),
			Length:   int(cfg.GetLength()),
			MaxSkew:  float64(cfg.GetMaxskew()),
			DotCount: int(cfg.GetDotCount()),
		},
	}
	return captcha.NewCaptcha(c), nil
}

func ProvideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

func ProvideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

func ProvideServiceMiddlewares(app *runtime.App, authorizer authz.Authorizer,
	skip security.SkipChecker) (container.ServerMiddlewareProvider,
	error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewBackend(authorizer, skip)
	// authz for backend
	provider.RegisterServerMiddleware("authz", m)
	provider.RegisterClientMiddleware("authz", middleware.Noop())

	return provider, nil
}

func ProvideGatewayMiddlewares(app *runtime.App, authenticator authn.Authenticator,
	skip security.SkipChecker) (container.ServerMiddlewareProvider,
	error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewGateway(authenticator, skip)
	// authz for backend
	provider.RegisterServerMiddleware("authn", m)
	provider.RegisterClientMiddleware("authn", middleware.Noop())

	return provider, nil
}

func ProvideClientMiddlewares(app *runtime.App) (container.ClientMiddlewareProvider,
	error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewClient()
	// authz for backend
	provider.RegisterClientMiddleware("propagation", m)
	provider.RegisterServerMiddleware("propagation", middleware.Noop())
	return provider, nil
}

func ProvideSkipChecker(app *runtime.App, cfg *conf.Config) security.SkipChecker {
	helper := log.NewHelper(log.With(app.Logger(), "module", "security.skip"))
	return func(ctx context.Context, req security.Request) bool {
		helper.Infow("kind", req.Kind(), "operation", req.GetOperation(), "method", req.GetMethod(), "path",
			req.GetRouteTemplate())
		if v, ok := policies[req.GetOperation()]; ok && v.Name == "public" {
			return true
		}
		return false
	}
}

var ProviderGatewaySet = wire.NewSet(
	ProvideClientMiddlewares,
	ProvideGatewayMiddlewares,
	ProvideSkipChecker,
)

var ProviderBackendSet = wire.NewSet(
	ProvideClientMiddlewares,
	ProvideServiceMiddlewares,
	ProvideSkipChecker,
)

var ProviderSet = wire.NewSet(
	// Instructions for wire to extract nested configs
	wire.FieldsOf(new(*conf.Config), "Bootstrap"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Captcha"),
	ProvideLogger,
	ProvideCache,
	ProvideAuthenticatorOptions,
	ProvideCredentialCreator,
	ProvideAuthenticator,
	ProvideAuthorizer,
	ProvideCaptcha,
	ProvideHasher,
)
