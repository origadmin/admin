/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/features/system/dto"
)

type userRepo struct {
	db *ent.Database
}

// NewUserRepo .
func NewUserRepo(db *ent.Database) dto.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Get(ctx context.Context, id int64, opts ...*dto.UserQueryOptions) (*types.User, error) {
	opt := &dto.UserQueryOptions{}
	if len(opts) > 0 {
		opt = opts[0]
	}

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

func (r *userRepo) Create(ctx context.Context, u *types.User, opts ...*dto.UserCreateOptions) (*types.User, error) {
	exist, err := r.db.User(ctx).Query().Where(user.UsernameEQ(u.Username)).Exist(ctx)
	if err != nil || exist {
		return nil, errors.New("user already exists")
	}

	entUser := dto.ConvertUserPBToUser(u)
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

func (r *userRepo) Update(ctx context.Context, u *types.User, opts ...*dto.UserUpdateOptions) (*types.User, error) {
	entUser := dto.ConvertUserPBToUser(u)
	update := r.db.User(ctx).UpdateOneID(u.Id).SetUser(entUser)

	// ... handle partial updates based on opts

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (r *userRepo) List(ctx context.Context, in *system.ListUsersRequest, opts ...*dto.UserQueryOptions) ([]*types.User, int32, error) {
	opt := &dto.UserQueryOptions{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	query := r.db.User(ctx).Query()

	if in.GetKeyword() != "" {
		query = query.Where(user.Or(user.UsernameContainsFold(in.GetKeyword()), user.PhoneContainsFold(in.GetKeyword()), user.EmailContainsFold(in.GetKeyword())))
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
