// Copyright 2024 OrigAdmin. All rights reserved.

package tools

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// AssertHTTPStatusCode checks response status code
func AssertHTTPStatusCode(t *testing.T, resp *http.Response, expected int) {
	t.Helper()
	require.Equal(t, expected, resp.StatusCode,
		"Expected status %d, got %d. Body: %s", expected, resp.StatusCode, getResponseBody(resp))
}

// AssertJSONBody decodes and validates JSON response
func AssertJSONBody(t *testing.T, resp *http.Response, target interface{}) {
	t.Helper()
	body := resp.Body
	defer body.Close()
	require.NoError(t, json.NewDecoder(body).Decode(target), "Failed to decode JSON response")
}

// getResponseBody reads response body for error messages
func getResponseBody(resp *http.Response) string {
	if resp == nil || resp.Body == nil {
		return ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := json.Marshal(resp)
	return string(bodyBytes)
}
