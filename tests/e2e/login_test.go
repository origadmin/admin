package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "origadmin/application/admin/api/v1/services/auth"
)

// TestConfig defines the configuration for the E2E tests.
var testConfig = struct {
	GatewayAddress string
}{
	GatewayAddress: "http://localhost:8000", // Default gateway address
}

// client is a shared HTTP client for all tests
var client = &http.Client{
	Timeout: 10 * time.Second,
}

// loginAndGetToken is a helper function to login and return the access token.
// It fails the test immediately if login fails.
func loginAndGetToken(t *testing.T) string {
	loginReq := &v1.LoginRequest{
		Username: "admin",    // Default root user
		Password: "admin123", // Default password
	}
	body, err := json.Marshal(loginReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", testConfig.GatewayAddress+"/api/v1/auth/login", bytes.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Login failed with status %d. Response: %s", resp.StatusCode, string(respBody))
	}

	var loginResp v1.LoginResponse
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	err = unmarshaler.Unmarshal(respBody, &loginResp)
	require.NoError(t, err)

	return loginResp.AccessToken
}

// doRequest is a helper to send an HTTP request with an optional token.
func doRequest(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)
		bodyReader = bytes.NewReader(jsonBody)
	} else {
		bodyReader = http.NoBody
	}

	req, err := http.NewRequest(method, testConfig.GatewayAddress+path, bodyReader)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	require.NoError(t, err)

	return resp
}

func TestLogin_E2E(t *testing.T) {
	t.Run("LoginSuccessWithoutCaptcha", func(t *testing.T) {
		// This test case explicitly verifies the login process itself,
		// duplicating some logic from loginAndGetToken but with assertions.
		loginReq := &v1.LoginRequest{
			Username: "admin",
			Password: "admin123",
		}
		body, err := json.Marshal(loginReq)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", testConfig.GatewayAddress+"/api/v1/auth/login", bytes.NewReader(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "response code should be 200 OK")

		if resp.StatusCode == http.StatusOK {
			var loginResp v1.LoginResponse
			unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
			err = unmarshaler.Unmarshal(respBody, &loginResp)
			require.NoError(t, err)

			assert.NotEmpty(t, loginResp.AccessToken)
			t.Logf("Successfully logged in via gateway.")
		}
	})
}
