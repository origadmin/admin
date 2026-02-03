package providers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	"github.com/google/wire"

	watcher "github.com/origadmin/casbin-watcher/v3"
	_ "github.com/origadmin/casbin-watcher/v3/drivers/nats"
	contribsecurity "github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/contrib/security/credential"
	securitymiddleware "github.com/origadmin/contrib/security/middleware"
	authzmiddleware "github.com/origadmin/contrib/security/middleware/authz"
	"github.com/origadmin/contrib/security/request"
	"github.com/origadmin/contrib/security/skip"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/extensions/configutil"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/security"
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
	"origadmin/application/admin/internal/helpers/pubsub"
)

const (
	// policyNamePublic defines the policy name for publicly accessible endpoints.
	policyNamePublic = "public"
	// policyNameAuthN defines the policy name for endpoints that require JWT authentication only.
	policyNameAuthN = "authn"
)

var (
	factory = securitymiddleware.NewFactory()
	// gatewaySkipMap contains service methods that should be skipped by the gateway authn middleware.
	gatewaySkipMap = make(map[string]struct{})
	// backendSkipMap contains service methods that should be skipped by the backend authz middleware.
	backendSkipMap = make(map[string]struct{})
)

func init() {
	// Pre-filter policies at startup to create fast lookup maps for skippers,
	// using locally defined constants for correctness.
	ps := security.RegisteredPolicies()
	for _, p := range ps {
		if p.Name == policyNamePublic {
			gatewaySkipMap[p.ServiceMethod] = struct{}{}
			backendSkipMap[p.ServiceMethod] = struct{}{}
		} else if p.Name == policyNameAuthN {
			backendSkipMap[p.ServiceMethod] = struct{}{}
		}
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
	wire.Bind(new(authz.Reloader), new(*casbin.Authorizer)), // Bind Authorizer to Reloader interface
	ProvideWatcher,
	ProvideAuthenticator,                                        // Provides *jwt.Authenticator
	wire.Bind(new(credential.Creator), new(*jwt.Authenticator)), // Binds the interface
	ProvideServiceMiddlewares,
	ProvideClientMiddlewares,
	ProvideSkipper,
	ProvidePublisher,
	pubsub.NewWatermillLogger,
	//wire.Bind(new(broker.Publisher), new(*nats.Publisher)), // Binds the interface
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
func ProvideGatewayMiddlewares(app *runtime.App, authenticator *jwt.Authenticator, skip contribsecurity.Skipper) (container.ServerMiddlewareProvider, error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewGateway(authenticator, skip, log.WithLogger(app.Logger()))
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
	if casbinConfig == nil || casbinConfig.GetCasbin() == nil || casbinConfig.GetType() != "casbin" {
		return nil, errors.New("casbin authorizer configuration not found")
	}
	// We remove WithWatcher(w) from here to handle it manually and explicitly.
	opts, err := casbin.NewOptions(casbinConfig, casbin.WithPolicyAdapter(adapter), casbin.WithLogger(app.Logger()))
	if err != nil {
		return nil, err
	}

	authorizer, err := casbin.New(opts, app.Logger())
	if err != nil {
		return nil, err
	}

	// Set up watcher for Stage 2 of the two-stage update flow:
	// After business data is synced to casbin_rule, watcher.Update() broadcasts
	// the update to all auth service instances, which then call LoadPolicy().
	if w != nil {
		enforcer := authorizer.GetEnforcer()
		helper := log.NewHelper(log.With(app.Logger(), "module", "casbin.watcher.setup"))

		// 1. Explicitly set the watcher on the enforcer instance.
		if err := enforcer.SetWatcher(w); err != nil {
			helper.Errorf("Failed to explicitly set watcher for enforcer: %v", err)
			return nil, fmt.Errorf("failed to explicitly set watcher for enforcer: %w", err)
		}
		helper.Info("Watcher explicitly set on Casbin Enforcer.")

		// 2. Explicitly set the callback to a function that logs and then reloads.
		// This is Stage 2: when a watcher update is received, reload policies from database.
		if err := w.SetUpdateCallback(func(msg string) {
			callbackHelper := log.NewHelper(log.With(app.Logger(), "module", "casbin.watcher.callback"))
			callbackHelper.Infof("Stage 2: Policy update notification received: %s. Reloading policies from database...", msg)
			if err := enforcer.LoadPolicy(); err != nil {
				callbackHelper.Errorf("Stage 2 failed: Failed to reload policy after watcher update: %v", err)
			} else {
				callbackHelper.Info("Stage 2 completed: Policy reloaded successfully from database to memory.")
			}
		}); err != nil {
			helper.Errorf("Failed to explicitly set watcher callback: %v", err)
			return nil, fmt.Errorf("failed to explicitly set watcher callback: %w", err)
		}
		helper.Info("Watcher callback explicitly set for Stage 2 policy reload.")
	} else {
		log.NewHelper(app.Logger()).Warn("Watcher is nil, two-stage update flow will not work. Only local instance will have updated policies.")
	}

	return authorizer, nil
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

func ruleSpec(ctx context.Context, p contribsecurity.Principal, req contribsecurity.Request) authz.RuleSpec {
	return authz.RuleSpec{
		Domain:   p.GetDomain(),
		Resource: req.GetOperation(),
		Action:   "ANY",
	}
}

// ProvideServiceMiddlewares creates backend-specific middlewares.
func ProvideServiceMiddlewares(app *runtime.App, authorizer *casbin.Authorizer, skip contribsecurity.Skipper) (container.ServerMiddlewareProvider, error) {
	provider, err := app.MiddlewareProvider()
	if err != nil {
		return nil, err
	}
	m := factory.NewBackend(authorizer, skip, log.WithLogger(app.Logger()), authzmiddleware.WithRuleSpec(ruleSpec))
	provider.RegisterServerMiddleware("authz", m)
	provider.RegisterClientMiddleware("authz", middleware.Noop())
	log.NewHelper(app.Logger()).Infof("registered %+v middlewares", provider.Names())
	return provider, nil
}

// ProvideSkipper creates a skipper for the backend.
func ProvideSkipper(app *runtime.App, _ *conf.Config) contribsecurity.Skipper {
	adminSkipper := skip.Principal(func(principal contribsecurity.Principal) bool {
		helper := log.NewHelper(log.With(app.Logger()))
		pid := strconv.Itoa(int(data.SystemUserID))
		if principal.GetID() == pid {
			helper.Infof("skip admin checker: %s", pid)
			return true
		}
		return false
	})
	pathSkipper := func(ctx context.Context, req contribsecurity.Request) bool {
		helper := log.NewHelper(log.With(app.Logger(), "kind", req.Kind(), "operation", req.GetOperation(), "method", req.GetMethod(), "path", req.GetRouteTemplate()))
		// Use the pre-filtered map for a fast lookup.
		if _, ok := backendSkipMap[req.GetOperation()]; ok {
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
func ProvidePublisher(c *conf.Config, wmLogger watermill.LoggerAdapter) (broker.Publisher, error) {
	brokerConfig := c.GetBrokers()
	if brokerConfig == nil {
		return nil, errors.New("broker configuration not found")
	}

	brokerUrl := brokerConfig.GetDefault().GetUrl()
	if brokerUrl == "" {
		return nil, errors.New("broker url not found")
	}

	// Parse broker URL to check for JetStream configuration
	// The URL format should be: nats://host:port?jetstream=true
	// or nats://host:port/topic?jetstream=true (topic is ignored for publisher)
	publisherConfig := nats.PublisherConfig{
		URL: brokerUrl,
	}

	// Check if JetStream is enabled in the URL
	if brokerConfig.GetDefault().GetType() == "nats" {
		// Use same JetStream configuration as watermill server
		// This ensures publisher and subscriber are compatible
		if strings.Contains(brokerUrl, "jetstream=true") {
			publisherConfig.JetStream = nats.JetStreamConfig{
				Disabled: false,
			}
		}
	}

	publisher, err := pubsub.NewPublisher(publisherConfig, wmLogger)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS publisher: %w", err)
	}
	return publisher, nil
}
