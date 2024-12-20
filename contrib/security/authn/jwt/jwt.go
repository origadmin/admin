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
	securityv1 "github.com/origadmin/runtime/gen/go/security/v1"
	"github.com/origadmin/runtime/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	middlewaresecurity "github.com/origadmin/runtime/agent/middleware/security"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/security"
	"github.com/origadmin/toolkits/storage/cache"
)

const (
	defaultIssuerDomain      = "localhost"
	defaultExpirationAccess  = 2 * 60 * time.Minute
	defaultExpirationRefresh = 14 * 24 * time.Hour
)

// Authenticator is a struct that implements the Authenticator interface.
type Authenticator struct {
	*Option
	// cache is the token cache service.
	cache security.TokenCacheService
}

func (obj *Authenticator) CreateIdentityClaims(_ context.Context, id string, refresh bool) (security.Claims, error) {
	expiration := getExpiration(obj, refresh)
	now := time.Now()
	// Create a new claims object with the base claims and the user ID.
	claims := &SecurityClaims{
		Claims: &securityv1.Claims{
			Sub:    id,
			Iss:    obj.issuer,
			Aud:    obj.audience,
			Exp:    timestamppb.New(now.Add(expiration)),
			Nbf:    timestamppb.New(now),
			Iat:    timestamppb.New(now),
			Jti:    obj.generateJTI(),
			Scopes: make(map[string]bool),
		},
		Extra: make(map[string]string),
	}

	// Add the extra keys to the claims.
	claims.Extra = maps.Clone(obj.extraClaims)

	// If the token is scoped, add the scope to the claims.
	if obj.scoped {
		claims.Claims.Scopes = maps.Clone(obj.scopes)
	}

	return claims, nil
}

func (obj *Authenticator) CreateIdentityClaimsContext(ctx context.Context, tokenType security.TokenType, id string) (context.Context, error) {
	// Create the claims.
	claims, err := obj.CreateIdentityClaims(ctx, id, false)
	if err != nil {
		return ctx, err
	}

	// Add the token to the context.
	ctx = middlewaresecurity.NewClaimsContext(ctx, claims)
	return ctx, nil
}

func (obj *Authenticator) Authenticate(ctx context.Context, tokenStr string) (security.Claims, error) {
	log.Debugf("Authenticating token string: %s", tokenStr)
	// If the token cache service is not nil, validate the token.
	if obj.cache != nil {
		log.Debugf("Validating token using cache service")
		ok, err := obj.cache.Validate(ctx, tokenStr)
		switch {
		case err != nil:
			log.Errorf("Error validating token: %v", err)
			return nil, ErrInvalidToken
		case !ok:
			log.Debugf("Token not found in cache")
			return nil, ErrTokenNotFound
		}
	}
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
	securityClaims, err := ToClaims(jwtToken.Claims, obj.extraClaims)
	if err != nil {
		log.Errorf("Error converting claims: %v", err)
		return nil, err
	}
	log.Debugf("Authentication successful")
	return securityClaims, nil
}

func (obj *Authenticator) AuthenticateContext(ctx context.Context, tokenType security.TokenType) (security.Claims, error) {
	log.Debugf("Entering AuthenticateContext with tokenType: %s", tokenType)
	// Get the token string from the context.
	tokenStr, err := middlewaresecurity.TokenFromTypeContext(ctx, tokenType, obj.schemeString())
	if err != nil {
		log.Errorf("Error getting token from context: %v", err)
	} else if tokenStr == "" {
		log.Debugf("Token string is empty")
	}
	if err != nil || tokenStr == "" {
		log.Errorf("Invalid token or token string is empty")
		return nil, ErrInvalidToken
	}
	log.Debugf("Token string retrieved from context: %s", tokenStr)
	// Authenticate the token string.
	log.Debugf("Authenticating token string")
	claims, err := obj.Authenticate(ctx, tokenStr)
	if err != nil {
		log.Errorf("Error authenticating token: %v", err)
	}
	log.Debugf("Authentication result: %+v", claims)
	return claims, err
}

func (obj *Authenticator) Verify(ctx context.Context, tokenStr string) (bool, error) {
	// Authenticate the token string.
	_, err := obj.Authenticate(ctx, tokenStr)
	// If there is an error, return false and the error.
	if err != nil {
		return false, err
	}
	// Otherwise, return true.
	return true, nil
}

func (obj *Authenticator) VerifyContext(ctx context.Context, tokenType security.TokenType) (bool, error) {
	// Get the token string from the context.
	tokenStr, err := middlewaresecurity.TokenFromTypeContext(ctx, tokenType, obj.schemeString())
	if err != nil || tokenStr == "" {
		return false, ErrInvalidToken
	}
	// Authenticate the token string.
	return obj.Verify(ctx, tokenStr)
}

// schemeString returns the scheme type as a string.
func (obj *Authenticator) schemeString() string {
	return obj.schemeType.String()
}

// CreateToken creates a token string from the claims.
func (obj *Authenticator) CreateToken(ctx context.Context, claims security.Claims) (string, error) {
	// Create a new token with the claims.
	jwtToken := jwtv5.NewWithClaims(obj.signingMethod, ClaimsToJwtClaims(claims))

	// Generate the token string.
	tokenStr, err := obj.generateToken(jwtToken)
	if err != nil || tokenStr == "" {
		return "", err
	}

	// If the token cache service is not nil, store the token.
	if obj.cache != nil {
		// Get the expiration time from the claims.
		exp := time.Duration(claims.GetExpiration().UnixMilli())
		if err := obj.cache.Store(ctx, tokenStr, exp); err != nil {
			return tokenStr, err
		}
	}
	return tokenStr, nil
}

func getExpiration(obj *Authenticator, refresh bool) time.Duration {
	if refresh {
		return obj.expirationRefresh
	}
	return obj.expirationAccess
}

// CreateTokenContext creates a token string from the claims and adds it to the context.
func (obj *Authenticator) CreateTokenContext(ctx context.Context, tokenType security.TokenType, claims security.Claims) (context.Context, error) {
	// Create the token string.
	tokenStr, err := obj.CreateToken(ctx, claims)
	if err != nil {
		return ctx, err
	}
	// Add the token string to the context.
	ctx = middlewaresecurity.TokenToTypeContext(ctx, tokenType, obj.schemeString(), tokenStr)
	return ctx, nil
}

// DestroyToken destroys the token string.
func (obj *Authenticator) DestroyToken(ctx context.Context, tokenStr string) error {
	// If the token cache service is not nil, remove the token.
	if obj.cache != nil {
		err := obj.cache.Remove(ctx, tokenStr)
		if err != nil && !errors.Is(err, cache.ErrNotFound) {
			return err
		}
	}
	return nil
}

// DestroyTokenContext destroys the token string from the context.
func (obj *Authenticator) DestroyTokenContext(ctx context.Context, token security.TokenType) error {
	// Get the token string from the context.
	tokenStr, err := middlewaresecurity.TokenFromTypeContext(ctx, token, obj.schemeString())
	if err != nil || tokenStr == "" {
		return ErrInvalidToken
	}
	// Destroy the token string.
	return obj.DestroyToken(ctx, tokenStr)
}

// Close closes the token cache service.
func (obj *Authenticator) Close(ctx context.Context) error {
	// If the token cache service is not nil, close it.
	if obj.cache != nil {
		return obj.cache.Close(ctx)
	}
	return nil
}

// parseToken parses the token string and returns the token.
func (obj *Authenticator) parseToken(token string) (*jwtv5.Token, error) {
	// If the key function is nil, return an error.
	if obj.keyFunc == nil {
		return nil, ErrMissingKeyFunc
	}
	// If the extra keys are nil, parse the token with the key function.
	if len(obj.extraClaims) == 0 && !obj.scoped {
		return jwtv5.ParseWithClaims(token, &jwtv5.RegisteredClaims{}, obj.keyFunc)
	}

	// Otherwise, parse the token with the key function and the extra keys.
	return jwtv5.Parse(token, obj.keyFunc)
}

// generateToken generates a signed token string from the token.
func (obj *Authenticator) generateToken(jwtToken *jwtv5.Token) (string, error) {
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

func (obj *Authenticator) generateJTI() string {
	if !obj.enabledJTI {
		return ""
	}
	if obj.genJTI != nil {
		return obj.genJTI()
	}
	// Encode the random byte slice in base64.
	return uniuri.New()
}

// NewAuthenticator creates a new Authenticator.
func NewAuthenticator(cfg *configv1.Security, ss ...Setting) (security.Authenticator, error) {
	// Get the JWT config from the security config.
	config := cfg.GetAuthn().GetJwt()
	if config == nil {
		return nil, errors.New("authenticator jwt config is empty")
	}
	option := settings.Apply(&Option{
		schemeType:        security.SchemeBearer,
		expirationAccess:  defaultExpirationAccess,
		expirationRefresh: defaultExpirationRefresh,
		issuer:            defaultIssuerDomain,
	}, ss)
	err := option.WithConfig(config)
	if err != nil {
		return nil, err
	}
	return &Authenticator{
		Option: option,
		cache:  option.cache,
	}, nil
}

var _ security.Authenticator = (*Authenticator)(nil)
