/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"bytes"
	"fmt"
	"sync"

	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/origadmin/runtime/context"
	jwtv1 "github.com/origadmin/runtime/gen/go/security/jwt/v1"
	securityv1 "github.com/origadmin/runtime/gen/go/security/v1"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/rand"
	"github.com/origadmin/toolkits/errors/httperr"
	"github.com/origadmin/toolkits/security"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/captcha"
	"origadmin/application/admin/helpers/resp"
	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dto"
	systemdto "origadmin/application/admin/internal/mods/system/dto"
)

type loginRepo struct {
	*LoginData
	captcha *captcha.Captcha
	bufpool *sync.Pool
}

func (repo loginRepo) TokenRefresh(ctx context.Context, in *dto.TokenRefreshRequest) (*dto.TokenRefreshResponse, error) {
	log.Debugf("Token refresh request received with data: %+v", in.GetData())
	return repo.refreshToken(ctx, in.GetData().GetRefreshToken())
}

func (repo loginRepo) Register(ctx context.Context, in *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	log.Debugf("Register request received with data: %+v", in.GetData())
	data := in.GetData()
	var err error
	createUser := new(dto.UserPB)
	createUser, _, err = dto.MakeCreateUser(createUser, data.GetUsername(), data.GetPassword(), dto.UserMutationOption{})
	if err != nil {
		return nil, err
	}
	if _, err := repo.User.Create(ctx, createUser); err != nil {
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
	if !repo.captcha.Store.Verify(data.CaptchaId, data.CaptchaCode, true) {
		log.Warnf("Invalid captcha id %s or code %s", data.CaptchaId, data.CaptchaCode)
		return nil, dto.ErrInvalidCaptchaID
	}

	if root := repo.cfg().RootUser; root.GetEnabled() {
		log.Debugf("Root userData is enabled, checking if username matches")
		// login by root
		username := root.Username
		if data.Username == username {
			log.Debugf("Username matches, checking password")
			if err := hash.Verify(root.Password, data.Password); err != nil {
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
	case userData.Status != systemdto.UserStatusActive:
		log.Warnf("User %s is not activated", data.Username)
		return nil, httperr.New("unknown", 400, "User status is not activated, please contact the administrator")
	default:
		log.Debugf("User found with ID %d and status %d", userData.Id, userData.Status)
	}

	// check password
	log.Debugf("Comparing password for userData %s", data.Username)
	if err := hash.Verify(userData.Password, data.Password); err != nil {
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
func (repo loginRepo) CaptchaAudio(ctx context.Context, id string, reload bool) (*dto.CaptchaAudioResponse, error) {
	var err error
	log.Debugf("Generating captcha audio with id %s and reload %v", id, reload)
	if reload && !repo.captcha.Reload(captcha.TypeAudio, id) {
		log.Warnf("Captcha id %s not found during reload, regenerating", id)
		id, err = repo.getCaptchaID()
		if err != nil {
			return nil, err
		}
	}
	content := repo.captcha.Store.Get(id, false)
	item, err := repo.captcha.DriverAudio.Driver.DrawCaptcha(content)
	if err != nil {
		return nil, err
	}
	buf := repo.getBuf()
	_, err = item.WriteTo(buf)
	if err != nil {
		return nil, err
	}
	response := new(dto.CaptchaAudioResponse)
	response.Headers = map[string]string{
		"Cache-Control": "no-cache, no-store, must-revalidate",
		"Pragma":        "no-cache",
		"Expires":       "0",
		"Content-Type":  captcha.MimeTypeAudio,
	}
	response.Audio = buf.Bytes()
	return response, nil
}

func (repo loginRepo) CaptchaImage(ctx context.Context, id string, reload bool) (*dto.CaptchaImageResponse, error) {
	log.Debugf("Generating captcha image with id %s and reload %v", id, reload)
	var err error
	if reload && !repo.captcha.Reload(captcha.TypeDigit, id) {
		log.Warnf("Captcha id %s not found during reload, regenerating", id)
		id, err = repo.getCaptchaID()
		if err != nil {
			return nil, err
		}
	}
	content := repo.captcha.Store.Get(id, false)
	item, err := repo.captcha.DriverDigit.Driver.DrawCaptcha(content)
	if err != nil {
		return nil, err
	}
	buf := repo.getBuf()
	_, err = item.WriteTo(buf)
	if err != nil {
		return nil, err
	}
	log.Debugf("Captcha image generated successfully")
	response := new(dto.CaptchaImageResponse)
	response.Headers = map[string]string{
		"Cache-Control": "no-cache, no-store, must-revalidate",
		"Pragma":        "no-cache",
		"Expires":       "0",
		"Content-Type":  captcha.MimeTypeImage,
	}
	response.Image = buf.Bytes()
	log.Debugf("Returning captcha image response with headers: %+v", response.Headers)
	return response, nil
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
	id, err := repo.getCaptchaID()
	if err != nil {
		return nil, err
	}
	return &dto.CaptchaIDResponse{
		Data: id,
	}, nil
}

func (repo loginRepo) Captcha(ctx context.Context, in *dto.CaptchaRequest) (*dto.CaptchaResponse, error) {
	var err error
	var id = in.Id
	if id == "" {
		id, err = repo.getCaptchaID()
		if err != nil {
			return nil, err
		}
	}
	driver, err := repo.getCaptchaDriver(in.Type)
	if err != nil {
		return nil, err
	}
	if in.Reload && in.Id != "" && !repo.captcha.Reload(in.Type, in.Id) {
		log.Warnf("Captcha id %s not found during reload, regenerating id", id)
		id, err = repo.getCaptchaID()
		if err != nil {
			return nil, err
		}
	}
	data, err := repo.getCaptchaData(driver, id)
	if err != nil {
		return nil, err
	}
	return &dto.CaptchaResponse{
		Id:   id,
		Type: in.Type,
		Data: data,
	}, nil
}

func (repo loginRepo) getCaptchaID() (string, error) {
	id, _, answ, err := repo.captcha.DriverDigit.Generate()
	if err != nil {
		return "", err
	}
	log.Debugf("Generated captcha with id %s and answer %s", id, answ)
	return id, nil
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

func (repo loginRepo) getCaptchaDriver(typ string) (captcha.Driver, error) {
	var driver captcha.Driver
	switch typ {
	case captcha.TypeAudio:
		driver = repo.captcha.DriverAudio.Driver
	default:
		driver = repo.captcha.DriverDigit.Driver
	}
	log.Debugf("Captcha audio generated successfully")
	return driver, nil
}

func (repo loginRepo) getCaptchaData(driver captcha.Driver, id string) (string, error) {
	content := repo.captcha.Store.Get(id, false)
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		return "", err
	}
	return item.EncodeB64string(), nil
}

func (repo loginRepo) getCaptchaAudio(id string) (string, error) {
	log.Debugf("Writing captcha audio to buffer with id %s", id)
	content := repo.captcha.Store.Get(id, false)
	item, err := repo.captcha.DriverAudio.Driver.DrawCaptcha(content)
	if err != nil {
		return "", err
	}
	log.Debugf("Captcha audio generated successfully")
	return item.EncodeB64string(), nil
}

func (repo loginRepo) getCaptchaImage(id string) (string, error) {
	log.Debugf("Writing captcha image to buffer with id %s", id)
	content := repo.captcha.Store.Get(id, false)
	item, err := repo.captcha.DriverDigit.Driver.DrawCaptcha(content)
	if err != nil {
		return "", err
	}
	log.Debugf("Captcha image generated successfully")
	return item.EncodeB64string(), nil
}

type LoginData struct {
	BasisConfig *configs.BasisConfig
	Tokenizer   security.RefreshTokenizer
	Resource    systemdto.ResourceRepo
	Role        systemdto.RoleRepo
	User        systemdto.UserRepo
}

func NewCaptcha(cfg *configs.Captcha) *captcha.Captcha {
	return captcha.NewCaptcha(&captcha.Config{
		DriverDigit: &captcha.DriverDigit{
			Height:   int(cfg.Height),
			Width:    int(cfg.Width),
			Length:   int(cfg.Length),
			MaxSkew:  0.7,
			DotCount: 120,
		},
	})
}

// NewLoginRepo .
func NewLoginRepo(data *LoginData, logger log.KLogger) dto.LoginRepo {
	var err error
	cfg := data.BasisConfig
	// todo: generate random password for root user if not exists
	if cfg.RootUser.RandomPassword {
		passwd := rand.GenerateRandom(12)
		cfg.RootUser.Password, err = hash.Generate(passwd)
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
		captcha:   NewCaptcha(cfg.Captcha),
	}
}

func BufPool() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
}
