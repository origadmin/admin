/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package server

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TestGenerateAccessToken tests JWT token generation functionality
func TestGenerateAccessToken(t *testing.T) {
	tests := []struct {
		name        string
		objectID    string
		expiresIn   time.Duration
		wantValid   bool
		description string
	}{
		{
			name:        "valid token with 1 hour expiry",
			objectID:    "obj-123",
			expiresIn:   1 * time.Hour,
			wantValid:   true,
			description: "should generate a valid token that passes validation",
		},
		{
			name:        "valid token with 5 minutes expiry",
			objectID:    "obj-456",
			expiresIn:   5 * time.Minute,
			wantValid:   true,
			description: "should generate a valid short-lived token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proxy := &ObjectStoreProxy{
				jwtSecret: "test-secret-key",
			}

			token, err := proxy.GenerateAccessToken(tt.objectID, tt.expiresIn)
			if err != nil {
				t.Fatalf("GenerateAccessToken() error = %v", err)
			}

			if token == "" {
				t.Fatal("GenerateAccessToken() returned empty token")
			}

			// Validate the token
			valid := proxy.validateAccessToken(token, tt.objectID)
			if valid != tt.wantValid {
				t.Errorf("validateAccessToken() = %v, want %v", valid, tt.wantValid)
			}

			// Verify token contains correct object_id claim
			parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				return []byte(proxy.jwtSecret), nil
			})
			if err != nil {
				t.Fatalf("Failed to parse token: %v", err)
			}

			if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
				objID, ok := claims["object_id"].(string)
				if !ok || objID != tt.objectID {
					t.Errorf("Token object_id = %v, want %v", objID, tt.objectID)
				}
			} else {
				t.Error("Invalid token claims")
			}
		})
	}
}

// TestValidateAccessToken tests JWT token validation functionality
func TestValidateAccessToken(t *testing.T) {
	proxy := &ObjectStoreProxy{
		jwtSecret: "test-secret-key",
	}

	// Generate a valid token
	validObjectID := "obj-valid-123"
	validToken, err := proxy.GenerateAccessToken(validObjectID, 1*time.Hour)
	if err != nil {
		t.Fatalf("Failed to generate test token: %v", err)
	}

	tests := []struct {
		name        string
		token       string
		objectID    string
		wantValid   bool
		description string
	}{
		{
			name:        "valid token for correct object",
			token:       validToken,
			objectID:    validObjectID,
			wantValid:   true,
			description: "should accept valid token for matching object",
		},
		{
			name:        "valid token for wrong object",
			token:       validToken,
			objectID:    "obj-wrong-456",
			wantValid:   false,
			description: "should reject token for different object",
		},
		{
			name:        "empty token",
			token:       "",
			objectID:    validObjectID,
			wantValid:   false,
			description: "should reject empty token",
		},
		{
			name:        "invalid token format",
			token:       "invalid.token.format",
			objectID:    validObjectID,
			wantValid:   false,
			description: "should reject malformed token",
		},
		{
			name:        "token with wrong secret",
			token:       generateTokenWithSecret(validObjectID, "wrong-secret", 1*time.Hour),
			objectID:    validObjectID,
			wantValid:   false,
			description: "should reject token signed with wrong secret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := proxy.validateAccessToken(tt.token, tt.objectID)
			if valid != tt.wantValid {
				t.Errorf("validateAccessToken() = %v, want %v", valid, tt.wantValid)
			}
		})
	}
}

// TestExpiredAccessToken tests that expired tokens are rejected
func TestExpiredAccessToken(t *testing.T) {
	proxy := &ObjectStoreProxy{
		jwtSecret: "test-secret-key",
	}

	objectID := "obj-expired-123"

	// Generate a token that expires in the past
	expiredToken, err := proxy.GenerateAccessToken(objectID, -1*time.Hour)
	if err != nil {
		t.Fatalf("Failed to generate expired token: %v", err)
	}

	valid := proxy.validateAccessToken(expiredToken, objectID)
	if valid {
		t.Error("validateAccessToken() should reject expired token")
	}
}

// generateTokenWithSecret creates a JWT token with a custom secret for testing
func generateTokenWithSecret(objectID, secret string, expiresIn time.Duration) string {
	claims := jwt.MapClaims{
		"object_id": objectID,
		"exp":       time.Now().Add(expiresIn).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

// BenchmarkGenerateAccessToken benchmarks token generation performance
func BenchmarkGenerateAccessToken(b *testing.B) {
	proxy := &ObjectStoreProxy{
		jwtSecret: "test-secret-key",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proxy.GenerateAccessToken("obj-123", 1*time.Hour)
	}
}

// BenchmarkValidateAccessToken benchmarks token validation performance
func BenchmarkValidateAccessToken(b *testing.B) {
	proxy := &ObjectStoreProxy{
		jwtSecret: "test-secret-key",
	}

	token, _ := proxy.GenerateAccessToken("obj-123", 1*time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = proxy.validateAccessToken(token, "obj-123")
	}
}
