/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"
	"strconv"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type viewRepo struct {
	db        *ent.Database
	Delimiter string
}

// NewViewRepo creates a new view repository.
func NewViewRepo(database *ent.Database) dto.ViewRepo {
	return &viewRepo{
		db:        database,
		Delimiter: "/",
	}
}

// Get retrieves a single view by its ID.
func (r *viewRepo) Get(ctx context.Context, id int64, opts ...*dto.ViewQueryOption) (*types.View, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.View(ctx).Query().Where(view.ID(id))

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, view.ValidColumn, view.FieldID,
			new(types.View))
		query = s.ViewQuery
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
		s := db.SelectFields(query, opt.ReadMask, view.ValidColumn, view.FieldID,
			new(types.View))
		query = s.ViewQuery
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertViewsToViewsPB(result), count, nil
}

// Create creates a new view and its associations within a transaction.
func (r *viewRepo) Create(ctx context.Context, in *types.View, opts ...*dto.ViewCreateOption) (*types.View, error) {
	if in == nil {
		return nil, errors.New("input view data cannot be nil")
	}

	var createdView *ent.View
	err := r.db.Tx(ctx, func(tx context.Context) error {
		var err error
		// Calculate TreePath before converting to ent object
		if in.ParentId > 0 {
			parent, err := r.db.View(tx).Get(ctx, in.ParentId)
			if err != nil {
				return err
			}
			in.TreePath = parent.TreePath + strconv.FormatInt(parent.ID, 10) + r.Delimiter
		}

		entView := dto.ConvertViewPBToView(in)
		create := r.db.View(tx).Create().SetViewSkipZero(entView)

		// Add resource associations if they exist
		if len(in.GetResourceIds()) > 0 {
			create.AddResourceIDs(in.GetResourceIds()...)
		}

		saved, err := create.Save(ctx)
		if err != nil {
			return err
		}

		// Eager load the associated resources to ensure the returned object is complete.
		createdView, err = r.db.View(tx).Query().Where(view.ID(saved.ID)).WithResources().Only(ctx)
		return err
	})

	if err != nil {
		return nil, err
	}

	return dto.ConvertViewToViewPB(createdView), nil
}

// Update updates an existing view and its associations within a transaction.
func (r *viewRepo) Update(ctx context.Context, in *types.View, opts ...*dto.ViewUpdateOption) (*types.View, error) {
	if in == nil {
		return nil, errors.New("input view data cannot be nil")
	}

	var updatedView *ent.View
	err := r.db.Tx(ctx, func(tx context.Context) error {
		var err error
		opt := repo.GetFirstOption(opts...)
		entView := dto.ConvertViewPBToView(in)

		// Step 1: Update scalar fields first.
		scalarUpdate := r.db.View(tx).UpdateOneID(in.Id)
		updateCols := db.UpdateFields(opt.UpdateMask, view.ValidColumn, in)
		if len(updateCols) > 0 {
			scalarUpdate.SetView(entView, updateCols...)
		} else {
			scalarUpdate.SetViewSkipZero(entView)
		}
		if err = scalarUpdate.Exec(ctx); err != nil {
			return err
		}

		// Step 2: Separately update the M2M relation.
		err = r.db.View(tx).UpdateOneID(in.Id).
			ClearResources().
			AddResourceIDs(in.GetResourceIds()...).
			Exec(ctx)
		if err != nil {
			return err
		}

		// Step 3: Eager load the complete, updated entity.
		updatedView, err = r.db.View(tx).Query().Where(view.ID(in.Id)).WithResources().Only(ctx)
		return err
	})

	if err != nil {
		return nil, err
	}

	return dto.ConvertViewToViewPB(updatedView), nil
}

// Delete deletes a view by its ID.
func (r *viewRepo) Delete(ctx context.Context, id int64) error {
	return r.db.View(ctx).DeleteOneID(id).Exec(ctx)
}
