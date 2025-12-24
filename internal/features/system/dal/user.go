/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/repo"
)

type userRepo struct {
	db *ent.Database
}

// NewUserRepo .
func NewUserRepo(db *ent.Database) dto.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Get(ctx context.Context, id int64, opts ...*dto.UserQueryOption) (*types.User, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.User(ctx).Query().Where(user.ID(id))

	if opt.WithRoles {
		query.WithRoles()
	}

	result, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) Create(ctx context.Context, u *types.User, password string, opts ...*dto.UserCreateOption) (*types.User, error) {
	exist, err := r.db.User(ctx).Query().Where(user.UsernameEQ(u.Username)).Exist(ctx)
	if err != nil || exist {
		return nil, errors.New("user already exists")
	}

	entUser := dto.ConvertUserPBToUser(u)
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	entUser.UUID = uuid.String()
	if password != "" {
		entUser.EncryptedPassword = password
	}
	create := r.db.User(ctx).Create().SetUser(entUser)

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	return r.db.User(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *userRepo) Update(ctx context.Context, u *types.User, opts ...*dto.UserUpdateOption) (*types.User, error) {
	entUser := dto.ConvertUserPBToUser(u)
	update := r.db.User(ctx).UpdateOneID(u.Id).SetUser(entUser)

	// ... handle partial updates based on opts

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (r *userRepo) List(ctx context.Context, opts ...*dto.UserQueryOption) ([]*types.User, int32, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.User(ctx).Query()

	if opt.Keyword != "" {
		query.Where(user.Or(user.UsernameContainsFold(opt.Keyword), user.PhoneContainsFold(opt.Keyword), user.EmailContainsFold(opt.Keyword)))
	}

	if opt.Page > 0 && opt.PageSize > 0 {
		query.Offset((opt.Page - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if opt.WithRoles {
		query.WithRoles()
	}

	result, err := query.All(ctx)
	return dto.ConvertUsersToUsersPB(result), int32(count), err
}

func (r *userRepo) AddRoleIDs(ctx context.Context, id int64, roleIDs []int64) error {
	return r.db.User(ctx).UpdateOneID(id).AddRoleIDs(roleIDs...).Exec(ctx)
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*types.User, error) {
	result, err := r.db.User(ctx).Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) GetRoleIDs(ctx context.Context, id int64) ([]int64, error) {
	ids, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().IDs(ctx)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (r *userRepo) ListResourceByUserID(ctx context.Context, id int64) ([]*types.Resource, error) {
	resources, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().QueryPermissions().QueryResources().All(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourcesToResourcesPB(resources), nil
}

func (r *userRepo) UpdateUserStatus(ctx context.Context, id int64, status int32) error {
	return r.db.User(ctx).UpdateOneID(id).SetStatus(int8(status)).Exec(ctx)
}
