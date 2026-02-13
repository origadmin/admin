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
var ProviderSet = wire.NewSet(NewObjectUseCase)

// ObjectUseCase is the use case for object storage.
type ObjectUseCase struct {
	repo dto.ObjectRepo
	log  *log.Helper
}

// NewObjectUseCase creates a new ObjectUseCase.
func NewObjectUseCase(repo dto.ObjectRepo, logger log.Logger) *ObjectUseCase {
	return &ObjectUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "objectstore.biz")),
	}
}

// UploadObject uploads a new object.
func (uc *ObjectUseCase) UploadObject(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error) {
	return uc.repo.Put(ctx, name, data, size)
}

// GetObject retrieves an object's metadata and content reader.
func (uc *ObjectUseCase) GetObject(ctx context.Context, id string) (io.ReadCloser, *types.Object, error) {
	return uc.repo.Get(ctx, id)
}

// DeleteObject deletes an object.
func (uc *ObjectUseCase) DeleteObject(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

// GetPresignedURL generates a presigned URL for downloading the object.
func (uc *ObjectUseCase) GetPresignedURL(ctx context.Context, id string, expires time.Duration) (string, error) {
	return uc.repo.GetPresignedURL(ctx, id, expires)
}

// InitiateMultipartUpload initiates a multipart upload.
func (uc *ObjectUseCase) InitiateMultipartUpload(ctx context.Context, name string, contentType string) (string, string, error) {
	return uc.repo.InitiateMultipartUpload(ctx, name, contentType)
}

// GetMultipartUploadURL generates a presigned URL for uploading a specific part.
func (uc *ObjectUseCase) GetMultipartUploadURL(ctx context.Context, objectID string, uploadID string, partNumber int32, expires time.Duration) (string, error) {
	return uc.repo.GetMultipartUploadURL(ctx, objectID, uploadID, partNumber, expires)
}

// ListParts lists the parts that have been uploaded for a specific multipart upload.
func (uc *ObjectUseCase) ListParts(ctx context.Context, objectID string, uploadID string) ([]*types.PartInfo, error) {
	return uc.repo.ListParts(ctx, objectID, uploadID)
}

// CompleteMultipartUpload completes a multipart upload.
func (uc *ObjectUseCase) CompleteMultipartUpload(ctx context.Context, objectID string, uploadID string, parts []*types.PartInfo) (*types.Object, error) {
	return uc.repo.CompleteMultipartUpload(ctx, objectID, uploadID, parts)
}

// AbortMultipartUpload aborts a multipart upload.
func (uc *ObjectUseCase) AbortMultipartUpload(ctx context.Context, objectID string, uploadID string) error {
	return uc.repo.AbortMultipartUpload(ctx, objectID, uploadID)
}
