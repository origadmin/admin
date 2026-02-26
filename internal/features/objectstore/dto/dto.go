/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto

import (
	"context"
	"io"
	"time"

	"origadmin/application/admin/api/v1/services/types"
)

// ObjectRepo defines the interface for object storage operations.
type ObjectRepo interface {
	Put(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error)
	Get(ctx context.Context, id string) (io.ReadCloser, *types.Object, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, prefix string, page, pageSize int32) ([]*types.Object, int32, error)
	GetPresignedURL(ctx context.Context, id string, expires time.Duration) (string, error)
	InitiateMultipartUpload(ctx context.Context, name string, contentType string) (string, string, error)
	GetMultipartUploadURL(ctx context.Context, objectID string, uploadID string, partNumber int32, expires time.Duration) (string, error)
	ListParts(ctx context.Context, objectID string, uploadID string) ([]*types.PartInfo, error)
	CompleteMultipartUpload(ctx context.Context, objectID string, uploadID string, parts []*types.PartInfo) (*types.Object, error)
	AbortMultipartUpload(ctx context.Context, objectID string, uploadID string) error
}
