/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"

	"github.com/LyricTian/captcha"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/goexts/generic"
	"github.com/google/uuid"
	"github.com/origadmin/runtime/context"
	jwtv1 "github.com/origadmin/runtime/gen/go/security/jwt/v1"
	securityv1 "github.com/origadmin/runtime/gen/go/security/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/rand"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/httperr"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/id"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dto"
	systemdto "origadmin/application/admin/internal/mods/system/dto"
)

type loginRepo struct {
	*LoginData
	bufpool *sync.Pool
}

func (repo loginRepo) TokenRefresh(ctx context.Context, in *dto.TokenRefreshRequest) (*dto.TokenRefreshResponse, error) {
	log.Debugf("Token refresh request received with data: %+v", in.GetData())
	return repo.refreshToken(ctx, in.GetData().GetRefreshToken())
}

func (repo loginRepo) Register(ctx context.Context, in *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	log.Debugf("Register request received with data: %+v", in.GetData())
	data := in.GetData()
	registerID := id.Gen()
	salt := rand.GenerateSalt()
	passwd, err := hash.Generate(data.GetPassword(), salt)
	if err != nil {
		return nil, err
	}
	if _, err := repo.User.Create(ctx, &system.User{
		Id:       registerID,
		Uuid:     generic.Must(uuid.NewRandom()).String(),
		Username: data.GetUsername(),
		Name:     "user_" + strconv.Itoa(int(registerID)),
		Password: passwd,
		Salt:     salt,
		Status:   1,
	}); err != nil {
		return nil, err
	}
	return &dto.RegisterResponse{
		Success: true,
		Data: &system.RegisterResponse_Data{
			Redirect: "",
		},
	}, nil
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

	if root := repo.cfg().RootUser; root.GetEnabled() {
		log.Debugf("Root userData is enabled, checking if username matches")
		// login by root
		username := root.Username
		if data.Username == username {
			log.Debugf("Username matches, checking password")
			if err := hash.Compare(root.Password, data.Password, root.Salt); err != nil {
				log.Warnf("Invalid password for root userData")
				return nil, dto.ErrInvalidPassword
			}

			userID := root.Id
			ctx = context.NewID(ctx, root.Id)
			log.Infof("Login by root successful, userData ID: %s", userID)
			return repo.genToken(ctx, userID)
		}
	}

	// get user info
	log.Debugf("Getting userData info for username %s", data.Username)
	userData, err := repo.User.GetByUserName(ctx, data.Username, user.FieldID, user.FieldPassword, user.FieldSalt, user.FieldStatus)
	if err != nil {
		log.Errorf("Error getting userData info: %v", err)
		return nil, err
	}
	switch {
	case userData == nil:
		log.Warnf("User not found with username %s", data.Username)
		return nil, dto.ErrInvalidUsername
	case userData.Status != systemdto.UserStatusActivated:
		log.Warnf("User %s is not activated", data.Username)
		return nil, httperr.New("unknown", 400, "User status is not activated, please contact the administrator")
	default:
		log.Debugf("User found with ID %s and status %s", userData.Id, userData.Status)
	}

	// check password
	log.Debugf("Comparing password for userData %s", data.Username)
	if err := hash.Compare(userData.Password, data.Password, userData.Salt); err != nil {
		log.Warnf("Invalid password for userData %s", data.Username)
		return nil, dto.ErrInvalidPassword
	}

	userUUID := userData.Uuid
	username := userData.Username
	ctx = context.NewID(ctx, userUUID)

	// set userData cache with role ids
	log.Debugf("Getting role IDs for userData %s", username)
	roleIDs, err := repo.User.GetRoleIDs(ctx, userData.Id)
	if err != nil {
		log.Errorf("Error getting role IDs: %v", err)
		return nil, kerr.Newf(404, "UNKNOWN", "failed to get userData role ids: %v", err)
	}

	log.Infof("User %s logged in successfully with role ids: %v", username, roleIDs)
	// generate token
	log.Debugf("Generating token for userData %s", username)
	return repo.genToken(ctx, userUUID)
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
	return &dto.CurrentMenusResponse{
		Data: resp.Proto2AnyPBArray(menus...),
	}, nil
}

func (repo loginRepo) CurrentUser(ctx context.Context, in *dto.CurrentUserRequest) (*dto.CurrentUserResponse, error) {
	current, err := repo.User.Current(ctx, in.GetData().GetUserId())
	if err != nil {
		return nil, err
	}
	return &dto.CurrentUserResponse{
		Data: resp.Any(current),
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
	claims, err := repo.Tokenizer.ParseClaims(ctx, token)
	if err != nil {
		return nil, err
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

	claims, err := repo.Tokenizer.CreateClaims(ctx, id)
	if err != nil {
		return nil, err
	}
	token, err := repo.Tokenizer.CreateToken(ctx, claims)
	if err != nil {
		return nil, err
	}
	refreshClaims, err := repo.Tokenizer.CreateRefreshClaims(ctx, id)
	if err != nil {
		return nil, err
	}
	refreshToken, err := repo.Tokenizer.CreateToken(ctx, refreshClaims)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Token: &jwtv1.Token{
			UserId:         id,
			AccessToken:    token,
			RefreshToken:   refreshToken,
			ExpirationTime: claims.GetExpiration(),
		},
	}, nil
}

func fromSecurityClaims(claims security.Claims) *securityv1.Claims {
	return &securityv1.Claims{
		Sub:    claims.GetSubject(),
		Iss:    claims.GetIssuer(),
		Aud:    claims.GetAudience(),
		Exp:    claims.GetExpiration(),
		Nbf:    claims.GetNotBefore(),
		Iat:    claims.GetIssuedAt(),
		Jti:    claims.GetID(),
		Scopes: claims.GetScopes(),
	}
}

func RefreshTokenizer(tokenizer security.Tokenizer) security.RefreshTokenizer {
	if rt, ok := tokenizer.(security.RefreshTokenizer); ok {
		return rt
	}
	return wrapRefreshTokenizer(tokenizer)
}

func (repo loginRepo) cfg() *configs.BasisConfig {
	return repo.LoginData.BasisConfig
}

type LoginData struct {
	BasisConfig *configs.BasisConfig
	Tokenizer   security.RefreshTokenizer
	Menu        systemdto.MenuRepo
	Role        systemdto.RoleRepo
	User        systemdto.UserRepo
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
