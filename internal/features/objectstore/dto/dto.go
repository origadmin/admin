/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto provides data transfer objects and repository interfaces for the objectstore module.
package dto

import (
	"context"
	"io"
	"time"

	"origadmin/application/admin/api/v1/services/types"
)

//go:generate abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/features/objectstore/biz,alias=biz
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="biz,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"

// ObjectRepo defines the interface for object storage data access.
type ObjectRepo interface {
	// Put stores data and returns the object metadata.
	Put(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error)

	// Get retrieves data for the given ID.
	Get(ctx context.Context, id string) (io.ReadCloser, *types.Object, error)

	// Delete removes the object with the given ID.
	Delete(ctx context.Context, id string) error

	// GetPresignedURL generates a presigned URL for downloading the object.
	GetPresignedURL(ctx context.Context, id string, expires time.Duration) (string, error)

	// InitiateMultipartUpload initiates a multipart upload and returns an upload ID.
	InitiateMultipartUpload(ctx context.Context, name string, contentType string) (uploadID string, objectID string, err error)

	// GetMultipartUploadURL generates a presigned URL for uploading a specific part.
	GetMultipartUploadURL(ctx context.Context, objectID string, uploadID string, partNumber int32, expires time.Duration) (string, error)

	// ListParts lists the parts that have been uploaded for a specific multipart upload.
	ListParts(ctx context.Context, objectID string, uploadID string) ([]*types.PartInfo, error)

	// CompleteMultipartUpload completes a multipart upload by assembling the parts.
	CompleteMultipartUpload(ctx context.Context, objectID string, uploadID string, parts []*types.PartInfo) (*types.Object, error)

	// AbortMultipartUpload aborts a multipart upload and cleans up resources.
	AbortMultipartUpload(ctx context.Context, objectID string, uploadID string) error
}
