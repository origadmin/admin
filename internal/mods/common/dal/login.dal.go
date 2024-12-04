/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "origadmin/application/admin/api/v1/services/common"
	"origadmin/application/admin/internal/mods/system/dto"
)

type loginRepo struct {
	//db *Data
}

func (d loginRepo) CaptchaID(ctx context.Context, in *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	log.Info("CaptchaID")
	return &pb.CaptchaIDResponse{}, nil
}

func (d loginRepo) Login(ctx context.Context, username, password string) (any, error) {
	//TODO implement me
	panic("implement me")
}

// NewLoginRepo .
func NewLoginRepo(logger log.Logger) dto.LoginRepo {
	return &loginRepo{
		//db: db,
	}
}
