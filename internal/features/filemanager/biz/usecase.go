/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	objclient "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/filemanager/dto"
)

// ProviderSet is biz providers for filemanager.
var ProviderSet = wire.NewSet(NewFileUseCase)

// FileUseCase is the use case for file management.
type FileUseCase struct {
	repo dto.FileRepo
	obj  objclient.ObjectStoreServiceClient
	log  *log.Helper
}

// NewFileUseCase creates a new FileUseCase.
func NewFileUseCase(repo dto.FileRepo, obj objclient.ObjectStoreServiceClient, logger log.Logger) *FileUseCase {
	return &FileUseCase{
		repo: repo,
		obj:  obj,
		log:  log.NewHelper(log.With(logger, "module", "filemanager.biz")),
	}
}

// UploadFile handles the logic for uploading a small file.
func (uc *FileUseCase) UploadFile(ctx context.Context, name string, data []byte, ownerID int64) (*types.FileMetadata, error) {
	// 1. Upload the file content to ObjectStore
	uploadReq := &objclient.UploadObjectRequest{
		Name: name,
		Data: data,
	}
	objResp, err := uc.obj.UploadObject(ctx, uploadReq)
	if err != nil {
		uc.log.Errorf("failed to upload object to objectstore: %v", err)
		return nil, err
	}

	// 2. Create the file metadata in our database
	fileMeta := &types.FileMetadata{
		Name:     name,
		ObjectId: objResp.Object.Id,
		OwnerId:  ownerID,
		Size:     objResp.Object.Size, // Correctly assign the size from the object store response
	}

	createdFile, err := uc.repo.Create(ctx, fileMeta)
	if err != nil {
		uc.log.Errorf("failed to create file metadata: %v", err)
		// TODO: Add compensation logic to delete the object from ObjectStore if metadata creation fails.
		return nil, err
	}

	return createdFile, nil
}

// GetFile retrieves file metadata and a download URL.
func (uc *FileUseCase) GetFile(ctx context.Context, id int64) (*types.FileMetadata, string, error) {
	uc.log.Info(">>>>>> [FINGERPRINT] Entering GetFile use case with new logic. <<<<<<")
	// 1. Get file metadata from our DB
	file, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, "", err
	}

	// 2. Construct the download URL pointing to the Gateway's dedicated download proxy endpoint.
	// We use a distinct path /api/v1/download/objects/... to avoid conflict with gRPC gateway routes.
	downloadURL := fmt.Sprintf("/api/v1/objects/%s", file.ObjectId)
	uc.log.Infof(">>>>>> [FINGERPRINT] Constructed download URL: %s <<<<<<", downloadURL)

	return file, downloadURL, nil
}

// InitiateMultipartUpload initiates a multipart upload for large files.
func (uc *FileUseCase) InitiateMultipartUpload(ctx context.Context, name, contentType, visibility string, size int64, ownerID int64) (string, error) {
	uc.log.Infof("InitiateMultipartUpload: name=%s, size=%d, visibility=%s", name, size, visibility)

	// 1. Initiate multipart upload in ObjectStore
	req := &objclient.InitiateMultipartUploadRequest{
		Name:        name,
		ContentType: contentType,
	}
	resp, err := uc.obj.InitiateMultipartUpload(ctx, req)
	if err != nil {
		uc.log.Errorf("failed to initiate multipart upload: %v", err)
		return "", err
	}

	// 2. Create file metadata with upload_id stored temporarily
	// Note: We'll update ObjectId after completing the upload
	fileMeta := &types.FileMetadata{
		Name:       name,
		ObjectId:   resp.UploadId, // Store upload_id temporarily in ObjectId
		OwnerId:    ownerID,
		Size:       size,
		Visibility: visibility,
	}

	_, err = uc.repo.Create(ctx, fileMeta)
	if err != nil {
		uc.log.Errorf("failed to create file metadata: %v", err)
		return "", err
	}

	return resp.UploadId, nil
}

// GetMultipartUploadUrl generates a presigned URL for uploading a part.
func (uc *FileUseCase) GetMultipartUploadUrl(ctx context.Context, uploadID string, partNumber int32) (string, error) {
	uc.log.Infof("GetMultipartUploadUrl: uploadID=%s, partNumber=%d", uploadID, partNumber)

	req := &objclient.GetMultipartUploadUrlRequest{
		UploadId:   uploadID,
		PartNumber: partNumber,
	}
	resp, err := uc.obj.GetMultipartUploadUrl(ctx, req)
	if err != nil {
		uc.log.Errorf("failed to get multipart upload URL: %v", err)
		return "", err
	}

	return resp.UploadUrl, nil
}

// CompleteMultipartUpload completes a multipart upload.
func (uc *FileUseCase) CompleteMultipartUpload(ctx context.Context, uploadID string, parts []*types.PartInfo) (*types.FileMetadata, error) {
	uc.log.Infof("CompleteMultipartUpload: uploadID=%s, partsCount=%d", uploadID, len(parts))

	// Convert to proto PartInfo
	protoParts := make([]*types.PartInfo, len(parts))
	for i, p := range parts {
		protoParts[i] = &types.PartInfo{
			PartNumber: p.PartNumber,
			Etag:       p.Etag,
		}
	}

	req := &objclient.CompleteMultipartUploadRequest{
		UploadId: uploadID,
		Parts:    protoParts,
	}
	resp, err := uc.obj.CompleteMultipartUpload(ctx, req)
	if err != nil {
		uc.log.Errorf("failed to complete multipart upload: %v", err)
		return nil, err
	}

	// Update file metadata with the actual object ID
	// Find the file by ObjectId (which contains uploadID temporarily)
	files, _, err := uc.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var fileID int64
	for _, f := range files {
		if f.ObjectId == uploadID {
			fileID = f.Id
			break
		}
	}

	if fileID == 0 {
		return nil, fmt.Errorf("file not found with uploadID: %s", uploadID)
	}

	// Update file with actual object ID
	fileMeta := &types.FileMetadata{
		Id:       fileID,
		ObjectId: resp.Object.Id, // Update with actual object ID from response
	}

	updated, err := uc.repo.Update(ctx, fileMeta)
	if err != nil {
		uc.log.Errorf("failed to update file metadata: %v", err)
		return nil, err
	}

	return updated, nil
}

// AbortMultipartUpload aborts a multipart upload.
func (uc *FileUseCase) AbortMultipartUpload(ctx context.Context, uploadID string) error {
	uc.log.Infof("AbortMultipartUpload: uploadID=%s", uploadID)

	req := &objclient.AbortMultipartUploadRequest{
		UploadId: uploadID,
	}
	_, err := uc.obj.AbortMultipartUpload(ctx, req)
	if err != nil {
		uc.log.Errorf("failed to abort multipart upload: %v", err)
		return err
	}

	return nil
}
