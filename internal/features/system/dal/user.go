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
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type userRepo struct {
	db *ent.Database
}

// NewUserRepo .
func NewUserRepo(database *ent.Database) dto.UserRepo {
	return &userRepo{db: database}
}

func (r *userRepo) Get(ctx context.Context, id int64, opts ...*dto.UserQueryOption) (*types.User, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.User(ctx).Query().Where(user.ID(id))

	if opt.WithRoles {
		query.WithRoles()
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, user.ValidColumn, user.FieldID, new(types.User))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
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
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	entUser.UUID = uid.String()
	if password != "" {
		entUser.EncryptedPassword = password
	}
	create := r.db.User(ctx).Create().SetUserSkipZero(entUser)
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
	opt := repo.GetFirstOption(opts...)
	entUser := dto.ConvertUserPBToUser(u)
	update := r.db.User(ctx).UpdateOneID(u.Id)

	updateCols := db.UpdateFields(opt.UpdateMask, user.ValidColumn, u)
	if len(updateCols) > 0 {
		// If a field mask is present, update only the specified fields, including zero values.
		update.SetUser(entUser, updateCols...)
	} else {
		// If no field mask, skip zero values to prevent accidental clearing of fields.
		update.SetUserSkipZero(entUser)
	}

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

	if opt.WithRoles {
		query.WithRoles()
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, user.ValidColumn, user.FieldID, new(types.User))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	// Sorting logic: only apply sorting from request if it's not already handled by a page token.
	if !opt.SortFromToken {
		if len(opt.OrderBy) > 0 {
			orders := db.OrderBy[user.OrderOption](opt.OrderBy)
			if len(orders) > 0 {
				query.Order(orders...)
			}
		}
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertUsersToUsersPB(result), count, err
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

func (r *userRepo) UpdateUserStatus(ctx context.Context, id int64, status int8) error {
	return r.db.User(ctx).UpdateOneID(id).SetStatus(enums.Status(status)).Exec(ctx)
}
