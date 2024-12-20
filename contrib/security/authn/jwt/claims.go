/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package jwt implements the functions, types, and interfaces for the module.
package jwt

import (
	"bytes"
	"strings"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
	"google.golang.org/protobuf/types/known/timestamppb"

	securityv1 "github.com/origadmin/runtime/gen/go/security/v1"
	"github.com/origadmin/toolkits/security"
)

var (
	ErrInvalidToken             = securityv1.ErrorSecurityErrorReasonBearerTokenMissing("invalid bearer token")
	ErrTokenNotFound            = securityv1.ErrorSecurityErrorReasonTokenNotFound("token not found")
	ErrTokenMalformed           = securityv1.ErrorSecurityErrorReasonBearerTokenMissing("token malformed")
	ErrTokenSignatureInvalid    = securityv1.ErrorSecurityErrorReasonSignTokenFailed("token signature invalid")
	ErrTokenExpired             = securityv1.ErrorSecurityErrorReasonTokenExpired("token expired")
	ErrTokenNotValidYet         = securityv1.ErrorSecurityErrorReasonTokenExpired("token not valid yet")
	ErrUnsupportedSigningMethod = securityv1.ErrorSecurityErrorReasonUnsupportedSigningMethod("unsupported signing method")
	ErrInvalidClaims            = securityv1.ErrorSecurityErrorReasonInvalidClaims("invalid Claims")
	ErrBearerTokenMissing       = securityv1.ErrorSecurityErrorReasonBearerTokenMissing("bearer token missing")
	ErrSignTokenFailed          = securityv1.ErrorSecurityErrorReasonSignTokenFailed("sign token failed")
	ErrMissingKeyFunc           = securityv1.ErrorSecurityErrorReasonMissingKeyFunc("missing key function")
	ErrGetKeyFailed             = securityv1.ErrorSecurityErrorReasonGetKeyFailed("get key failed")
	ErrInvalidSubject           = securityv1.ErrorSecurityErrorReasonInvalidSubject("invalid subject")
	ErrInvalidIssuer            = securityv1.ErrorSecurityErrorReasonInvalidIssuer("invalid issuer")
	ErrInvalidAudience          = securityv1.ErrorSecurityErrorReasonInvalidAudience("invalid audience")
	ErrInvalidExpiration        = securityv1.ErrorSecurityErrorReasonInvalidExpiration("invalid expiration")
	//ErrInvalidNotBefore         = securityv1.ErrorSecurityErrorReasonInvalidNotBefore("invalid not before")
	//ErrInvalidIssuedAt          = securityv1.ErrorSecurityErrorReasonInvalidIssuedAt("invalid issued at")
)

type SecurityClaims struct {
	*securityv1.Claims
	Extra map[string]string
}

func (s *SecurityClaims) GetSubject() string {
	return s.Claims.Sub
}

func (s *SecurityClaims) GetIssuer() string {
	return s.Claims.Iss
}

func (s *SecurityClaims) GetAudience() []string {
	return s.Claims.Aud
}

func (s *SecurityClaims) GetExpiration() time.Time {
	return s.Claims.Exp.AsTime()
}

func (s *SecurityClaims) GetNotBefore() time.Time {
	return s.Claims.Nbf.AsTime()
}

func (s *SecurityClaims) GetIssuedAt() time.Time {
	return s.Claims.Iat.AsTime()
}

func (s *SecurityClaims) GetJWTID() string {
	return s.Claims.Jti
}

func (s *SecurityClaims) GetExtra() map[string]string {
	return s.Extra
}

func (s *SecurityClaims) GetScopes() map[string]bool {
	return s.Claims.Scopes
}

func ClaimsToJwtClaims(raw security.Claims) jwtv5.Claims {
	mapClaims := jwtv5.MapClaims{
		"sub": raw.GetSubject(),
	}

	if iss := raw.GetIssuer(); iss != "" {
		mapClaims["iss"] = raw.GetIssuer()
	}
	if aud := raw.GetAudience(); len(aud) > 0 {
		mapClaims["aud"] = aud
	}
	if exp := raw.GetExpiration(); !exp.IsZero() {
		mapClaims["exp"] = exp.UnixMilli()
	}

	extras := raw.GetExtra()
	for key, val := range extras {
		mapClaims[key] = val
	}

	var buffer bytes.Buffer
	count := len(raw.GetScopes())
	idx := 0
	for scope := range raw.GetScopes() {
		buffer.WriteString(scope)
		if idx != count-1 {
			buffer.WriteString(" ")
		}
		idx++
	}
	str := buffer.String()
	if len(str) > 0 {
		mapClaims["scope"] = buffer.String()
	}

	return mapClaims
}

func MapToClaims(rawClaims jwtv5.MapClaims, extras map[string]string) (security.Claims, error) {
	//claims := security.claims{
	//	Scopes: make(ScopeSet),
	//}
	claims := &securityv1.Claims{
		Scopes: make(map[string]bool),
	}

	// optional Subject
	if subjectClaim, err := rawClaims.GetSubject(); err == nil {
		claims.Sub = subjectClaim
	} else {
		return nil, ErrInvalidSubject
	}
	// optional Issuer
	if issuerClaim, err := rawClaims.GetIssuer(); err == nil {
		claims.Iss = issuerClaim
	} else {
		return nil, ErrInvalidIssuer
	}
	// optional Audience
	if audienceClaim, err := rawClaims.GetAudience(); err == nil {
		claims.Aud = append(claims.Aud, audienceClaim...)
	} else {
		return nil, ErrInvalidAudience
	}
	// optional Expiration
	if expClaim, err := rawClaims.GetExpirationTime(); err == nil {
		if expClaim != nil {
			claims.Exp = timestamppb.New(expClaim.Time)
		}
	} else {
		return nil, ErrInvalidExpiration
	}
	// optional scopes
	if scopeKey, ok := rawClaims["scope"]; ok {
		if scope, ok := scopeKey.(string); ok {
			scopes := strings.Split(scope, " ")
			for _, s := range scopes {
				claims.Scopes[s] = true
			}
		}
	}

	return &SecurityClaims{
		Claims: claims,
		Extra:  extras,
	}, nil
}

func RegisteredToClaims(rawClaims *jwtv5.RegisteredClaims) (security.Claims, error) {
	Claims := &securityv1.Claims{
		Scopes: make(map[string]bool),
	}

	// optional Subject
	if subjectClaim, err := rawClaims.GetSubject(); err == nil {
		Claims.Sub = subjectClaim
	} else {
		return nil, ErrInvalidSubject
	}
	// optional Issuer
	if issuerClaim, err := rawClaims.GetIssuer(); err == nil {
		Claims.Iss = issuerClaim
	} else {
		return nil, ErrInvalidIssuer
	}
	// optional Audience
	if audienceClaim, err := rawClaims.GetAudience(); err == nil {
		Claims.Aud = append(Claims.Aud, audienceClaim...)
	} else {
		return nil, ErrInvalidAudience
	}
	// optional Expiration
	if expClaim, err := rawClaims.GetExpirationTime(); err == nil {
		if expClaim != nil {
			Claims.Exp = timestamppb.New(expClaim.Time)
		}
	} else {
		return nil, ErrInvalidExpiration
	}
	// optional scopes
	//if scopeKey, ok := rawClaims.Scope["scope"]; ok {
	//	if scope, ok := scopeKey.(string); ok {
	//		scopes := strings.Split(scope, " ")
	//		for _, s := range scopes {
	//			Claims.Scopes[s] = true
	//		}
	//	}
	//}

	return &SecurityClaims{
		Claims: Claims,
	}, nil
}

func ProtoClaimsToClaims(rawClaims *securityv1.Claims) security.Claims {
	return &SecurityClaims{
		Claims: rawClaims,
	}
}

func ToClaims(rawClaims jwtv5.Claims, extras map[string]string) (security.Claims, error) {
	if Claims, ok := rawClaims.(*jwtv5.RegisteredClaims); ok {
		return RegisteredToClaims(Claims)
	}
	if Claims, ok := rawClaims.(jwtv5.MapClaims); ok {
		return MapToClaims(Claims, extras)
	}
	return nil, ErrInvalidClaims
}
