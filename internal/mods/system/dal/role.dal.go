/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dal is the data access object
package dal

import (
	"errors"

	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/rand"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/db"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dto"
)

type roleRepo struct {
	gen *rand.Rand
	db  *Data
}

func (repo roleRepo) Get(ctx context.Context, id int64, options ...dto.RoleQueryOption) (*dto.RolePB, error) {
	var option dto.RoleQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.Role(ctx).Query().Where(role.ID(id))
	query = roleQueryOptions(query, option)
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRole2PB(result), nil
}

func (repo roleRepo) Create(ctx context.Context, rolePB *dto.RolePB, options ...dto.RoleUpdateOption) (*dto.RolePB, error) {
	var option dto.RoleUpdateOption
	if len(options) > 0 {
		option = options[0]
	}
	obj := dto.ConvertRolePB2Object(rolePB)
	if obj.Keyword == "" {
		obj.Keyword = "system:role:" + repo.gen.RandString(12)
	}
	exist, err := repo.db.Role(ctx).Query().Where(role.KeywordEqualFold(rolePB.Keyword)).Exist(ctx)
	if err != nil || exist {
		return nil, errors.New("role keyword already exists")
	}
	err = repo.db.Tx(ctx, func(ctx context.Context) error {
		create := repo.db.Role(ctx).Create()
		create.SetRole(obj, option.Fields...)
		saved, err := create.Save(ctx)
		if err != nil {
			return err
		}
		rolePB = dto.ConvertRole2PB(saved)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return rolePB, nil
}

func (repo roleRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.Tx(ctx, func(ctx context.Context) error {
		err := repo.db.Role(ctx).DeleteOneID(id).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (repo roleRepo) Update(ctx context.Context, rolePB *dto.RolePB, options ...dto.RoleUpdateOption) (*dto.RolePB, error) {
	var option dto.RoleUpdateOption
	if len(options) > 0 {
		option = options[0]
	}
	update := repo.db.Role(ctx).UpdateOneID(rolePB.Id)
	if len(rolePB.PermissionIds) > 0 {
		update.ClearPermissions()
		update.AddPermissionIDs(rolePB.PermissionIds...)
	}
	if len(rolePB.Permissions) > 0 {
		update.ClearPermissions()
		update.AddPermissions(dto.ConvertPermissionsPB2Object(rolePB.Permissions)...)
	}
	saved, err := update.SetRoleWithZero(dto.ConvertRolePB2Object(rolePB), option.Fields...).Save(ctx)
	if err != nil {
		return nil, err
	}
	rolePB = dto.ConvertRole2PB(saved)
	return rolePB, nil
}

func (repo roleRepo) List(ctx context.Context, in *pb.ListRolesRequest, options ...dto.RoleQueryOption) ([]*dto.RolePB, int32, error) {
	var option dto.RoleQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	query := repo.db.Role(ctx).Query()
	if option.IncludePermissions {
		query = query.WithPermissions()
	}
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

// NewRoleRepo .
func NewRoleRepo(db *Data, logger log.KLogger) dto.RoleRepo {
	return &roleRepo{
		gen: rand.DigitAndLowerCase,
		db:  db,
	}
}

func rolePageQuery(ctx context.Context, query *ent.RoleQuery, in *pb.ListRolesRequest, option dto.RoleQueryOption) ([]*dto.RolePB, int32, error) {
	if in.OnlyCount {
		count, err := query.Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		return nil, int32(count), nil
	}

	query = roleQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	query = db.PaginationQuery(query, in, !in.NoPaging)
	result, err := query.All(ctx)
	return dto.ConvertRoles(result), int32(count), err
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
