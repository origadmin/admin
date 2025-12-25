/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"

	"github.com/origadmin/toolkits/crypto/rand"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type roleRepo struct {
	db  *ent.Database
	gen rand.Generator
}

// NewRoleRepo .
func NewRoleRepo(database *ent.Database) dto.RoleRepo {
	generator := rand.NewGenerator(rand.KindDigit | rand.KindLowerCase)
	return &roleRepo{db: database, gen: generator}
}

func (r *roleRepo) Get(ctx context.Context, id int64, opts ...*dto.RoleQueryOption) (*types.Role, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Role(ctx).Query().Where(role.ID(id))

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, role.ValidColumn, role.FieldID, new(types.Role))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	result, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(result), nil
}

func (r *roleRepo) Create(ctx context.Context, rl *types.Role, opts ...*dto.RoleCreateOption) (*types.Role, error) {
	if rl.Keyword == "" {
		randString, err := r.gen.RandString(12)
		if err != nil {
			randString = ""
		}
		rl.Keyword = "system:role:" + randString
	}
	exist, err := r.db.Role(ctx).Query().Where(role.KeywordEqualFold(rl.Keyword)).Exist(ctx)
	if err != nil || exist {
		return nil, errors.New("role keyword already exists")
	}

	entRole := dto.ConvertRolePBToRole(rl)
	create := r.db.Role(ctx).Create().SetRole(entRole)

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(saved), nil
}

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	return r.db.Role(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *roleRepo) Update(ctx context.Context, rl *types.Role, opts ...*dto.RoleUpdateOption) (*types.Role, error) {
	opt := repo.GetFirstOption(opts...)
	entRole := dto.ConvertRolePBToRole(rl)
	update := r.db.Role(ctx).UpdateOneID(rl.Id)

	updateCols := db.UpdateFields(opt.UpdateMask, role.ValidColumn, rl)
	if len(updateCols) > 0 {
		update.SetRole(entRole, updateCols...)
	} else {
		update.SetRole(entRole)
	}

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(saved), nil
}

func (r *roleRepo) List(ctx context.Context, opts ...*dto.RoleQueryOption) ([]*types.Role, int32, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Role(ctx).Query()

	if opt.Keyword != "" {
		query.Where(role.NameContainsFold(opt.Keyword))
	}

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, role.ValidColumn, role.FieldID, new(types.Role))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	if opt.OrderBy != nil {
		orders := db.OrderBy[role.OrderOption](opt.OrderBy)
		if len(orders) > 0 {
			query.Order(orders...)
		}
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertRolesToRolesPB(result), count, err
}

func (r *roleRepo) GetPermissions(ctx context.Context, id int64) ([]*types.Permission, error) {
	permissions, err := r.db.Role(ctx).Query().Where(role.ID(id)).QueryPermissions().All(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertPermissionsToPermissionsPB(permissions), nil
}

func (r *roleRepo) UpdatePermissions(ctx context.Context, id int64, permissionIDs []int64) error {
	_, err := r.db.Role(ctx).UpdateOneID(id).ClearPermissions().AddPermissionIDs(permissionIDs...).Save(ctx)
	return err
}
