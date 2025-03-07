/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dto"
)

type permissionRepo struct {
	db *Data
}

func (repo permissionRepo) Get(ctx context.Context, id int64, options ...dto.PermissionQueryOption) (*dto.PermissionPB, error) {
	var option dto.PermissionQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.Permission(ctx).Query().Where(permission.ID(id))
	query = permissionQueryOptions(query, option)
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermission2PB(result), nil
}

func (repo permissionRepo) Create(ctx context.Context, permission *dto.PermissionPB, options ...dto.PermissionQueryOption) (*dto.PermissionPB, error) {
	var option dto.PermissionQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	obj := dto.ConvertPermissionPB2Object(permission)
	create := repo.db.Permission(ctx).Create()
	if len(permission.ResourceIds) > 0 {
		create.AddResourceIDs(permission.ResourceIds...)
	}
	if len(permission.Resources) > 0 {
		create.AddResources(dto.ConvertResourcesPB2Object(permission.Resources)...)
	}
	create.SetPermission(obj, option.Fields...)
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermission2PB(saved), nil
}

func (repo permissionRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.Permission(ctx).DeleteOneID(id).Exec(ctx)
}

func (repo permissionRepo) Update(ctx context.Context, permission *dto.PermissionPB, options ...dto.PermissionQueryOption) (*dto.PermissionPB, error) {
	var option dto.PermissionQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	update := repo.db.Permission(ctx).UpdateOneID(permission.Id)
	obj := dto.ConvertPermissionPB2Object(permission)
	if len(permission.ResourceIds) > 0 {
		update.ClearResources()
		update.AddResourceIDs(permission.ResourceIds...)
	}
	if len(permission.Resources) > 0 {
		update.ClearResources()
		update.AddResources(dto.ConvertResourcesPB2Object(permission.Resources)...)
	}
	update.SetPermission(obj, option.Fields...)
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermission2PB(saved), nil
}

func (repo permissionRepo) List(ctx context.Context, in *dto.ListPermissionsRequest, options ...dto.PermissionQueryOption) ([]*dto.PermissionPB, int32, error) {
	var option dto.PermissionQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	query := repo.db.Permission(ctx).Query()
	if option.IncludeResources {
		query = query.WithResources()
	}
	if option.IncludeRoles {
		query = query.WithRoles()
	}
	if len(in.DataScopes) > 0 {
		query = query.Where(permission.DataScopeIn(in.DataScopes...))
	}
	return permissionPageQuery(ctx, query, in, option)
}

// NewPermissionRepo .
func NewPermissionRepo(db *Data, logger log.KLogger) dto.PermissionRepo {
	return &permissionRepo{
		db: db,
	}
}

func permissionPageQuery(ctx context.Context, query *ent.PermissionQuery, in *pb.ListPermissionsRequest, option dto.PermissionQueryOption) ([]*dto.PermissionPB, int32, error) {
	if in.OnlyCount {
		count, err := query.Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		return nil, int32(count), nil
	}

	query = permissionQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	query = permissionQueryPage(query, in)
	result, err := query.All(ctx)
	return dto.ConvertPermissions(result), int32(count), err
}

func permissionQueryPage(query *ent.PermissionQuery, in *pb.ListPermissionsRequest) *ent.PermissionQuery {
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

func permissionQueryOptions(query *ent.PermissionQuery, option dto.PermissionQueryOption) *ent.PermissionQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).PermissionQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).PermissionQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(permissionOrderBy(option.OrderFields)...)
	}
	return query
}

func permissionOrderBy(fields []string, opts ...sql.OrderTermOption) []permission.OrderOption {
	var orders []permission.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
