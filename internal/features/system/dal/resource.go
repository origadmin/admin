/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"strings"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/resource"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type resourceRepo struct {
	db        *ent.Database
	Delimiter string
}

// NewResourceRepo .
func NewResourceRepo(database *ent.Database) dto.ResourceRepo {
	return &resourceRepo{
		db:        database,
		Delimiter: "/",
	}
}

func (r *resourceRepo) Get(ctx context.Context, id int64, opts ...*dto.ResourceQueryOption) (*types.Resource, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Resource(ctx).Query().Where(resource.ID(id))

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, resource.ValidColumn, resource.FieldID, new(types.Resource))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	result, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(result), nil
}

func (r *resourceRepo) Create(ctx context.Context, res *types.Resource, opts ...*dto.ResourceCreateOption) (*types.Resource, error) {
	entResource := dto.ConvertResourcePBToResource(res)
	create := r.db.Resource(ctx).Create().
		SetResourceSkipZero(entResource).
		SetName(res.Name).
		SetSyncStatus("Modified").
		SetVersionID("").
		SetLastSyncVersionID("")

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (r *resourceRepo) CreateFromPolicy(ctx context.Context, input *dto.ResourceFromPolicyInput) (*types.Resource, error) {
	create := r.db.Resource(ctx).Create().
		SetKeyword(input.Resource.Keyword).
		SetSequence(int(input.Resource.Sequence))

	// Set Type if provided
	if input.Resource.Type != "" {
		create.SetType(input.Resource.Type)
	}

	// Set TreePath if provided
	if input.Resource.TreePath != "" {
		create.SetTreePath(input.Resource.TreePath)
	}

	// Set ParentID if provided
	if input.Resource.ParentId != 0 {
		create.SetParentID(input.Resource.ParentId)
	}

	// Set Name, I18n, ServiceName if provided
	if input.Resource.Name != "" {
		create.SetName(input.Resource.Name)
	}

	if input.Resource.I18N != "" {
		create.SetI18n(input.Resource.I18N)
	}

	if input.Resource.ServiceName != "" {
		create.SetServiceName(input.Resource.ServiceName)
	}

	// Only set policy-related fields if Policy is provided
	if input.Policy != nil {
		policy := input.Policy

		// Extract method and path from GatewayPath if not already provided
		method := input.Resource.Method
		path := input.Resource.Path
		if policy.GatewayPath != "" && (method == "" || path == "") {
			if parts := strings.SplitN(policy.GatewayPath, ":", 2); len(parts) == 2 {
				if method == "" {
					method = parts[0]
				}
				if path == "" {
					path = parts[1]
				}
			}
		}

		// Ensure path has the correct prefix
		if path != "" && !strings.HasPrefix(path, conf.APIPrefix) {
			path = conf.APIPrefix + path
		}

		create.SetPath(path).
			SetMethod(method).
			SetOperation(policy.ServiceMethod).
			SetPolicy(policy.Name).
			SetVersionID(policy.VersionID).
			SetLastSyncVersionID(policy.VersionID).
			SetSyncStatus("Synced")
	}

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (r *resourceRepo) Delete(ctx context.Context, id int64) error {
	return r.db.Resource(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *resourceRepo) Update(ctx context.Context, res *types.Resource, opts ...*dto.ResourceUpdateOption) (*types.Resource, error) {
	opt := repo.GetFirstOption(opts...)
	entResource := dto.ConvertResourcePBToResource(res)
	update := r.db.Resource(ctx).UpdateOneID(res.Id)

	updateCols := db.UpdateFields(opt.UpdateMask, resource.ValidColumn, res)
	if len(updateCols) > 0 {
		// If a field mask is present, update only the specified fields, including zero values.
		update.SetResource(entResource, updateCols...)
	} else {
		// If no field mask, skip zero values to prevent accidental clearing of fields.
		update.SetResourceSkipZero(entResource)
	}

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(saved), nil
}

func (r *resourceRepo) List(ctx context.Context, opts ...*dto.ResourceQueryOption) ([]*types.Resource, int32, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.Resource(ctx).Query()

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.Keyword != "" {
		query.Where(resource.KeywordContains(opt.Keyword))
	}

	if opt.Operation != "" {
		query.Where(resource.Operation(opt.Operation))
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, resource.ValidColumn, resource.FieldID, new(types.Resource))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	if opt.OrderBy != nil {
		orders := db.OrderBy[resource.OrderOption](opt.OrderBy)
		if len(orders) > 0 {
			query.Order(orders...)
		}
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertResourcesToResourcesPB(result), count, err
}
