package identity

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "origadmin/application/admin/api/v1/services/identity"
	"origadmin/application/admin/tests/e2e"
)

func TestLogin_E2E(t *testing.T) {
	t.Run("LoginSuccessWithoutCaptcha", func(t *testing.T) {
		// This test case explicitly verifies the login process itself.
		loginReq := &v1.LoginRequest{
			Username: "admin",
			Password: "admin123", // Corrected admin password
		}

		resp := e2e.DoRequest(t, "POST", "/api/v1/auth/login", loginReq, "")
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
