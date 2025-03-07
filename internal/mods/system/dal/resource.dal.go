/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dto"
)

type resourceRepo struct {
	db *Data
}

func (repo resourceRepo) Get(ctx context.Context, id int64, options ...dto.ResourceQueryOption) (*dto.ResourcePB, error) {
	var option dto.ResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.Resource(ctx).Query().Where(resource.ID(id))
	query = resourceQueryOptions(query, option)
	if option.IncludePermissions {
		query.WithPermissions()
	}
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResource2PB(result), nil
}

func (repo resourceRepo) Create(ctx context.Context, resource *dto.ResourcePB, options ...dto.ResourceQueryOption) (*dto.ResourcePB, error) {
	var option dto.ResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	obj := dto.ConvertResourcePB2Object(resource)
	if obj.ParentID > 0 {
		parent, err := repo.db.Resource(ctx).Get(ctx, obj.ParentID)
		if err != nil {
			return nil, err
		}
		obj.TreePath = parent.TreePath + strconv.Itoa(int(parent.ID)) + TreePathDelimiter
	}

	create := repo.db.Resource(ctx).Create()
	create.SetResource(obj, option.Fields...)
	if len(resource.PermissionIds) > 0 {
		create.AddPermissionIDs(resource.PermissionIds...)
	}
	if len(resource.Permissions) > 0 {
		create.AddPermissions(dto.ConvertPermissionsPB2Object(resource.Permissions)...)
	}
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResource2PB(saved), nil
}

func (repo resourceRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.Resource(ctx).DeleteOneID(id).Exec(ctx)
}

func (repo resourceRepo) Update(ctx context.Context, resource *dto.ResourcePB, options ...dto.ResourceQueryOption) (*dto.ResourcePB, error) {
	err := repo.db.Tx(ctx, func(ctx context.Context) error {
		update := repo.db.Resource(ctx).UpdateOneID(resource.Id)
		update.SetResource(dto.ConvertResourcePB2Object(resource))
		if len(resource.PermissionIds) > 0 {
			update.AddPermissionIDs(resource.PermissionIds...)
		}
		if len(resource.Permissions) > 0 {
			update.AddPermissions(dto.ConvertPermissionsPB2Object(resource.Permissions)...)
		}
		saved, err := update.Save(ctx)
		if err != nil {
			return err
		}
		resource = dto.ConvertResource2PB(saved)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (repo resourceRepo) List(ctx context.Context, in *dto.ListResourcesRequest, options ...dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	var option dto.ResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	query := repo.db.Resource(ctx).Query()
	return resourcePageQuery(ctx, query, in, option)
}

// NewResourceRepo .
func NewResourceRepo(db *Data, logger log.KLogger) dto.ResourceRepo {
	return &resourceRepo{
		db: db,
	}
}

func resourcePageQuery(ctx context.Context, query *ent.ResourceQuery, in *pb.ListResourcesRequest, option dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	if in.OnlyCount {
		count, err := query.Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		return nil, int32(count), nil
	}

	query = resourceQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	query = resourceQueryPage(query, in)
	result, err := query.All(ctx)
	return dto.ConvertResources(result), int32(count), err
}

func resourceQueryPage(query *ent.ResourceQuery, in *pb.ListResourcesRequest) *ent.ResourceQuery {
	if in.NoPaging {
		pageSize := in.PageSize
		if pageSize > 0 {
			query = query.Limit(int(pageSize))
		}
		return query
	}

	pageSize := in.PageSize
	if pageSize > 0 {
		query = query.Limit(int(pageSize))
	}
	current := in.Current
	if current > 0 {
		query = query.Offset(int((current - 1) * pageSize))
	}
	return query
}

func resourceQueryOptions(query *ent.ResourceQuery, option dto.ResourceQueryOption) *ent.ResourceQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).ResourceQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).ResourceQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(resourceOrderBy(option.OrderFields)...)
	}
	return query
}

func resourceOrderBy(fields []string, opts ...sql.OrderTermOption) []resource.OrderOption {
	var orders []resource.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
