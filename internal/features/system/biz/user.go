/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/dto"
)

// UserUseCase is a User use case.
type UserUseCase struct {
	repo   dto.UserRepo
	hasher hash.Crypto
}

// NewUserUseCase new a User use case.
func NewUserUseCase(repo dto.UserRepo, hasher hash.Crypto) *UserUseCase {
	return &UserUseCase{repo: repo, hasher: hasher}
}

func (uc *UserUseCase) ListUserResources(ctx context.Context, id int64) ([]*types.Resource, error) {
	return uc.repo.ListResourceByUserID(ctx, id)
}

func (uc *UserUseCase) UpdateUserRoles(ctx context.Context, id int64, roleIDs []int64) error {
	return uc.repo.AddRoleIDs(ctx, id, roleIDs)
}

func (uc *UserUseCase) UpdateUserStatus(ctx context.Context, id int64, status int8) error {
	return uc.repo.UpdateUserStatus(ctx, id, status)
}

func (uc *UserUseCase) ResetUserPassword(ctx context.Context, id int64, password string) error {
	// TODO
	return nil
}

func (uc *UserUseCase) ListUsers(ctx context.Context, in *system.ListUsersRequest) ([]*types.User, int32, error) {
	queryOpt := dto.ListUsersRequestToQueryOption(in)
	return uc.repo.List(ctx, queryOpt)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*types.User, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *UserUseCase) CreateUser(ctx context.Context, in *types.User, password string) (*types.User, error) {
	hashedPassword, err := uc.hasher.Hash(password)
	if err != nil {
		return nil, err
	}
	fmt.Println("Create new user username:", in.Username, "password:", password)

	return uc.repo.Create(ctx, in, hashedPassword)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, in *types.User) (*types.User, error) {
	return uc.repo.Update(ctx, in)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
