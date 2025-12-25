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
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type resourceRepo struct {
	db        *ent.Database
	Delimiter string
}

// NewResourceRepo .
func NewResourceRepo(database *ent.Database) dto.ResourceRepo {
	return &resourceRepo{
		db:        database,
		Delimiter: "/",
	}
}

func (r *resourceRepo) Get(ctx context.Context, id int64, opts ...*dto.ResourceQueryOption) (*types.Resource, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Resource(ctx).Query().Where(resource.ID(id))

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, resource.ValidColumn, resource.FieldID, new(types.Resource))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
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
		res.TreePath = parent.TreePath + strconv.FormatInt(parent.ID, 10) + r.Delimiter
	}

	entResource := dto.ConvertResourcePBToResource(res)
	create := r.db.Resource(ctx).Create().SetResource(entResource)

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
	opt := repo.GetFirstOption(opts...)
	entResource := dto.ConvertResourcePBToResource(res)
	update := r.db.Resource(ctx).UpdateOneID(res.Id)

	updateCols := db.UpdateFields(opt.UpdateMask, resource.ValidColumn, res)
	if len(updateCols) > 0 {
		update.SetResource(entResource, updateCols...)
	} else {
		update.SetResource(entResource)
	}

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (r *resourceRepo) List(ctx context.Context, opts ...*dto.ResourceQueryOption) ([]*types.Resource, int32, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Resource(ctx).Query()

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, resource.ValidColumn, resource.FieldID, new(types.Resource))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	if opt.OrderBy != nil {
		orders := db.OrderBy[resource.OrderOption](opt.OrderBy)
		if len(orders) > 0 {
			query.Order(orders...)
		}
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertResourcesToResourcesPB(result), count, err
}
