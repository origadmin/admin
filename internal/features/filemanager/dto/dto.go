/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto provides data transfer objects and repository contracts for the filemanager module.
package dto

import (
	"context"

	"origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/helpers/repo"
)

//go:generate abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/data/entity/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="ent,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"
//go:abgen:convert="source=ent.File,target=types.FileMetadata,direction=both"

// FileRepo defines the interface for file metadata data access.
type FileRepo interface {
	// Create creates a new file metadata record.
	Create(ctx context.Context, file *types.FileMetadata, opts ...*FileCreateOption) (*types.FileMetadata, error)

	// Get retrieves a file metadata record by its ID.
	Get(ctx context.Context, id int64, opts ...*FileQueryOption) (*types.FileMetadata, error)

	// List retrieves a list of file metadata records based on query options.
	List(ctx context.Context, opts ...*FileQueryOption) ([]*types.FileMetadata, int, error)

	// Update updates a file metadata record.
	Update(ctx context.Context, file *types.FileMetadata, opts ...*FileUpdateOption) (*types.FileMetadata, error)

	// Delete removes a file metadata record by its ID.
	Delete(ctx context.Context, id int64) error
}

// FileQueryOption specifies options for querying files.
type FileQueryOption struct {
	repo.QueryOption
	OwnerID    int64
	Visibility string
}

// FileCreateOption specifies options for creating a file.
type FileCreateOption struct {
	// Add any create-specific options here
}

// FileUpdateOption specifies options for updating a file.
type FileUpdateOption struct {
	repo.UpdateOption
	// Add any update-specific options here
}

// ListFilesRequestToQueryOption converts an API request to a query option object.
func ListFilesRequestToQueryOption(req *filemanager.ListFilesRequest) *FileQueryOption {
	if req == nil {
		return &FileQueryOption{}
	}
	return &FileQueryOption{
		QueryOption: repo.QueryOptionFromRequest(req),
		OwnerID:     req.GetOwnerId(),
		Visibility:  req.GetVisibility(),
	}
}
