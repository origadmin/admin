/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal is the data access object
package dal

import (
	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dto"
)

type roleRepo struct {
	db *Data
}

func (m roleRepo) Get(ctx context.Context, id string, options ...dto.RoleQueryOption) (*dto.RolePB, error) {
	var option dto.RoleQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := m.db.Role(ctx).Query().Where(role.ID(id))
	query = roleQueryOptions(query, option)
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRole2PB(result), nil
}

func (m roleRepo) Create(ctx context.Context, role *dto.RolePB, options ...dto.RoleQueryOption) (*dto.RolePB, error) {
	var option dto.RoleQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	create := m.db.Role(ctx).Create()
	create.SetRole(dto.RoleObject(role), option.Fields...)
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRole2PB(saved), nil
}

func (m roleRepo) Delete(ctx context.Context, id string) error {
	return m.db.Role(ctx).DeleteOneID(id).Exec(ctx)
}

func (m roleRepo) Update(ctx context.Context, role *dto.RolePB, options ...dto.RoleQueryOption) (*dto.RolePB, error) {
	update := m.db.Role(ctx).UpdateOneID(role.Id)
	update.SetRole(dto.RoleObject(role))
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRole2PB(saved), nil
}

func (m roleRepo) List(ctx context.Context, in *pb.ListRolesRequest, options ...dto.RoleQueryOption) ([]*dto.RolePB, int64, error) {
	var option dto.RoleQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	query := m.db.Role(ctx).Query()
	if v := option.InIDs; len(v) > 0 {
		query = query.Where(role.IDIn(v...))
	}
	if v := option.Name; len(v) > 0 {
		query = query.Where(role.NameContains(v))
	}
	if v := option.Status; v > 0 {
		query = query.Where(role.StatusEQ(v))
	}
	if v := option.UpdateTimeGT; v != nil {
		query = query.Where(role.UpdateTimeGT(*v))
	}

	return rolePageQuery(ctx, query, in, option)
}

// NewRoleDal .
func NewRoleDal(db *Data, logger log.Logger) dto.RoleRepo {
	return &roleRepo{
		db: db,
	}
}

func rolePageQuery(ctx context.Context, query *ent.RoleQuery, in *pb.ListRolesRequest, option dto.RoleQueryOption) ([]*dto.RolePB, int64, error) {
	if in.OnlyCount {
		count, err := query.Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		return nil, int64(count), nil
	}

	query = roleQueryPage(query, in)
	query = roleQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertRoles(result), int64(count), err
}

func roleQueryPage(query *ent.RoleQuery, in *pb.ListRolesRequest) *ent.RoleQuery {
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

func roleQueryOptions(query *ent.RoleQuery, option dto.RoleQueryOption) *ent.RoleQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).RoleQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).RoleQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(roleOrderBy(option.OrderFields)...)
	}
	return query
}

func roleOrderBy(fields []string, opts ...sql.OrderTermOption) []role.OrderOption {
	var orders []role.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}