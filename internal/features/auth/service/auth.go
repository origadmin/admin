package service

import (
	"context"

	v1 "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/auth/biz"
	"origadmin/application/admin/internal/helpers/captcha"
	"origadmin/application/admin/internal/helpers/idutil"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"

	securityv1 "github.com/origadmin/contrib/api/gen/go/security/v1"
	"github.com/origadmin/contrib/security/credential"
	securityPrincipal "github.com/origadmin/contrib/security/principal"
)

// AuthService is a service for authentication.
type AuthService struct {
	v1.UnimplementedAuthServiceServer
	uc        *biz.AuthUseCase
	captchaUC *biz.CaptchaUseCase
	creator   credential.Creator
	log       *log.Helper
}

// NewAuthService creates a new authentication service.
func NewAuthService(uc *biz.AuthUseCase, cuc *biz.CaptchaUseCase, creator credential.Creator, logger log.Logger) *AuthService {
	return &AuthService{uc: uc, captchaUC: cuc, creator: creator, log: log.NewHelper(logger)}
}

// Login authenticates a user and returns a token pair.
func (s *AuthService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// Verify captcha
	if !s.captchaUC.VerifyCaptcha(ctx, req.GetCaptchaId(), req.GetCaptchaCode()) {
		return nil, errors.New(400, "CAPTCHA_INVALID", "invalid captcha")
	}

	user, err := s.uc.VerifyUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// Update login info before creating the token.
	var loginIP string
	if tr, ok := transport.FromServerContext(ctx); ok {
		loginIP = tr.RequestHeader().Get("X-Real-IP")
	}
	if err := s.uc.UpdateLoginInfo(ctx, user.Id, loginIP); err != nil {
		// Log the error but don't block the login process.
		s.log.Errorf("failed to update login info for user %d: %v", user.Id, err)
	}

	// Create a principal for the user, including their roles.
	var roleKeywords []string
	for _, role := range user.GetRoles() {
		roleKeywords = append(roleKeywords, role.Keyword)
	}
	p := securityPrincipal.New(
		idutil.FormatUserID(user.Id),
		securityPrincipal.WithRoles(roleKeywords),
	)

	// Create a credential (which contains the token).
	credResp, err := s.creator.CreateCredential(ctx, p)
	if err != nil {
		return nil, err
	}

	token := credResp.Payload().GetToken()
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
	captchaType := req.GetCaptchaType()
	if captchaType == "" {
		captchaType = captcha.TypeDigit
	}

	if captchaType == captcha.TypeAudio {
		b64s, err := s.captchaUC.GetCaptchaAudio(ctx, req.GetCaptchaId())
		if err != nil {
			return nil, err
		}
		return &v1.GetCaptchaResponse{
			CaptchaId:   req.GetCaptchaId(),
			CaptchaData: b64s,
			MimeType:    captcha.MimeTypeAudio,
		}, nil
	}

	id, b64s, _, err := s.captchaUC.GenerateCaptcha(ctx, captchaType)
	if err != nil {
		return nil, err
	}
	return &v1.GetCaptchaResponse{
		CaptchaId:   id,
		CaptchaData: b64s,
		MimeType:    captcha.MimeTypeImage,
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
	refreshToken := req.GetRefreshToken()
	if refreshToken == "" {
		return nil, types.ErrorAuthErrorReasonTokenMissing("refresh token is missing")
	}

	refresher, ok := s.creator.(credential.Refresher)
	if !ok {
		return nil, types.ErrorAuthErrorReasonUnspecified("credential creator does not support refresh")
	}

	credResp, err := refresher.RefreshCredential(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	token := credResp.Response().GetPayload().GetToken()
	if token == nil {
		return nil, types.ErrorAuthErrorReasonTokenInvalid("token is missing in response")
	}

	return &v1.RefreshTokenResponse{
		AccessToken:  token.GetAccessToken(),
		TokenType:    token.GetTokenType(),
		ExpiresIn:    token.GetExpiresIn(),
		RefreshToken: token.GetRefreshToken(),
	}, nil
}

// Authenticate is for internal use by the gateway to verify user access via gRPC.
func (s *AuthService) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	return &v1.AuthenticateResponse{}, nil
}
