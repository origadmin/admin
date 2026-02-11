/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/permission"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type permissionRepo struct {
	db  *ent.Database
	log *log.Helper
}

// NewPermissionRepo .
func NewPermissionRepo(database *ent.Database, logger log.Logger) dto.PermissionRepo {
	return &permissionRepo{
		db:  database,
		log: log.NewHelper(log.With(logger, "module", "dal.permission")),
	}
}

func (r *permissionRepo) Get(ctx context.Context, id int64, opts ...*dto.PermissionQueryOption) (*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "Get", "id", id)
	opt := repo.FirstOrDefault(opts...)
	query := r.db.Permission(ctx).Query().Where(permission.ID(id))

	if opt.WithResources {
		query.WithResources()
	}
	if opt.WithRoles {
		query.WithRoles()
	}
	if opt.WithViews {
		query.WithViews()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, permission.ValidColumn, permission.FieldID, new(types.Permission))
		query = s.PermissionQuery
	}

	result, err := query.Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Get", "err", err)
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(result), nil
}

func (r *permissionRepo) Create(ctx context.Context, p *types.Permission, opts ...*dto.PermissionCreateOption) (*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "Create", "permission", p)
	opt := repo.FirstOrDefault(opts...)
	entPermission := dto.ConvertPermissionPBToPermission(p)
	create := r.db.Permission(ctx).Create().SetPermissionSkipZero(entPermission)

	if opt.WithResourceIDs != nil {
		create.AddResourceIDs(opt.WithResourceIDs...)
	}
	if opt.WithViewIDs != nil {
		create.AddViewIDs(opt.WithViewIDs...)
	}

	saved, err := create.Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Create", "err", err)
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(saved), nil
}

func (r *permissionRepo) Delete(ctx context.Context, id int64) error {
	r.log.WithContext(ctx).Debugw("msg", "Delete", "id", id)
	return r.db.Tx(ctx, func(txCtx context.Context) error {
		// Clear all associations before deletion (deletes intermediate table records)
		err := r.db.Permission(txCtx).UpdateOneID(id).
			ClearResources().
			ClearRoles().
			ClearPositions().
			ClearViews().
			Exec(txCtx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Delete.ClearAssociations", "err", err)
			return err
		}

		// Delete the Permission entity
		err = r.db.Permission(txCtx).DeleteOneID(id).Exec(txCtx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Delete.Exec", "err", err)
		}
		return err
	})
}

func (r *permissionRepo) Update(ctx context.Context, p *types.Permission, opts ...*dto.PermissionUpdateOption) (*types.Permission, error) {
	r.log.WithContext(ctx).Debugw("msg", "Update", "permission", p)
	var updatedPermission *ent.Permission
	err := r.db.Tx(ctx, func(tx context.Context) error {
		opt := repo.FirstOrDefault(opts...)
		entPermission := dto.ConvertPermissionPBToPermission(p)
		update := r.db.Permission(tx).UpdateOneID(p.Id)

		updateCols := db.UpdateFields(opt.UpdateMask, permission.ValidColumn, p)
		if len(updateCols) > 0 {
			// If a field mask is present, update only the specified fields, including zero values.
			update.SetPermission(entPermission, updateCols...)
		} else {
			// If no field mask, skip zero values to prevent accidental clearing of fields.
			update.SetPermissionSkipZero(entPermission)
		}
		update.ClearResources()
		if len(opt.WithResourceIDs) > 0 {
			update.AddResourceIDs(opt.WithResourceIDs...)
		}
		update.ClearViews()
		if len(opt.WithViewIDs) > 0 {
			update.AddViewIDs(opt.WithViewIDs...)
		}

		var err error
		updatedPermission, err = update.Save(ctx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Update.Save", "err", err)
		}
		return err
	})
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Update.Tx", "err", err)
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(updatedPermission), nil
}

func (r *permissionRepo) List(ctx context.Context, opts ...*dto.PermissionQueryOption) ([]*types.Permission, int32, error) {
	r.log.WithContext(ctx).Debugw("msg", "List", "opts", opts)
	opt := repo.FirstOrDefault(opts...)
	query := r.db.Permission(ctx).Query()

	if len(opt.DataScopes) > 0 {
		query.Where(permission.DataScopeIn(opt.DataScopes...))
	}

	if opt.WithResources {
		query.WithResources()
	}
	if opt.WithRoles {
		query.WithRoles()
	}
	if opt.WithViews {
		query.WithViews()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, permission.ValidColumn, permission.FieldID, new(types.Permission))
		query = s.PermissionQuery
	}

	// db.Find no longer requires a callback.
	// The generic type parameter [predicate.Permission] tells db.Find what kind of predicate to build.
	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "List", "err", err)
		return nil, 0, err
	}

	return dto.ConvertPermissionsToPermissionsPB(result), count, err
}
