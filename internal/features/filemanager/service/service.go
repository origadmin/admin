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

func (s *FileManagerService) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	ownerID, err := contextutil.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	fileMeta, err := s.uc.UploadFile(ctx, req.Name, req.Data, ownerID, req.ContentType, req.Visibility)
	if err != nil {
		return nil, err
	}
	return &pb.UploadFileResponse{FileMetadata: fileMeta}, nil
}

func (s *FileManagerService) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	fileMeta, downloadURL, err := s.uc.GetFile(ctx, req.Id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.NotFound("FILE_NOT_FOUND", "file not found")
		}
		return nil, err
	}
	return &pb.GetFileResponse{FileMetadata: fileMeta, DownloadUrl: downloadURL}, nil
}

func (s *FileManagerService) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	s.log.Infof("ListFiles called: page=%d, pageSize=%d", req.Page, req.PageSize)
	files, total, err := s.uc.ListFiles(ctx, req.Page, req.PageSize, req.OwnerId, req.Visibility)
	if err != nil {
		return nil, err
	}
	// FIX: Explicitly return page and page_size
	return &pb.ListFilesResponse{
		Files:    files,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *FileManagerService) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UpdateFileResponse, error) {
	fileMeta := &types.FileMetadata{Id: req.Id, Name: req.Name, Visibility: req.Visibility}
	updated, err := s.uc.UpdateFile(ctx, fileMeta)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFileResponse{FileMetadata: updated}, nil
}

func (s *FileManagerService) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	err := s.uc.DeleteFile(ctx, req.Id)
	return &pb.DeleteFileResponse{}, err
}

func (s *FileManagerService) InitiateMultipartUpload(ctx context.Context, req *pb.InitiateMultipartUploadRequest) (*pb.InitiateMultipartUploadResponse, error) {
	ownerID, err := contextutil.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	uploadID, err := s.uc.InitiateMultipartUpload(ctx, req.Name, req.ContentType, req.Visibility, req.Size, ownerID)
	if err != nil {
		return nil, err
	}
	return &pb.InitiateMultipartUploadResponse{UploadId: uploadID}, nil
}

func (s *FileManagerService) GetMultipartUploadUrl(ctx context.Context, req *pb.GetMultipartUploadUrlRequest) (*pb.GetMultipartUploadUrlResponse, error) {
	url, err := s.uc.GetMultipartUploadUrl(ctx, req.UploadId, req.PartNumber)
	if err != nil {
		return nil, err
	}
	return &pb.GetMultipartUploadUrlResponse{UploadUrl: url}, nil
}

func (s *FileManagerService) CompleteMultipartUpload(ctx context.Context, req *pb.CompleteMultipartUploadRequest) (*pb.CompleteMultipartUploadResponse, error) {
	ownerID, err := contextutil.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	parts := make([]*types.PartInfo, len(req.Parts))
	for i, p := range req.Parts {
		parts[i] = &types.PartInfo{PartNumber: p.PartNumber, Etag: p.Etag}
	}
	fileMeta, err := s.uc.CompleteMultipartUpload(ctx, req.UploadId, parts, ownerID)
	if err != nil {
		return nil, err
	}
	return &pb.CompleteMultipartUploadResponse{FileMetadata: fileMeta}, nil
}

func (s *FileManagerService) AbortMultipartUpload(ctx context.Context, req *pb.AbortMultipartUploadRequest) (*pb.AbortMultipartUploadResponse, error) {
	err := s.uc.AbortMultipartUpload(ctx, req.UploadId)
	return &pb.AbortMultipartUploadResponse{}, err
}
