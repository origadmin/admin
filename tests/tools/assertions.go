// Copyright 2024 OrigAdmin. All rights reserved.

package tools

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// AssertHTTPStatusCode checks response status code
func AssertHTTPStatusCode(t *testing.T, resp *http.Response, expected int) {
	t.Helper()
	require.Equal(t, expected, resp.StatusCode,
		"Expected status %d, got %d. Body: %s", expected, resp.StatusCode, getResponseBodyForError(resp))
}

// AssertJSONBody decodes and validates JSON response
func AssertJSONBody(t *testing.T, resp *http.Response, target interface{}) {
	t.Helper()
	body := resp.Body
	defer body.Close()
	require.NoError(t, json.NewDecoder(body).Decode(target), "Failed to decode JSON response")
}

// getResponseBodyForError reads response body for error messages without closing it
func getResponseBodyForError(resp *http.Response) string {
	if resp == nil || resp.Body == nil {
		return ""
	}
	// Read body without closing it, so caller can still read it
	bodyBytes, _ := io.ReadAll(resp.Body)
	// Reset the body so it can be read again
	resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	return string(bodyBytes)
}
