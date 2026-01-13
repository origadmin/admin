/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/errors"
	"origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/features/system/biz"
	"origadmin/application/admin/internal/features/system/dto"
	"origadmin/application/admin/internal/helpers/db"
)

type ViewService struct {
	system.UnimplementedViewServiceServer
	uc *biz.ViewUseCase
}

func NewViewService(uc *biz.ViewUseCase) *ViewService {
	return &ViewService{uc: uc}
}

func (s *ViewService) ListViews(ctx context.Context, req *system.ListViewsRequest) (*system.ListViewsResponse, error) {
	queryOpt := dto.ListViewsRequestToQueryOption(req)
	views, total, err := s.uc.ListViews(ctx, queryOpt)
	if err != nil {
		return nil, err
	}

	pageSize := db.GetPageSize(req)
	resp := &system.ListViewsResponse{
		Views:    views,
		Total:    total,
		PageSize: int32(pageSize),
	}

	if req.GetPagingMode() == db.PagingModeCursor {
		// Only generate a next page token if the number of results equals the page size,
		// which implies there might be more data.
		if len(views) > 0 && len(views) == pageSize {
			nextToken, err := db.GenerateNextPageToken(views, &queryOpt.QueryOption)
			if err != nil {
				return nil, errors.InternalServer("TOKEN_GENERATION_FAILED", err.Error())
			}
			resp.NextPageToken = nextToken
		}
	} else {
		resp.Page = req.GetPage()
	}

	return resp, nil
}

func (s *ViewService) GetView(ctx context.Context, req *system.GetViewRequest) (*system.GetViewResponse, error) {
	view, err := s.uc.GetView(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("VIEW_NOT_FOUND", "View not found")
		}
		return nil, err
	}
	return &system.GetViewResponse{View: view}, nil
}

func (s *ViewService) CreateView(ctx context.Context, req *system.CreateViewRequest) (*system.CreateViewResponse, error) {
	opts := dto.CreateViewOptionsFromRequest(req)
	view, err := s.uc.CreateView(ctx, req.GetView(), opts)
	if err != nil {
		return nil, err
	}
	return &system.CreateViewResponse{View: view}, nil
}

func (s *ViewService) UpdateView(ctx context.Context, req *system.UpdateViewRequest) (*system.UpdateViewResponse, error) {
	opts := dto.UpdateViewOptionsFromRequest(req)
	view, err := s.uc.UpdateView(ctx, req.GetView(), opts)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("VIEW_NOT_FOUND", "View not found")
		}
		return nil, err
	}
	return &system.UpdateViewResponse{View: view}, nil
}

func (s *ViewService) DeleteView(ctx context.Context, req *system.DeleteViewRequest) (*system.DeleteViewResponse, error) {
	err := s.uc.DeleteView(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NotFound("VIEW_NOT_FOUND", "View not found")
		}
		return nil, err
	}
	return &system.DeleteViewResponse{}, nil
}
