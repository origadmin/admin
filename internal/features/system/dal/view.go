/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type viewRepo struct {
	db *ent.Database
}

// NewViewRepo creates a new view repository.
func NewViewRepo(database *ent.Database) dto.ViewRepo {
	return &viewRepo{db: database}
}

// Get retrieves a single view by its ID.
func (r *viewRepo) Get(ctx context.Context, id int64, opts ...*dto.ViewQueryOption) (*types.View, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.View(ctx).Query().Where(view.ID(id))

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, view.ValidColumn, view.FieldID, new(types.View))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	result, err := query.Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertViewToViewPB(result), nil
}

// List retrieves a list of views based on query options.
func (r *viewRepo) List(ctx context.Context, opts ...*dto.ViewQueryOption) ([]*types.View, int32, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.View(ctx).Query()

	// Apply filters
	if opt.Keyword != "" {
		query.Where(view.Or(
			view.NameContains(opt.Keyword),
			view.KeywordContains(opt.Keyword),
		))
	}
	if opt.Scope != "" {
		query.Where(view.ScopeEQ(opt.Scope))
	}

	if opt.ReadMask != nil {
		selectCols := db.SelectFields(opt.ReadMask, view.ValidColumn, view.FieldID, new(types.View))
		if len(selectCols) > 0 {
			query.Select(selectCols...)
		}
	}

	if opt.OrderBy != nil {
		orders := db.OrderBy[view.OrderOption](opt.OrderBy)
		if len(orders) > 0 {
			query.Order(orders...)
		}
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertViewsToViewsPB(result), count, nil
}

// Create creates a new view.
func (r *viewRepo) Create(ctx context.Context, in *types.View, opts ...*dto.ViewCreateOption) (*types.View, error) {
	entView := dto.ConvertViewPBToView(in)
	// After template modification, SetView is now the method that includes zero values.
	create := r.db.View(ctx).Create().SetView(entView)
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertViewToViewPB(saved), nil
}

// Update updates an existing view. It supports partial updates via FieldMask.
func (r *viewRepo) Update(ctx context.Context, in *types.View, opts ...*dto.ViewUpdateOption) (*types.View, error) {
	opt := repo.GetFirstOption(opts...)
	entView := dto.ConvertViewPBToView(in)
	update := r.db.View(ctx).UpdateOneID(in.Id)

	// After template modification, SetView is now the method that includes zero values.
	updateCols := db.UpdateFields(opt.UpdateMask, view.ValidColumn, in)
	if len(updateCols) > 0 {
		// The primary key should never be in the update list.
		// UpdateFields already ensures this.
		update.SetView(entView, updateCols...)
	} else {
		update.SetView(entView)
	}

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertViewToViewPB(saved), nil
}

// Delete deletes a view by its ID.
func (r *viewRepo) Delete(ctx context.Context, id int64) error {
	return r.db.View(ctx).DeleteOneID(id).Exec(ctx)
}
