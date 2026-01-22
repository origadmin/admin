package providers

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/wire"

	watcher "github.com/origadmin/casbin-watcher/v3"
	_ "github.com/origadmin/casbin-watcher/v3/drivers/nats"
	authnv1 "github.com/origadmin/contrib/api/gen/go/security/authn/v1"
	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/contrib/security/credential"
	secmiddleware "github.com/origadmin/contrib/security/middleware"
	"github.com/origadmin/contrib/security/request"
	"github.com/origadmin/contrib/security/skip"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/extensions/configutil"
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
	"origadmin/application/admin/internal/helpers/captcha"
)

var (
	factory  = secmiddleware.NewFactory()
	policies = make(map[string]security.Policy)
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
	jwtConfig, _, err := configutil.Normalize(authnConfig.GetActive(), authnConfig.GetDefault(),
		authnConfig.GetConfigs())
	if err != nil {
		return nil, fmt.Errorf("failed to normalize JWT authenticator configuration: %w", err)
	}
	return jwt.NewOptions(jwtConfig)
}

// ProvideCredentialCreator creates the JWT authenticator instance.
// It returns a credential.Creator interface, which is implemented by *jwt.Authenticator.
func ProvideCredentialCreator(opts *jwt.Options, logger log.Logger) (credential.Creator, error) {
	return jwt.New(opts, logger)
}

// ProvideAuthorizer creates the Casbin authorizer.
func ProvideAuthorizer(app *runtime.App, c *conf.Config, adapter *data.CasbinAdapter,
	w *watcher.Watcher) (*casbin.Authorizer, error) {
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
	opts, err := casbin.NewOptions(casbinConfig, casbin.WithPolicyAdapter(adapter), casbin.WithWatcher(w))
	if err != nil {
		return nil, err
	}

	return casbin.New(opts, app.Logger())
}

func ProvideWatcher(app *runtime.App, c *conf.Config) (*watcher.Watcher, error) {
	brokerConfig := c.GetBrokers()
	if brokerConfig == nil {
		return nil, errors.New("broker configuration not found")
	}

	brokerUrl := brokerConfig.GetDefault().GetUrl()
	if brokerUrl == "" {
		return nil, errors.New("broker url not found")
	}

	// Create a NATS watcher
	w, err := watcher.NewWatcher(app.Context(), brokerUrl)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// ProvideAuthenticator creates the Casbin authorizer.
func ProvideAuthenticator(app *runtime.App, c *conf.Config) (*jwt.Authenticator, error) {
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
	opts, err := jwt.NewOptions(jwtConfig)
	if err != nil {
		return nil, err
	}
	return jwt.New(opts, app.Logger())
}

func ProvideCache(app *runtime.App) (container.CacheProvider, error) {
	cacheProvider, err := app.CacheProvider()
	if err != nil {
		return nil, err
	}
	return cacheProvider, nil
}

func ProvideCaptcha(app *runtime.App, p container.CacheProvider, cfg *confpb.Bootstrap) (*captcha.Captcha,
	error) {
	captchaConfig := cfg.GetCaptcha()
	if captchaConfig == nil {
		captchaConfig = &confpb.Captcha{}
	}
	if captchaConfig.CacheName == "" {
		captchaConfig.CacheName = "default"
	}
	if captchaConfig.Height == 0 {
		captchaConfig.Height = 80
	}
	if captchaConfig.Width == 0 {
		captchaConfig.Width = 240
	}
	if captchaConfig.Length == 0 {
		captchaConfig.Length = 6
	}
	if captchaConfig.MaxSkew == 0 {
		captchaConfig.MaxSkew = 0.7
	}
	if captchaConfig.DotCount == 0 {
		captchaConfig.DotCount = 80
	}

	cache, err := p.Cache(captchaConfig.CacheName)
	if err != nil {
		return nil, err
	}

	c := &captcha.Config{
		Store:   captcha.NewStore(cache),
		Captcha: captchaConfig,
	}
	return captcha.NewCaptcha(c), nil
}

func ProvideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

func ProvideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

func ProvideServiceMiddlewares(app *runtime.App, authorizer *casbin.Authorizer,
	skip security.Skipper) (container.ServerMiddlewareProvider,
	error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewBackend(authorizer, skip)
	// authz for backend
	provider.RegisterServerMiddleware("authz", m)
	provider.RegisterClientMiddleware("authz", middleware.Noop())
	helper := log.NewHelper(app.Logger())
	helper.Infof("registered %+v middlewares", provider.Names())
	return provider, nil
}

func ProvideGatewayMiddlewares(app *runtime.App, authenticator *jwt.Authenticator,
	skip security.Skipper) (container.ServerMiddlewareProvider,
	error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewGateway(authenticator, skip)
	// authz for backend
	provider.RegisterServerMiddleware("authn", m)
	provider.RegisterClientMiddleware("authn", middleware.Noop())
	helper := log.NewHelper(app.Logger())
	helper.Infof("registered %+v middlewares", provider.Names())
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
	helper := log.NewHelper(app.Logger())
	helper.Infof("registered %+v client middlewares", provider.Names())
	return provider, nil
}

func ProvideGatewaySkipper(app *runtime.App, _ *conf.Config) security.Skipper {
	skips := make(map[string]struct{})
	for _, v := range policies {
		if v.Name == "public" {
			skips[v.ServiceMethod] = struct{}{}
		}
	}

	return func(ctx context.Context, req security.Request) bool {
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
		if _, ok := skips[req.GetOperation()]; ok {
			helper.Infof("skip gateway checker: %s", req.GetOperation())
			return true
		}
		helper.Infof("unskipped request: %s", req.GetOperation())
		return false
	}
}

func ProvideSkipper(app *runtime.App, _ *conf.Config) security.Skipper {
	adminSkipper := skip.Principal(func(principal security.Principal) bool {
		helper := log.NewHelper(log.With(app.Logger()))
		pid := strconv.Itoa(int(data.SystemUserID)) // Conveapp int64 to string for comparison
		if principal.GetID() == pid {
			helper.Infof("skip admin checker: %s", pid)
			return true
		}
		return false
	})
	skips := make(map[string]struct{})
	for _, v := range policies {
		if v.Name == "public" {
			skips[v.ServiceMethod] = struct{}{}
		}
		if v.Name == "jwt-auth" {
			skips[v.ServiceMethod] = struct{}{}
		}
	}
	pathSkipper := func(ctx context.Context, req security.Request) bool {
		helper := log.NewHelper(log.With(app.Logger(), "kind", req.Kind(), "operation", req.GetOperation(), "method", req.GetMethod(), "path",
			req.GetRouteTemplate()))
		if _, ok := skips[req.GetOperation()]; ok {
			helper.Infof("skip checker: %s", req.GetOperation())
			return true
		}
		return false
	}
	skippers := skip.Composite(adminSkipper, pathSkipper)
	return func(ctx context.Context, req security.Request) bool {
		return skippers(ctx, req)
	}
}

var ProviderGatewaySet = wire.NewSet(
	ProvideClientMiddlewares,
	ProvideGatewayMiddlewares,
	ProvideGatewaySkipper,
)

var ProviderBackendSet = wire.NewSet(
	ProvideClientMiddlewares,
	ProvideServiceMiddlewares,
	ProvideSkipper,
)

var ProviderSet = wire.NewSet(
	// CORRECTED: Added FieldsOf for Security to ensure it's provided to the authenticator.
	wire.FieldsOf(new(*conf.Config), "Bootstrap"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Security"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Captcha"),
	ProvideLogger,
	ProvideCache,
	ProvideAuthenticatorOptions,
	ProvideCredentialCreator,
	ProvideWatcher,
	ProvideAuthenticator,
	ProvideAuthorizer,
	ProvideCaptcha,
	ProvideHasher,
)
