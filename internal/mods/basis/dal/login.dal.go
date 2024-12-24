/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/LyricTian/captcha"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/origadmin/runtime/context"
	jwtv1 "github.com/origadmin/runtime/gen/go/security/jwt/v1"
	securityv1 "github.com/origadmin/runtime/gen/go/security/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/rand"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/httperr"
	"github.com/origadmin/toolkits/security"
	"google.golang.org/protobuf/types/known/anypb"

	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/basis/dto"
	systemdto "origadmin/application/admin/internal/mods/system/dto"
)

type loginRepo struct {
	*LoginData
	bufpool *sync.Pool
	//Menu          systemdto.MenuRepo
	//Role          systemdto.RoleRepo
	//User          systemdto.UserRepo
	//Authenticator security.Authenticator
}

func (repo loginRepo) TokenRefresh(ctx context.Context, in *dto.TokenRefreshRequest) (*dto.TokenRefreshResponse, error) {
	log.Debugf("Token refresh request received with data: %+v", in.GetData())
	return repo.refreshToken(ctx, in.GetData().GetRefreshToken())
}

func (repo loginRepo) Register(ctx context.Context, in *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	log.Debugf("Register request received with data: %+v", in.GetData())
	return nil, errors.New("not implemented")
}

func (repo loginRepo) Login(ctx context.Context, in *dto.LoginRequest) (*dto.LoginResponse, error) {
	log.Debugf("Login request received with data: %+v", in.GetData())
	data := in.GetData()

	// verify captcha
	log.Debugf("Verifying captcha with id %s and code %s", data.CaptchaId, data.CaptchaCode)
	if !captcha.VerifyString(data.CaptchaId, data.CaptchaCode) {
		log.Warnf("Invalid captcha id %s or code %s", data.CaptchaId, data.CaptchaCode)
		return nil, dto.ErrInvalidCaptchaID
	}

	//ctx = context.NewTag(ctx, logging.TagKeyLogin)

	if root := repo.cfg().RootUser; root.GetEnabled() {
		log.Debugf("Root user is enabled, checking if username matches")
		// login by root
		username := root.Username
		if data.Username == username {
			log.Debugf("Username matches, checking password")
			if err := hash.Compare(root.Password, data.Password, root.Salt); err != nil {
				log.Warnf("Invalid password for root user")
				return nil, dto.ErrInvalidPassword
			}

			userID := root.Id
			ctx = context.NewID(ctx, root.Id)
			log.Infof("Login by root successful, user ID: %s", userID)
			return repo.genToken(ctx, userID)
		}
	}

	// get user info
	log.Debugf("Getting user info for username %s", data.Username)
	user, err := repo.User.GetByUserName(ctx, data.Username, "id", "password", "salt", "status")
	if err != nil {
		log.Errorf("Error getting user info: %v", err)
		return nil, err
	}
	switch {
	case user == nil:
		log.Warnf("User not found with username %s", data.Username)
		return nil, dto.ErrInvalidUsername
	case user.Status != systemdto.UserStatusActivated:
		log.Warnf("User %s is not activated", data.Username)
		return nil, httperr.New("unknown", 400, "User status is not activated, please contact the administrator")
	default:
		log.Debugf("User found with ID %s and status %s", user.Id, user.Status)
	}

	// check password
	log.Debugf("Comparing password for user %s", data.Username)
	if err := hash.Compare(user.Password, data.Password, user.Salt); err != nil {
		log.Warnf("Invalid password for user %s", data.Username)
		return nil, dto.ErrInvalidPassword
	}

	userID := user.Id
	username := user.Username
	ctx = context.NewID(ctx, userID)

	// set user cache with role ids
	log.Debugf("Getting role IDs for user %s", username)
	roleIDs, err := repo.User.GetRoleIDs(ctx, userID)
	if err != nil {
		log.Errorf("Error getting role IDs: %v", err)
		return nil, kerr.Newf(404, "UNKNOWN", "failed to get user role ids: %v", err)
	}

	log.Infof("User %s logged in successfully with role ids: %v", username, roleIDs)
	// generate token
	log.Debugf("Generating token for user %s", username)
	return repo.genToken(ctx, userID)
}

func (repo loginRepo) CaptchaImage(ctx context.Context, id string, reload bool) (*dto.CaptchaImageResponse, error) {
	log.Debugf("Generating captcha image with id %s and reload %v", id, reload)
	if reload && !captcha.Reload(id) {
		log.Warnf("Captcha id %s not found during reload", id)
		return nil, dto.ErrCaptchaIDNotFound
	}
	buf := repo.getBuf()
	captchaConfig := repo.cfg().Captcha
	log.Debugf("Writing captcha image to buffer with width %d and height %d", int(captchaConfig.Width), int(captchaConfig.Height))
	err := captcha.WriteImage(buf, id, int(repo.cfg().Captcha.Width), int(captchaConfig.Height))
	if err != nil {
		log.Errorf("Error writing captcha image: %v", err)
		if errors.Is(err, captcha.ErrNotFound) {
			log.Warnf("Captcha id %s not found", id)
			return nil, dto.ErrCaptchaIDNotFound
		}
		return nil, kerr.Newf(400, "captcha write image", "failed to generate captcha image: %v", err)
	}
	log.Debugf("Captcha image generated successfully")
	response := new(dto.CaptchaImageResponse)
	response.Headers = map[string]string{
		"Cache-Control": "no-cache, no-store, must-revalidate",
		"Pragma":        "no-cache",
		"Expires":       "0",
		"Content-Type":  "image/png",
	}
	response.Image = buf.Bytes()
	log.Debugf("Returning captcha image response with headers: %+v", response.Headers)
	return response, nil
}

func (repo loginRepo) CurrentMenus(ctx context.Context, in *dto.CurrentMenusRequest) (*dto.CurrentMenusResponse, error) {
	menus, err := repo.User.ListMenuByUserID(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	var anyMenus []*anypb.Any
	for i, menu := range menus {
		anyMenus[i] = resp.Any(menu)
	}
	return &dto.CurrentMenusResponse{
		Data: anyMenus,
	}, nil
}

func (repo loginRepo) CurrentUser(ctx context.Context, in *dto.CurrentUserRequest) (*dto.CurrentUserResponse, error) {
	user, err := repo.User.Current(ctx, in.GetData().GetUserId())
	if err != nil {
		return nil, err
	}
	return &dto.CurrentUserResponse{
		Data: resp.Any(user),
	}, nil
}

func (repo loginRepo) Logout(ctx context.Context, in *dto.LogoutRequest) (*dto.LogoutResponse, error) {
	return &dto.LogoutResponse{}, nil
}

func (repo loginRepo) CaptchaID(ctx context.Context, in *dto.CaptchaIDRequest) (*dto.CaptchaIDResponse, error) {
	return &dto.CaptchaIDResponse{
		Data: captcha.NewLen(int(repo.cfg().Captcha.Length)),
	}, nil
}

func (repo loginRepo) FreeBuf(buf *bytes.Buffer) {
	repo.putBuf(buf)
}

func (repo loginRepo) getBuf() *bytes.Buffer {
	return repo.bufpool.Get().(*bytes.Buffer)
}

func (repo loginRepo) putBuf(buf *bytes.Buffer) {
	buf.Reset()
	repo.bufpool.Put(buf)
}

func (repo loginRepo) refreshToken(ctx context.Context, token string) (*dto.TokenRefreshResponse, error) {
	claims, err := repo.Authenticator.Authenticate(ctx, token)
	if err != nil {
		return nil, err
	}
	if claims.GetExpiration().Before(time.Now()) {
		if err := repo.Authenticator.DestroyToken(ctx, token); err != nil {
			log.Errorf("Error destroying token: %v", err)
			return nil, err
		}
		return nil, fmt.Errorf("token expired")
	}
	genToken, err := repo.genToken(ctx, claims.GetSubject())
	if err != nil {
		return nil, err
	}
	return &dto.TokenRefreshResponse{
		Token: genToken.Token,
	}, nil
}

func (repo loginRepo) genToken(ctx context.Context, id string) (*dto.LoginResponse, error) {
	claims, err := repo.Authenticator.CreateIdentityClaims(ctx, id, false)
	if err != nil {
		return nil, err
	}
	token, err := repo.Authenticator.CreateToken(ctx, claims)
	if err != nil {
		return nil, err
	}
	refreshClaims, err := repo.Authenticator.CreateIdentityClaims(ctx, id, true)
	if err != nil {
		return nil, err
	}
	refreshToken, err := repo.Authenticator.CreateToken(ctx, refreshClaims)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Token: &jwtv1.Token{
			UserId:         id,
			AccessToken:    token,
			RefreshToken:   refreshToken,
			ExpirationTime: claims.GetExpiration().Unix(),
			//Claims:         fromSecurityClaims(claims),
		},
	}, nil
}

func fromSecurityClaims(claims security.Claims) *securityv1.Claims {
	return &securityv1.Claims{
		Sub:    claims.GetSubject(),
		Iss:    claims.GetIssuer(),
		Aud:    claims.GetAudience(),
		Exp:    claims.GetExpiration().Unix(),
		Nbf:    claims.GetNotBefore().Unix(),
		Iat:    claims.GetIssuedAt().Unix(),
		Jti:    claims.GetJWTID(),
		Scopes: claims.GetScopes(),
	}
}

func (repo loginRepo) cfg() *configs.BasisConfig {
	return repo.LoginData.BasisConfig
}

type LoginData struct {
	BasisConfig   *configs.BasisConfig
	Authenticator security.Authenticator
	Menu          systemdto.MenuRepo
	Role          systemdto.RoleRepo
	User          systemdto.UserRepo
}

// NewLoginRepo .
func NewLoginRepo(data *LoginData, logger log.KLogger) dto.LoginRepo {
	var err error
	cfg := data.BasisConfig
	// todo: generate random password for root user if not exists
	if cfg.RootUser.RandomPassword {
		passwd := rand.GenerateRandom(12)
		cfg.RootUser.Salt = rand.GenerateSalt()
		cfg.RootUser.Password, err = hash.Generate(passwd, cfg.RootUser.Salt)
		if err == nil {
			fmt.Println("Root user password:", passwd)
		} else {
			log.Errorf("Error generating password: %v", err)
			cfg.RootUser.RandomPassword = false
		}
	}
	if cfg.RootUser.Id == "" {
		cfg.RootUser.Id = cfg.RootUser.Username
	}
	//authenticator, err := jwt.NewAuthenticator(&configv1.Security{})
	//if err != nil {
	//	panic(err)
	//}
	return &loginRepo{
		bufpool:   BufPool(),
		LoginData: data,
	}
}

func BufPool() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
}
