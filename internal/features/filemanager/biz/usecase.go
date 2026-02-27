/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc/metadata"

	objclient "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/filemanager/dto"
	"origadmin/application/admin/internal/helpers/repo"
)

// ProviderSet is biz providers for filemanager.
var ProviderSet = wire.NewSet(NewFileUseCase)

type FileUseCase struct {
	repo dto.FileRepo
	obj  objclient.ObjectStoreServiceClient
	log  *log.Helper
}

func NewFileUseCase(repo dto.FileRepo, obj objclient.ObjectStoreServiceClient, logger log.Logger) *FileUseCase {
	return &FileUseCase{
		repo: repo,
		obj:  obj,
		log:  log.NewHelper(log.With(logger, "module", "filemanager.biz")),
	}
}

func (uc *FileUseCase) UploadFile(ctx context.Context, name string, data []byte, ownerID int64, mimeType, visibility string) (*types.FileMetadata, error) {
	uploadReq := &objclient.UploadObjectRequest{Name: name, Data: data}
	objResp, err := uc.obj.UploadObject(ctx, uploadReq)
	if err != nil {
		return nil, err
	}

	serverHash := ""
	if u, err := url.Parse(objResp.Object.Url); err == nil {
		serverHash = u.Query().Get("sha256")
	}

	fileMeta := &types.FileMetadata{
		Name: name, ObjectId: objResp.Object.Id, OwnerId: ownerID, Size: objResp.Object.Size, MimeType: mimeType, Visibility: visibility, Sha256: serverHash,
	}
	return uc.repo.Create(ctx, fileMeta)
}

func (uc *FileUseCase) InitiateMultipartUpload(ctx context.Context, name, contentType, visibility string, size int64, ownerID int64) (string, error) {
	req := &objclient.InitiateMultipartUploadRequest{Name: name, ContentType: contentType}
	resp, err := uc.obj.InitiateMultipartUpload(ctx, req)
	if err != nil {
		return "", err
	}

	fileMeta := &types.FileMetadata{
		Name: name, ObjectId: resp.UploadId, OwnerId: ownerID, Size: size, Visibility: visibility, MimeType: contentType,
	}
	_, err = uc.repo.Create(ctx, fileMeta)
	if err != nil {
		return "", err
	}
	return resp.UploadId, nil
}

func (uc *FileUseCase) CompleteMultipartUpload(ctx context.Context, uploadID string, parts []*types.PartInfo, ownerID int64) (*types.FileMetadata, error) {
	// 1. Recover Metadata
	files, _, _ := uc.repo.List(ctx)
	var foundFile *types.FileMetadata
	for _, f := range files {
		if f.ObjectId == uploadID {
			foundFile = f
			break
		}
	}
	if foundFile == nil {
		return nil, fmt.Errorf("metadata not found for session %s", uploadID)
	}

	// 2. Physical Merge
	protoParts := make([]*types.PartInfo, len(parts))
	for i, p := range parts {
		protoParts[i] = &types.PartInfo{PartNumber: p.PartNumber, Etag: p.Etag}
	}
	resp, err := uc.obj.CompleteMultipartUpload(ctx, &objclient.CompleteMultipartUploadRequest{UploadId: uploadID, ObjectId: uploadID, Parts: protoParts})
	if err != nil {
		return nil, err
	}

	// 3. INTEGRITY AUDIT
	serverHash := ""
	if u, err := url.Parse(resp.Object.Url); err == nil {
		serverHash = u.Query().Get("sha256")
	}

	clientHash := ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if val := md.Get("x-file-sha256"); len(val) > 0 {
			clientHash = val[0]
		}
	}

	fmt.Printf("\n[VERIFICATION] File: %s\n", foundFile.Name)
	fmt.Printf("[VERIFICATION] Client SHA-256: %s\n", clientHash)
	fmt.Printf("[VERIFICATION] Server SHA-256: %s\n", serverHash)

	// STRICT LOGIC: If we expect a hash but get empty or wrong, FAIL.
	if clientHash == "" {
		fmt.Printf("[VERIFICATION] RESULT: FAILED (Client hash is missing!)\n\n")
		return nil, fmt.Errorf("integrity check failed: client hash missing")
	}

	if serverHash != clientHash {
		fmt.Printf("[VERIFICATION] RESULT: FAILED (Hash Mismatch!)\n\n")
		return nil, fmt.Errorf("integrity check failed: mismatch")
	}
	fmt.Printf("[VERIFICATION] RESULT: PASSED\n\n")

	// 4. Update and Persist
	foundFile.ObjectId = resp.Object.Id
	foundFile.Sha256 = serverHash
	return uc.repo.Update(ctx, foundFile)
}

func (uc *FileUseCase) GetFile(ctx context.Context, id int64) (*types.FileMetadata, string, error) {
	file, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, "", err
	}
	downloadURL := fmt.Sprintf("/api/v1/obs/objects/%s/download", file.ObjectId)
	return file, downloadURL, nil
}

func (uc *FileUseCase) ListFiles(ctx context.Context, page, pageSize int32, ownerID int64, visibility string) ([]*types.FileMetadata, int32, error) {
	opts := &dto.FileQueryOption{QueryOption: repo.QueryOption{Page: int(page), PageSize: int(pageSize)}, OwnerID: ownerID, Visibility: visibility}
	files, total, err := uc.repo.List(ctx, opts)
	if err != nil {
		return nil, 0, err
	}
	return files, int32(total), nil
}

func (uc *FileUseCase) UpdateFile(ctx context.Context, fileMeta *types.FileMetadata) (*types.FileMetadata, error) {
	return uc.repo.Update(ctx, fileMeta)
}
func (uc *FileUseCase) DeleteFile(ctx context.Context, id int64) error {
	file, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	_, _ = uc.obj.DeleteObject(ctx, &objclient.DeleteObjectRequest{Id: file.ObjectId})
	return uc.repo.Delete(ctx, id)
}
func (uc *FileUseCase) GetMultipartUploadUrl(ctx context.Context, uploadID string, partNumber int32) (string, error) {
	// FIX: Must pass ObjectId to satisfy ObjectStore's route requirements
	resp, err := uc.obj.GetMultipartUploadUrl(ctx, &objclient.GetMultipartUploadUrlRequest{
		UploadId:   uploadID,
		ObjectId:   uploadID, // Fallback to uploadID as objectID for local storage
		PartNumber: partNumber,
	})
	if err != nil {
		return "", err
	}
	return resp.UploadUrl, nil
}
func (uc *FileUseCase) AbortMultipartUpload(ctx context.Context, uploadID string) error {
	_, err := uc.obj.AbortMultipartUpload(ctx, &objclient.AbortMultipartUploadRequest{UploadId: uploadID})
	return err
}
