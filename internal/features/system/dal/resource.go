/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strconv"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/features/system/dto"
)

type resourceRepo struct {
	db        *ent.Database
	Delimiter string
}

// NewResourceRepo .
func NewResourceRepo(db *ent.Database) dto.ResourceRepo {
	return &resourceRepo{
		db:        db,
		Delimiter: "/",
	}
}

func (r *resourceRepo) Get(ctx context.Context, id int64, opts ...dto.ResourceQueryOption) (*types.Resource, error) {
	result, err := r.db.Resource(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(result), nil
}

func (r *resourceRepo) Create(ctx context.Context, res *types.Resource, opts ...dto.ResourceMutationOption) (*types.Resource, error) {
	if res.ParentId > 0 {
		parent, err := r.db.Resource(ctx).Get(ctx, res.ParentId)
		if err != nil {
			return nil, err
		}
		res.TreePath = parent.TreePath + strconv.FormatInt(int64(parent.ID), 10) + r.Delimiter
	}

	create := r.db.Resource(ctx).Create().
		SetName(res.Name).
		SetParentID(res.ParentId).
		SetTreePath(res.TreePath)

	//if len(res.PermissionIds) > 0 {
	//	create.AddPermissionIDs(res.PermissionIds...)
	//}

	// ... set other fields

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (r *resourceRepo) Delete(ctx context.Context, id int64) error {
	return r.db.Resource(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *resourceRepo) Update(ctx context.Context, res *types.Resource, opts ...dto.ResourceMutationOption) (*types.Resource, error) {
	update := r.db.Resource(ctx).UpdateOneID(res.Id)

	//if len(res.PermissionIds) > 0 {
	//	update.ClearPermissions().AddPermissionIDs(res.PermissionIds...)
	//}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (r *resourceRepo) List(ctx context.Context, in *system.ListResourcesRequest, opts ...dto.ResourceQueryOption) ([]*types.Resource, int32, error) {
	query := r.db.Resource(ctx).Query()

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	//query = db.QueryPage(query, in)

	result, err := query.All(ctx)
	return dto.ConvertResourcesToResourcesPB(result), int32(count), err
}
