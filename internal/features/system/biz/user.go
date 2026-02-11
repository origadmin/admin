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
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dto"
)

// UserUseCase is a User use case.
type UserUseCase struct {
	repo dto.UserRepo
	log  *kratosLog.Helper
}

// NewUserUseCase new a User use case.
func NewUserUseCase(repo dto.UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "system.biz.user")),
	}
}

func (uc *UserUseCase) ListUserResources(ctx context.Context, id int64) ([]*types.Resource, error) {
	return uc.repo.ListResourceByUserID(ctx, id)
}

func (uc *UserUseCase) ListUserRoles(ctx context.Context, id int64) ([]*types.Role, error) {
	return uc.repo.ListRoleByUserID(ctx, id)
}

func (uc *UserUseCase) ListUserViews(ctx context.Context, id int64) ([]*types.View, error) {
	return uc.repo.ListViewByUserID(ctx, id)
}

func (uc *UserUseCase) ListUserPermissions(ctx context.Context, id int64) ([]*types.Permission, error) {
	return uc.repo.ListPermissionByUserID(ctx, id)
}

func (uc *UserUseCase) UpdateUserRoles(ctx context.Context, id int64, roleIDs []int64) ([]*types.Role, error) {
	return uc.repo.AddRoleIDs(ctx, id, roleIDs)
}

func (uc *UserUseCase) UpdateUserStatus(ctx context.Context, id int64, status int8) error {
	return uc.repo.UpdateUserStatus(ctx, id, status)
}

// UpdateUserPassword updates a user's password hash directly.
// Note: The password verification and hashing should be handled by the identity service.
func (uc *UserUseCase) UpdateUserPassword(ctx context.Context, userID int64, hashedPassword string) error {
	// Update the password in the repository
	return uc.repo.ChangeUserPassword(ctx, userID, hashedPassword)
}

func (uc *UserUseCase) ListUsers(ctx context.Context, opts ...*dto.UserQueryOption) ([]*types.User, int32, error) {
	return uc.repo.List(ctx, opts...)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id int64, opts ...*dto.UserQueryOption) (*types.User, error) {
	return uc.repo.Get(ctx, id, opts...)
}

// CreateUser creates a new user, ensuring essential fields have valid default values.
// The password argument must be already hashed.
func (uc *UserUseCase) CreateUser(ctx context.Context, in *types.User, hashedPassword string, opts ...*dto.UserCreateOption) (*types.User, error) {
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

	fmt.Println("Create new user username:", in.Username, "hashedPassword:", hashedPassword)

	return uc.repo.Create(ctx, in, hashedPassword, opts...)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, in *types.User, opts ...*dto.UserUpdateOption) (*types.User, error) {
	return uc.repo.Update(ctx, in, opts...)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
