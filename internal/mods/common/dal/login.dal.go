/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"bytes"
	"context"
	"sync"

	"github.com/LyricTian/captcha"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/httperr"

	"origadmin/application/admin/internal/configs"
	"origadmin/application/admin/internal/mods/common/dto"
	systemdto "origadmin/application/admin/internal/mods/system/dto"
)

type loginRepo struct {
	bufpool *sync.Pool
	Menu    systemdto.MenuRepo
	Captcha *configs.Captcha
}

func (repo loginRepo) Login(ctx context.Context, username, password string) (*dto.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo loginRepo) CaptchaImage(ctx context.Context, id string, reload bool) (*dto.CaptchaImageResponse, error) {
	log.Debugf("Generating captcha image with id %s and reload %v", id, reload)
	if reload && !captcha.Reload(id) {
		log.Warnf("Captcha id %s not found during reload", id)
		return nil, httperr.NotFound("Captcha id not found")
	}
	buf := repo.getBuf()
	log.Debugf("Writing captcha image to buffer with width %d and height %d", int(repo.Captcha.Width), int(repo.Captcha.Height))
	err := captcha.WriteImage(buf, id, int(repo.Captcha.Width), int(repo.Captcha.Height))
	if err != nil {
		log.Errorf("Error writing captcha image: %v", err)
		if errors.Is(err, captcha.ErrNotFound) {
			log.Warnf("Captcha id %s not found", id)
			return nil, httperr.NotFound("Captcha id not found")
		}
		return nil, err
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

// NewLoginRepo .
func NewLoginRepo(cfg *configs.Captcha, sysMenu systemdto.MenuRepo, logger log.Logger) dto.LoginRepo {
	return &loginRepo{
		bufpool: &sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
		Captcha: cfg,
		Menu:    sysMenu,
	}
}
