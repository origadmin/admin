/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/mods/system/dto"
)

type loginDal struct {
	db *Data
}

func (l loginDal) Login(ctx context.Context, username, password string) (any, error) {
	//TODO implement me
	panic("implement me")
}

// NewLoginDal .
func NewLoginDal(db *Data, logger log.Logger) dto.LoginRepo {
	return &loginDal{
		db: db,
	}
}
