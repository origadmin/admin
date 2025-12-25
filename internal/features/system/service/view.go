package service

import (
	"context"

	"github.com/origadmin/runtime/errors"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/data/entity/ent"
)

// ListViews handles the RPC for listing views.
func (s *SystemService) ListViews(ctx context.Context, req *system.ListViewsRequest) (*system.ListViewsResponse, error) {
	views, total, err := s.View.ListViews(ctx, req)
	if err != nil {
		return nil, err
	}
	return &system.ListViewsResponse{
		Views:    views,
		Total:    total,
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}, nil
}

// GetView handles the RPC for getting a single view.
func (s *SystemService) GetView(ctx context.Context, req *system.GetViewRequest) (*system.GetViewResponse, error) {
	view, err := s.View.GetView(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("VIEW_NOT_FOUND", "View not found")
		}
		return nil, err
	}
	return &system.GetViewResponse{View: view}, nil
}

// CreateView handles the RPC for creating a new view.
func (s *SystemService) CreateView(ctx context.Context, req *system.CreateViewRequest) (*system.CreateViewResponse, error) {
	view, err := s.View.CreateView(ctx, req.GetView())
	if err != nil {
		return nil, err
	}
	return &system.CreateViewResponse{View: view}, nil
}

// UpdateView handles the RPC for updating an existing view.
func (s *SystemService) UpdateView(ctx context.Context, req *system.UpdateViewRequest) (*system.UpdateViewResponse, error) {
	view, err := s.View.UpdateView(ctx, req.GetView())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("VIEW_NOT_FOUND", "View not found")
		}
		return nil, err
	}
	return &system.UpdateViewResponse{View: view}, nil
}

// DeleteView handles the RPC for deleting a view.
func (s *SystemService) DeleteView(ctx context.Context, req *system.DeleteViewRequest) (*system.DeleteViewResponse, error) {
	err := s.View.DeleteView(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("VIEW_NOT_FOUND", "View not found")
		}
		return nil, err
	}
	return &system.DeleteViewResponse{}, nil
}
