/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

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
func NewViewRepo(db *ent.Database) dto.ViewRepo {
	return &viewRepo{db: db}
}

// Get retrieves a single view by its ID.
func (r *viewRepo) Get(ctx context.Context, id int64, opts ...*dto.ViewQueryOption) (*types.View, error) {
	opt := repo.GetFirstOption(opts...)
	query := r.db.View(ctx).Query().Where(view.ID(id))

	query = viewQueryOptions(query, opt)

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

	query = viewQueryOptions(query, opt)

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertViewsToViewsPB(result), count, nil
}

// Create creates a new view.
func (r *viewRepo) Create(ctx context.Context, in *types.View, opts ...*dto.ViewCreateOption) (*types.View, error) {
	entView := dto.ConvertViewPBToView(in)
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

	// The default behavior of SetView now includes zero values.
	if opt.UpdateMask == nil || !opt.UpdateMask.IsValid(in) {
		update.SetView(entView) // Use the new default SetView
	} else {
		var updateCols []string
		for _, path := range opt.UpdateMask.GetPaths() {
			if view.ValidColumn(path) {
				updateCols = append(updateCols, path)
			}
		}
		if len(updateCols) > 0 {
			update.SetView(entView, updateCols...) // Use the new default SetView
		}
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

// viewQueryOptions applies common query options to the ViewQuery.
func viewQueryOptions(query *ent.ViewQuery, option *dto.ViewQueryOption) *ent.ViewQuery {
	if option == nil {
		return query
	}

	// Handle FieldMask for field selection
	if option.ReadMask != nil {
		// Ensure the mask is valid before using it
		option.ReadMask.Normalize()
		if option.ReadMask.IsValid(new(types.View)) {
			var selectCols []string
			for _, path := range option.ReadMask.GetPaths() {
				// Directly validate against the ent schema definition.
				if view.ValidColumn(path) {
					selectCols = append(selectCols, path)
				}
			}
			// Only apply select if valid columns were found
			if len(selectCols) > 0 {
				// Always include the ID for entity hydration
				selectCols = append(selectCols, view.FieldID)
				query.Select(selectCols...)
			}
		}
	}

	// Handle OrderBy
	if len(option.OrderBy) > 0 {
		query.Order(viewOrderBy(option.OrderBy)...)
	}
	return query
}

func viewOrderBy(fields []string, opts ...sql.OrderTermOption) []view.OrderOption {
	var orders []view.OrderOption
	for _, field := range fields {
		// Here you might also want a map for sorting fields if they differ from DB columns
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
