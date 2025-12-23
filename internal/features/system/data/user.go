/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data is the data access object
package data

import (
	"context"
	"errors"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/data/ent"
	"origadmin/application/admin/internal/features/system/data/ent/user"
	"origadmin/application/admin/internal/features/system/dto"
)

type userRepo struct {
	db *ent.Client
}

func (repo *userRepo) Get(ctx context.Context, id int64, options ...dto.UserQueryOption) (*types.User, error) {
	result, err := repo.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (repo *userRepo) GetByUsername(ctx context.Context, username string, fields ...string) (*types.User, error) {
	result, err := repo.db.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (repo *userRepo) Create(ctx context.Context, u *types.User, options ...dto.UserMutationOption) (*types.User, error) {
	exist, err := repo.db.User.Query().Where(user.UsernameEQ(u.Username)).Exist(ctx)
	if err != nil || exist {
		return nil, errors.New("user already exists")
	}

	create := repo.db.User.Create().
		SetUsername(u.Username).
		SetPassword(u.Password)

	// ... set other fields

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (repo *userRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.User.DeleteOneID(id).Exec(ctx)
}

func (repo *userRepo) Update(ctx context.Context, u *types.User, options ...dto.UserMutationOption) (*types.User, error) {
	update := repo.db.User.UpdateOneID(u.Id)

	if len(u.RoleIds) > 0 {
		update.ClearRoles().AddRoleIDs(u.RoleIds...)
	}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (repo *userRepo) List(ctx context.Context, in *system.ListUsersRequest, options ...dto.UserQueryOption) ([]*types.User, int32, error) {
	query := repo.db.User.Query()

	if in.Title != "" {
		query = query.Where(user.Or(user.UsernameContainsFold(in.Title), user.PhoneContainsFold(in.Title), user.EmailContainsFold(in.Title)))
	}
	//if in.Status != nil {
	//	query = query.Where(user.StatusEQ(int8(*in.Status)))
	//}

	if in.OnlyCount {
		count, err := query.Count(ctx)
		return nil, int32(count), err
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	//query = db.QueryPage(query, in)

	result, err := query.All(ctx)
	return dto.ConvertUsersToUsersPB(result), int32(count), err
}

func (repo *userRepo) AddRoleIDs(ctx context.Context, id int64, roleIDs []int64, options ...dto.UserMutationOption) error {
	return repo.db.User.UpdateOneID(id).AddRoleIDs(roleIDs...).Exec(ctx)
}

func (repo *userRepo) GetRoleIDs(ctx context.Context, id int64) ([]int64, error) {
	ids, err := repo.db.User.Query().Where(user.ID(id)).QueryRoles().IDs(ctx)
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, i := range ids {
		result = append(result, int64(i))
	}
	return result, nil
}

func (repo *userRepo) ListResourceByUserID(ctx context.Context, id int64, options ...dto.UserQueryOption) ([]*types.Resource, error) {
	resources, err := repo.db.User.Query().Where(user.ID(id)).QueryRoles().QueryPermissions().QueryResources().All(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourcesToResourcesPB(resources), nil
}

func (repo *userRepo) UpdateUserStatus(ctx context.Context, id int64, status int32, options ...dto.UserQueryOption) error {
	return repo.db.User.UpdateOneID(id).SetStatus(int8(status)).Exec(ctx)
}

func (repo *userRepo) Current(ctx context.Context, id int64) (*types.User, error) {
	return repo.Get(ctx, id)
}

// NewUserRepo .
func NewUserRepo(d *Data) (dto.UserRepo, error) {
	return &userRepo{db: d.db}, nil
}
