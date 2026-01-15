package biz

import (
	"context"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/features/system/dto"
)

// ViewUseCase is a View use case.
type ViewUseCase struct {
	repo dto.ViewRepo
}

// NewViewUseCase new a View use case.
func NewViewUseCase(repo dto.ViewRepo) *ViewUseCase {
	return &ViewUseCase{repo: repo}
}

// ListViews retrieves a list of views.
func (uc *ViewUseCase) ListViews(ctx context.Context, opts ...*dto.ViewQueryOption) ([]*types.View, int32, error) {
	return uc.repo.List(ctx, opts...)
}

// GetView retrieves a single view by its ID.
func (uc *ViewUseCase) GetView(ctx context.Context, id int64, opts ...*dto.ViewQueryOption) (*types.View, error) {
	return uc.repo.Get(ctx, id, opts...)
}

// CreateView creates a new view, ensuring essential fields have valid default values.
func (uc *ViewUseCase) CreateView(ctx context.Context, in *types.View, opts ...*dto.ViewCreateOption) (*types.View, error) {
	// The backend must always enforce data integrity, regardless of frontend behavior.
	if in.Type == "" {
		in.Type = string(enums.ViewTypePage)
	}
	if in.Status == 0 {
		in.Status = int32(enums.StatusActive)
	}

	return uc.repo.Create(ctx, in, opts...)
}

// UpdateView updates an existing view.
func (uc *ViewUseCase) UpdateView(ctx context.Context, in *types.View, opts ...*dto.ViewUpdateOption) (*types.View, error) {
	return uc.repo.Update(ctx, in, opts...)
}

// DeleteView deletes a view by its ID.
func (uc *ViewUseCase) DeleteView(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
