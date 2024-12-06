/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"bytes"
	"sync"

	"github.com/LyricTian/captcha"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/httperr"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/basis/dto"
	systemdto "origadmin/application/admin/internal/mods/system/dto"
)

type loginRepo struct {
	bufpool  *sync.Pool
	Captcha  *configs.Captcha
	RootUser *configs.RootUser
	Menu     systemdto.MenuRepo
	Role     systemdto.RoleRepo
	User     systemdto.UserRepo
}

func (repo loginRepo) Login(ctx context.Context, in *dto.LoginRequest) (*dto.LoginResponse, error) {
	data := in.GetData()
	// verify captcha
	if !captcha.VerifyString(data.CaptchaId, data.CaptchaCode) {
		return nil, dto.ErrInvalidCaptchaID
	}

	//ctx = context.NewTag(ctx, logging.TagKeyLogin)

	if root := repo.RootUser; root.GetEnabled() {
		// login by root
		username := root.Username
		if data.Username == username {
			if data.Password != root.Password {
				return nil, dto.ErrInvalidUsernameOrPassword
			}

			userID := root.Id
			ctx = context.NewID(ctx, root.Id)
			log.Info("Login by root")
			return repo.genToken(ctx, userID)
		}
	}

	// get user info
	user, err := repo.User.GetByUserName(ctx, data.Username, "id", "password", "salt", "status")
	switch {
	case err != nil:
		return nil, err
	case user == nil:
		return nil, dto.ErrInvalidUsernameOrPassword
	case user.Status != systemdto.UserStatusActivated:
		return nil, httperr.New("unknown", 400, "User status is not activated, please contact the administrator")
	default:

	}

	// check password
	if err := hash.Compare(user.Password, data.Password, user.Salt); err != nil {
		return nil, dto.ErrInvalidUsernameOrPassword
	}

	userID := user.Id
	username := user.Username
	ctx = context.NewID(ctx, userID)

	// set user cache with role ids
	roleIDs, err := repo.User.GetRoleIDs(ctx, userID)
	if err != nil {
		return nil, err
	}

	//userCache := helpers.UserCache{
	//	Root:     config.C().General.Root.ID == userID,
	//	ID:       userID,
	//	Username: username,
	//	RoleIDs:  roleIDs,
	//}.String()
	//err = repo.Cache.Set(ctx, config.CacheNSForUser,
	//	userID, userCache,
	//	time.Duration(config.C().Dictionary.UserCacheExp)*time.Hour)
	//if err != nil {
	//	logging.FromAbnormal(ctx).LogAttrs(ctx, slog.LevelError, "Failed to set cache", slog.Any("error", err))
	//}

	log.Infof("User %s logged in successfully with role ids: %v", username, roleIDs)
	// generate token
	return repo.genToken(ctx, userID)
}

func (repo loginRepo) CaptchaImage(ctx context.Context, id string, reload bool) (*dto.CaptchaImageResponse, error) {
	log.Debugf("Generating captcha image with id %s and reload %v", id, reload)
	if reload && !captcha.Reload(id) {
		log.Warnf("Captcha id %s not found during reload", id)
		return nil, dto.ErrCaptchaIDNotFound
	}
	buf := repo.getBuf()
	log.Debugf("Writing captcha image to buffer with width %d and height %d", int(repo.Captcha.Width), int(repo.Captcha.Height))
	err := captcha.WriteImage(buf, id, int(repo.Captcha.Width), int(repo.Captcha.Height))
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
	//TODO implement me
	panic("implement me")
}

func (repo loginRepo) CurrentUser(ctx context.Context, in *dto.CurrentUserRequest) (*dto.CurrentUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo loginRepo) Logout(ctx context.Context, in *dto.LogoutRequest) (*dto.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo loginRepo) CaptchaID(ctx context.Context, in *dto.CaptchaIDRequest) (*dto.CaptchaIDResponse, error) {
	return &dto.CaptchaIDResponse{
		Data: captcha.NewLen(int(repo.Captcha.Length)),
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

func (repo loginRepo) genToken(ctx context.Context, id string) (*dto.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewLoginRepo .
func NewLoginRepo(cfg *configs.Captcha, sysMenu systemdto.MenuRepo, sysRole systemdto.RoleRepo, sysUser systemdto.UserRepo, logger log.Logger) dto.LoginRepo {
	return &loginRepo{
		bufpool: BufPool(),
		Captcha: cfg,
		Menu:    sysMenu,
		Role:    sysRole,
		User:    sysUser,
	}
}

func BufPool() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
}
