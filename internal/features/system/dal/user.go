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
	"origadmin/application/admin/internal/helpers/contextutil"
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
	opt := repo.FirstOrDefault(opts...)
	query := r.db.User(ctx).Query().Where(user.ID(id))

	if opt.WithRoles {
		query.WithRoles()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, user.ValidColumn, user.FieldID,
			new(types.User))
		query = s.UserQuery
	}

	result, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) Create(ctx context.Context, u *types.User, password string, opts ...*dto.UserCreateOption) (*types.User, error) {
	opt := repo.FirstOrDefault(opts...)
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

	if opt.WithRoleIDs != nil {
		create.AddRoleIDs(opt.WithRoleIDs...)
	}

	if contextutil.IsSystemUser(ctx) {
		create.SetIsSystem(true)
	}

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	// User uses soft-delete, so we only set delete_time.
	// Associations (user_roles, user_positions, user_departments) are NOT cleared
	// to allow for potential restoration of the user with their original associations.
	return r.db.User(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *userRepo) Restore(ctx context.Context, id int64) error {
	// We must use `Update` which bypasses the soft-delete interceptor
	// to restore a soft-deleted record.
	result, err := r.db.User(ctx).
		Update().
		Where(user.ID(id)).
		ClearDeleteTime().
		Save(ctx)
	if err != nil {
		return err
	}
	// If no rows were affected, the user doesn't exist (including soft-deleted ones)
	if result == 0 {
		return &ent.NotFoundError{}
	}
	return nil
}

func (r *userRepo) Update(ctx context.Context, u *types.User, opts ...*dto.UserUpdateOption) (*types.User, error) {
	var updatedUser *ent.User
	err := r.db.Tx(ctx, func(tx context.Context) error {
		opt := repo.FirstOrDefault(opts...)
		entUser := dto.ConvertUserPBToUser(u)
		update := r.db.User(tx).UpdateOneID(u.Id)

		updateCols := db.UpdateFields(opt.UpdateMask, user.ValidColumn, u)
		if len(updateCols) > 0 {
			// If a field mask is present, update only the specified fields, including zero values.
			update.SetUser(entUser, updateCols...)
		} else {
			// If no field mask, skip zero values to prevent accidental clearing of fields.
			update.SetUserSkipZero(entUser)
		}
		update.ClearRoles()
		if len(opt.WithRoleIDs) > 0 {
			update.AddRoleIDs(opt.WithRoleIDs...)
		}

		var err error
		updatedUser, err = update.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserToUserPB(updatedUser), nil
}

func (r *userRepo) List(ctx context.Context, opts ...*dto.UserQueryOption) ([]*types.User, int32, error) {
	opt := repo.FirstOrDefault(opts...)
	query := r.db.User(ctx).Query()

	if opt.Keyword != "" {
		query.Where(user.Or(user.UsernameContainsFold(opt.Keyword), user.PhoneContainsFold(opt.Keyword), user.EmailContainsFold(opt.Keyword)))
	}

	if opt.WithRoles {
		query.WithRoles()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, user.ValidColumn, user.FieldID,
			new(types.User))
		query = s.UserQuery
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertUsersToUsersPB(result), count, err
}

func (r *userRepo) AddRoleIDs(ctx context.Context, id int64, roleIDs []int64) ([]*types.Role, error) {
	var roles []*ent.Role
	err := r.db.Tx(ctx, func(tx context.Context) error {
		u, err := r.db.User(tx).UpdateOneID(id).ClearRoles().AddRoleIDs(roleIDs...).Save(ctx)
		if err != nil {
			return err
		}
		roles, err = u.QueryRoles().All(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return dto.ConvertRolesToRolesPB(roles), nil
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
