package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	authv1 "origadmin/application/admin/api/v1/services/auth"
)

const (
	baseURL        = "http://localhost:8000"
	adminUser      = "admin"
	adminPass      = "admin"
	requestTimeout = 5 * time.Second
)

// doRequest performs an HTTP request and returns the response.
func doRequest(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	t.Helper()

	var bodyReader io.Reader
	if body != nil {
		var jsonBytes []byte
		var err error

		// Check if the body is a protobuf message
		if pb, ok := body.(proto.Message); ok {
			jsonBytes, err = protojson.Marshal(pb)
			require.NoError(t, err)
		} else {
			// Fallback to standard JSON marshaling
			jsonBytes, err = json.Marshal(body)
			require.NoError(t, err)
		}
		bodyReader = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest(method, baseURL+path, bodyReader)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{Timeout: requestTimeout}
	resp, err := client.Do(req)
	require.NoError(t, err)

	return resp
}

// login performs a login request and returns the auth token.
func login(t *testing.T, username, password string) string {
	t.Helper()
	reqBody := &authv1.LoginRequest{
		Username: username,
		Password: password,
	}
	resp := doRequest(t, "POST", "/api/v1/auth/login", reqBody, "")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "Login failed")

	var loginResp authv1.LoginResponse
	bodyBytes, _ := io.ReadAll(resp.Body)
	err := protojson.Unmarshal(bodyBytes, &loginResp)
	require.NoError(t, err)
	require.NotEmpty(t, loginResp.AccessToken)

	return loginResp.AccessToken
}

// loginAndGetToken is a convenience helper to log in as the default admin.
func loginAndGetToken(t *testing.T) string {
	t.Helper()
	return login(t, adminUser, adminPass)
}
