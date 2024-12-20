/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package Option implements the functions, types, and interfaces for the module.
package security

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/goexts/generic/settings"
	"github.com/origadmin/contrib/transport/gins"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	middlewarev1 "github.com/origadmin/runtime/gen/go/middleware/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/middleware"
	"github.com/origadmin/toolkits/security"
)

type Option struct {
	Authenticator   security.Authenticator
	Authorizer      security.Authorizer
	Skipper         func(path string) bool
	Parser          func(c *gin.Context, subject string) (security.UserClaims, error)
	Scheme          string
	HeaderAuthorize string
}

type Setting = func(m *Option)

func New(cfg *configv1.Security, opts ...Setting) gins.HandlerFunc {
	option := settings.Apply(&Option{
		Parser: func(c *gin.Context, subject string) (security.UserClaims, error) {
			return nil, errors.New("not implemented")
		},
		Scheme:          "Bearer",
		HeaderAuthorize: "Authorization",
	}, opts)
	if cfg == nil {
		return EmptyMiddleware
	}
	//ms := middleware.NewClient(cfg)
	authn := NewAuthN(cfg, option)
	authz := NewAuthZ(cfg, option)

	return func(c *gins.Context) {
		if option.Skipper != nil {
			if option.Skipper(c.Request.URL.Path) {
				c.Request = c.Request.WithContext(NewSkipContext(c.Request.Context()))
				c.Next()
				return
			}
		}

		if !c.IsAborted() {
			authn(c)
		}
		if !c.IsAborted() {
			authz(c)
		}
	}
}

func WithAuthenticator(authenticator security.Authenticator) Setting {
	return func(m *Option) {
		m.Authenticator = authenticator
	}
}

func WithAuthorizer(authorizer security.Authorizer) Setting {
	return func(m *Option) {
		m.Authorizer = authorizer
	}
}

func WithSkipper(paths []string) Setting {
	log.Debugf("Setting up skipper with paths: %v", paths)
	return func(m *Option) {
		log.Debugf("Configuring skipper for option: %+v", m)
		m.Skipper = func(path string) bool {
			log.Debugf("Checking if path '%s' should be skipped", path)
			for _, p := range paths {
				if strings.HasPrefix(path, p) {
					log.Debugf("Path '%s' matches skipper path '%s', skipping", path, p)
					return true
				}
			}
			log.Debugf("Path '%s' does not match any skipper paths, not skipping", path)
			return false
		}
	}
}

//func WithMiddlewares(middlewares []middleware.Middleware) Setting {
//	return func(m *Option) {
//		m.Middlewares = middlewares
//	}
//}

func OuterMiddlewares(ms ...middleware.Middleware) func(middleware.Handler) middleware.Handler {
	return func(handler middleware.Handler) middleware.Handler {
		for i := range ms {
			ms[i](handler)
		}
		return handler
	}
}

// AdapterMiddleware is a Gin middleware that adapts to Kratos middleware
func AdapterMiddleware(handler middleware.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		tr, ok := transport.FromServerContext(c.Request.Context())
		if !ok {
			return
		}
		// Call the Kratos middleware handler
		resp, err := handler(c.Request.Context(), tr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

func NewClient(cfg *middlewarev1.Middleware, ss ...middleware.OptionSetting) []middleware.Middleware {
	return middleware.NewClient(cfg, ss...)
}

var EmptyMiddleware = func(c *gins.Context) { c.Next() }
