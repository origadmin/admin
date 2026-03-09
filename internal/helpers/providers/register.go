/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"
	"fmt"
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
	"github.com/origadmin/runtime/log"
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
				res := &component.ModuleConfig{Active: authn.GetActive()}
				if d := authn.GetDefault(); d != nil {
					name := d.GetName()
					res.Entries = append(res.Entries, component.ConfigEntry{Name: component.DefaultName, Value: d})
					if name != "" {
						res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: d})
					}
					if res.Active == "" {
						res.Active = name
					}
				}
				for _, cfg := range authn.GetConfigs() {
					name := cfg.GetName()
					if name == "" {
						name = cfg.GetType()
					}
					res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: cfg})
				}
				// Fallback: If no Active/Default but we have entries, pick the first one as default
				if res.Active == "" && len(res.Entries) > 0 {
					res.Active = res.Entries[0].Name
				}
				// Map logic Default if not already mapped
				if res.Active != "" {
					foundDefault := false
					for _, e := range res.Entries {
						if e.Name == component.DefaultName {
							foundDefault = true
							break
						}
					}
					if !foundDefault {
						for _, e := range res.Entries {
							if e.Name == res.Active {
								res.Entries = append(res.Entries, component.ConfigEntry{Name: component.DefaultName, Value: e.Value})
								break
							}
						}
					}
				}
				return res, nil
			}
			return nil, nil
		}))

	// 3. Register Ent Database (Data Layer)
	// This provider is defined locally in data.go
	engine.Register(CategoryEnt, NewEnt,
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			// Ent database is a logical wrapper, it just needs to be activated.
			return &component.ModuleConfig{
				Entries: []component.ConfigEntry{{Name: component.DefaultName, Value: nil}},
				Active:  component.DefaultName,
			}, nil
		}))

	// 4. Register Casbin Adapter (Service Support)
	engine.Register(component.CategoryStorage,
		func(ctx context.Context, h component.Handle) (any, error) {
			// DEPENDENCY: Now correctly depends on CategoryEnt (The Data Layer)
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
				res := &component.ModuleConfig{Active: authz.GetActive()}
				if d := authz.GetDefault(); d != nil {
					name := d.GetName()
					res.Entries = append(res.Entries, component.ConfigEntry{Name: component.DefaultName, Value: d})
					if name != "" {
						res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: d})
					}
					if res.Active == "" {
						res.Active = name
					}
				}
				for _, cfg := range authz.GetConfigs() {
					name := cfg.GetName()
					if name == "" {
						name = cfg.GetType()
					}
					res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: cfg})
				}
				// Fallback: If no Active/Default but we have entries, pick the first one as default
				if res.Active == "" && len(res.Entries) > 0 {
					res.Active = res.Entries[0].Name
				}
				// Map logic Default if not already mapped
				if res.Active != "" {
					foundDefault := false
					for _, e := range res.Entries {
						if e.Name == component.DefaultName {
							foundDefault = true
							break
						}
					}
					if !foundDefault {
						for _, e := range res.Entries {
							if e.Name == res.Active {
								res.Entries = append(res.Entries, component.ConfigEntry{Name: component.DefaultName, Value: e.Value})
								break
							}
						}
					}
				}
				return res, nil
			}
			return nil, nil
		}))

	// 6. Register Captcha (Security Layer)
	// Defined in captcha.go
	engine.Register(component.CategorySecurity, NewCaptcha)

	// 7. Register NATS Publisher
	// Defined in publisher.go
	engine.Register(CategoryPublisher, NewPublisher,
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: component.DefaultName, Value: root}}}, nil
		}))

	// 8. Register NATS Watcher
	// Defined in watcher.go
	engine.Register(CategoryWatcher, NewWatcher,
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: component.DefaultName, Value: root}}}, nil
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
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: component.DefaultName, Value: nil}}}, nil
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
			return &component.ModuleConfig{Entries: []component.ConfigEntry{{Name: component.DefaultName, Value: nil}}}, nil
		}))
}

func registerMiddlewares() {
	// Comprehensive Middleware Resolver: Aggregates all sources
	middlewareResolver := func(source any, _ component.Category) (*component.ModuleConfig, error) {
		res := &component.ModuleConfig{}
		b, ok := source.(*confpb.Bootstrap)
		if !ok {
			return nil, nil
		}

		// 1. Load standard middlewares from middlewares.yaml
		if b.GetMiddlewares() != nil {
			for _, cfg := range b.GetMiddlewares().GetConfigs() {
				name := cfg.GetName()
				if name == "" {
					name = cfg.GetType()
				}
				res.Entries = append(res.Entries, component.ConfigEntry{Name: name, Value: cfg})
			}
		}

		// 2. Load Authn if present
		if b.GetSecurity() != nil && b.GetSecurity().GetAuthn() != nil {
			authn := b.GetSecurity().GetAuthn()
			if d := authn.GetDefault(); d != nil {
				res.Entries = append(res.Entries, component.ConfigEntry{Name: "authn", Value: d})
			} else if len(authn.GetConfigs()) > 0 {
				res.Entries = append(res.Entries, component.ConfigEntry{Name: "authn", Value: authn.GetConfigs()[0]})
			}
		}

		// 3. Load Authz if present
		if b.GetSecurity() != nil && b.GetSecurity().GetAuthz() != nil {
			authz := b.GetSecurity().GetAuthz()
			if d := authz.GetDefault(); d != nil {
				res.Entries = append(res.Entries, component.ConfigEntry{Name: "authz", Value: d})
			} else if len(authz.GetConfigs()) > 0 {
				res.Entries = append(res.Entries, component.ConfigEntry{Name: "authz", Value: authz.GetConfigs()[0]})
			}
		}

		// 4. Always include propagation for clients/servers if security is active
		if b.GetSecurity() != nil {
			res.Entries = append(res.Entries, component.ConfigEntry{
				Name:  "propagation",
				Value: &middlewarev1.Middleware{Name: "propagation", Type: "propagation", Enabled: true},
			})
		}

		return res, nil
	}

	// Register all providers for the SAME category with the SAME comprehensive resolver
	opts := []engine.RegisterOption{
		engine.WithScopes(component.ServerScope, component.ClientScope),
		engine.WithResolverOption(middlewareResolver),
	}

	engine.Register(component.CategoryMiddleware, NewDefaultMiddleware, opts...)
	engine.Register(component.CategoryMiddleware, NewAuthnMiddleware, append(opts, engine.WithTag(GatewayTag))...)
	engine.Register(component.CategoryMiddleware, NewAuthzMiddleware, append(opts, engine.WithTag(FeatureTag))...)

	// 4. Explicitly Register Propagation Middleware
	// Ensure it has a dedicated entry in the engine even if not explicitly in config
	engine.Register(component.CategoryMiddleware, NewDefaultMiddleware,
		engine.WithScopes(component.ServerScope, component.ClientScope),
		engine.WithResolverOption(func(root any, _ component.Category) (*component.ModuleConfig, error) {
			return &component.ModuleConfig{
				Entries: []component.ConfigEntry{{
					Name:  "propagation",
					Value: &middlewarev1.Middleware{Name: "propagation", Type: "propagation", Enabled: true},
				}},
				Active: "propagation",
			}, nil
		}))
}
