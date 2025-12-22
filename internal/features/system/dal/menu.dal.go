/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strings"

	"entgo.io/ent/dialect/sql"

	"github.com/origadmin/runtime"
	"origadmin/application/admin/internal/data"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/resource"
	"origadmin/application/admin/internal/features/system/dto" // Corrected import path
	"origadmin/application/admin/internal/helpers/db"
)

type menuRepo struct {
	data *data.Data
	db   *ent.Database
}

func (repo menuRepo) Get(ctx context.Context, id int64, options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
	var option dto.MenuQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.Resource(ctx).Query().Where(resource.ID(id))
	query = menuQueryOptions(query, option)
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToMenuPB(result), nil
}

func (repo menuRepo) Create(ctx context.Context, menuPB *dto.MenuPB,
	options ...dto.MenuQueryOption) (*dto.MenuPB, error) {
	var option dto.MenuQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	err := repo.db.Tx(ctx, func(ctx context.Context) error {
		create := repo.db.Resource(ctx).Create()
		create.SetResource(dto.ConvertMenuPBToResource(menuPB), option.Fields...)
		saved, err := create.Save(ctx)
		if err != nil {
			return err
		}
		menuPB = dto.ConvertResourceToMenuPB(saved)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return menuPB, nil
}

func (repo menuRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.Tx(ctx, func(ctx context.Context) error {
		return repo.db.Resource(ctx).DeleteOneID(id).Exec(ctx)
	})
}

func (repo menuRepo) Update(ctx context.Context, menuPB *dto.MenuPB, options ...dto.MenuQueryOption) (*dto.MenuPB,
	error) {
	err := repo.db.Tx(ctx, func(ctx context.Context) error {
		update := repo.db.Resource(ctx).UpdateOneID(menuPB.Id)
		update.SetResource(dto.ConvertMenuPBToResource(menuPB))
		saved, err := update.Save(ctx)
		if err != nil {
			return err
		}
		menuPB = dto.ConvertResourceToMenuPB(saved)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return menuPB, nil
}

func (repo menuRepo) List(ctx context.Context, in *dto.ListMenusRequest, options ...dto.MenuQueryOption) ([]*dto.MenuPB, int32, error) {
	var option dto.MenuQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	query := repo.db.Resource(ctx).Query()
	//if option.IncludeResources {
	//	query = query.WithResources()
	//}
	//if v := option.UserID; v > 0 {
	//	query = query.Where(resource.HasRolesWith(role.HasUsersWith(user.ID(v))))
	//}
	//if v := option.RoleID; v > 0 {
	//	query = query.Where(resource.HasRolesWith(role.ID(v)))
	//}
	if v := option.InIDs; len(v) > 0 {
		query = query.Where(resource.IDIn(v...))
	}
	//if v := option.Name; len(v) > 0 {
	//	query = query.Where(resource.ParentPathContains(v))
	//}
	if v := option.Status; v > 0 {
		query = query.Where(resource.StatusEQ(v))
	}
	if v := option.ParentID; v > 0 {
		query = query.Where(resource.ParentID(v))
	}
	//if v := option.ParentPathPrefix; len(v) > 0 {
	//	query = query.Where(resource.ParentPathHasPrefix(v))
	//}

	return menuPageQuery(ctx, query, in, option)
}

// NewMenuRepo .
func NewMenuRepo(r *runtime.App, d *data.Data) dto.MenuRepo {
	return &menuRepo{
		data: d,
		db:   d.DB(),
	}
}

func menuPageQuery(ctx context.Context, query *ent.ResourceQuery, in *dto.ListMenusRequest,
	option dto.MenuQueryOption) ([]*dto.MenuPB, int32, error) {
	if in.OnlyCount {
		count, err := query.Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		return nil, int32(count), nil
	}
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	query = db.QueryPage(query, in)
	query = menuQueryOptions(query, option)
	result, err := query.Clone().All(ctx)
	menus := make([]*dto.MenuPB, len(result))
	for i, r := range result {
		menus[i] = dto.ConvertResourceToMenuPB(r)
	}
	return menus, int32(count), err
}

func menuQueryOptions(query *ent.ResourceQuery, option dto.MenuQueryOption) *ent.ResourceQuery {
	//if len(option.SelectFields) > 0 {
	//	query = query.Select(option.SelectFields...).(*ent.ResourceQuery)
	//}
	//if len(option.OmitFields) > 0 {
	//	query = query.Omit(option.OmitFields...).(*ent.ResourceQuery)
	//}
	if len(option.OrderFields) > 0 {
		query = query.Order(menuOrderBy(option.OrderFields)...)
	}
	return query
}

func menuOrderBy(fields []string, opts ...sql.OrderTermOption) []resource.OrderOption {
	var orders []resource.OrderOption
	for _, field := range fields {
		parts := strings.Split(field, ",")
		fieldName := parts[0]
		var orderOpt sql.OrderTermOption

		if len(parts) > 1 {
			switch strings.ToLower(parts[1]) {
			case "desc":
				orderOpt = sql.OrderDesc()
			default:
				orderOpt = sql.OrderAsc()
			}
		} else {
			orderOpt = sql.OrderAsc()
		}

		orders = append(orders, sql.OrderByField(fieldName, orderOpt).ToFunc())
	}
	return orders
}
