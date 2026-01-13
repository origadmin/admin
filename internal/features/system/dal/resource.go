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
	opt := repo.FirstOrDefault(opts...)
	query := r.db.Resource(ctx).Query().Where(resource.ID(id))

	if opt.WithPermissions {
		query.WithPermissions()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, resource.ValidColumn, resource.FieldID, new(types.Resource))
		query = s.ResourceQuery
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
	policy := input.Policy

	// Extract method and path from GatewayPath, e.g., "GET:/api/v1/users/{id}"
	var method, path string
	if policy.GatewayPath != "" {
		if parts := strings.SplitN(policy.GatewayPath, ":", 2); len(parts) == 2 {
			method = parts[0]
			path = parts[1]
		}
	}

	// Ensure path has the correct prefix
	if path != "" && !strings.HasPrefix(path, conf.APIPrefix) {
		path = conf.APIPrefix + path
	}

	create := r.db.Resource(ctx).Create().
		SetKeyword(input.Keyword).
		SetPath(path).
		SetMethod(method).
		SetOperation(policy.ServiceMethod).
		SetPolicy(policy.Name).
		SetVersionID(policy.VersionID).
		SetLastSyncVersionID(policy.VersionID).
		SetSyncStatus("Synced").
		SetSequence(input.Sequence)

	if input.DisplayName != "" {
		create.SetName(input.DisplayName)
	}

	if input.I18n != "" {
		create.SetI18n(input.I18n)
	}

	if input.ServiceName != "" {
		create.SetServiceName(input.ServiceName)
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
	var updatedResource *ent.Resource
	err := r.db.Tx(ctx, func(tx context.Context) error {
		opt := repo.FirstOrDefault(opts...)
		entResource := dto.ConvertResourcePBToResource(res)
		update := r.db.Resource(tx).UpdateOneID(res.Id)

		updateCols := db.UpdateFields(opt.UpdateMask, resource.ValidColumn, res)
		if len(updateCols) > 0 {
			// If a field mask is present, update only the specified fields, including zero values.
			update.SetResource(entResource, updateCols...)
		} else {
			// If no field mask, skip zero values to prevent accidental clearing of fields.
			update.SetResourceSkipZero(entResource)
		}

		var err error
		updatedResource, err = update.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return dto.ConvertResourceToResourcePB(updatedResource), nil
}

func (r *resourceRepo) List(ctx context.Context, opts ...*dto.ResourceQueryOption) ([]*types.Resource, int32, error) {
	opt := repo.FirstOrDefault(opts...)
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
		s := db.SelectFields(query, opt.ReadMask, resource.ValidColumn, resource.FieldID, new(types.Resource))
		query = s.ResourceQuery
	}

	// Sorting logic: only apply sorting from request if it's not already handled by a page token.
	if !opt.SortFromToken {
		if len(opt.OrderBy) > 0 {
			orders := db.OrderBy[resource.OrderOption](opt.OrderBy)
			if len(orders) > 0 {
				query.Order(orders...)
			}
		}
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertResourcesToResourcesPB(result), count, err
}
