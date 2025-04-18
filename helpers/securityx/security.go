/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package securityx implements the functions, types, and interfaces for the module.
package securityx

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/transport"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	msecurity "github.com/origadmin/runtime/agent/middleware/security"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/contrib/security/authn/jwt"
	"origadmin/application/admin/contrib/security/authz/casbin"
	"origadmin/application/admin/internal/configs"
)

func NewAuthenticator(bootstrap *configs.Bootstrap, ss ...jwt.Setting) (security.Authenticator, error) {
	tokenizer, err := jwt.NewTokenizer(bootstrap.GetSecurity(), ss...)
	if err != nil {
		return nil, err
	}

	authenticator := jwt.NewAuthenticator(tokenizer)
	return authenticator, nil
}

func NewTokenizer(bootstrap *configs.Bootstrap, ss ...jwt.Setting) (security.Tokenizer, error) {
	tokenizer, err := jwt.NewTokenizer(bootstrap.GetSecurity(), ss...)
	if err != nil {
		return nil, err
	}
	return tokenizer, nil
}

func NewAuthorizer(bootstrap *configs.Bootstrap, ss ...casbin.Setting) (security.Authorizer, error) {
	authorizer, err := casbin.NewAuthorizer(bootstrap.GetSecurity(), ss...)
	if err != nil {
		return nil, err
	}
	return authorizer, nil
}

type DataProvider interface {
	QueryRoles(ctx context.Context, subject string) ([]string, error)
	QueryPermissions(ctx context.Context, subject string) ([]string, error)
}

type SecurityConfig struct {
	// Authenticator is the authenticator used for the authorization header.
	Authenticator security.Authenticator
	// Authorizer is the authorizer used for the authorization header.
	Authorizer security.Authorizer
}

type Option = func(config *SecurityConfig)

func WithAuthenticator(authenticator security.Authenticator) Option {
	return func(config *SecurityConfig) {
		config.Authenticator = authenticator
	}
}

func WithAuthorizer(authorizer security.Authorizer) Option {
	return func(config *SecurityConfig) {
		config.Authorizer = authorizer
	}
}

type SecurityBridge struct {
	// TokenSource is the type of the token.
	TokenSource security.TokenSource
	// Scheme is the scheme used for the authorization header.
	Scheme security.Scheme
	// AuthenticationHeader is the header used for the authorization header.
	AuthenticationHeader string
	// Authenticator is the authenticator used for the authorization header.
	Authenticator security.Authenticator
	// Authorizer is the authorizer used for the authorization header.
	Authorizer security.Authorizer
	// SkipKey is the key used to skip authentication.
	SkipKey string
	// PublicPaths are the public paths that do not require authentication.
	PublicPaths []string
	// Provider is the role/permission data from the database.
	Provider DataProvider
	// Skipper is the function used to skip authentication.
	Skipper func(string) bool
	// IsRoot is the function used to check if the request is root.
	IsRoot func(ctx context.Context, claims security.Claims) bool
	// TokenParser is the parser used to parse the token from the context.
	TokenParser func(ctx context.Context) string
	// PolicyParser is the parser used to parse the policy from the context.
	PolicyParser func(ctx context.Context, claims security.Claims) (security.Policy, error)
}

func (obj SecurityBridge) SkipFromContext(ctx context.Context) (context.Context, bool) {
	if msecurity.IsSkipped(ctx, obj.SkipKey) {
		log.Debugf("NewAuthN: skipping request due to skip key")
		return ctx, true
	}
	if obj.Skipper == nil {
		log.Debugf("NewAuthNServer: skipper is nil")
		return ctx, false
	}
	if tr, ok := transport.FromServerContext(ctx); ok {
		log.Debugf("NewAuthNServer ServerContext: checking skipper for operation: %+v", tr.Operation())
		if obj.Skipper(tr.Operation()) {
			log.Debugf("NewAuthNServer: skipping request")
			ctx := msecurity.WithSkipContextServer(msecurity.NewSkipContext(ctx), obj.SkipKey)
			return ctx, true
		}
	}
	if tr, ok := transport.FromClientContext(ctx); ok {
		log.Debugf("NewAuthNServer ClientContext: checking skipper for operation: %+v", tr.Operation())
		if obj.Skipper(tr.Operation()) {
			log.Debugf("NewAuthNServer: skipping request")
			ctx := msecurity.WithSkipContextClient(msecurity.NewSkipContext(ctx), obj.SkipKey)
			return ctx, true
		}
	}
	return ctx, false
}

func (obj SecurityBridge) schemeString() string {
	return obj.Scheme.String()
}
func tokenParser(ctx context.Context, fns []func(ctx context.Context) string) string {
	for i := range fns {
		if s := fns[i](ctx); s != "" {
			return s
		}
	}
	return ""
}
func (obj SecurityBridge) aggregateTokenParsers(outer ...func(ctx context.Context) string) func(ctx context.Context) string {
	fns := []func(ctx context.Context) string{
		security.TokenFromContext,
	}
	for i := range outer {
		if outer[i] == nil {
			continue
		}
		fns = append(fns, outer[i])
	}
	return func(ctx context.Context) string {
		return tokenParser(ctx, fns)
	}
}

func (obj SecurityBridge) Build() middleware.KMiddleware {
	if obj.TokenParser == nil {
		obj.TokenParser = obj.aggregateTokenParsers(
			//FromTransportClient(obj.AuthenticationHeader, obj.Scheme.String()),
			FromTransportServer(obj.AuthenticationHeader, obj.Scheme.String()),
		)
	}
	return func(handler middleware.KHandler) middleware.KHandler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Debugf("NewAuthN: handling request: %+v", req)
			if ctx, ok := obj.SkipFromContext(ctx); ok {
				return handler(ctx, req)
			}
			log.Debugf("NewAuthN: parsing token from context")
			token := obj.TokenParser(ctx)
			if token == "" {
				log.Errorf("NewAuthN: missing token, returning error")
				return nil, msecurity.ErrInvalidToken
			}

			log.Debugf("NewAuthN: authenticating token")
			claims, err := obj.Authenticator.Authenticate(ctx, token)
			if err != nil {
				log.Errorf("NewAuthN: authentication failed: %s", err.Error())
				return nil, err
			}
			log.Debugf("NewAuthN: setting claims to context user-id: %+v", claims.GetSubject())

			ctx = AppendUserID(ctx, claims.GetSubject())
			log.Debugf("NewAuthN: found context metadata: %+v", ctx)
			if obj.IsRoot(ctx, claims) {
				ctx = security.WithRootContext(ctx)
				log.Debugf("NewAuthN: setting root to context")
				return handler(ctx, req)
			}

			policy, err := obj.PolicyParser(ctx, claims)
			if err != nil {
				return nil, err
			}
			if ok, err := obj.Authorizer.Authorized(ctx, policy, policy.GetAction(), policy.GetObject()); err != nil {
				log.Errorf("NewAuthN: authorization failed")
				return nil, msecurity.ErrInvalidAuthorization
			} else if !ok {
				log.Errorf("NewAuthN: authorization check failed")
				return nil, msecurity.ErrInvalidAuthorization
			} else {
				log.Debugf("NewAuthN: authorization successful, proceeding with request")
			}
			ctx = obj.TokenTo(ctx, token)
			log.Debugf("NewAuthN: calling next handler")
			return handler(ctx, req)
		}
	}
}

func (obj SecurityBridge) TokenTo(ctx context.Context, token string) context.Context {
	return msecurity.TokenToContext(ctx, obj.TokenSource, obj.schemeString(), token)
}
func (obj SecurityBridge) policyParser(ctx context.Context, claims security.Claims) (security.Policy, error) {
	if obj.PolicyParser != nil {
		return obj.PolicyParser(ctx, claims)
	}
	log.Debugf("PolicyParser: parsing policy for subject: %s", claims.GetSubject())
	roles, err := obj.Provider.QueryRoles(ctx, claims.GetSubject())
	if err != nil {
		log.Errorf("PolicyParser: failed to query roles for subject: %s, error: %s", claims.GetSubject(), err.Error())
		return nil, err
	}
	log.Debugf("PolicyParser: queried roles for subject: %s, roles: %v", claims.GetSubject(), roles)

	permissions, err := obj.Provider.QueryPermissions(ctx, claims.GetSubject())
	if err != nil {
		log.Errorf("PolicyParser: failed to query permissions for subject: %s, error: %s", claims.GetSubject(), err.Error())
		return nil, err
	}

	req, ok := transhttp.RequestFromServerContext(ctx)
	if !ok {
		log.Errorf("PolicyParser: failed to get request from server context, error: %s", msecurity.ErrInvalidToken.Error())
		return nil, msecurity.ErrInvalidToken
	}
	policy := security.RegisteredPolicy{
		Subject:     claims.GetSubject(),
		Object:      req.URL.Path,
		Action:      req.Method,
		Domain:      claims.GetIssuer(),
		Roles:       roles,
		Permissions: permissions,
	}
	log.Debugf("PolicyParser: created policy for subject: %s, policy: %+v", claims.GetSubject(), policy)
	return &policy, nil
}

func FromTransportClient(authorize string, scheme string) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		if tr, ok := transport.FromClientContext(ctx); ok {
			token := tr.RequestHeader().Get(authorize)
			splits := strings.SplitN(token, " ", 2)
			if len(splits) > 1 && strings.EqualFold(splits[0], scheme) {
				return splits[1]
			}
		}
		return ""
	}
}

func FromTransportServer(authorize string, scheme string) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		if tr, ok := transport.FromServerContext(ctx); ok {
			token := tr.RequestHeader().Get(authorize)
			splits := strings.SplitN(token, " ", 2)
			if len(splits) > 1 && strings.EqualFold(splits[0], scheme) {
				return splits[1]
			}
		}
		return ""
	}
}

func DefaultBridge() *SecurityBridge {
	bridge := SecurityBridge{
		TokenSource:          security.TokenSourceHeader,
		Scheme:               security.SchemeBearer,
		AuthenticationHeader: security.HeaderAuthorize,
		SkipKey:              msecurity.MetadataSecuritySkipKey,
		PublicPaths:          nil,
		Skipper: func(path string) bool {
			return false
		},
		IsRoot: func(ctx context.Context, claims security.Claims) bool {
			return claims.GetSubject() == "root" || claims.GetSubject() == "admin"
		},
		TokenParser: nil,
	}
	return &bridge
}
