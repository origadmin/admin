/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package jwt

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	middlewaresecurity "github.com/origadmin/runtime/agent/middleware/security"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	securityv1 "github.com/origadmin/runtime/gen/go/security/v1"
	"github.com/origadmin/toolkits/security"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	HeaderAuthorize = "Authorization"
)

type headerCarrier http.Header

func (hc headerCarrier) Get(key string) string { return http.Header(hc).Get(key) }

func (hc headerCarrier) Set(key string, value string) { http.Header(hc).Set(key, value) }

// Add append value to key-values pair.
func (hc headerCarrier) Add(key string, value string) {
	http.Header(hc).Add(key, value)
}

// Values returns a slice of values associated with the passed key.
func (hc headerCarrier) Values(key string) []string {
	return http.Header(hc).Values(key)
}

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}

func newTokenHeader(headerKey string, token string) *headerCarrier {
	header := &headerCarrier{}
	header.Set(headerKey, fmt.Sprintf("%s %s", security.SchemeBearer.String(), token))
	return header
}

type Transport struct {
	kind      transport.Kind
	endpoint  string
	operation string
	reqHeader transport.Header
}

func (tr *Transport) Kind() transport.Kind {
	return tr.kind
}

func (tr *Transport) Endpoint() string {
	return tr.endpoint
}

func (tr *Transport) Operation() string {
	return tr.operation
}

func (tr *Transport) RequestHeader() transport.Header {
	return tr.reqHeader
}

func (tr *Transport) ReplyHeader() transport.Header {
	return nil
}

func generateJwtKey(key, sub string) string {
	mapClaims := jwtv5.MapClaims{}
	mapClaims["sub"] = sub
	claims := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, mapClaims)
	token, _ := claims.SignedString([]byte(key))
	return token
}

var ErrMissingBearerToken = ErrBearerTokenMissing

func TestServer(t *testing.T) {
	testKey := "testKey"

	token := generateJwtKey(testKey, "fly")

	tests := []struct {
		name      string
		ctx       context.Context
		alg       string
		exceptErr error
		key       string
	}{
		{
			name:      "normal",
			ctx:       transport.NewServerContext(context.Background(), &Transport{reqHeader: newTokenHeader(HeaderAuthorize, token)}),
			alg:       "HS256",
			exceptErr: nil,
			key:       testKey,
		},
		{
			name:      "miss token",
			ctx:       transport.NewServerContext(context.Background(), &Transport{reqHeader: headerCarrier{}}),
			alg:       "HS256",
			exceptErr: ErrMissingBearerToken,
			key:       testKey,
		},
		{
			name: "token invalid",
			ctx: transport.NewServerContext(context.Background(), &Transport{
				reqHeader: newTokenHeader(HeaderAuthorize, "12313123"),
			}),
			alg:       "HS256",
			exceptErr: ErrInvalidToken,
			key:       testKey,
		},
		{
			name:      "method invalid",
			ctx:       transport.NewServerContext(context.Background(), &Transport{reqHeader: newTokenHeader(HeaderAuthorize, token)}),
			alg:       "ES384",
			exceptErr: ErrInvalidToken,
			key:       testKey,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testToken security.Claims
			next := func(ctx context.Context, req interface{}) (interface{}, error) {
				t.Log(req)
				testToken = middlewaresecurity.ClaimsFromContext(ctx)
				t.Log(testToken)
				return "reply", nil
			}
			cfg := &configv1.Security{
				Authn: &configv1.AuthNConfig{
					Jwt: &configv1.AuthNConfig_JWTConfig{
						Algorithm:     test.alg,
						SigningKey:    testKey,
						OldSigningKey: "",
						ExpireTime:    nil,
						RefreshTime:   nil,
						CacheName:     "",
					},
				},
			}
			authenticator, err := NewAuthenticator(cfg) //WithKey([]byte(testKey)),
			//WithSigningMethod(test.alg),

			assert.Nil(t, err)
			server, _ := middlewaresecurity.NewAuthN(cfg,
				middlewaresecurity.WithAuthenticator(authenticator),
				middlewaresecurity.WithSkipper())
			handle := server(next)
			ctx := middlewaresecurity.WithSkipContextServer(test.ctx, middlewaresecurity.MetadataSecuritySkipKey)
			_, err2 := handle(ctx, test.name)
			if !errors.Is(test.exceptErr, err2) {
				t.Errorf("except error %v, but got %v", test.exceptErr, err2)
			}
			if test.exceptErr == nil {
				if testToken == nil {
					t.Errorf("except testToken not nil, but got nil")
				}
			}
		})
	}
}

func TestClient(t *testing.T) {
	testKey := "testKey"

	tests := []struct {
		name        string
		expectError error
	}{
		{
			name:        "normal",
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			next := func(ctx context.Context, req interface{}) (interface{}, error) {
				if header, ok := transport.FromClientContext(ctx); ok {
					t.Log("token: ", header.RequestHeader().Get(HeaderAuthorize))
				}
				return "reply", nil
			}
			cfg := &configv1.Security{
				Authn: &configv1.AuthNConfig{
					Jwt: &configv1.AuthNConfig_JWTConfig{
						Algorithm:     "HS256",
						SigningKey:    testKey,
						OldSigningKey: "",
						ExpireTime:    nil,
						RefreshTime:   nil,
						CacheName:     "",
					},
				},
			}
			authenticator, err := NewAuthenticator(cfg)
			assert.Nil(t, err)

			principal := SecurityClaims{
				Claims: &securityv1.Claims{
					Sub:    "user_name",
					Scopes: make(map[string]bool),
				},
			}
			principal.Scopes["local:admin:user_name"] = true
			principal.Scopes["tenant:admin:user_name"] = true
			auth, _ := middlewaresecurity.NewAuthN(cfg,
				middlewaresecurity.WithAuthenticator(authenticator),
			)
			header := newTokenHeader(HeaderAuthorize, generateJwtKey(testKey, "fly"))
			ctx := transport.NewClientContext(context.Background(), &Transport{reqHeader: header})
			handle := auth(next)
			_, err2 := handle(ctx, "ok")
			if !errors.Is(test.expectError, err2) {
				t.Errorf("except error %v, but got %v", test.expectError, err2)
			}
		})
	}
}

func TestAuth(t *testing.T) {
	//cache := memory.NewCache(memory.Selector{CleanupInterval: time.Second})
	//c:=security.WithCache(cache)
	//store := Memory
	ctx := context.Background()
	//middlewaresecurity.WithStorage(store)
	cfg := &configv1.Security{
		Authn: &configv1.AuthNConfig{
			Jwt: &configv1.AuthNConfig_JWTConfig{
				Algorithm:     "HS256",
				SigningKey:    "abc123",
				OldSigningKey: "",
				ExpireTime:    nil,
				RefreshTime:   nil,
				CacheName:     "",
			},
		},
	}
	jwtAuth, err := NewAuthenticator(cfg, WithCache(security.DefaultTokenCacheService()))
	assert.Nil(t, err)
	if err != nil {
		t.Fatal(err)
	}
	userID := "test"
	claims := &securityv1.Claims{
		Sub: userID,
		Iss: "test",
		Aud: []string{"test"},
		Exp: timestamppb.New(time.Now().Add(time.Hour)),
		Nbf: timestamppb.New(time.Now()),
		Iat: timestamppb.New(time.Now()),
		Jti: "not need",
		Scopes: map[string]bool{
			"test": true,
		},
	}
	token, err := jwtAuth.CreateToken(ctx, &SecurityClaims{Claims: claims})
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log("token: ", token)
	resultClaims, err := jwtAuth.Authenticate(ctx, token)
	assert.Nil(t, err)
	fmt.Println("error", err)
	fmt.Println("result_id:", resultClaims.GetSubject())
	fmt.Println("user_id: ", userID)
	assert.Equal(t, userID, resultClaims.GetSubject())
	var ok bool
	ok, err = jwtAuth.Verify(ctx, token)
	assert.Nil(t, err)
	assert.True(t, ok)
	err = jwtAuth.DestroyToken(ctx, token)
	assert.Nil(t, err)
	ok, err = jwtAuth.Verify(ctx, token)
	assert.NotNil(t, err)
	assert.False(t, ok)
	t.Log("token: ", token)
	resultClaims, err = jwtAuth.Authenticate(ctx, token)
	assert.NotNil(t, err)
	assert.EqualError(t, err, ErrTokenNotFound.Error())
	assert.Empty(t, resultClaims)

	err = jwtAuth.Close(ctx)
	assert.Nil(t, err)
}
