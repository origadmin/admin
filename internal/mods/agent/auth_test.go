/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package agent implements the functions, types, and interfaces for the module.
package agent

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/context"
	configv1 "github.com/origadmin/runtime/gen/go/config/v1"
	"github.com/origadmin/runtime/log"
	kslog "github.com/origadmin/slog-kratos"
	"github.com/origadmin/toolkits/security"
	"github.com/stretchr/testify/assert"

	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/configs"
)

var bridge = securityx.SecurityBridge{
	Authenticator:        mockAuthenticator{},
	Authorizer:           mockAuthorizer{},
	Scheme:               security.SchemeBearer,
	AuthenticationHeader: security.HeaderAuthorize,
	Skipper: func(path string) bool {
		fmt.Println("path:", path)
		return path == "/login"
	},
	IsRoot: func(ctx context.Context, claims security.Claims) bool {
		return claims.GetSubject() == "admin"
	},
	Data: mockData{},
	//TokenParser: func(ctx context.Context) string {
	//	if tr, ok := transport.FromServerContext(ctx); ok {
	//		return tr.RequestHeader().Get("Authorization")
	//	}
	//	return ""
	//},
	// Initialize other fields...
}

func init() {
	fmt.Println("debug mode")
	slog.SetLogLoggerLevel(slog.LevelDebug)
	l := log.With(kslog.NewLogger(),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(l)
}

// Test authenticator
func TestAuthenticator(t *testing.T) {
	// Initialize test bootstrap configuration
	cfg := &configs.Bootstrap{
		Security: &configv1.Security{
			Authn: &configv1.AuthNConfig{
				Type: "jwt",
				Jwt: &configv1.AuthNConfig_JWTConfig{
					Algorithm:  "HS512",
					SigningKey: "test_key",
				},
			},
		},
		/* Fill in test configuration */}

	auth, err := securityx.NewAuthenticator(cfg)
	if err != nil {
		t.Fatalf("Failed to create authenticator: %v", err)
	}
	tk, err := securityx.NewTokenizer(cfg)
	if err != nil {
		t.Fatalf("Failed to create tokenizer: %v", err)
	}
	// Test valid token
	t.Run("valid token", func(t *testing.T) {
		cc, err := tk.CreateClaims(context.Background(), "valid_token")
		if err != nil {
			t.Fatalf("Failed to create claims: %v", err)
		}
		ct, err := tk.CreateToken(context.Background(), cc)
		if err != nil {
			t.Fatalf("Failed to create token: %v", err)
		}
		claims, err := auth.Authenticate(context.Background(), ct)
		if err != nil {
			t.Errorf("Expected authentication to succeed, but failed: %v", err)
		}
		t.Logf("Authentication result: %+v", claims)
		assert.Equal(t, "valid_token", claims.GetSubject())
	})

	// Test invalid token
	t.Run("invalid token", func(t *testing.T) {
		_, err := auth.Authenticate(context.Background(), "invalid_token")
		if err == nil {
			t.Error("Expected authentication to fail, but succeeded")
		}
	})
}

// Test security middleware chain
func TestSecurityMiddlewareChain(t *testing.T) {
	// Create test middleware chain
	chain := selector.Server(bridge.Build()).Match(func(ctx context.Context, operation string) bool {
		fmt.Println("operation:", operation)
		return operation == "/protected"
	}).Build()

	// Create test router
	router := transhttp.NewServer()
	router.HandleFunc("/protected", func(writer http.ResponseWriter, request *http.Request) {
		_, err := chain(func(ctx context.Context, req interface{}) (interface{}, error) {
			//debug.PrintStack()
			// Test logic can be added here
			writer.WriteHeader(200)
			writer.Write([]byte("Hello, World!"))
			return nil, nil
		})(testContextWithToken(request), nil)
		if err != nil {
			writer.WriteHeader(401)
			writer.Write([]byte("Authentication failed"))
		}
	})

	// Test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	t.Run("unauthorized request", func(t *testing.T) {
		resp, _ := http.Get(ts.URL + "/protected")
		if resp.StatusCode != 401 {
			t.Errorf("Expected status code 401, but got %d", resp.StatusCode)
		}
	})

	t.Run("authorized request", func(t *testing.T) {
		req, _ := http.NewRequest("GET", ts.URL+"/protected", nil)
		req.Header.Add("Authorization", "Bearer valid_token")

		resp, _ := http.DefaultClient.Do(req)
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
		}
	})
}

// Test skipper logic
func TestSkipperLogic(t *testing.T) {

	testCases := []struct {
		path     string
		expected bool
	}{
		{"/login", true},
		{"/api/secret", false},
	}

	for _, tc := range testCases {
		got := bridge.Skipper(tc.path)
		if got != tc.expected {
			t.Errorf("Path %s expected %v but got %v", tc.path, tc.expected, got)
		}
	}
}

// Test is root logic
func TestIsRoot(t *testing.T) {
	testClaims := &security.RegisteredClaims{
		Subject: "admin",
	}

	if !bridge.IsRoot(context.Background(), testClaims) {
		t.Error("Admin user should be identified as root")
	}
}

// Create test context
func testContextWithToken(req *http.Request) context.Context {
	ctx := context.Background()
	return transport.NewServerContext(ctx, &mockTransport{md: headerCarrier(req.Header), req: req})
}

type mockTransport struct {
	md  headerCarrier
	req *http.Request
}

func (m mockTransport) Kind() transport.Kind {
	//TODO implement me
	panic("implement me")
}

func (m mockTransport) Endpoint() string {
	//TODO implement me
	panic("implement me")
}

func (m mockTransport) Operation() string {
	return "/protected"
}

func (m mockTransport) RequestHeader() transport.Header {
	fmt.Println("request header:", m.md, m.md.Get("Authorization"))
	return m.md
}

func (m mockTransport) ReplyHeader() transport.Header {
	return m.md
}

type headerCarrier http.Header

// Get returns the value associated with the passed key.
func (hc headerCarrier) Get(key string) string {
	return http.Header(hc).Get(key)
}

// Set stores the key-value pair.
func (hc headerCarrier) Set(key string, value string) {
	http.Header(hc).Set(key, value)
}

// Add append value to key-values pair.
func (hc headerCarrier) Add(key string, value string) {
	http.Header(hc).Add(key, value)
}

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of values associated with the passed key.
func (hc headerCarrier) Values(key string) []string {
	return http.Header(hc).Values(key)
}
