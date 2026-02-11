/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"errors"
	"strconv"

	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
	"origadmin/application/admin/internal/helpers/repo"
)

type viewRepo struct {
	db        *ent.Database
	log       *log.Helper
	Delimiter string
}

// NewViewRepo creates a new view repository.
func NewViewRepo(database *ent.Database, logger log.Logger) dto.ViewRepo {
	return &viewRepo{
		db:        database,
		Delimiter: "/",
		log:       log.NewHelper(log.With(logger, "module", "dal.view")),
	}
}

// Get retrieves a single view by its ID.
func (r *viewRepo) Get(ctx context.Context, id int64, opts ...*dto.ViewQueryOption) (*types.View, error) {
	r.log.WithContext(ctx).Debugw("msg", "Get", "id", id)
	opt := repo.FirstOrDefault(opts...)
	query := r.db.View(ctx).Query().Where(view.ID(id))

	if opt.WithResources {
		query.WithResources()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, view.ValidColumn, view.FieldID,
			new(types.View))
		query = s.ViewQuery
	}

	result, err := query.Only(ctx)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Get", "err", err)
		return nil, err
	}
	return dto.ConvertViewToViewPB(result), nil
}

// List retrieves a list of views based on query options.
func (r *viewRepo) List(ctx context.Context, opts ...*dto.ViewQueryOption) ([]*types.View, int32, error) {
	r.log.WithContext(ctx).Debugw("msg", "List", "opts", opts)
	opt := repo.FirstOrDefault(opts...)
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

	if opt.WithResources {
		query.WithResources()
	}

	if opt.ReadMask != nil {
		s := db.SelectFields(query, opt.ReadMask, view.ValidColumn, view.FieldID,
			new(types.View))
		query = s.ViewQuery
	}

	result, count, err := db.Find(ctx, query, &opt.QueryOption)
	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "List", "err", err)
		return nil, 0, err
	}

	return dto.ConvertViewsToViewsPB(result), count, nil
}

// Create creates a new view and its associations within a transaction.
func (r *viewRepo) Create(ctx context.Context, in *types.View, opts ...*dto.ViewCreateOption) (*types.View, error) {
	r.log.WithContext(ctx).Debugw("msg", "Create", "in", in)
	if in == nil {
		return nil, errors.New("input view data cannot be nil")
	}
	opt := repo.FirstOrDefault(opts...)
	var createdView *ent.View
	err := r.db.Tx(ctx, func(tx context.Context) error {
		var err error
		// Calculate TreePath before converting to ent object
		if in.ParentId > 0 {
			parent, err := r.db.View(tx).Get(ctx, in.ParentId)
			if err != nil {
				r.log.WithContext(ctx).Errorw("msg", "Create.GetParent", "err", err)
				return err
			}
			in.TreePath = parent.TreePath + strconv.FormatInt(parent.ID, 10) + r.Delimiter
		}

		entView := dto.ConvertViewPBToView(in)
		create := r.db.View(tx).Create().SetViewSkipZero(entView)

		// Add resource associations if they exist
		if opt.WithResourceIDs != nil {
			create.AddResourceIDs(opt.WithResourceIDs...)
		}

		saved, err := create.Save(ctx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Create.Save", "err", err)
			return err
		}

		// Eager load the associated resources to ensure the returned object is complete.
		createdView, err = r.db.View(tx).Query().Where(view.ID(saved.ID)).WithResources().Only(ctx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Create.EagerLoad", "err", err)
		}
		return err
	})

	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Create.Tx", "err", err)
		return nil, err
	}

	return dto.ConvertViewToViewPB(createdView), nil
}

// Update updates an existing view and its associations within a transaction.
func (r *viewRepo) Update(ctx context.Context, in *types.View, opts ...*dto.ViewUpdateOption) (*types.View, error) {
	r.log.WithContext(ctx).Debugw("msg", "Update", "in", in)
	if in == nil {
		return nil, errors.New("input view data cannot be nil")
	}

	var updatedView *ent.View
	err := r.db.Tx(ctx, func(tx context.Context) error {
		var err error
		opt := repo.FirstOrDefault(opts...)
		entView := dto.ConvertViewPBToView(in)

		// 1. Start a single update builder using the transaction client
		updateBuilder := r.db.View(tx).UpdateOneID(in.Id)

		// 2. Chain scalar field updates
		updateCols := db.UpdateFields(opt.UpdateMask, view.ValidColumn, in)
		if len(updateCols) > 0 {
			updateBuilder.SetView(entView, updateCols...)
		} else {
			updateBuilder.SetViewSkipZero(entView)
		}

		// 3. Unconditionally replace resource associations
		updateBuilder.ClearResources()
		if len(opt.WithResourceIDs) > 0 {
			updateBuilder.AddResourceIDs(opt.WithResourceIDs...)
		}

		// 4. Execute a single Save operation
		if _, err = updateBuilder.Save(ctx); err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Update.Save", "err", err)
			return err
		}

		// 5. Eager load the complete, updated entity for the return value
		updatedView, err = r.db.View(tx).Query().Where(view.ID(in.Id)).WithResources().Only(ctx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Update.EagerLoad", "err", err)
		}
		return err
	})

	if err != nil {
		r.log.WithContext(ctx).Errorw("msg", "Update.Tx", "err", err)
		return nil, err
	}

	return dto.ConvertViewToViewPB(updatedView), nil
}

// Delete deletes a view by its ID.
func (r *viewRepo) Delete(ctx context.Context, id int64) error {
	r.log.WithContext(ctx).Debugw("msg", "Delete", "id", id)
	return r.db.Tx(ctx, func(txCtx context.Context) error {
		// Clear all associations before deletion
		err := r.db.View(txCtx).UpdateOneID(id).
			ClearResources().
			ClearPermissions().
			ClearChildren().
			Exec(txCtx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Delete.ClearAssociations", "err", err)
			return err
		}

		// Delete the View entity
		err = r.db.View(txCtx).DeleteOneID(id).Exec(txCtx)
		if err != nil {
			r.log.WithContext(ctx).Errorw("msg", "Delete.Exec", "err", err)
		}
		return err
	})
}
