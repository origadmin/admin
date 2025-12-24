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

func (uc *UserUseCase) ListUserResources(ctx context.Context, id int64) ([]*types.Resource, error) {
	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return nil, fmt.Errorf("ListUserResources not implemented")
}

func (uc *UserUseCase) UpdateUserRoles(ctx context.Context, id int64, roleIDs []int64) error {
	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return fmt.Errorf("UpdateUserRoles not implemented")
}

func (uc *UserUseCase) UpdateUserStatus(ctx context.Context, id int64, status int32) error {
	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return fmt.Errorf("UpdateUserStatus not implemented")
}

func (uc *UserUseCase) ResetUserPassword(ctx context.Context, id int64, password string) error {
	// TODO
	return nil
}

func (uc *UserUseCase) ListUsers(ctx context.Context, in *system.ListUsersRequest) ([]*types.User, int32, error) {
	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return nil, 0, fmt.Errorf("ListUsers not implemented")
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*types.User, error) {
	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return nil, fmt.Errorf("GetUser not implemented")
}

func (uc *UserUseCase) CreateUser(ctx context.Context, in *types.User, password string) (*types.User, error) {
	hashedPassword, err := uc.hasher.Hash(password)
	if err != nil {
		return nil, err
	}
	in.Password = hashedPassword

	fmt.Println("Create new user username:", in.Username, "password:", password)

	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return nil, fmt.Errorf("CreateUser not implemented")
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, in *types.User) (*types.User, error) {
	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return nil, fmt.Errorf("UpdateUser not implemented")
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	// This method is not in the biz.UserRepo interface, so it needs to be implemented in the data layer
	// and added to the interface. For now, we'll return an error.
	return fmt.Errorf("DeleteUser not implemented")
}

// NewUserUseCase new a User use case.
func NewUserUseCase(repo dto.UserRepo, hasher hash.Crypto) (*UserUseCase, error) {
	return &UserUseCase{repo: repo, hasher: hasher}, nil
}
