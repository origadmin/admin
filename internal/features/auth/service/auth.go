package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/internal/features/auth/biz"
	"origadmin/application/admin/internal/pkg/token"
)

// AuthService is a service for authentication.
type AuthService struct {
	v1.UnimplementedAuthServer
	uc *biz.AuthUseCase
	tm token.Manager
}

// NewAuthService creates a new authentication service.
func NewAuthService(uc *biz.AuthUseCase, tm token.Manager) *AuthService {
	return &AuthService{uc: uc, tm: tm}
}

// Login authenticates a user and returns a token pair.
func (s *AuthService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	userID, err := s.uc.VerifyUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := s.tm.Generate(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &v1.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}

// Register creates a new user account.
func (s *AuthService) Register(ctx context.Context, req *v1.RegisterRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Logout invalidates the user's session.
func (s *AuthService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// RefreshToken provides a new access token.
func (s *AuthService) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	return &v1.RefreshTokenResponse{}, nil
}

// GetCaptcha generates a new captcha.
func (s *AuthService) GetCaptcha(ctx context.Context, req *v1.GetCaptchaRequest) (*v1.GetCaptchaResponse, error) {
	return &v1.GetCaptchaResponse{}, nil
}

// Authenticate is for internal use by the gateway to verify user access via gRPC.
func (s *AuthService) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	return &v1.AuthenticateResponse{}, nil
}
