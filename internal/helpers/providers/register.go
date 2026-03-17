package providers

import (
	_ "github.com/origadmin/casbin-watcher/v3/drivers/nats"
	"github.com/origadmin/contrib/security/authn"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/contracts/component"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/helpers/captcha"
	"origadmin/application/admin/internal/helpers/debounce"
	"origadmin/application/admin/internal/helpers/middleware"
	"origadmin/application/admin/internal/helpers/pubsub"
)

const (
	CategoryAuthn     component.Category = "authn"
	CategoryAuthz     component.Category = "authz"
	CategoryWatcher   component.Category = "watcher"
	CategoryPublisher component.Category = "publisher"
	CategoryEnt                          = data.CategoryEnt
)

const (
	GatewayTag = "gateway"
	FeatureTag = "feature"
)

func init() {
	// 1. Register filter policies
	middleware.RegisterFilterPolicies()

	// 2. Register Infrastructure
	registerInfrastructure()

	// 3. Register Security and Middlewares
	registerSecurityComponents()

	// 4. Register Middlewares
	registerMiddlewares()
}

func registerInfrastructure() {
	// Register Ent Database
	runtime.Register(CategoryEnt, data.NewEnt,
		runtime.WithResolver(data.EntResolver))

	// Register Authn
	runtime.Register(CategoryAuthn, jwt.Provider,
		runtime.WithResolver(authn.ConfigResolver),
	)

	// Register Casbin Adapter
	runtime.Register(runtime.CategoryStorage, data.CasbinAdapterProvider,
		runtime.WithResolver(data.CasbinAdapterResolver))

	// Register Project Authorizer
	runtime.Register(CategoryAuthz, casbin.Provider,
		runtime.WithResolver(authz.ConfigResolver),
		runtime.WithRequirement(middleware.AuthzRequirementResolver),
	)

	// Register Captcha
	runtime.Register(runtime.CategorySecurity, captcha.NewCaptchaHandle, runtime.WithResolver(captcha.Resolver))

	// Register NATS Publisher
	runtime.Register(CategoryPublisher, pubsub.NewPublisherHandle, runtime.WithResolver(pubsub.Resolver))

	// Register NATS Watcher
	runtime.Register(CategoryWatcher, pubsub.NewWatcherHandle, runtime.WithResolver(pubsub.Resolver))

	// Register Debounce
	runtime.Register(debounce.CategoryDebounce, debounce.Provider, runtime.WithResolver(debounce.Resolver))
}

func registerSecurityComponents() {
	// Register Gateway Skipper
	runtime.Register(runtime.CategorySkipper, middleware.GatewaySkipperProvider,
		runtime.WithTag(GatewayTag),
	)

	// Register Backend Skipper
	runtime.Register(runtime.CategorySkipper, middleware.BackendSkipperProvider,
		runtime.WithTag(FeatureTag),
	)
}

func registerMiddlewares() {
	opts := []runtime.RegisterOption{
		runtime.WithScopes(runtime.ServerScope, runtime.ClientScope),
	}

	runtime.Register(runtime.CategoryMiddleware, middleware.NewAuthnMiddleware, append(opts,
		runtime.WithEntries(middleware.NameAuthnMiddleware), runtime.WithTag(GatewayTag))...)
	runtime.Register(runtime.CategoryMiddleware, middleware.NewAuthzMiddleware, append(opts,
		runtime.WithEntries(middleware.NameAuthzMiddleware), runtime.WithTag(FeatureTag))...)
	runtime.Register(runtime.CategoryMiddleware, middleware.NewPropagationMiddleware, append(opts,
		runtime.WithEntries(middleware.NamePropagationMiddleware))...)
}
