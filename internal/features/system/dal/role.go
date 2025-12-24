/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"

	"github.com/origadmin/toolkits/crypto/rand"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/role"
	"origadmin/application/admin/internal/features/system/dto"
)

type roleRepo struct {
	db  *ent.Database
	gen rand.Generator
}

// NewRoleRepo .
func NewRoleRepo(db *ent.Database) (dto.RoleRepo, error) {
	generator, err := rand.NewGenerator(rand.KindDigit | rand.KindLowerCase)
	if err != nil {
		return nil, err
	}
	return &roleRepo{db: db, gen: generator}, nil
}

func (r *roleRepo) Get(ctx context.Context, id int64, opts ...dto.RoleQueryOption) (*types.Role, error) {
	result, err := r.db.Role(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(result), nil
}

func (r *roleRepo) Create(ctx context.Context, rl *types.Role, opts ...dto.RoleMutationOption) (*types.Role, error) {
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

	create := r.db.Role(ctx).Create().
		SetName(rl.Name).
		SetKeyword(rl.Keyword)

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(saved), nil
}

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	return r.db.Role(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *roleRepo) Update(ctx context.Context, rl *types.Role, opts ...dto.RoleMutationOption) (*types.Role, error) {
	update := r.db.Role(ctx).UpdateOneID(rl.Id)
	//if len(rl.PermissionIds) > 0 {
	//	update.ClearPermissions().AddPermissionIDs(rl.PermissionIds...)
	//}

	// ... set other fields

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertRoleToRolePB(saved), nil
}

func (r *roleRepo) List(ctx context.Context, in *system.ListRolesRequest, opts ...dto.RoleQueryOption) ([]*types.Role, int32, error) {
	query := r.db.Role(ctx).Query()

	if in.GetKeyword() != "" {
		query = query.Where(role.NameContainsFold(in.GetKeyword()))
	}
	//if in.Status != nil {
	//	query = query.Where(role.StatusEQ(*in.Status))
	//}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	//query = db.QueryPage(query, in)

	result, err := query.All(ctx)
	return dto.ConvertRolesToRolesPB(result), int32(count), err
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
