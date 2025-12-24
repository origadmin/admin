/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strconv"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/resource"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/repo"
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

func (r *resourceRepo) Get(ctx context.Context, id int64, opts ...*dto.ResourceQueryOption) (*types.Resource, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Resource(ctx).Query().Where(resource.ID(id))

	if opt.WithPermissions {
		query.WithPermissions()
	}

	result, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(result), nil
}

func (r *resourceRepo) Create(ctx context.Context, res *types.Resource, opts ...*dto.ResourceCreateOption) (*types.Resource, error) {
	if res.ParentId > 0 {
		parent, err := r.db.Resource(ctx).Get(ctx, res.ParentId)
		if err != nil {
			return nil, err
		}
		res.TreePath = parent.TreePath + strconv.FormatInt(int64(parent.ID), 10) + r.Delimiter
	}

	entResource := dto.ConvertResourcePBToResource(res)
	create := r.db.Resource(ctx).Create().SetResource(entResource)

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

func (r *resourceRepo) Update(ctx context.Context, res *types.Resource, opts ...*dto.ResourceUpdateOption) (*types.Resource, error) {
	entResource := dto.ConvertResourcePBToResource(res)
	update := r.db.Resource(ctx).UpdateOneID(res.Id).SetResource(entResource)

	// ... handle partial updates based on opts ...

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (r *resourceRepo) List(ctx context.Context, opts ...*dto.ResourceQueryOption) ([]*types.Resource, int32, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Resource(ctx).Query()

	if opt.Page > 0 && opt.PageSize > 0 {
		query.Offset((opt.Page - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if opt.WithPermissions {
		query.WithPermissions()
	}

	result, err := query.All(ctx)
	return dto.ConvertResourcesToResourcesPB(result), int32(count), err
}
