/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"

	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/rand"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type roleRepo struct {
	db  *ent.Database
	log *log.Helper
	gen rand.Rand
}

// NewRoleRepo .
func NewRoleRepo(database *ent.Database, logger log.Logger) dto.RoleRepo {
	generator := rand.NewGenerator(rand.KindDigit | rand.KindLowerCase)
	return &roleRepo{
		db:  database,
		gen: generator,
		log: log.NewHelper(log.With(logger, "module", "dal.role")),
	}
}

func (r *roleRepo) Get(ctx context.Context, id int64, opts ...*dto.RoleQueryOption) (*types.Role, error) {
	r.log.WithContext(ctx).Debugw("msg", "Get", "id", id)
	opt := repo.FirstOrDefault(opts...)
	query := r.db.Role(ctx).Query().Where(role.ID(id))

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, role.ValidColumn, role.FieldID,
			new(types.Role))
		query = s.RoleQuery
	}

	result, err := query.Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Get", "err", err)
		return nil, err
	}
	return dto.ConvertRoleToRolePB(result), nil
}

func (r *roleRepo) Create(ctx context.Context, rl *types.Role, opts ...*dto.RoleCreateOption) (*types.Role, error) {
	r.log.WithContext(ctx).Debugw("msg", "Create", "role", rl)
	opt := repo.FirstOrDefault(opts...)
	if rl.Keyword == "" {
		randString, err := r.gen.String(12)
		if err != nil {
			randString = ""
		}
		rl.Keyword = "system:role:" + randString
	}
	exist, err := r.db.Role(ctx).Query().Where(role.KeywordEqualFold(rl.Keyword)).Exist(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Create.Exist", "err", err)
		return nil, err
	}
	if exist {
		err = errors.New("role keyword already exists")
		r.log.WithContext(ctx).Warnw("msg", "Create.Exist", "err", err)
		return nil, err
	}

	entRole := dto.ConvertRolePBToRole(rl)
	create := r.db.Role(ctx).Create().SetRoleSkipZero(entRole)

	if opt.WithPermissionIDs != nil {
		create.AddPermissionIDs(opt.WithPermissionIDs...)
	}

	saved, err := create.Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Create.Save", "err", err)
		return nil, err
	}
	return dto.ConvertRoleToRolePB(saved), nil
}

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	r.log.WithContext(ctx).Debugw("msg", "Delete", "id", id)
	return r.db.Tx(ctx, func(txCtx context.Context) error {
		// Clear all associations before deletion (deletes intermediate table records)
		err := r.db.Role(txCtx).UpdateOneID(id).
			ClearPermissions().
			ClearUsers().
			Exec(txCtx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Delete.ClearAssociations", "err", err)
			return err
		}

		// Delete the Role entity
		err = r.db.Role(txCtx).DeleteOneID(id).Exec(txCtx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Delete.Exec", "err", err)
		}
		return err
	})
}

func (r *roleRepo) Update(ctx context.Context, rl *types.Role, opts ...*dto.RoleUpdateOption) (*types.Role, error) {
	r.log.WithContext(ctx).Debugw("msg", "Update", "role", rl)
	var updatedRole *ent.Role
	err := r.db.Tx(ctx, func(tx context.Context) error {
		opt := repo.FirstOrDefault(opts...)

		entRole := dto.ConvertRolePBToRole(rl)
		update := r.db.Role(tx).UpdateOneID(rl.Id)

		updateCols := db.UpdateFields(opt.UpdateMask, role.ValidColumn, rl)
		if len(updateCols) > 0 {
			update.SetRole(entRole, updateCols...)
		} else {
			update.SetRoleSkipZero(entRole)
		}
		update.ClearPermissions()
		if len(opt.WithPermissionIDs) > 0 {
			update.AddPermissionIDs(opt.WithPermissionIDs...)
		}

		var err error
		updatedRole, err = update.Save(ctx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Update.Save", "err", err)
		}
		return err
	})
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Update.Tx", "err", err)
		return nil, err
	}
	return dto.ConvertRoleToRolePB(updatedRole), nil
}

func (r *roleRepo) List(ctx context.Context, opts ...*dto.RoleQueryOption) ([]*types.Role, int32, error) {
	r.log.WithContext(ctx).Debugw("msg", "List", "opts", opts)
	opt := repo.FirstOrDefault(opts...)
	query := r.db.Role(ctx).Query()

	if opt.Keyword != "" {
		query.Where(role.NameContainsFold(opt.Keyword))
	}

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, role.ValidColumn, role.FieldID,
			new(types.Role))
		query = s.RoleQuery
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "List", "err", err)
		return nil, 0, err
	}

	return dto.ConvertRolesToRolesPB(result), count, err
}

func (r *roleRepo) GetPermissions(ctx context.Context, id int64) ([]*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "GetPermissions", "id", id)
	permissions, err := r.db.Role(ctx).Query().Where(role.ID(id)).QueryPermissions().All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "GetPermissions", "err", err)
		return nil, err
	}
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

func (r *roleRepo) UpdatePermissions(ctx context.Context, id int64, permissionIDs []int64) error {
	r.log.WithContext(ctx).Debugw("msg", "UpdatePermissions", "id", id, "permissionIDs", permissionIDs)
	_, err := r.db.Role(ctx).UpdateOneID(id).ClearPermissions().AddPermissionIDs(permissionIDs...).Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "UpdatePermissions", "err", err)
	}
	return err
}
