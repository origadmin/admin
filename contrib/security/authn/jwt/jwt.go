/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package jwt implements the functions, types, and interfaces for the module.
package jwt

import (
	"context"
	"errors"
	"maps"
	"time"

	"github.com/dchest/uniuri"
	"github.com/goexts/generic/settings"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	securityv1 "github.com/origadmin/runtime/gen/go/security/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/security"
)

const (
	defaultIssuerDomain      = "localhost"
	defaultExpirationAccess  = 2 * 60 * time.Minute
	defaultExpirationRefresh = 14 * 24 * time.Hour
)

// Tokenizer is a struct that implements the Tokenizer interface.
type Tokenizer struct {
	Option            *Option
	keyFunc           func(token *jwtv5.Token) (any, error)
	signingMethod     jwtv5.SigningMethod
	expirationAccess  time.Duration
	expirationRefresh time.Duration
}

func (obj *Tokenizer) ParseClaims(_ context.Context, tokenStr string) (security.Claims, error) {
	log.Debugf("Authenticating token string: %s", tokenStr)
	// Parse the token string.
	log.Debugf("Parsing token string")
	jwtToken, err := obj.parseToken(tokenStr)

	// If the token is nil, return an error.
	if jwtToken == nil {
		log.Errorf("Failed to parse token: token is nil")
		return nil, ErrInvalidToken
	}

	// If there is an error, return the appropriate error.
	if err != nil {
		log.Errorf("Error parsing token: %v", err)
		switch {
		case errors.Is(err, jwtv5.ErrTokenMalformed):
			log.Debugf("Token is malformed")
			return nil, ErrTokenMalformed
		case errors.Is(err, jwtv5.ErrTokenSignatureInvalid):
			log.Debugf("Token signature is invalid")
			return nil, ErrTokenSignatureInvalid
		case errors.Is(err, jwtv5.ErrTokenExpired) || errors.Is(err, jwtv5.ErrTokenNotValidYet):
			log.Debugf("Token is expired or not valid yet")
			return nil, ErrTokenExpired
		default:
			log.Debugf("Unknown error parsing token")
			return nil, ErrInvalidToken
		}
	}

	// If the token is not valid, return an error.
	if !jwtToken.Valid {
		log.Errorf("Token is not valid")
		return nil, ErrInvalidToken
	}

	// If the signing method is not supported, return an error.
	if jwtToken.Method != obj.signingMethod {
		log.Errorf("Unsupported signing method: %s", jwtToken.Method)
		return nil, ErrUnsupportedSigningMethod
	}

	// If the claims are nil, return an error.
	if jwtToken.Claims == nil {
		log.Errorf("Claims are nil")
		return nil, ErrInvalidClaims
	}

	// Convert the claims to security.Claims.
	log.Debugf("Converting claims to security.Claims")
	securityClaims, err := ToClaims(jwtToken.Claims, obj.Option.extraClaims)
	if err != nil {
		log.Errorf("Error converting claims: %v", err)
		return nil, err
	}
	log.Debugf("Authentication successful")
	return securityClaims, nil
}

func (obj *Tokenizer) DestroyRefreshToken(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func (obj *Tokenizer) CreateRefreshClaims(_ context.Context, id string) (security.Claims, error) {
	expiration := obj.expirationRefresh
	now := time.Now()
	// Create a new claims object with the base claims and the user ID.
	claims := &SecurityClaims{
		Claims: &securityv1.Claims{
			Sub:    id,
			Iss:    obj.Option.issuer,
			Aud:    obj.Option.audience,
			Exp:    now.Add(expiration).Unix(),
			Nbf:    now.Unix(),
			Iat:    now.Unix(),
			Jti:    obj.generateJTI(),
			Scopes: make(map[string]bool),
		},
		Extra: make(map[string]string),
	}

	// Add the extra keys to the claims.
	claims.Extra = maps.Clone(obj.Option.extraClaims)

	// If the token is scoped, add the scope to the claims.
	if obj.Option.scoped {
		claims.Claims.Scopes = maps.Clone(obj.Option.scopes)
	}

	return claims, nil
}

func (obj *Tokenizer) CreateClaims(_ context.Context, id string) (security.Claims, error) {
	expiration := obj.expirationAccess
	now := time.Now()
	// Create a new claims object with the base claims and the user ID.
	claims := &SecurityClaims{
		Claims: &securityv1.Claims{
			Sub:    id,
			Iss:    obj.Option.issuer,
			Aud:    obj.Option.audience,
			Exp:    now.Add(expiration).Unix(),
			Nbf:    now.Unix(),
			Iat:    now.Unix(),
			Jti:    obj.generateJTI(),
			Scopes: make(map[string]bool),
		},
		Extra: make(map[string]string),
	}

	// Add the extra keys to the claims.
	claims.Extra = maps.Clone(obj.Option.extraClaims)

	// If the token is scoped, add the scope to the claims.
	if obj.Option.scoped {
		claims.Claims.Scopes = maps.Clone(obj.Option.scopes)
	}

	return claims, nil
}

func (obj *Tokenizer) Validate(ctx context.Context, tokenStr string) (bool, error) {
	// Authenticate the token string.
	_, err := obj.ParseClaims(ctx, tokenStr)
	// If there is an error, return false and the error.
	if err != nil {
		return false, err
	}
	// Otherwise, return true.
	return true, nil
}

// CreateToken creates a token string from the claims.
func (obj *Tokenizer) CreateToken(ctx context.Context, claims security.Claims) (string, error) {
	// Create a new token with the claims.
	jwtToken := jwtv5.NewWithClaims(obj.signingMethod, ClaimsToJwtClaims(claims))

	// Generate the token string.
	tokenStr, err := obj.generateToken(jwtToken)
	if err != nil || tokenStr == "" {
		return "", err
	}
	return tokenStr, nil
}

// parseToken parses the token string and returns the token.
func (obj *Tokenizer) parseToken(token string) (*jwtv5.Token, error) {
	// If the key function is nil, return an error.
	if obj.keyFunc == nil {
		return nil, ErrMissingKeyFunc
	}
	// If the extra keys are nil, parse the token with the key function.
	if len(obj.Option.extraClaims) == 0 && !obj.Option.scoped {
		return jwtv5.ParseWithClaims(token, &jwtv5.RegisteredClaims{}, obj.keyFunc)
	}

	// Otherwise, parse the token with the key function and the extra keys.
	return jwtv5.Parse(token, obj.keyFunc)
}

// generateToken generates a signed token string from the token.
func (obj *Tokenizer) generateToken(jwtToken *jwtv5.Token) (string, error) {
	// If the key function is nil, return an error.
	if obj.keyFunc == nil {
		return "", ErrMissingKeyFunc
	}

	// Get the key from the key function.
	key, err := obj.keyFunc(jwtToken)
	if err != nil {
		return "", ErrGetKeyFailed
	}

	// Generate the token string.
	strToken, err := jwtToken.SignedString(key)
	if err != nil {
		return "", ErrSignTokenFailed
	}

	return strToken, nil
}

func (obj *Tokenizer) generateJTI() string {
	if !obj.Option.enabledJTI {
		return ""
	}
	if obj.Option.genJTI != nil {
		return obj.Option.genJTI()
	}
	// Encode the random byte slice in base64.
	return uniuri.New()
}

func (obj *Tokenizer) WithConfig(config *configv1.AuthNConfig_JWTConfig) error {
	// If the signing key is empty, return an error.
	signingKey := config.SigningKey
	if signingKey == "" && (obj.signingMethod == nil || obj.keyFunc == nil) {
		return errors.New("signing key is empty")
	}

	// Get the signing method and key function from the signing key.
	signingMethod, keyFunc, err := getSigningMethodAndKeyFunc(config.Algorithm, config.SigningKey)
	if err != nil {
		return err
	}
	obj.signingMethod = signingMethod
	obj.keyFunc = keyFunc
	if config.ExpireTime > 0 {
		obj.expirationAccess = time.Duration(config.ExpireTime) * time.Second
	}
	if config.RefreshTime > 0 {
		obj.expirationRefresh = time.Duration(config.RefreshTime) * time.Second
	}

	return nil
}

// NewTokenizer creates a new Tokenizer.
func NewTokenizer(cfg *configv1.Security, ss ...Setting) (security.RefreshTokenizer, error) {
	// Get the JWT config from the security config.
	config := cfg.GetAuthn().GetJwt()
	if config == nil {
		return nil, errors.New("authenticator jwt config is empty")
	}
	option := settings.Apply(&Option{
		issuer: defaultIssuerDomain,
	}, ss)
	tokenizer := &Tokenizer{
		Option:            option,
		expirationAccess:  defaultExpirationAccess,
		expirationRefresh: defaultExpirationRefresh,
	}
	err := tokenizer.WithConfig(config)
	if err != nil {
		return nil, err
	}
	return tokenizer, nil
}

var _ security.RefreshTokenizer = (*Tokenizer)(nil)
