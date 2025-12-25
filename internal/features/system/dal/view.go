package dal

import (
	"context"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/view"
	"origadmin/application/admin/internal/features/system/dto"
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
func (r *viewRepo) Get(ctx context.Context, id int64) (*types.View, error) {
	result, err := r.db.View(ctx).Query().Where(view.ID(id)).Only(ctx)
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

	// Get the total count before applying pagination.
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if opt.Page > 0 && opt.PageSize > 0 {
		query.Offset((opt.Page - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	result, err := query.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return dto.ConvertViewsToViewsPB(result), int32(count), nil
}

// Create creates a new view.
func (r *viewRepo) Create(ctx context.Context, in *types.View, opts ...*dto.ViewCreateOption) (*types.View, error) {
	create := r.db.View(ctx).Create().
		SetKeyword(in.Keyword).
		SetName(in.Name).
		SetScope(in.Scope).
		SetType(in.Type).
		SetNillableComponent(&in.Component).
		SetNillablePath(&in.Path).
		SetNillableIcon(&in.Icon).
		SetVisible(in.Visible).
		SetSequence(int(in.Sequence))

	if in.ParentId > 0 {
		create.SetParentID(in.ParentId)
	}

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertViewToViewPB(saved), nil
}

// Update updates an existing view.
func (r *viewRepo) Update(ctx context.Context, in *types.View, opts ...*dto.ViewUpdateOption) (*types.View, error) {
	update := r.db.View(ctx).UpdateOneID(in.Id).
		SetKeyword(in.Keyword).
		SetName(in.Name).
		SetScope(in.Scope).
		SetType(in.Type).
		SetNillableComponent(&in.Component).
		SetNillablePath(&in.Path).
		SetNillableIcon(&in.Icon).
		SetVisible(in.Visible).
		SetSequence(int(in.Sequence))

	if in.ParentId > 0 {
		update.SetParentID(in.ParentId)
	} else {
		update.ClearParent()
	}

	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertViewToViewPB(saved), nil
}

// Delete deletes a view by its ID (soft delete).
func (r *viewRepo) Delete(ctx context.Context, id int64) error {
	return r.db.View(ctx).DeleteOneID(id).Exec(ctx)
}
