// Copyright 2024 OrigAdmin. All rights reserved.

package tools

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestHTTPClient 测试用的HTTP客户端
type TestHTTPClient struct {
	client  *http.Client
	baseURL string
}

// NewTestHTTPClient 创建测试HTTP客户端
func NewTestHTTPClient(baseURL string) *TestHTTPClient {
	return &TestHTTPClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

// Request 发送HTTP请求
func (c *TestHTTPClient) Request(t *testing.T, method, path string, body interface{}, token string) *http.Response {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)
		bodyReader = bytes.NewReader(jsonBody)
	} else {
		bodyReader = http.NoBody
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := c.client.Do(req)
	require.NoError(t, err)

	return resp
}

// Login 登录获取Token
func (c *TestHTTPClient) Login(t *testing.T, username, password string) string {
	req := map[string]interface{}{
		"username": username,
		"password": password,
	}

	resp := c.Request(t, "POST", "/api/v1/auth/login", req, "")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &result)
	require.NoError(t, err)

	return result["access_token"].(string)
}
