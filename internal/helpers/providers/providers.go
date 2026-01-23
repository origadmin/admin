package providers

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	"github.com/google/wire"

	watcher "github.com/origadmin/casbin-watcher/v3"
	_ "github.com/origadmin/casbin-watcher/v3/drivers/nats"
	"github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/contrib/security/credential"
	securitymiddleware "github.com/origadmin/contrib/security/middleware"
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
	"origadmin/application/admin/internal/broker"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/helpers/captcha"
)

var (
	factory  = securitymiddleware.NewFactory()
	policies = make(map[string]security.Policy)
)

func init() {
	ps := security.RegisteredPolicies()
	for _, p := range ps {
		policies[p.ServiceMethod] = p
	}
}

// ProviderCommonSet provides common dependencies that are safe for all modules.
var ProviderCommonSet = wire.NewSet(
	ProvideLogger,
	ProvideCache,
	ProvideHasher,
	ProvideCaptcha,
	wire.FieldsOf(new(*conf.Config), "Bootstrap"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Security"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Captcha"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Brokers"),
)

// ProviderGatewaySet provides gateway-specific dependencies.
var ProviderGatewaySet = wire.NewSet(
	ProviderCommonSet,
	ProvideAuthenticator,
	ProvideGatewayMiddlewares,
	ProvideClientMiddlewares,
	ProvideGatewaySkipper,
)

// ProviderBackendSet provides backend-specific dependencies.
var ProviderBackendSet = wire.NewSet(
	ProviderCommonSet,
	ProvideAuthorizer,
	ProvideWatcher,
	ProvideAuthenticator,                                        // Provides *jwt.Authenticator
	wire.Bind(new(credential.Creator), new(*jwt.Authenticator)), // Binds the interface
	ProvideServiceMiddlewares,
	ProvideClientMiddlewares,
	ProvideSkipper,
	ProvidePublisher,
	wire.Bind(new(broker.Publisher), new(*nats.Publisher)), // Binds the interface
)

// ProvideLogger provides a logger instance from the runtime App.
func ProvideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

// ProvideCache provides a cache provider from the runtime App.
func ProvideCache(app *runtime.App) (container.CacheProvider, error) {
	return app.CacheProvider()
}

// ProvideHasher provides a password hasher.
func ProvideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

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
func ProvideGatewayMiddlewares(app *runtime.App, authenticator *jwt.Authenticator, skip security.Skipper) (container.ServerMiddlewareProvider, error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewGateway(authenticator, skip)
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
	m := factory.NewClient()
	provider.RegisterClientMiddleware("propagation", m)
	provider.RegisterServerMiddleware("propagation", middleware.Noop())
	log.NewHelper(app.Logger()).Infof("registered %+v client middlewares", provider.Names())
	return provider, nil
}

// ProvideGatewaySkipper creates a skipper for the gateway.
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

// ProvideAuthorizer creates the Casbin authorizer.
func ProvideAuthorizer(app *runtime.App, c *conf.Config, adapter *data.CasbinAdapter, w *watcher.Watcher) (*casbin.Authorizer, error) {
	securityConfig := c.GetBootstrap().GetSecurity()
	if securityConfig == nil {
		return nil, errors.New("security configuration not found")
	}
	authzConfig := securityConfig.GetAuthz()
	if authzConfig == nil {
		return nil, errors.New("authz configuration not found")
	}

	casbinConfig, _, err := configutil.Normalize(authzConfig.GetActive(), authzConfig.GetDefault(), authzConfig.GetConfigs())
	if err != nil {
		return nil, fmt.Errorf("failed to normalize Casbin authorizer configuration: %w", err)
	}
	if casbinConfig == nil || casbinConfig.GetCasbin() == nil || casbinConfig.GetType() == "casbin" {
		return nil, errors.New("casbin authorizer configuration not found")
	}
	opts, err := casbin.NewOptions(casbinConfig, casbin.WithPolicyAdapter(adapter), casbin.WithWatcher(w))
	if err != nil {
		return nil, err
	}

	return casbin.New(opts, app.Logger())
}

// ProvideWatcher creates a new casbin watcher.
func ProvideWatcher(app *runtime.App, c *conf.Config) (*watcher.Watcher, error) {
	brokerConfig := c.GetBrokers()
	if brokerConfig == nil {
		return nil, errors.New("broker configuration not found")
	}

	brokerUrl := brokerConfig.GetDefault().GetUrl()
	if brokerUrl == "" {
		return nil, errors.New("broker url not found")
	}

	w, err := watcher.NewWatcher(app.Context(), brokerUrl)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// ProvideServiceMiddlewares creates backend-specific middlewares.
func ProvideServiceMiddlewares(app *runtime.App, authorizer *casbin.Authorizer, skip security.Skipper) (container.ServerMiddlewareProvider, error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewBackend(authorizer, skip)
	provider.RegisterServerMiddleware("authz", m)
	provider.RegisterClientMiddleware("authz", middleware.Noop())
	log.NewHelper(app.Logger()).Infof("registered %+v middlewares", provider.Names())
	return provider, nil
}

// ProvideSkipper creates a skipper for the backend.
func ProvideSkipper(app *runtime.App, _ *conf.Config) security.Skipper {
	adminSkipper := skip.Principal(func(principal security.Principal) bool {
		helper := log.NewHelper(log.With(app.Logger()))
		pid := strconv.Itoa(int(data.SystemUserID))
		if principal.GetID() == pid {
			helper.Infof("skip admin checker: %s", pid)
			return true
		}
		return false
	})
	skips := make(map[string]struct{})
	for _, v := range policies {
		if v.Name == "public" || v.Name == "jwt-auth" {
			skips[v.ServiceMethod] = struct{}{}
		}
	}
	pathSkipper := func(ctx context.Context, req security.Request) bool {
		helper := log.NewHelper(log.With(app.Logger(), "kind", req.Kind(), "operation", req.GetOperation(), "method", req.GetMethod(), "path", req.GetRouteTemplate()))
		if _, ok := skips[req.GetOperation()]; ok {
			helper.Infof("skip checker: %s", req.GetOperation())
			return true
		}
		return false
	}
	return skip.Composite(adminSkipper, pathSkipper)
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

// ProvidePublisher creates a Watermill message.Publisher based on broker configuration.
func ProvidePublisher(c *conf.Config) (*nats.Publisher, error) {
	brokerConfig := c.GetBrokers()
	if brokerConfig == nil {
		return nil, errors.New("broker configuration not found")
	}

	brokerUrl := brokerConfig.GetDefault().GetUrl()
	if brokerUrl == "" {
		return nil, errors.New("broker url not found")
	}

	wmLogger := watermill.NewStdLogger(false, false)

	publisher, err := nats.NewPublisher(
		nats.PublisherConfig{
			URL: brokerUrl,
		},
		wmLogger,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS publisher: %w", err)
	}
	return publisher, nil
}
