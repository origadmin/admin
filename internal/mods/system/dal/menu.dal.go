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
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dto"
)

type menuRepo struct {
	db *Data
}

func (m menuRepo) Get(ctx context.Context, id string, options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
	var option dto.MenuQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := m.db.Menu(ctx).Query().Where(menu.ID(id))
	query = menuQueryOptions(query, option)
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertMenu2PB(result), nil
}

func (m menuRepo) Create(ctx context.Context, menu *dto.MenuPB, options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
	var option dto.MenuQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	create := m.db.Menu(ctx).Create()
	create.SetMenu(dto.MenuObject(menu), option.Fields...)
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertMenu2PB(saved), nil
}

func (m menuRepo) Delete(ctx context.Context, id string) error {
	return m.db.Menu(ctx).DeleteOneID(id).Exec(ctx)
}

func (m menuRepo) Update(ctx context.Context, menu *dto.MenuPB, options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
	update := m.db.Menu(ctx).UpdateOneID(menu.Id)
	update.SetMenu(dto.MenuObject(menu))
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertMenu2PB(saved), nil
}

func (m menuRepo) List(ctx context.Context, in *dto.ListMenusRequest, options ...dto.MenuQueryOption) ([]*dto.MenuPB, int32, error) {
	var option dto.MenuQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	query := m.db.Menu(ctx).Query()
	if option.IncludeResources {
		query = query.WithResources()
	}
	if v := option.UserID; len(v) > 0 {
		query = query.Where(menu.HasRolesWith(role.HasUsersWith(user.ID(v))))
	}
	if v := option.RoleID; len(v) > 0 {
		query = query.Where(menu.HasRolesWith(role.ID(v)))
	}
	if v := option.InIDs; len(v) > 0 {
		query = query.Where(menu.IDIn(v...))
	}
	if v := option.Name; len(v) > 0 {
		query = query.Where(menu.ParentPathContains(v))
	}
	if v := option.Status; v > 0 {
		query = query.Where(menu.StatusEQ(v))
	}
	if v := option.ParentID; len(v) > 0 {
		query = query.Where(menu.ParentID(v))
	}
	if v := option.ParentPathPrefix; len(v) > 0 {
		query = query.Where(menu.ParentPathHasPrefix(v))
	}

	return menuPageQuery(ctx, query, in, option)
}

// NewMenuRepo .
func NewMenuRepo(db *Data, logger log.Logger) dto.MenuRepo {
	return &menuRepo{
		db: db,
	}
}

func menuPageQuery(ctx context.Context, query *ent.MenuQuery, in *pb.ListMenusRequest, option dto.MenuQueryOption) ([]*dto.MenuPB, int32, error) {
	if in.OnlyCount {
		count, err := query.Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		return nil, int32(count), nil
	}

	query = menuQueryPage(query, in)
	query = menuQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertMenus(result), int32(count), err
}

func menuQueryPage(query *ent.MenuQuery, in *pb.ListMenusRequest) *ent.MenuQuery {
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

func menuQueryOptions(query *ent.MenuQuery, option dto.MenuQueryOption) *ent.MenuQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).MenuQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).MenuQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(menuOrderBy(option.OrderFields)...)
	}
	return query
}

func menuOrderBy(fields []string, opts ...sql.OrderTermOption) []menu.OrderOption {
	var orders []menu.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
