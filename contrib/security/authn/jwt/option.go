/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package jwt implements the functions, types, and interfaces for the module.
package jwt

import (
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/origadmin/toolkits/errors"
)

type Option struct {
	enabledJTI    bool
	genJTI        func() string
	issuer        string
	audience      []string
	scoped        bool
	scopes        map[string]bool
	extraClaims   map[string]string
	signingMethod jwtv5.SigningMethod
	keyFunc       func(token *jwtv5.Token) (any, error)
}

// Setting is a function type for setting the Tokenizer.
type Setting = func(*Option)

func (option *Option) ApplyDefaults() error {
	return nil
}

// GetKeyFunc returns a function that retrieves the key for a given token.
// The returned function takes a jwtv5.Token as an argument and returns the key as a string.
func GetKeyFunc(key string) func(token *jwtv5.Token) (any, error) {
	// Return a function that checks if the token's algorithm is empty.
	// If it is, return an error. Otherwise, return the key.
	return func(token *jwtv5.Token) (any, error) {
		if token.Method.Alg() == "" {
			// Return an error if the token's algorithm is empty.
			return nil, ErrInvalidToken
		}
		// Return the key if the token's algorithm is not empty.
		return key, nil
	}
}

// GetKeyFuncWithAlg returns a function that retrieves the key for a given token
// with a specific algorithm.
// The returned function takes a jwtv5.Token as an argument and returns the key as a byte slice.
func GetKeyFuncWithAlg(alg, key string) func(token *jwtv5.Token) (any, error) {
	// Return a function that checks if the token's algorithm matches the provided algorithm.
	// If it does not, return an error. Otherwise, return the key as a byte slice.
	return func(token *jwtv5.Token) (any, error) {
		if token.Method.Alg() == "" || alg != token.Method.Alg() {
			// Return an error if the token's algorithm does not match the provided algorithm.
			return nil, ErrInvalidToken
		}
		// jwtv5 requires the key to be a byte slice.
		return []byte(key), nil
	}
}

// GetAlgorithmSigningMethod returns the signing method for a given algorithm.
func GetAlgorithmSigningMethod(algorithm string) jwtv5.SigningMethod {
	// Use a switch statement to map the algorithm to its corresponding signing method.
	switch algorithm {
	case "HS256":
		// Return the signing method for HS256.
		return jwtv5.SigningMethodHS256
	case "HS384":
		// Return the signing method for HS384.
		return jwtv5.SigningMethodHS384
	case "HS512":
		// Return the signing method for HS512.
		return jwtv5.SigningMethodHS512
	case "RS256":
		// Return the signing method for RS256.
		return jwtv5.SigningMethodRS256
	case "RS384":
		// Return the signing method for RS384.
		return jwtv5.SigningMethodRS384
	case "RS512":
		// Return the signing method for RS512.
		return jwtv5.SigningMethodRS512
	case "ES256":
		// Return the signing method for ES256.
		return jwtv5.SigningMethodES256
	case "ES384":
		// Return the signing method for ES384.
		return jwtv5.SigningMethodES384
	case "ES512":
		// Return the signing method for ES512.
		return jwtv5.SigningMethodES512
	case "EdDSA":
		// Return the signing method for EdDSA.
		return jwtv5.SigningMethodEdDSA
	default:
		// Return nil if the algorithm is not recognized.
		return nil
	}
}

// WithExtraClaims returns a Setting function that sets the extra keys for an Tokenizer.
func WithExtraClaims(extras map[string]string) Setting {
	// Return a function that sets the extra keys for an Tokenizer.
	return func(option *Option) {
		// Set the extra keys for the Tokenizer.
		option.extraClaims = extras
	}
}

// WithSigningMethod returns a Setting function that sets the signing method for an Tokenizer.
// The signing method is used to sign and verify tokens.
func WithSigningMethod(signingMethod jwtv5.SigningMethod) Setting {
	// Return a function that sets the signing method for an Tokenizer.
	return func(option *Option) {
		// Set the signing method for the Tokenizer.
		option.signingMethod = signingMethod
	}
}

// WithKeyFunc returns a Setting function that sets the key function for an Tokenizer.
// The key function is used to retrieve the key for a given token.
func WithKeyFunc(keyFunc func(token *jwtv5.Token) (any, error)) Setting {
	// Return a function that sets the key function for an Tokenizer.
	return func(option *Option) {
		// Set the key function for the Tokenizer.
		option.keyFunc = keyFunc
	}
}

// WithJTI returns a Setting function that sets the JTI generator function for an Tokenizer.
func WithJTI(fn func() string) Setting {
	return func(option *Option) {
		option.genJTI = fn
		option.enabledJTI = true
	}
}

// WithIssuer returns a Setting function that sets the issuer for an Tokenizer.
func WithIssuer(issuer string) Setting {
	return func(option *Option) {
		option.issuer = issuer
	}
}

// WithAudience returns a Setting function that sets the audience for an Tokenizer.
func WithAudience(audience []string) Setting {
	return func(option *Option) {
		option.audience = audience
	}
}

// WithScopes returns a Setting function that sets the scoped flag for an Tokenizer.
// The scoped flag determines whether the Tokenizer should use scoped tokens.
func WithScopes(scopes map[string]bool) Setting {
	return func(option *Option) {
		option.scopes = scopes
		option.scoped = true
	}
}

func getSigningMethodAndKeyFunc(algorithm string, signingKey string) (jwtv5.SigningMethod, func(*jwtv5.Token) (any, error), error) {
	signingMethod := GetAlgorithmSigningMethod(algorithm)
	if signingMethod == nil {
		return nil, nil, errors.New("invalid signing method")
	}

	keyFunc := GetKeyFuncWithAlg(algorithm, signingKey)
	if keyFunc == nil {
		return nil, nil, errors.New("invalid key function")
	}

	return signingMethod, keyFunc, nil
}
