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
	"origadmin/application/admin/internal/features/system/dal"
)

// UserUseCase is a User use case.
type UserUseCase struct {
	repo   dal.UserRepo
	hasher hash.Crypto
}

func (uc *UserUseCase) ListUserResources(ctx context.Context, id int64) ([]*types.Resource, error) {
	result, err := uc.repo.ListResourceByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *UserUseCase) UpdateUserRoles(ctx context.Context, id int64, roleIDs []int64) error {
	err := uc.repo.AddRoleIDs(ctx, id, roleIDs)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) UpdateUserStatus(ctx context.Context, id int64, status int32) error {
	err := uc.repo.UpdateUserStatus(ctx, id, status)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) ResetUserPassword(ctx context.Context, id int64, password string) error {
	// TODO
	return nil
}

func (uc *UserUseCase) ListUsers(ctx context.Context, in *system.ListUsersRequest) ([]*types.User, int32, error) {
	result, total, err := uc.repo.List(ctx, in)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*types.User, error) {
	result, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *UserUseCase) CreateUser(ctx context.Context, in *types.User, password string) (*types.User, error) {
	hashedPassword, err := uc.hasher.Hash(password)
	if err != nil {
		return nil, err
	}
	in.Password = hashedPassword

	fmt.Println("Create new user username:", in.Username, "password:", password)

	result, err := uc.repo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, in *types.User) (*types.User, error) {
	result, err := uc.repo.Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

// NewUserUseCase new a User use case.
func NewUserUseCase(repo dal.UserRepo, hasher hash.Crypto) (*UserUseCase, error) {
	return &UserUseCase{repo: repo, hasher: hasher}, nil
}
