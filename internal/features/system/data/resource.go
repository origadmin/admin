/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"context"
	"strconv"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/data/ent"
	"origadmin/application/admin/internal/features/system/dto"
)

type resourceRepo struct {
	db        *ent.Client
	Delimiter string
}

func (repo *resourceRepo) Get(ctx context.Context, id int64, options ...dto.ResourceQueryOption) (*types.Resource, error) {
	result, err := repo.db.Resource.Get(ctx, (id))
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(result), nil
}

func (repo *resourceRepo) Create(ctx context.Context, r *types.Resource, options ...dto.ResourceQueryOption) (*types.Resource, error) {
	if r.ParentId > 0 {
		parent, err := repo.db.Resource.Get(ctx, r.ParentId)
		if err != nil {
			return nil, err
		}
		r.TreePath = parent.TreePath + strconv.Itoa(int(parent.ID)) + repo.Delimiter
	}

	create := repo.db.Resource.Create().
		SetName(r.Name).
		SetParentID(r.ParentId).
		SetTreePath(r.TreePath)

	if len(r.PermissionIds) > 0 {
		create.AddPermissionIDs(r.PermissionIds...)
	}

	// ... set other fields

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (repo *resourceRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.Resource.DeleteOneID(id).Exec(ctx)
}

func (repo *resourceRepo) Update(ctx context.Context, r *types.Resource, options ...dto.ResourceQueryOption) (*types.Resource, error) {
	update := repo.db.Resource.UpdateOneID(r.Id)

	if len(r.PermissionIds) > 0 {
		update.ClearPermissions().AddPermissionIDs(r.PermissionIds...)
	}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (repo *resourceRepo) List(ctx context.Context, in *system.ListResourcesRequest, options ...dto.ResourceQueryOption) ([]*types.Resource, int32, error) {
	query := repo.db.Resource.Query()

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
	return dto.ConvertResourcesToResourcesPB(result), int32(count), err
}

// NewResourceRepo .
func NewResourceRepo(d *Data) (dto.ResourceRepo, error) {
	return &resourceRepo{
		db:        d.db,
		Delimiter: "/",
	}, nil
}
