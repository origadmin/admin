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
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dto"
)

type userRepo struct {
	db *Data
}

func (repo userRepo) GetByUserName(ctx context.Context, username string, fields ...string) (*dto.UserPB, error) {
	query := repo.db.User(ctx).Query().Where(user.UsernameEQ(username))
	var option dto.UserQueryOption
	if len(fields) > 0 {
		option.SelectFields = fields
	}
	query = userQueryOptions(query, option)
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUser2PB(result), nil
}

func (repo userRepo) GetRoleIDs(ctx context.Context, id string) ([]string, error) {
	return repo.db.User(ctx).Query().Where(user.ID(id)).QueryRoles().IDs(ctx)
}

func (repo userRepo) Get(ctx context.Context, id string, options ...dto.UserQueryOption) (*dto.UserPB, error) {
	var option dto.UserQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.User(ctx).Query().Where(user.ID(id))
	query = userQueryOptions(query, option)
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUser2PB(result), nil
}

func (repo userRepo) Create(ctx context.Context, user *dto.UserPB, options ...dto.UserQueryOption) (*dto.UserPB, error) {
	var option dto.UserQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	create := repo.db.User(ctx).Create()
	create.SetUser(dto.UserObject(user), option.Fields...)
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUser2PB(saved), nil
}

func (repo userRepo) Delete(ctx context.Context, id string) error {
	return repo.db.User(ctx).DeleteOneID(id).Exec(ctx)
}

func (repo userRepo) Update(ctx context.Context, user *dto.UserPB, options ...dto.UserQueryOption) (*dto.UserPB, error) {
	update := repo.db.User(ctx).UpdateOneID(user.Id)
	update.SetUser(dto.UserObject(user))
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertUser2PB(saved), nil
}

func (repo userRepo) List(ctx context.Context, in *pb.ListUsersRequest, options ...dto.UserQueryOption) ([]*dto.UserPB, int32, error) {
	var option dto.UserQueryOption
	if len(options) > 0 {
		option = options[0]
	}

	query := repo.db.User(ctx).Query()
	if v := option.Username; len(v) > 0 {
		query = query.Where(user.UsernameContains(v))
	}
	if v := option.Name; len(v) > 0 {
		query = query.Where(user.NameContains(v))
	}
	if v := option.Status; v > 0 {
		query = query.Where(user.StatusEQ(v))
	}

	return userPageQuery(ctx, query, in, option)
}

// NewUserRepo .
func NewUserRepo(db *Data, logger log.Logger) dto.UserRepo {
	return &userRepo{
		db: db,
	}
}

func userPageQuery(ctx context.Context, query *ent.UserQuery, in *pb.ListUsersRequest, option dto.UserQueryOption) ([]*dto.UserPB, int32, error) {
	if in.OnlyCount {
		count, err := query.Count(ctx)
		if err != nil {
			return nil, 0, err
		}
		return nil, int32(count), nil
	}

	query = userQueryPage(query, in)
	query = userQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertUsers(result), int32(count), err
}

func userQueryPage(query *ent.UserQuery, in *pb.ListUsersRequest) *ent.UserQuery {
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

func userQueryOptions(query *ent.UserQuery, option dto.UserQueryOption) *ent.UserQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).UserQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).UserQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(userOrderBy(option.OrderFields)...)
	}
	return query
}

func userOrderBy(fields []string, opts ...sql.OrderTermOption) []user.OrderOption {
	var orders []user.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
