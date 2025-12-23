/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"context"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/data/ent"
	"origadmin/application/admin/internal/features/system/data/ent/permission"
	"origadmin/application/admin/internal/features/system/dto"
)

type permissionRepo struct {
	db *ent.Client
}

func (repo *permissionRepo) Get(ctx context.Context, id int64, options ...dto.PermissionQueryOption) (*types.Permission, error) {
	result, err := repo.db.Permission.Get(ctx, (id))
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(result), nil
}

func (repo *permissionRepo) Create(ctx context.Context, p *types.Permission, options ...dto.PermissionQueryOption) (*types.Permission, error) {
	create := repo.db.Permission.Create().
		SetName(p.Name)

	if len(p.ResourceIds) > 0 {
		create.AddResourceIDs(p.ResourceIds...)
	}

	// ... set other fields

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(saved), nil
}

func (repo *permissionRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.Permission.DeleteOneID((id)).Exec(ctx)
}

func (repo *permissionRepo) Update(ctx context.Context, p *types.Permission, options ...dto.PermissionQueryOption) (*types.Permission, error) {
	update := repo.db.Permission.UpdateOneID((p.Id))

	if len(p.ResourceIds) > 0 {
		update.ClearResources().AddResourceIDs(p.ResourceIds...)
	}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(saved), nil
}

func (repo *permissionRepo) List(ctx context.Context, in *system.ListPermissionsRequest, options ...dto.PermissionQueryOption) ([]*types.Permission, int32, error) {
	query := repo.db.Permission.Query()

	if len(in.DataScopes) > 0 {
		query = query.Where(permission.DataScopeIn(in.DataScopes...))
	}

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
	return dto.ConvertPermissionsToPermissionsPB(result), int32(count), err
}

// NewPermissionRepo .
func NewPermissionRepo(d *Data) (dto.PermissionRepo, error) {
	return &permissionRepo{db: d.db}, nil
}
