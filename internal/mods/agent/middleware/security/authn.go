/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	middlewaresecurity "github.com/origadmin/runtime/agent/middleware/security"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/security"
)

func NewAuthN(config *configv1.Security, option *Option) gin.HandlerFunc {
	log.Debugf("NewAuthN: creating server authenticator Option with config: %+v", config)
	if option == nil || option.Authenticator == nil {
		log.Errorf("NewAuthN: option or authenticator is nil, returning error")
		return EmptyMiddleware
	}

	log.Debugf("NewAuthN: applying defaults and creating token parser")
	//tokenParsers := aggregateTokenParsers(
	//	option.TokenParser,
	//	FromTransportClient(option.HeaderAuthorize, option.Scheme),
	//	FromTransportServer(option.HeaderAuthorize, option.Scheme))
	tokenParser := FromHeader(option.HeaderAuthorize, option.Scheme)
	return func(c *gin.Context) {
		log.Debugf("NewAuthN: handling request: %+v", c.Request)
		if IsSkipped(c) {
			log.Debugf("NewAuthN: skipping request due to skip key")
			c.Next()
			return
		}

		log.Debugf("NewAuthN: parsing token from context")
		var err error
		token := tokenParser(c)
		if token == "" {
			log.Errorf("NewAuthN: missing token, returning error")
			c.AbortWithError(401, errors.New("missing token"))
			return
		}

		log.Debugf("NewAuthN: authenticating token")
		claims, err := option.Authenticator.Authenticate(c, token)
		if err != nil {
			log.Errorf("NewAuthN: authentication failed: %option", err.Error())
			c.AbortWithError(401, fmt.Errorf("authentication failed: %v", err))
			return
		}

		log.Debugf("NewAuthN: setting claims to context")
		c.Request = NewClaimsContext(c, claims)
		log.Debugf("NewAuthN: calling next handler")
		c.Next()
	}
}

func FromHeader(header string, scheme string) func(ctx *gin.Context) string {
	return func(ctx *gin.Context) string {
		token := ctx.GetHeader(header)
		splits := strings.SplitN(token, " ", 2)
		if len(splits) > 1 && strings.EqualFold(splits[0], scheme) {
			return splits[1]
		}
		return ""
	}
}

func NewClaimsContext(c *gin.Context, claims security.Claims) *http.Request {
	return c.Request.WithContext(middlewaresecurity.NewClaimsContext(c.Request.Context(), claims))

}
