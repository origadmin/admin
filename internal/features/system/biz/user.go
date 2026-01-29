/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"
	"fmt"
	"strconv"

	kratosLog "github.com/go-kratos/kratos/v2/log"

	"github.com/origadmin/contrib/security"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dto"
)

// UserUseCase is a User use case.
type UserUseCase struct {
	repo   dto.UserRepo
	hasher hash.Crypto
	log    *kratosLog.Helper
}

// NewUserUseCase new a User use case.
func NewUserUseCase(repo dto.UserRepo, hasher hash.Crypto, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		hasher: hasher,
		log:    log.NewHelper(log.With(logger, "module", "system.biz.user")),
	}
}

func (uc *UserUseCase) ListUserResources(ctx context.Context, id int64) ([]*types.Resource, error) {
	return uc.repo.ListResourceByUserID(ctx, id)
}

func (uc *UserUseCase) UpdateUserRoles(ctx context.Context, id int64, roleIDs []int64) ([]*types.Role, error) {
	return uc.repo.AddRoleIDs(ctx, id, roleIDs)
}

func (uc *UserUseCase) UpdateUserStatus(ctx context.Context, id int64, status int8) error {
	return uc.repo.UpdateUserStatus(ctx, id, status)
}

func (uc *UserUseCase) ResetUserPassword(ctx context.Context, id int64, password string) error {
	// TODO
	return nil
}

func (uc *UserUseCase) ListUsers(ctx context.Context, opts ...*dto.UserQueryOption) ([]*types.User, int32, error) {
	return uc.repo.List(ctx, opts...)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64, opts ...*dto.UserQueryOption) (*types.User, error) {
	return uc.repo.Get(ctx, id, opts...)
}

// CreateUser creates a new user, ensuring essential fields have valid default values.
func (uc *UserUseCase) CreateUser(ctx context.Context, in *types.User, password string, opts ...*dto.UserCreateOption) (*types.User, error) {
	// The backend must always enforce data integrity, regardless of frontend behavior.
	if in.Status == 0 {
		in.Status = int32(enums.StatusEnabled)
	}

	// Automatically set audit fields from the context
	p, ok := security.FromContext(ctx)
	if ok {
		authorID, _ := strconv.ParseInt(p.GetID(), 10, 64)
		in.CreateAuthor = authorID
		in.UpdateAuthor = authorID
	}

	hashedPassword, err := uc.hasher.Hash(password)
	if err != nil {
		return nil, err
	}
	fmt.Println("Create new user username:", in.Username, "password:", password)

	return uc.repo.Create(ctx, in, hashedPassword, opts...)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, in *types.User, opts ...*dto.UserUpdateOption) (*types.User, error) {
	return uc.repo.Update(ctx, in, opts...)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
