/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/casbin/casbin/v3/persist"

	_ "github.com/origadmin/casbin-watcher/v3/drivers/nats"
	authnv1 "github.com/origadmin/contrib/api/gen/go/security/authn/v1"
	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	contribsecurity "github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn/jwt"
	securitycasbin "github.com/origadmin/contrib/security/authz/casbin"
	"github.com/origadmin/contrib/security/request"
	"github.com/origadmin/contrib/security/skip"
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/config/middleware/v1"
	"github.com/origadmin/runtime/contracts/component"
	"github.com/origadmin/runtime/engine"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/helpers/configutil"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/runtime/security"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/data/entity/ent"

	// Side-effect imports for policy registration
	_ "origadmin/application/admin/api/v1/services/admin"
	_ "origadmin/application/admin/api/v1/services/filemanager"
	_ "origadmin/application/admin/api/v1/services/identity"
	_ "origadmin/application/admin/api/v1/services/notification"
	_ "origadmin/application/admin/api/v1/services/objectstore"
	_ "origadmin/application/admin/api/v1/services/system"

	"origadmin/application/admin/internal/data"
	"github.com/origadmin/runtime"
)

const (
	// Sub-categories for project-specific infrastructure
	CategoryAuthn     component.Category = "infrastructure/authn"
	CategoryAuthz     component.Category = "infrastructure/authz"
	CategoryWatcher   component.Category = "infrastructure/watcher"
	CategoryPublisher component.Category = "infrastructure/publisher"
	CategoryEnt       component.Category = "infrastructure/ent"
)

var (
	gatewaySkipMap = make(map[string]struct{})
	backendSkipMap = make(map[string]struct{})
)

const (
	policyNamePublic = "public"
	policyNameAuthN  = "authn"

	GatewayTag = "gateway"
	FeatureTag = "feature"
)

func init() {
	fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] providers.init() called\n")
	// 1. Pre-filter policies
	ps := security.RegisteredPolicies()
	for _, p := range ps {
		if p.Name == policyNamePublic {
			gatewaySkipMap[p.ServiceMethod] = struct{}{}
			backendSkipMap[p.ServiceMethod] = struct{}{}
		} else if p.Name == policyNameAuthN {
			backendSkipMap[p.ServiceMethod] = struct{}{}
		}
	}

	// 2. Register Authn Factory
	engine.Register(CategoryAuthn,
		func(ctx context.Context, h component.Handle) (any, error) {
			cfg, err := comp.AsConfig[authnv1.Authenticator](h)
			if err != nil {
				return nil, err
			}
			if cfg.GetType() == "jwt" {
				jwtOpts, err := jwt.NewOptions(cfg)
				if err != nil {
					return nil, err
				}
				logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))
				return jwt.New(jwtOpts, logger)
			}
			return nil, fmt.Errorf("engine: unsupported authn type: %s", cfg.GetType())
		},
		engine.WithResolverOption(func(source any, _ component.Category) (*component.ModuleConfig, error) {
			if b, ok := source.(*confpb.Bootstrap); ok && b.GetSecurity() != nil && b.GetSecurity().GetAuthn() != nil {
				authn := b.GetSecurity().GetAuthn()

				// Apply authoritative normalization logic: Default -> Active -> First
				def, configs, err := configutil.Normalize(authn.GetActive(), authn.GetDefault(), authn.GetConfigs())
				if err != nil {
					return nil, err
				}

				res := &component.ModuleConfig{Active: extractName(def)}
				for _, cfg := range configs {
					if name := extractName(cfg); name != "" {
						res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: cfg})
					}
				}
				return res, nil
			}
			return nil, nil
		}))

	// 3. Register Ent Database (Data Layer)
	engine.Register(CategoryEnt, NewEnt,
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{
				Entries: []component.ConfigEntry{{Name: "default", Value: nil}},
				Active:  "default",
			}, nil
		}))

	// 4. Register Casbin Adapter (Service Support)
	engine.Register(component.CategoryStorage,
		func(ctx context.Context, h component.Handle) (any, error) {
			dbInst, err := comp.GetDefault[*ent.Database](ctx, h.Locator().In(CategoryEnt))
			if err != nil {
				return nil, err
			}
			logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))
			return data.NewAdapter(ctx, dbInst, logger)
		},
		engine.WithTag("casbin"),
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: "casbin", Value: nil}}, Active: "casbin"}, nil
		}))

	// 5. Register Project Authorizer
	engine.Register(CategoryAuthz,
		func(ctx context.Context, h component.Handle) (any, error) {
			cfg, err := comp.AsConfig[authzv1.Authorizer](h)
			if err != nil {
				return nil, err
			}
			if cfg.GetType() == "casbin" {
				inst, _ := comp.GetDefault[persist.Watcher](ctx, h.Locator().In(CategoryWatcher))
				logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))
				casbinOpts := []securitycasbin.Option{securitycasbin.WithLogger(logger), securitycasbin.WithWatcher(inst)}
				copts, err := securitycasbin.NewOptions(cfg, casbinOpts...)
				if err != nil {
					return nil, err
				}
				return securitycasbin.New(copts, logger)
			}
			return nil, fmt.Errorf("engine: unsupported authz type: %s", cfg.GetType())
		},
		engine.WithResolverOption(func(source any, _ component.Category) (*component.ModuleConfig, error) {
			if b, ok := source.(*confpb.Bootstrap); ok && b.GetSecurity() != nil && b.GetSecurity().GetAuthz() != nil {
				authz := b.GetSecurity().GetAuthz()

				// Apply authoritative normalization logic: Default -> Active -> First
				def, configs, err := configutil.Normalize(authz.GetActive(), authz.GetDefault(), authz.GetConfigs())
				if err != nil {
					return nil, err
				}

				res := &component.ModuleConfig{Active: extractName(def)}
				for _, cfg := range configs {
					if name := extractName(cfg); name != "" {
						res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: cfg})
					}
				}
				return res, nil
			}
			return nil, nil
		}))

	// 6. Register Captcha (Security Layer)
	engine.Register(component.CategorySecurity, NewCaptcha)

	// 7. Register NATS Publisher
	engine.Register(CategoryPublisher, NewPublisher,
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: "default", Value: root}}}, nil
		}))

	// 8. Register NATS Watcher
	engine.Register(CategoryWatcher, NewWatcher,
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: "default", Value: root}}}, nil
		}))

	registerSecurityComponents()
	registerMiddlewares()
}

// ProvideAuthenticator bridges the engine-managed Authenticator to wire.
func ProvideAuthenticator(app *runtime.App) (*jwt.Authenticator, error) {
	return comp.GetDefault[*jwt.Authenticator](app.Context(), app.Container().In(CategoryAuthn))
}

// ProvideAuthorizer bridges the engine-managed Authorizer to wire.
func ProvideAuthorizer(app *runtime.App) (*securitycasbin.Authorizer, error) {
	return comp.GetDefault[*securitycasbin.Authorizer](app.Context(), app.Container().In(CategoryAuthz))
}

func registerSecurityComponents() {
	// Register Gateway Skipper
	engine.Register(component.CategorySkipper,
		func(ctx context.Context, h component.Handle) (any, error) {
			return contribsecurity.Skipper(func(ctx context.Context, req contribsecurity.Request) bool {
				if req, err := request.NewFromServerContext(ctx); err == nil {
					if _, ok := gatewaySkipMap[req.GetOperation()]; ok {
						return true
					}
				}
				return false
			}), nil
		},
		engine.WithTag(GatewayTag),
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: "default", Value: nil}}}, nil
		}))

	// Register Backend Skipper
	engine.Register(component.CategorySkipper,
		func(ctx context.Context, h component.Handle) (any, error) {
			adminSkipper := skip.Principal(func(principal contribsecurity.Principal) bool {
				id := data.GetSystemUserID()
				if id != 0 && principal.GetID() == strconv.FormatInt(id, 10) {
					return true
				}
				return false
			})
			pathSkipper := func(ctx context.Context, req contribsecurity.Request) bool {
				if _, ok := backendSkipMap[req.GetOperation()]; ok {
					return true
				}
				return false
			}
			return skip.Composite(adminSkipper, pathSkipper), nil
		},
		engine.WithTag(FeatureTag),
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: "default", Value: nil}}}, nil
		}))
}

func registerMiddlewares() {
	// Unified Middleware Resolver: Aggregates all sources and ensures system mandatory ones (propagation)
	middlewareResolver := func(source any, _ component.Category) (*component.ModuleConfig, error) {
		res := &component.ModuleConfig{}
		b, ok := source.(*confpb.Bootstrap)
		if !ok {
			fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] middlewareResolver: type mismatch, source=%T\n", source)
			return nil, nil
		}

		fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] middlewareResolver called\n")

		// 1. Load standard middlewares from middlewares.yaml
		if b.GetMiddlewares() != nil {
			for _, cfg := range b.GetMiddlewares().GetConfigs() {
				name := extractName(cfg)
				if name != "" {
					fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] middlewareResolver: adding %s from config\n", name)
					res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: cfg})
				}
			}
		}

		// 2. Load Authn if present
		if b.GetSecurity() != nil && b.GetSecurity().GetAuthn() != nil {
			authn := b.GetSecurity().GetAuthn()
			def, _, err := configutil.Normalize(authn.GetActive(), authn.GetDefault(), authn.GetConfigs())
			if err == nil {
				fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] middlewareResolver: adding authn\n")
				res.Entries = append(res.Entries, component.ConfigEntry{Name: "authn", Value: def})
			}
		}

		// 3. Load Authz if present
		if b.GetSecurity() != nil && b.GetSecurity().GetAuthz() != nil {
			authz := b.GetSecurity().GetAuthz()
			def, _, err := configutil.Normalize(authz.GetActive(), authz.GetDefault(), authz.GetConfigs())
			if err == nil {
				fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] middlewareResolver: adding authz\n")
				res.Entries = append(res.Entries, component.ConfigEntry{Name: "authz", Value: def})
			}
		}

		// 4. Mandatory: Always include propagation for clients/servers
		fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] middlewareResolver: FORCE adding propagation\n")
		res.Entries = append(res.Entries, component.ConfigEntry{
			Name:  "propagation",
			Value: &middlewarev1.Middleware{Name: "propagation", Type: "propagation", Enabled: true},
		})

		return res, nil
	}

	// Combined Provider for ALL Middlewares to avoid stack interference
	unifiedProvider := func(ctx context.Context, h component.Handle) (any, error) {
		name := h.Name()
		tag := h.Tag()
		fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] unifiedProvider: name=%s, tag=%s\n", name, tag)

		// A. Specialized: Propagation
		if name == "propagation" {
			return NewPropagationMiddleware(ctx, h)
		}

		// B. Specialized: Authn
		if name == "authn" {
			return NewAuthnMiddleware(ctx, h)
		}

		// C. Specialized: Authz
		if name == "authz" {
			return NewAuthzMiddleware(ctx, h)
		}

		// D. General: Default (logging, tracing, etc.)
		return NewDefaultMiddleware(ctx, h)
	}

	// Register the unified provider
	engine.Register(component.CategoryMiddleware, unifiedProvider,
		engine.WithScopes(component.ServerScope, component.ClientScope),
		engine.WithResolverOption(middlewareResolver),
		engine.WithDefaultEntry("propagation"),
	)
}

// NewPropagationMiddleware handles the specialized creation of propagation middleware.
func NewPropagationMiddleware(ctx context.Context, h component.Handle) (any, error) {
	tag := h.Tag()
	scope := h.Locator().Scope()
	fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] NewPropagationMiddleware: tag=%s, scope=%s\n", tag, scope)

	// Logic driven by middlewareFactory extension:
	if scope == component.ServerScope {
		if tag == GatewayTag {
			fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] NewPropagationMiddleware: returning Noop for Gateway Server\n")
			return middleware.Noop(), nil
		}
		fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] NewPropagationMiddleware: returning Backend for Service Server\n")
		return middlewareFactory.NewPropagationBackend(), nil
	}

	fmt.Fprintf(os.Stderr, "[STDOUT_DEBUG] NewPropagationMiddleware: returning Client for %s Client\n", tag)
	return middlewareFactory.NewPropagationClient(), nil
}

func extractName(item any) string {
	if item == nil {
		return ""
	}
	// Use formalized interfaces for identification
	if n, ok := item.(component.Named); ok {
		if name := n.GetName(); name != "" {
			return name
		}
	}
	if t, ok := item.(component.Typed); ok {
		if name := t.GetType(); name != "" {
			return name
		}
	}
	if d, ok := item.(component.Dialectal); ok {
		if name := d.GetDialect(); name != "" {
			return name
		}
	}
	if d, ok := item.(component.Driver); ok {
		if name := d.GetDriver(); name != "" {
			return name
		}
	}
	return ""
}
