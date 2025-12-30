package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "origadmin/application/admin/api/v1/services/auth"
	securityv1 "github.com/origadmin/contrib/api/gen/go/security/v1"
	"github.com/origadmin/contrib/security/credential"
	securityPrincipal "github.com/origadmin/contrib/security/principal"
	"origadmin/application/admin/internal/features/auth/biz"
	"origadmin/application/admin/internal/helpers/captcha"
)

// AuthService is a service for authentication.
type AuthService struct {
	v1.UnimplementedAuthServiceServer
	uc         *biz.AuthUseCase
	captcha    *captcha.Captcha
	creator    credential.Creator
}

// NewAuthService creates a new authentication service.
func NewAuthService(uc *biz.AuthUseCase, captcha *captcha.Captcha, creator credential.Creator) *AuthService {
	return &AuthService{uc: uc, captcha: captcha, creator: creator}
}

// Login authenticates a user and returns a token pair.
func (s *AuthService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// Verify captcha
	if !s.captcha.Verify(req.GetCaptchaId(), req.GetCaptchaCode(), true) {
		return nil, errors.New(400, "CAPTCHA_INVALID", "invalid captcha")
	}

	userID, err := s.uc.VerifyUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// Create a principal for the user.
	p := securityPrincipal.New(fmt.Sprint(userID))

	// Create a credential (which contains the token).
	credResp, err := s.creator.CreateCredential(ctx, p)
	if err != nil {
		return nil, err
	}

	token := credResp.Response().GetPayload().GetToken()
	if token == nil {
		return nil, securityv1.ErrorTokenInvalid("token is missing")
	}

	return &v1.LoginResponse{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		TokenType:    token.GetTokenType(),
		ExpiresIn:    token.GetExpiresIn(),
	}, nil
}

// GetCaptcha generates a new captcha.
func (s *AuthService) GetCaptcha(ctx context.Context, req *v1.GetCaptchaRequest) (*v1.GetCaptchaResponse, error) {
	_, id, b64s, err := s.captcha.GenerateDigit()
	if err != nil {
		return nil, err
	}
	return &v1.GetCaptchaResponse{
		CaptchaId:    id,
		CaptchaImage: b64s,
	}, nil
}

// Register creates a new user account.
func (s *AuthService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return &v1.RegisterResponse{}, nil
}

// Logout invalidates the user's session.
func (s *AuthService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	return &v1.LogoutResponse{}, nil
}

// RefreshToken provides a new access token.
func (s *AuthService) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	return &v1.RefreshTokenResponse{}, nil
}

// Authenticate is for internal use by the gateway to verify user access via gRPC.
func (s *AuthService) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	return &v1.AuthenticateResponse{}, nil
}
