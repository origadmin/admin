/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/permission"
	"origadmin/application/admin/internal/features/system/dto"
)

type permissionRepo struct {
	db *ent.Database
}

// NewPermissionRepo .
func NewPermissionRepo(db *ent.Database) dto.PermissionRepo {
	return &permissionRepo{db: db}
}

func (r *permissionRepo) Get(ctx context.Context, id int64, opts ...*dto.PermissionQueryOptions) (*types.Permission, error) {
	opt := &dto.PermissionQueryOptions{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	query := r.db.Permission(ctx).Query().Where(permission.ID(id))

	if opt.WithResources {
		query.WithResources()
	}
	if opt.WithRoles {
		query.WithRoles()
	}

	result, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(result), nil
}

func (r *permissionRepo) Create(ctx context.Context, p *types.Permission, opts ...*dto.PermissionCreateOptions) (*types.Permission, error) {
	entPermission := dto.ConvertPermissionPBToPermission(p)
	create := r.db.Permission(ctx).Create().SetPermission(entPermission)

	// ... set other fields

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(saved), nil
}

func (r *permissionRepo) Delete(ctx context.Context, id int64) error {
	return r.db.Permission(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *permissionRepo) Update(ctx context.Context, p *types.Permission, opts ...*dto.PermissionUpdateOptions) (*types.Permission, error) {
	entPermission := dto.ConvertPermissionPBToPermission(p)
	update := r.db.Permission(ctx).UpdateOneID(p.Id).SetPermission(entPermission)

	// ... handle partial updates based on opts ...

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(saved), nil
}

func (r *permissionRepo) List(ctx context.Context, in *system.ListPermissionsRequest, opts ...*dto.PermissionQueryOptions) ([]*types.Permission, int32, error) {
	opt := &dto.PermissionQueryOptions{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	query := r.db.Permission(ctx).Query()

	if len(in.DataScopes) > 0 {
		query = query.Where(permission.DataScopeIn(in.DataScopes...))
	}

	if opt.Page > 0 && opt.PageSize > 0 {
		query.Offset((opt.Page - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if opt.WithResources {
		query.WithResources()
	}
	if opt.WithRoles {
		query.WithRoles()
	}

	result, err := query.All(ctx)
	return dto.ConvertPermissionsToPermissionsPB(result), int32(count), err
}
