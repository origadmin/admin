/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/google/wire"

	pb "origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/internal/features/filemanager/biz"
	"origadmin/application/admin/internal/helpers/contextutil"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewFileManagerService)

// FileManagerService implements the FileManagerService gRPC service.
type FileManagerService struct {
	pb.UnimplementedFileManagerServiceServer
	uc *biz.FileUseCase
}

// NewFileManagerService creates a new FileManagerService.
func NewFileManagerService(uc *biz.FileUseCase) *FileManagerService {
	return &FileManagerService{
		uc: uc,
	}
}

// UploadFile uploads a small file.
func (s *FileManagerService) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	ownerID, err := contextutil.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	fileMeta, err := s.uc.UploadFile(ctx, req.Name, req.Data, ownerID)
	if err != nil {
		return nil, err
	}
	return &pb.UploadFileResponse{
		FileMetadata: fileMeta,
	}, nil
}

// GetFile gets file metadata.
func (s *FileManagerService) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	fileMeta, err := s.uc.GetFile(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetFileResponse{
		FileMetadata: fileMeta,
	}, nil
}

// DeleteFile deletes a file.
func (s *FileManagerService) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	// TODO: Implement delete logic in usecase
	return &pb.DeleteFileResponse{}, nil
}

// InitiateMultipartUpload initiates a multipart upload.
func (s *FileManagerService) InitiateMultipartUpload(ctx context.Context, req *pb.InitiateMultipartUploadRequest) (*pb.InitiateMultipartUploadResponse, error) {
	// TODO: Implement multipart upload
	return nil, nil
}

// GetMultipartUploadUrl gets a presigned URL for a part.
func (s *FileManagerService) GetMultipartUploadUrl(ctx context.Context, req *pb.GetMultipartUploadUrlRequest) (*pb.GetMultipartUploadUrlResponse, error) {
	// TODO: Implement multipart upload
	return nil, nil
}

// CompleteMultipartUpload completes a multipart upload.
func (s *FileManagerService) CompleteMultipartUpload(ctx context.Context, req *pb.CompleteMultipartUploadRequest) (*pb.CompleteMultipartUploadResponse, error) {
	// TODO: Implement multipart upload
	return nil, nil
}

// AbortMultipartUpload aborts a multipart upload.
func (s *FileManagerService) AbortMultipartUpload(ctx context.Context, req *pb.AbortMultipartUploadRequest) (*pb.AbortMultipartUploadResponse, error) {
	// TODO: Implement multipart upload
	return nil, nil
}
