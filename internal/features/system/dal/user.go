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

func (r *userRepo) Get(ctx context.Context, id int64, opts ...dto.UserQueryOption) (*types.User, error) {
	result, err := r.db.User(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) Create(ctx context.Context, u *types.User, opts ...dto.UserMutationOption) (*types.User, error) {
	exist, err := r.db.User(ctx).Query().Where(user.UsernameEQ(u.Username)).Exist(ctx)
	if err != nil || exist {
		return nil, errors.New("user already exists")
	}

	create := r.db.User(ctx).Create().
		SetUsername(u.Username).
		SetPassword(u.Password)

	// ... set other fields

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	return r.db.User(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *userRepo) Update(ctx context.Context, u *types.User, opts ...dto.UserMutationOption) (*types.User, error) {
	update := r.db.User(ctx).UpdateOneID(u.Id)

	//if len(u.RoleIds) > 0 {
	//	update.ClearRoles().AddRoleIDs(u.RoleIds...)
	//}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (r *userRepo) List(ctx context.Context, in *system.ListUsersRequest, opts ...dto.UserQueryOption) ([]*types.User, int32, error) {
	query := r.db.User(ctx).Query()

	if in.GetKeyword() != "" {
		query = query.Where(user.Or(user.UsernameContainsFold(in.GetKeyword()), user.PhoneContainsFold(in.GetKeyword()), user.EmailContainsFold(in.GetKeyword())))
	}
	//if in.Status != nil {
	//	query = query.Where(user.StatusEQ(int8(*in.Status)))
	//}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	//query = db.QueryPage(query, in)

	result, err := query.All(ctx)
	return dto.ConvertUsersToUsersPB(result), int32(count), err
}

func (r *userRepo) AddRoleIDs(ctx context.Context, id int64, roleIDs []int64, opts ...dto.UserMutationOption) error {
	return r.db.User(ctx).UpdateOneID(id).AddRoleIDs(roleIDs...).Exec(ctx)
}

func (r *userRepo) GetByUsername(ctx context.Context, username string, fields ...string) (*types.User, error) {
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

func (r *userRepo) ListResourceByUserID(ctx context.Context, id int64, opts ...dto.UserQueryOption) ([]*types.Resource, error) {
	resources, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().QueryPermissions().QueryResources().All(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourcesToResourcesPB(resources), nil
}

func (r *userRepo) UpdateUserStatus(ctx context.Context, id int64, status int32, opts ...dto.UserQueryOption) error {
	return r.db.User(ctx).UpdateOneID(id).SetStatus(int8(status)).Exec(ctx)
}

func (r *userRepo) Current(ctx context.Context, id int64) (*types.User, error) {
	return r.Get(ctx, id)
}
