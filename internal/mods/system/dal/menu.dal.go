/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"github.com/origadmin/runtime/log"

	"origadmin/application/admin/internal/mods/system/dto"
)

type menuRepo struct {
	db *Data
}

//
//func (repo menuRepo) Get(ctx context.Context, id int64, options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
//	var option dto.MenuQueryOption
//	if len(options) > 0 {
//		option = options[0]
//	}
//	query := repo.db.Menu(ctx).Query().Where(menu.ID(id))
//	query = menuQueryOptions(query, option)
//	result, err := query.First(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return dto.ConvertMenu2PB(result), nil
//}
//
//func (repo menuRepo) Create(ctx context.Context, menuPB *dto.MenuPB, options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
//	var option dto.MenuQueryOption
//	if len(options) > 0 {
//		option = options[0]
//	}
//	err := repo.db.Tx(ctx, func(ctx context.Context) error {
//		create := repo.db.Menu(ctx).Create()
//		create.SetMenu(dto.ConvertMenuPB2Object(menuPB), option.Fields...)
//		saved, err := create.Save(ctx)
//		if err != nil {
//			return err
//		}
//		menuPB = dto.ConvertMenu2PB(saved)
//		return nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return menuPB, nil
//}
//
//func (repo menuRepo) Delete(ctx context.Context, id int64) error {
//	return repo.db.Tx(ctx, func(ctx context.Context) error {
//		return repo.db.Menu(ctx).DeleteOneID(id).Exec(ctx)
//	})
//}
//
//func (repo menuRepo) Update(ctx context.Context, menuPB *dto.MenuPB, options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
//	err := repo.db.Tx(ctx, func(ctx context.Context) error {
//		update := repo.db.Menu(ctx).UpdateOneID(menuPB.Id)
//		update.SetMenu(dto.ConvertMenuPB2Object(menuPB))
//		saved, err := update.Save(ctx)
//		if err != nil {
//			return err
//		}
//		menuPB = dto.ConvertMenu2PB(saved)
//		return nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return menuPB, nil
//}
//
//func (repo menuRepo) List(ctx context.Context, in *dto.ListMenusRequest, options ...dto.MenuQueryOption) ([]*dto.MenuPB, int32, error) {
//	var option dto.MenuQueryOption
//	if len(options) > 0 {
//		option = options[0]
//	}
//
//	query := repo.db.Menu(ctx).Query()
//	if option.IncludeResources {
//		query = query.WithResources()
//	}
//	if v := option.UserID; v > 0 {
//		query = query.Where(menu.HasRolesWith(role.HasUsersWith(user.ID(v))))
//	}
//	if v := option.RoleID; v > 0 {
//		query = query.Where(menu.HasRolesWith(role.ID(v)))
//	}
//	if v := option.InIDs; len(v) > 0 {
//		query = query.Where(menu.IDIn(v...))
//	}
//	if v := option.Name; len(v) > 0 {
//		query = query.Where(menu.ParentPathContains(v))
//	}
//	if v := option.Status; v > 0 {
//		query = query.Where(menu.StatusEQ(v))
//	}
//	if v := option.ParentID; v > 0 {
//		query = query.Where(menu.ParentID(v))
//	}
//	if v := option.ParentPathPrefix; len(v) > 0 {
//		query = query.Where(menu.ParentPathHasPrefix(v))
//	}
//
//	return menuPageQuery(ctx, query, in, option)
//}

// NewMenuRepo .
func NewMenuRepo(db *Data, logger log.KLogger) dto.MenuRepo {
	return &menuRepo{
		db: db,
	}
}

//
//func menuPageQuery(ctx context.Context, query *ent.MenuQuery, in *pb.ListMenusRequest, option dto.MenuQueryOption) ([]*dto.MenuPB, int32, error) {
//	if in.OnlyCount {
//		count, err := query.Count(ctx)
//		if err != nil {
//			return nil, 0, err
//		}
//		return nil, int32(count), nil
//	}
//	count, err := query.Clone().Count(ctx)
//	if err != nil {
//		return nil, 0, err
//	}
//	query = menuQueryPage(query, in)
//	query = menuQueryOptions(query, option)
//	result, err := query.Clone().All(ctx)
//	return dto.ConvertMenus(result), int32(count), err
//}
//
//func menuQueryPage(query *ent.MenuQuery, in *pb.ListMenusRequest) *ent.MenuQuery {
//	if in.NoPaging {
//		pageSize := in.PageSize
//		if pageSize > 0 {
//			query = query.Limit(int(pageSize))
//		}
//		return query
//	}
//
//	pageSize := in.PageSize
//	if pageSize > 0 {
//		query = query.Limit(int(pageSize))
//	}
//	current := in.Current
//	if current > 0 {
//		query = query.Offset(int((current - 1) * pageSize))
//	}
//	return query
//}
//
//func menuQueryOptions(query *ent.MenuQuery, option dto.MenuQueryOption) *ent.MenuQuery {
//	if len(option.SelectFields) > 0 {
//		query = query.Select(option.SelectFields...).MenuQuery
//	}
//	if len(option.OmitFields) > 0 {
//		query = query.Omit(option.OmitFields...).MenuQuery
//	}
//	if len(option.OrderFields) > 0 {
//		query = query.Order(menuOrderBy(option.OrderFields)...)
//	}
//	return query
//}
//
//func menuOrderBy(fields []string, opts ...sql.OrderTermOption) []menu.OrderOption {
//	var orders []menu.OrderOption
//	for _, field := range fields {
//		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
//	}
//	return orders
//}
