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
