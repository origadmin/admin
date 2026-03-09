/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"

	contribsecurity "github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn/jwt"
	securityauthz "github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/contrib/security/authz/casbin"
	securitymiddleware "github.com/origadmin/contrib/security/middleware"
	"github.com/origadmin/contrib/security/middleware/authz"
	middlewarev1 "github.com/origadmin/runtime/api/gen/go/config/middleware/v1"
	"github.com/origadmin/runtime/contracts/component"
	"github.com/origadmin/runtime/engine"
	enginecontext "github.com/origadmin/runtime/engine/context"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	_ "github.com/origadmin/runtime/middleware"
)

var (
	// middlewareFactory is the global security middleware factory.
	middlewareFactory = securitymiddleware.NewFactory()
)

// NewDefaultMiddleware is the engine provider for standard runtime middlewares.
// It bridges to the runtime/middleware factory for both server and client scopes.
func NewDefaultMiddleware(ctx context.Context, h component.Handle) (any, error) {
	cfg, err := comp.AsConfig[middlewarev1.Middleware](h)
	if err != nil {
		return nil, err
	}

	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))

	// Instantiate middleware based on the current scope
	if h.Locator().Scope() == component.ClientScope {
		m, ok := middleware.NewClient(cfg, log.WithLogger(logger))
		if ok {
			return m, nil
		}
	} else {
		m, ok := middleware.NewServer(cfg, log.WithLogger(logger))
		if ok {
			return m, nil
		}
	}

	return nil, nil
}

// NewAuthnMiddleware is the specific provider for Authn middleware.
func NewAuthnMiddleware(ctx context.Context, h component.Handle) (any, error) {
	if h.Name() != "authn" {
		return nil, nil
	}
	authnInst, err := comp.GetDefault[*jwt.Authenticator](ctx, h.Locator().In(CategoryAuthn))
	if err != nil {
		return nil, err
	}

	// Dynamically get the skipper with the same tag as the middleware
	skipper, _ := comp.GetDefault[contribsecurity.Skipper](ctx, h.Locator().In(component.CategorySkipper,
		engine.WithInTags(GatewayTag)))

	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))

	return middlewareFactory.NewAuthnGateway(authnInst, skipper, authz.WithLogger(logger)), nil
}

func ruleSpec(_ enginecontext.Context, p contribsecurity.Principal, req contribsecurity.Request) securityauthz.RuleSpec {
	return securityauthz.RuleSpec{Domain: p.GetDomain(), Resource: req.GetOperation(), Action: "ANY"}
}

// NewAuthzMiddleware is the specific provider for Authz middleware.
func NewAuthzMiddleware(ctx context.Context, h component.Handle) (any, error) {
	if h.Name() != "authz" {
		return nil, nil
	}
	authzInst, err := comp.GetDefault[*casbin.Authorizer](ctx, h.Locator().In(CategoryAuthz))
	if err != nil {
		return nil, err
	}

	// Dynamically get the skipper with the same tag as the middleware
	skipper, _ := comp.GetDefault[contribsecurity.Skipper](ctx, h.Locator().In(component.CategorySkipper,
		engine.WithInTags(FeatureTag)))

	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))

	return middlewareFactory.NewAuthzBackend(authzInst, skipper, casbin.WithLogger(logger),
		authz.WithRuleSpec(ruleSpec)), nil
}
