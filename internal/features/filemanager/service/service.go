/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc/metadata"

	pb "origadmin/application/admin/api/v1/services/filemanager"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/filemanager/biz"
	"origadmin/application/admin/internal/helpers/contextutil"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewFileManagerService)

// FileManagerService implements the FileManagerService gRPC service.
type FileManagerService struct {
	pb.UnimplementedFileManagerServiceServer
	uc  *biz.FileUseCase
	log *log.Helper
}

// NewFileManagerService creates a new FileManagerService.
func NewFileManagerService(uc *biz.FileUseCase, logger log.Logger) *FileManagerService {
	return &FileManagerService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "filemanager.service")),
	}
}

// UploadFile uploads a small file.
func (s *FileManagerService) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.log.Infof(">>>>>> [METADATA] Incoming Metadata: %+v", md)

	s.log.Infof("UploadFile called: name=%s, size=%d", req.Name, len(req.Data))
	ownerID, err := contextutil.GetUserID(ctx)
	if err != nil {
		s.log.Errorf("UploadFile failed to get user ID: %v", err)
		return nil, err
	}

	fileMeta, err := s.uc.UploadFile(ctx, req.Name, req.Data, ownerID)
	if err != nil {
		s.log.Errorf("UploadFile failed: %v", err)
		return nil, err
	}
	return &pb.UploadFileResponse{
		FileMetadata: fileMeta,
	}, nil
}

// GetFile gets file metadata.
func (s *FileManagerService) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	s.log.Infof("GetFile called: id=%d", req.Id)
	fileMeta, downloadURL, err := s.uc.GetFile(ctx, req.Id)
	if err != nil {
		s.log.Errorf("GetFile failed: %v", err)
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.NotFound("FILE_NOT_FOUND", "file not found")
		}
		return nil, err
	}
	return &pb.GetFileResponse{
		FileMetadata: fileMeta,
		DownloadUrl:  downloadURL,
	}, nil
}

// ListFiles lists files.
func (s *FileManagerService) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	s.log.Infof("ListFiles called: page=%d, pageSize=%d", req.Page, req.PageSize)
	files, total, err := s.uc.ListFiles(ctx, req.Page, req.PageSize, req.OwnerId, req.Visibility)
	if err != nil {
		s.log.Errorf("ListFiles failed: %v", err)
		return nil, err
	}
	return &pb.ListFilesResponse{
		Files:      files,
		TotalCount: total,
	}, nil
}

// DeleteFile deletes a file.
func (s *FileManagerService) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	s.log.Infof("DeleteFile called: id=%d", req.Id)

	err := s.uc.DeleteFile(ctx, req.Id)
	if err != nil {
		s.log.Errorf("DeleteFile failed: %v", err)
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.NotFound("FILE_NOT_FOUND", "file not found")
		}
		return nil, err
	}

	return &pb.DeleteFileResponse{}, nil
}

// InitiateMultipartUpload initiates a multipart upload.
func (s *FileManagerService) InitiateMultipartUpload(ctx context.Context, req *pb.InitiateMultipartUploadRequest) (*pb.InitiateMultipartUploadResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.log.Infof(">>>>>> [METADATA] InitiateMultipartUpload: %+v", md)

	s.log.Infof("InitiateMultipartUpload called: name=%s, size=%d, visibility=%s", req.Name, req.Size, req.Visibility)
	ownerID, err := contextutil.GetUserID(ctx)
	if err != nil {
		s.log.Errorf("InitiateMultipartUpload failed to get user ID: %v", err)
		return nil, err
	}

	uploadID, err := s.uc.InitiateMultipartUpload(ctx, req.Name, req.ContentType, req.Visibility, req.Size, ownerID)
	if err != nil {
		s.log.Errorf("InitiateMultipartUpload failed: %v", err)
		return nil, err
	}

	return &pb.InitiateMultipartUploadResponse{
		UploadId: uploadID,
	}, nil
}

// GetMultipartUploadUrl gets a presigned URL for a part.
func (s *FileManagerService) GetMultipartUploadUrl(ctx context.Context, req *pb.GetMultipartUploadUrlRequest) (*pb.GetMultipartUploadUrlResponse, error) {
	s.log.Infof("GetMultipartUploadUrl called: uploadID=%s, partNumber=%d", req.UploadId, req.PartNumber)

	uploadURL, err := s.uc.GetMultipartUploadUrl(ctx, req.UploadId, req.PartNumber)
	if err != nil {
		s.log.Errorf("GetMultipartUploadUrl failed: %v", err)
		return nil, err
	}

	return &pb.GetMultipartUploadUrlResponse{
		UploadUrl: uploadURL,
	}, nil
}

// CompleteMultipartUpload completes a multipart upload.
func (s *FileManagerService) CompleteMultipartUpload(ctx context.Context, req *pb.CompleteMultipartUploadRequest) (*pb.CompleteMultipartUploadResponse, error) {
	s.log.Infof("CompleteMultipartUpload called: uploadID=%s, partsCount=%d", req.UploadId, len(req.Parts))

	// Convert proto PartInfo to types PartInfo
	parts := make([]*types.PartInfo, len(req.Parts))
	for i, p := range req.Parts {
		parts[i] = &types.PartInfo{
			PartNumber: p.PartNumber,
			Etag:       p.Etag,
		}
	}

	fileMeta, err := s.uc.CompleteMultipartUpload(ctx, req.UploadId, parts)
	if err != nil {
		s.log.Errorf("CompleteMultipartUpload failed: %v", err)
		return nil, err
	}

	return &pb.CompleteMultipartUploadResponse{
		FileMetadata: fileMeta,
	}, nil
}

// AbortMultipartUpload aborts a multipart upload.
func (s *FileManagerService) AbortMultipartUpload(ctx context.Context, req *pb.AbortMultipartUploadRequest) (*pb.AbortMultipartUploadResponse, error) {
	s.log.Infof("AbortMultipartUpload called: uploadID=%s", req.UploadId)

	err := s.uc.AbortMultipartUpload(ctx, req.UploadId)
	if err != nil {
		s.log.Errorf("AbortMultipartUpload failed: %v", err)
		return nil, err
	}

	return &pb.AbortMultipartUploadResponse{}, nil
}
