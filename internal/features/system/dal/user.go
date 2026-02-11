/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/permission"
	"origadmin/application/admin/internal/data/entity/ent/resource"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/userprofile"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/contextutil"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type userRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewUserRepo .
func NewUserRepo(database *ent.Database, logger log.Logger) dto.UserRepo {
	return &userRepo{
		db:  database,
		log: log.NewHelper(log.With(logger, "module", "dal.user")),
	}
}

// ChangeUserPassword resets a user's password (admin operation)
func (r *userRepo) ChangeUserPassword(ctx context.Context, id int64, hashedPassword string) error {
	r.log.WithContext(ctx).Debugw("msg", "ChangeUserPassword", "id", id)
	err := r.db.User(ctx).UpdateOneID(id).SetEncryptedPassword(hashedPassword).Exec(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ChangeUserPassword", "err", err)
	}
	return err
}

func (r *userRepo) UpdateUserProfile(ctx context.Context, id int64, profile *types.UserProfile) error {
	r.log.WithContext(ctx).Debugw("msg", "UpdateUserProfile", "id", id)
	// Update or create profile
	profileEnt, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryProfile().Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		r.log.WithContext(ctx).Errorw("msg", "UpdateUserProfile.QueryProfile", "err", err)
		return err
	}

	if profileEnt != nil {
		_, err = r.db.UserProfile(ctx).UpdateOneID(profileEnt.ID).
			SetName(profile.Name).
			SetAvatar(profile.Avatar).
			SetGender(userprofile.Gender(profile.Gender)).
			SetDepartment(profile.Department).
			SetRemark(profile.Remark).
			Save(ctx)
	} else {
		_, err = r.db.UserProfile(ctx).Create().
			SetName(profile.Name).
			SetAvatar(profile.Avatar).
			SetGender(userprofile.Gender(profile.Gender)).
			SetDepartment(profile.Department).
			SetRemark(profile.Remark).
			SetUserID(id).
			Save(ctx)
	}
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "UpdateUserProfile.Save", "err", err)
	}
	return err
}

func (r *userRepo) GetUserProfile(ctx context.Context, id int64) (*types.UserProfile, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetUserProfile", "id", id)
	result, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryProfile().Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			r.log.WithContext(ctx).Warnw("msg", "GetUserProfile.QueryProfile", "err", err)
		}
		// Return default profile if not exists
		return &types.UserProfile{
			Name:       "",
			Avatar:     "",
			Gender:     "",
			Department: "",
			Remark:     "",
		}, nil
	}

	return &types.UserProfile{
		Name:       result.Name,
		Avatar:     result.Avatar,
		Gender:     string(result.Gender),
		Department: result.Department,
		Remark:     result.Remark,
	}, nil
}

func (r *userRepo) UpdateUserSetting(ctx context.Context, id int64, setting *types.UserSetting) error {
	r.log.WithContext(ctx).Debugw("msg", "UpdateUserSetting", "id", id)
	settingEnt, err := r.db.User(ctx).Query().Where(user.ID(id)).QuerySetting().Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		r.log.WithContext(ctx).Errorw("msg", "UpdateUserSetting.QuerySetting", "err", err)
		return err
	}

	// Convert map[string]string to map[string]interface{}
	preferences := make(map[string]interface{})
	for k, v := range setting.Preferences {
		preferences[k] = v
	}

	if settingEnt != nil {
		_, err = r.db.UserSetting(ctx).UpdateOneID(settingEnt.ID).
			SetTheme(setting.Theme).
			SetLanguage(setting.Language).
			SetTimezone(setting.Timezone).
			SetPreferences(preferences).
			Save(ctx)
	} else {
		_, err = r.db.UserSetting(ctx).Create().
			SetTheme(setting.Theme).
			SetLanguage(setting.Language).
			SetTimezone(setting.Timezone).
			SetPreferences(preferences).
			SetUserID(id).
			Save(ctx)
	}
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "UpdateUserSetting.Save", "err", err)
	}
	return err
}

func (r *userRepo) GetUserSetting(ctx context.Context, id int64) (*types.UserSetting, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetUserSetting", "id", id)
	result, err := r.db.User(ctx).Query().Where(user.ID(id)).QuerySetting().Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			r.log.WithContext(ctx).Warnw("msg", "GetUserSetting.QuerySetting", "err", err)
		}
		// Return default setting if not exists
		return &types.UserSetting{
			Theme:       "",
			Language:    "",
			Timezone:    "",
			Preferences: make(map[string]string),
		}, nil
	}

	// Convert map[string]interface{} to map[string]string
	preferences := make(map[string]string)
	for k, v := range result.Preferences {
		if str, ok := v.(string); ok {
			preferences[k] = str
		}
	}

	return &types.UserSetting{
		Theme:       result.Theme,
		Language:    result.Language,
		Timezone:    result.Timezone,
		Preferences: preferences,
	}, nil
}

func (r *userRepo) DeleteRoleIDs(ctx context.Context, id int64, roleIDs []int64) error {
	r.log.WithContext(ctx).Debugw("msg", "DeleteRoleIDs", "id", id, "roleIDs", roleIDs)
	_, err := r.db.User(ctx).UpdateOneID(id).RemoveRoleIDs(roleIDs...).Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "DeleteRoleIDs", "err", err)
	}
	return err
}

func (r *userRepo) AddResourceIDs(ctx context.Context, id int64, resourceIDs []int64) ([]*types.Resource, error) {
	r.log.WithContext(ctx).Debugw("msg", "AddResourceIDs", "id", id, "resourceIDs", resourceIDs)
	// Note: Resources are associated with Permissions, not directly with Users
	// This method adds resources to the user's permissions
	var resources []*ent.Resource
	err := r.db.Tx(ctx, func(tx context.Context) error {
		// Get user's current roles
		roles, err := r.db.User(tx).Query().Where(user.ID(id)).QueryRoles().All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddResourceIDs.QueryRoles", "err", err)
			return err
		}

		// Add resources to all permissions of user's roles
		for _, role := range roles {
			perms, err := role.QueryPermissions().All(tx)
			if err != nil {
				r.log.WithContext(ctx).Errorw("msg", "AddResourceIDs.QueryPermissions", "err", err)
				return err
			}
			for _, perm := range perms {
				_, err = perm.Update().AddResourceIDs(resourceIDs...).Save(tx)
				if err != nil {
					r.log.WithContext(ctx).Errorw("msg", "AddResourceIDs.AddResourceIDs", "err", err)
					return err
				}
			}
		}

		// Return the added resources
		resources, err = r.db.Resource(tx).Query().Where(resource.IDIn(resourceIDs...)).All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddResourceIDs.QueryResources", "err", err)
		}
		return err
	})
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "AddResourceIDs.Tx", "err", err)
		return nil, err
	}
	return dto.ConvertResourcesToResourcesPB(resources), nil
}

func (r *userRepo) GetResourceIDs(ctx context.Context, id int64) ([]int64, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetResourceIDs", "id", id)
	// Get all resource IDs from user's roles -> permissions -> resources
	ids, err := r.db.User(ctx).Query().Where(user.ID(id)).
		QueryRoles().
		QueryPermissions().
		QueryResources().
		IDs(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetResourceIDs", "err", err)
		return nil, err
	}
	return ids, nil
}

func (r *userRepo) DeleteResourceIDs(ctx context.Context, id int64, resourceIDs []int64) error {
	r.log.WithContext(ctx).Debugw("msg", "DeleteResourceIDs", "id", id, "resourceIDs", resourceIDs)
	// Remove resources from user's permissions
	return r.db.Tx(ctx, func(tx context.Context) error {
		roles, err := r.db.User(tx).Query().Where(user.ID(id)).QueryRoles().All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "DeleteResourceIDs.QueryRoles", "err", err)
			return err
		}

		for _, role := range roles {
			perms, err := role.QueryPermissions().All(tx)
			if err != nil {
				r.log.WithContext(ctx).Errorw("msg", "DeleteResourceIDs.QueryPermissions", "err", err)
				return err
			}
			for _, perm := range perms {
				_, err = perm.Update().RemoveResourceIDs(resourceIDs...).Save(tx)
				if err != nil {
					r.log.WithContext(ctx).Errorw("msg", "DeleteResourceIDs.RemoveResourceIDs", "err", err)
					return err
				}
			}
		}
		return nil
	})
}

func (r *userRepo) AddViewIDs(ctx context.Context, id int64, viewIDs []int64) ([]*types.View, error) {
	r.log.WithContext(ctx).Debugw("msg", "AddViewIDs", "id", id, "viewIDs", viewIDs)
	// Add views to user's permissions
	var views []*ent.View
	err := r.db.Tx(ctx, func(tx context.Context) error {
		roles, err := r.db.User(tx).Query().Where(user.ID(id)).QueryRoles().All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddViewIDs.QueryRoles", "err", err)
			return err
		}

		for _, role := range roles {
			perms, err := role.QueryPermissions().All(tx)
			if err != nil {
				r.log.WithContext(ctx).Errorw("msg", "AddViewIDs.QueryPermissions", "err", err)
				return err
			}
			for _, perm := range perms {
				_, err = perm.Update().AddViewIDs(viewIDs...).Save(tx)
				if err != nil {
					r.log.WithContext(ctx).Errorw("msg", "AddViewIDs.AddViewIDs", "err", err)
					return err
				}
			}
		}

		views, err = r.db.View(tx).Query().Where(func(s *sql.Selector) {
			sql.In(s.C(view.FieldID), viewIDs)
		}).All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddViewIDs.QueryViews", "err", err)
		}
		return err
	})
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "AddViewIDs.Tx", "err", err)
		return nil, err
	}
	return dto.ConvertViewsToViewsPB(views), nil
}

func (r *userRepo) GetViewIDs(ctx context.Context, id int64) ([]int64, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetViewIDs", "id", id)
	ids, err := r.db.User(ctx).Query().Where(user.ID(id)).
		QueryRoles().
		QueryPermissions().
		QueryViews().
		IDs(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetViewIDs", "err", err)
		return nil, err
	}
	return ids, nil
}

func (r *userRepo) DeleteViewIDs(ctx context.Context, id int64, viewIDs []int64) error {
	r.log.WithContext(ctx).Debugw("msg", "DeleteViewIDs", "id", id, "viewIDs", viewIDs)
	return r.db.Tx(ctx, func(tx context.Context) error {
		roles, err := r.db.User(tx).Query().Where(user.ID(id)).QueryRoles().All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "DeleteViewIDs.QueryRoles", "err", err)
			return err
		}

		for _, role := range roles {
			perms, err := role.QueryPermissions().All(tx)
			if err != nil {
				r.log.WithContext(ctx).Errorw("msg", "DeleteViewIDs.QueryPermissions", "err", err)
				return err
			}
			for _, perm := range perms {
				_, err = perm.Update().RemoveViewIDs(viewIDs...).Save(tx)
				if err != nil {
					r.log.WithContext(ctx).Errorw("msg", "DeleteViewIDs.RemoveViewIDs", "err", err)
					return err
				}
			}
		}
		return nil
	})
}

func (r *userRepo) AddPermissionIDs(ctx context.Context, id int64, permissionIDs []int64) ([]*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "AddPermissionIDs", "id", id, "permissionIDs", permissionIDs)
	// Add permissions to user's roles
	var permissions []*ent.Permission
	err := r.db.Tx(ctx, func(tx context.Context) error {
		roles, err := r.db.User(tx).Query().Where(user.ID(id)).QueryRoles().All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddPermissionIDs.QueryRoles", "err", err)
			return err
		}

		for _, role := range roles {
			_, err = role.Update().AddPermissionIDs(permissionIDs...).Save(tx)
			if err != nil {
				r.log.WithContext(ctx).Errorw("msg", "AddPermissionIDs.AddPermissionIDs", "err", err)
				return err
			}
		}

		permissions, err = r.db.Permission(tx).Query().Where(func(s *sql.Selector) {
			sql.In(s.C(permission.FieldID), permissionIDs)
		}).All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddPermissionIDs.QueryPermissions", "err", err)
		}
		return err
	})
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "AddPermissionIDs.Tx", "err", err)
		return nil, err
	}
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

func (r *userRepo) GetPermissionIDs(ctx context.Context, id int64) ([]int64, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetPermissionIDs", "id", id)
	ids, err := r.db.User(ctx).Query().Where(user.ID(id)).
		QueryRoles().
		QueryPermissions().
		IDs(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetPermissionIDs", "err", err)
		return nil, err
	}
	return ids, nil
}

func (r *userRepo) DeletePermissionIDs(ctx context.Context, id int64, permissionIDs []int64) error {
	r.log.WithContext(ctx).Debugw("msg", "DeletePermissionIDs", "id", id, "permissionIDs", permissionIDs)
	return r.db.Tx(ctx, func(tx context.Context) error {
		roles, err := r.db.User(tx).Query().Where(user.ID(id)).QueryRoles().All(tx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "DeletePermissionIDs.QueryRoles", "err", err)
			return err
		}

		for _, role := range roles {
			_, err = role.Update().RemovePermissionIDs(permissionIDs...).Save(tx)
			if err != nil {
				r.log.WithContext(ctx).Errorw("msg", "DeletePermissionIDs.RemovePermissionIDs", "err", err)
				return err
			}
		}
		return nil
	})
}

func (r *userRepo) GetByEmail(ctx context.Context, mail string) (*types.User, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetByEmail", "mail", mail)
	result, err := r.db.User(ctx).Query().Where(user.EmailEQ(mail)).Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetByEmail", "err", err)
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) GetByPhone(ctx context.Context, s string) (*types.User, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetByPhone", "phone", s)
	result, err := r.db.User(ctx).Query().Where(user.PhoneEQ(s)).Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetByPhone", "err", err)
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) GetUserAndPassword(ctx context.Context, id int64) (*types.User, string, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetUserAndPassword", "id", id)
	result, err := r.db.User(ctx).Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetUserAndPassword", "err", err)
		return nil, "", err
	}
	return dto.ConvertUserToUserPB(result), result.EncryptedPassword, nil
}

func (r *userRepo) Get(ctx context.Context, id int64, opts ...*dto.UserQueryOption) (*types.User, error) {
	r.log.WithContext(ctx).Debugw("msg", "Get", "id", id)
	opt := repo.FirstOrDefault(opts...)
	query := r.db.User(ctx).Query().Where(user.ID(id))

	if opt.WithRoles {
		query.WithRoles()
	}
	if opt.WithProfile {
		query.WithProfile()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, user.ValidColumn, user.FieldID,
			new(types.User))
		query = s.UserQuery
	}

	result, err := query.Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Get", "err", err)
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) Create(ctx context.Context, u *types.User, password string, opts ...*dto.UserCreateOption) (*types.User, error) {
	r.log.WithContext(ctx).Debugw("msg", "Create", "user", u)
	opt := repo.FirstOrDefault(opts...)
	exist, err := r.db.User(ctx).Query().Where(user.UsernameEQ(u.Username)).Exist(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Create.Exist", "err", err)
		return nil, err
	}
	if exist {
		err = errors.New("user already exists")
		r.log.WithContext(ctx).Warnw("msg", "Create.Exist", "err", err)
		return nil, err
	}

	entUser := dto.ConvertUserPBToUser(u)
	uid, err := uuid.NewRandom()
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Create.NewRandom", "err", err)
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
		r.log.WithContext(ctx).Errorw("msg", "Create.Save", "err", err)
		return nil, err
	}
	return dto.ConvertUserToUserPB(saved), nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	r.log.WithContext(ctx).Debugw("msg", "Delete", "id", id)
	// User uses soft-delete, so we only set delete_time.
	// Associations (user_roles, user_positions, user_departments) are NOT cleared
	// to allow for potential restoration of the user with their original associations.
	err := r.db.User(ctx).DeleteOneID(id).Exec(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Delete", "err", err)
	}
	return err
}

func (r *userRepo) Restore(ctx context.Context, id int64) error {
	r.log.WithContext(ctx).Debugw("msg", "Restore", "id", id)
	// We must use `Update` which bypasses the soft-delete interceptor
	// to restore a soft-deleted record.
	result, err := r.db.User(ctx).
		Update().
		Where(user.ID(id)).
		ClearDeleteTime().
		Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Restore.Save", "err", err)
		return err
	}
	// If no rows were affected, the user doesn't exist (including soft-deleted ones)
	if result == 0 {
		err = &ent.NotFoundError{}
		r.log.WithContext(ctx).Warnw("msg", "Restore.NotFound", "err", err)
		return err
	}
	return nil
}

func (r *userRepo) Update(ctx context.Context, u *types.User, opts ...*dto.UserUpdateOption) (*types.User, error) {
	r.log.WithContext(ctx).Debugw("msg", "Update", "user", u)
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
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Update.Save", "err", err)
		}
		return err
	})
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Update.Tx", "err", err)
		return nil, err
	}
	return dto.ConvertUserToUserPB(updatedUser), nil
}

func (r *userRepo) List(ctx context.Context, opts ...*dto.UserQueryOption) ([]*types.User, int32, error) {
	r.log.WithContext(ctx).Debugw("msg", "List", "opts", opts)
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
		r.log.WithContext(ctx).Errorw("msg", "List", "err", err)
		return nil, 0, err
	}

	return dto.ConvertUsersToUsersPB(result), count, err
}

func (r *userRepo) AddRoleIDs(ctx context.Context, id int64, roleIDs []int64) ([]*types.Role, error) {
	r.log.WithContext(ctx).Debugw("msg", "AddRoleIDs", "id", id, "roleIDs", roleIDs)
	var roles []*ent.Role
	err := r.db.Tx(ctx, func(tx context.Context) error {
		u, err := r.db.User(tx).UpdateOneID(id).ClearRoles().AddRoleIDs(roleIDs...).Save(ctx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddRoleIDs.Save", "err", err)
			return err
		}
		roles, err = u.QueryRoles().All(ctx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "AddRoleIDs.QueryRoles", "err", err)
		}
		return err
	})
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "AddRoleIDs.Tx", "err", err)
		return nil, err
	}
	return dto.ConvertRolesToRolesPB(roles), nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*types.User, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetByUsername", "username", username)
	result, err := r.db.User(ctx).Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetByUsername", "err", err)
		return nil, err
	}
	return dto.ConvertUserToUserPB(result), nil
}

func (r *userRepo) GetRoleIDs(ctx context.Context, id int64) ([]int64, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetRoleIDs", "id", id)
	ids, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().IDs(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetRoleIDs", "err", err)
		return nil, err
	}
	return ids, nil
}

func (r *userRepo) ListRoleByUserID(ctx context.Context, id int64) ([]*types.Role, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListRoleByUserID", "id", id)
	roles, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListRoleByUserID", "err", err)
		return nil, err
	}
	return dto.ConvertRolesToRolesPB(roles), nil
}

func (r *userRepo) ListResourceByUserID(ctx context.Context, id int64) ([]*types.Resource, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListResourceByUserID", "id", id)
	resources, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().QueryPermissions().QueryResources().All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListResourceByUserID", "err", err)
		return nil, err
	}
	return dto.ConvertResourcesToResourcesPB(resources), nil
}

func (r *userRepo) ListViewByUserID(ctx context.Context, id int64) ([]*types.View, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListViewByUserID", "id", id)
	views, err := r.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().QueryPermissions().QueryViews().All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListViewByUserID", "err", err)
		return nil, err
	}
	return dto.ConvertViewsToViewsPB(views), nil
}

func (r *userRepo) ListPermissionByUserID(ctx context.Context, id int64) ([]*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "ListPermissionByUserID", "id", id)
	permissions, err := r.db.User(ctx).Query().
		Where(user.ID(id)).
		QueryRoles().
		QueryPermissions().
		WithResources().
		WithViews().
		All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "ListPermissionByUserID", "err", err)
		return nil, err
	}
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

func (r *userRepo) UpdateUserStatus(ctx context.Context, id int64, status int8) error {
	r.log.WithContext(ctx).Debugw("msg", "UpdateUserStatus", "id", id, "status", status)
	err := r.db.User(ctx).UpdateOneID(id).SetStatus(enums.Status(status)).Exec(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "UpdateUserStatus", "err", err)
	}
	return err
}
