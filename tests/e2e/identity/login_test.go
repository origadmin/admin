package identity

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/tests/e2e"
	"origadmin/application/admin/tests/tools"
)

func TestLogin_E2E(t *testing.T) {
	t.Run("LoginSuccessWithoutCaptcha", func(t *testing.T) {
		// This test case explicitly verifies the login process itself.
		loginReq := &v1.LoginRequest{
			Username: "admin",
			Password: "admin123", // Corrected admin password
		}

		resp := e2e.DoRequest(t, "POST", "/auth/login", loginReq, "")
		defer resp.Body.Close()

		// Use AssertHTTPStatusCode for detailed failure reporting (logs URL and Body)
		tools.AssertHTTPStatusCode(t, resp, http.StatusOK)

		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var loginResp v1.LoginResponse
		unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
		err = unmarshaler.Unmarshal(respBody, &loginResp)
		require.NoError(t, err, "Failed to unmarshal login response")

		require.NotEmpty(t, loginResp.AccessToken, "Access token should not be empty")
		t.Logf("Successfully logged in via gateway.")
	})
}
