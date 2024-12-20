/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/origadmin/runtime/agent/middleware/security"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
)

func NewAuthZ(config *configv1.Security, option *Option) gin.HandlerFunc {
	log.Debugf("NewAuthZ: creating new server Option with config %+v", config)
	if option.Authorizer == nil {
		log.Errorf("NewAuthZ: authorizer is nil")
		return EmptyMiddleware
	}
	log.Debugf("NewAuthZ: authorizer is not nil, proceeding with Option creation")
	return func(c *gin.Context) {
		log.Debugf("NewAuthZ: handling request %+v", c.Request)
		if IsSkipped(c) {
			log.Debugf("NewAuthZ: skipping request due to skip key")
			c.Next()
			return
		}

		var (
			allowed bool
			err     error
		)

		claims := security.ClaimsFromContext(c)
		if claims == nil {
			log.Errorf("NewAuthZ: claims are nil")
			c.AbortWithError(401, errors.New("claims are nil"))
			return
		}
		log.Debugf("NewAuthZ: claims are not nil, subject: %s", claims.GetSubject())
		if option.Parser == nil {
			log.Errorf("NewAuthZ: parser is nil")
			c.AbortWithError(401, errors.New("parser is nil"))
			return
		}
		userClaims, err := option.Parser(c, claims.GetSubject())
		if err != nil {
			log.Errorf("NewAuthZ: error parsing user claims: %v", err)
			c.AbortWithError(401, fmt.Errorf("error parsing user claims: %v", err))
			return
		}
		log.Debugf("NewAuthZ: user claims: %+v", userClaims)

		if userClaims.GetSubject() == "" || userClaims.GetAction() == "" || userClaims.GetObject() == "" {
			log.Errorf("NewAuthZ: invalid user claims")
			c.AbortWithError(401, errors.New("invalid user claims"))
			return
		}

		//var project []string
		//if domains := claims.GetDomain(); domains != nil {
		//	project = domains
		//}
		// todo add domain project
		log.Debugf("NewAuthZ: checking authorization for user claims %+v", userClaims)
		allowed, err = option.Authorizer.Authorized(c, userClaims)
		if err != nil {
			log.Errorf("NewAuthZ: authorization error: %v", err)
			c.AbortWithError(401, fmt.Errorf("authorization error: %v", err))
			return
		}
		if !allowed {
			log.Errorf("NewAuthZ: authorization denied")
			c.AbortWithError(403, errors.New("authorization denied"))
			return
		}

		log.Debugf("NewAuthZ: authorization successful, proceeding with request")
		c.Next()
		return
	}
}
