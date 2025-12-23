/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data is the data access object
package data

import (
	"context"
	"errors"

	"github.com/origadmin/toolkits/crypto/rand"

	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/system/data/ent"
	"origadmin/application/admin/internal/features/system/data/ent/role"
	"origadmin/application/admin/internal/features/system/dto"
)

type roleRepo struct {
	gen rand.Generator
	db  *ent.Client
}

func (repo *roleRepo) Get(ctx context.Context, id int64, options ...dto.RoleQueryOption) (*types.Role, error) {
	result, err := repo.db.Role.Get(ctx, (id))
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(result), nil
}

func (repo *roleRepo) Create(ctx context.Context, r *types.Role, options ...dto.RoleUpdateOption) (*types.Role, error) {
	if r.Keyword == "" {
		randString, err := repo.gen.RandString(12)
		if err != nil {
			randString = ""
		}
		r.Keyword = "system:role:" + randString
	}
	exist, err := repo.db.Role.Query().Where(role.KeywordEqualFold(r.Keyword)).Exist(ctx)
	if err != nil || exist {
		return nil, errors.New("role keyword already exists")
	}

	create := repo.db.Role.Create().
		SetName(r.Name).
		SetKeyword(r.Keyword)

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(saved), nil
}

func (repo *roleRepo) Delete(ctx context.Context, id int64) error {
	return repo.db.Role.DeleteOneID((id)).Exec(ctx)
}

func (repo *roleRepo) Update(ctx context.Context, r *types.Role, options ...dto.RoleUpdateOption) (*types.Role, error) {
	update := repo.db.Role.UpdateOneID((r.Id))
	if len(r.PermissionIds) > 0 {
		update.ClearPermissions().AddPermissionIDs(r.PermissionIds...)
	}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(saved), nil
}

func (repo *roleRepo) List(ctx context.Context, in *system.ListRolesRequest, options ...dto.RoleQueryOption) ([]*types.Role, int32, error) {
	query := repo.db.Role.Query()

	//if in.Name != nil {
	//	query = query.Where(role.NameContains(*in.Name))
	//}
	//if in.Status != nil {
	//	query = query.Where(role.StatusEQ(*in.Status))
	//}

	if in.OnlyCount {
		count, err := query.Count(ctx)
		return nil, int32(count), err
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	//query = db.QueryPage(query, in)

	result, err := query.All(ctx)
	return dto.ConvertRolesToRolesPB(result), int32(count), err
}

// NewRoleRepo .
func NewRoleRepo(d *Data) (dto.RoleRepo, error) {
	return &roleRepo{
		gen: rand.NewGenerator(rand.KindDigit | rand.KindLowerCase),
		db:  d.db,
	}, nil
}
