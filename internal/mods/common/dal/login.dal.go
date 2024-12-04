/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "origadmin/application/admin/api/v1/services/common"
	"origadmin/application/admin/internal/mods/common/dto"
)

type loginRepo struct {
	//db *Data
}

func (d loginRepo) Login(ctx context.Context, username, password string) (*dto.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d loginRepo) CaptchaImage(ctx context.Context, in *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d loginRepo) CurrentMenus(ctx context.Context, in *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d loginRepo) CurrentUser(ctx context.Context, in *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d loginRepo) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d loginRepo) CaptchaID(ctx context.Context, in *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewLoginRepo .
func NewLoginRepo(logger log.Logger) dto.LoginRepo {
	return &loginRepo{
		//db: db,
	}
}
