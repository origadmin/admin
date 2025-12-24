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

func (r *permissionRepo) Get(ctx context.Context, id int64, opts ...dto.PermissionQueryOption) (*types.Permission, error) {
	result, err := r.db.Permission(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(result), nil
}

func (r *permissionRepo) Create(ctx context.Context, p *types.Permission, opts ...dto.PermissionMutationOption) (*types.Permission, error) {
	create := r.db.Permission(ctx).Create().
		SetName(p.Name)

	//if len(p.ResourceIds) > 0 {
	//	create.AddResourceIDs(p.ResourceIds...)
	//}

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

func (r *permissionRepo) Update(ctx context.Context, p *types.Permission, opts ...dto.PermissionMutationOption) (*types.Permission, error) {
	update := r.db.Permission(ctx).UpdateOneID(p.Id)

	//if len(p.ResourceIds) > 0 {
	//	update.ClearResources().AddResourceIDs(p.ResourceIds...)
	//}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionToPermissionPB(saved), nil
}

func (r *permissionRepo) List(ctx context.Context, in *system.ListPermissionsRequest, opts ...dto.PermissionQueryOption) ([]*types.Permission, int32, error) {
	query := r.db.Permission(ctx).Query()

	if len(in.DataScopes) > 0 {
		query = query.Where(permission.DataScopeIn(in.DataScopes...))
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	//query = db.QueryPage(query, in)

	result, err := query.All(ctx)
	return dto.ConvertPermissionsToPermissionsPB(result), int32(count), err
}
