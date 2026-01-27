package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	authv1 "origadmin/application/admin/api/v1/services/auth"
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

// TestMain is the entry point for all tests in this package.
func TestMain(m *testing.M) {
	// Here you can add global setup logic, e.g., starting a server.
	// For now, we just run the tests.
	exitCode := m.Run()
	// Here you can add global teardown logic.
	os.Exit(exitCode)
}

// loginAndGetToken is a helper function to login as the default admin and return the access token.
// It fails the test immediately if login fails.
func loginAndGetToken(t *testing.T) string {
	// This helper uses the default admin credentials.
	return login(t, "admin", "admin123")
}

// login is a generic login helper.
func login(t *testing.T, username, password string) string {
	loginReq := &authv1.LoginRequest{
		Username: username,
		Password: password,
	}

	resp := doRequest(t, "POST", "/api/v1/auth/login", loginReq, "")
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		t.Fatalf("Login failed for %s: %d %s", username, resp.StatusCode, string(bodyBytes))
	}

	var loginResp authv1.LoginResponse
	respBody, _ := io.ReadAll(resp.Body)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	err := unmarshaler.Unmarshal(respBody, &loginResp)
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
