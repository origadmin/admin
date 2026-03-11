/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"io"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/objectstore/dto"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewObjectStoreUseCase)

// ObjectStoreUseCase is the use case for object storage.
type ObjectStoreUseCase struct {
	repo dto.ObjectRepo
	log  *log.Helper
}

// NewObjectStoreUseCase creates a new ObjectStoreUseCase.
func NewObjectStoreUseCase(repo dto.ObjectRepo, logger log.Logger) *ObjectStoreUseCase {
	return &ObjectStoreUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "objectstore.biz")),
	}
}

// UploadObject uploads a new object.
func (uc *ObjectStoreUseCase) UploadObject(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error) {
	return uc.repo.Put(ctx, name, data, size)
}

// GetObject retrieves an object's metadata.
func (uc *ObjectStoreUseCase) GetObject(ctx context.Context, id string) (*types.Object, error) {
	rc, obj, err := uc.repo.Get(ctx, id)
	if err == nil && rc != nil {
		_ = rc.Close()
	}
	return obj, err
}

// ListObjects retrieves a list of objects.
func (uc *ObjectStoreUseCase) ListObjects(ctx context.Context, prefix string, page, pageSize int32) ([]*types.Object, int32, error) {
	return uc.repo.List(ctx, prefix, page, pageSize)
}

// DownloadObject retrieves an object content reader.
func (uc *ObjectStoreUseCase) DownloadObject(ctx context.Context, id string) (io.ReadCloser, *types.Object, error) {
	return uc.repo.Get(ctx, id)
}

// DeleteObject deletes an object.
func (uc *ObjectStoreUseCase) DeleteObject(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

// GetPresignedURL generates a presigned URL.
func (uc *ObjectStoreUseCase) GetPresignedURL(ctx context.Context, id string, expires time.Duration) (string, error) {
	return uc.repo.GetPresignedURL(ctx, id, expires)
}

// InitiateMultipartUpload initiates a multipart upload.
func (uc *ObjectStoreUseCase) InitiateMultipartUpload(ctx context.Context, name string, contentType string) (string, string, error) {
	return uc.repo.InitiateMultipartUpload(ctx, name, contentType)
}

// GetMultipartUploadUrl generates a URL for a part.
func (uc *ObjectStoreUseCase) GetMultipartUploadUrl(ctx context.Context, objectID string, uploadID string, partNumber int32, expires time.Duration) (string, error) {
	return uc.repo.GetMultipartUploadURL(ctx, objectID, uploadID, partNumber, expires)
}

// ListParts lists uploaded parts.
func (uc *ObjectStoreUseCase) ListParts(ctx context.Context, objectID string, uploadID string) ([]*types.PartInfo, error) {
	return uc.repo.ListParts(ctx, objectID, uploadID)
}

// UploadPart uploads a part.
func (uc *ObjectStoreUseCase) UploadPart(ctx context.Context, objectID string, uploadID string, partNumber int32, data io.Reader) (string, error) {
	return uc.repo.UploadPart(ctx, objectID, uploadID, partNumber, data)
}

// CompleteMultipartUpload completes a multipart upload.
func (uc *ObjectStoreUseCase) CompleteMultipartUpload(ctx context.Context, objectID string, uploadID string, parts []*types.PartInfo) (*types.Object, error) {
	return uc.repo.CompleteMultipartUpload(ctx, objectID, uploadID, parts)
}

// AbortMultipartUpload aborts a multipart upload.
func (uc *ObjectStoreUseCase) AbortMultipartUpload(ctx context.Context, objectID string, uploadID string) error {
	return uc.repo.AbortMultipartUpload(ctx, objectID, uploadID)
}
