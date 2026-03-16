/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package middleware

import (
	"context"
	"fmt"

	contribsecurity "github.com/origadmin/contrib/security"
	"github.com/origadmin/contrib/security/authn/jwt"
	securityauthz "github.com/origadmin/contrib/security/authz"
	securitycasbin "github.com/origadmin/contrib/security/authz/casbin"
	securitymiddleware "github.com/origadmin/contrib/security/middleware"
	"github.com/origadmin/contrib/security/middleware/authz"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/contracts/component"
	enginecontext "github.com/origadmin/runtime/engine/context"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
)

const (
	MiddlewareAuthn       = "authn"
	MiddlewareAuthz       = "authz"
	MiddlewarePropagation = "propagation"
	GatewayTag            = "gateway"
	FeatureTag            = "feature"
)

const (
	CategoryAuthn component.Category = "authn"
	CategoryAuthz component.Category = "authz"
)

var (
	middlewareFactory = securitymiddleware.NewFactory()
)

// NewPropagationMiddleware is the provider for principal propagation middleware.
func NewPropagationMiddleware(ctx context.Context, h component.Handle) (any, error) {
	if h.Name() != MiddlewarePropagation {
		return nil, nil
	}
	tag := h.Tag()
	scope := h.Scope()
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(runtime.CategoryLogger))

	if scope == runtime.ServerScope {
		if tag == GatewayTag {
			return middleware.Noop(), nil
		}
		return middlewareFactory.NewPropagationBackend(log.WithLogger(logger)), nil
	}
	return middlewareFactory.NewPropagationClient(log.WithLogger(logger)), nil
}

// NewAuthnMiddleware is the specific provider for Authn middleware.
func NewAuthnMiddleware(ctx context.Context, h component.Handle) (any, error) {
	if h.Name() != MiddlewareAuthn {
		return nil, nil
	}
	tag := h.Tag()
	if tag != GatewayTag {
		return middleware.Noop(), nil
	}

	authnH := h.Locator().In(CategoryAuthn)
	authnInst, err := comp.GetDefault[*jwt.Authenticator](ctx, authnH)
	if err != nil {
		return nil, fmt.Errorf("engine: failed to get authn instance: %w", err)
	}

	if authnInst == nil {
		return nil, fmt.Errorf("engine: authn instance is nil")
	}

	// Correct functional domain matching:
	// Use WithInTags(tag) to narrow down the space to the correct functional domain,
	// then use GetDefault to retrieve the default skipper within that domain.
	skipperH := h.Locator().In(runtime.CategorySkipper).WithInTags(GatewayTag)
	skipper, _ := comp.GetDefault[contribsecurity.Skipper](ctx, skipperH)
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(runtime.CategoryLogger))

	return middlewareFactory.NewAuthnGateway(authnInst, skipper,
		log.WithLogger(logger)), nil
}

func ruleSpec(_ enginecontext.Context, p contribsecurity.Principal, req contribsecurity.Request) securityauthz.RuleSpec {
	return securityauthz.RuleSpec{Domain: "*", Resource: req.GetOperation(), Action: "ANY"}
}

// NewAuthzMiddleware is the specific provider for Authz middleware.
func NewAuthzMiddleware(ctx context.Context, h component.Handle) (any, error) {
	if h.Name() != MiddlewareAuthz {
		return nil, nil
	}
	tag := h.Tag()
	if tag != FeatureTag {
		return middleware.Noop(), nil
	}

	authzH := h.Locator().In(CategoryAuthz)
	authzInst, err := comp.GetDefault[*securitycasbin.Authorizer](ctx, authzH)
	if err != nil {
		return nil, fmt.Errorf("engine: failed to get authz instance: %w", err)
	}

	if authzInst == nil {
		return nil, fmt.Errorf("engine: authz instance is nil")
	}

	// Narrow down to the specific functional domain (tag) and get the default skipper.
	skipperH := h.Locator().In(runtime.CategorySkipper).WithInTags(FeatureTag)
	skipper, _ := comp.GetDefault[contribsecurity.Skipper](ctx, skipperH)
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(runtime.CategoryLogger))

	return middlewareFactory.NewAuthzBackend(authzInst, skipper,
		log.WithLogger(logger),
		authz.WithRuleSpec(ruleSpec)), nil
}
