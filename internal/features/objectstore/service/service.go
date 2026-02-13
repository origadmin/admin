/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"bytes"
	"context"
	"time"

	"github.com/google/wire"

	pb "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/features/objectstore/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewObjectStoreService)

// ObjectStoreService implements the ObjectStoreService gRPC service.
type ObjectStoreService struct {
	pb.UnimplementedObjectStoreServiceServer
	uc *biz.ObjectUseCase
}

// NewObjectStoreService creates a new ObjectStoreService.
func NewObjectStoreService(uc *biz.ObjectUseCase) *ObjectStoreService {
	if uc == nil {
		panic("use case cannot be nil")
	}
	return &ObjectStoreService{
		uc: uc,
	}
}

// UploadObject uploads a new object.
func (s *ObjectStoreService) UploadObject(ctx context.Context, req *pb.UploadObjectRequest) (*pb.UploadObjectResponse, error) {
	// In a real scenario, streaming might be preferred, but for this simple example we use bytes.
	reader := bytes.NewReader(req.Data)
	size := int64(len(req.Data))

	// Directly use the PB object returned by Biz layer
	obj, err := s.uc.UploadObject(ctx, req.Name, reader, size)
	if err != nil {
		return nil, err
	}

	return &pb.UploadObjectResponse{
		Object: obj,
	}, nil
}

// GetObject gets an object metadata.
// Note: This API currently returns metadata. Downloading content would typically be a separate stream API or HTTP download.
func (s *ObjectStoreService) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	// For metadata only, we might need a separate method in StorageBackend or optimize Get.
	// Here we use Get but close the stream immediately.
	rc, obj, err := s.uc.GetObject(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return &pb.GetObjectResponse{
		Object: obj,
	}, nil
}

// DeleteObject deletes an object.
func (s *ObjectStoreService) DeleteObject(ctx context.Context, req *pb.DeleteObjectRequest) (*pb.DeleteObjectResponse, error) {
	if err := s.uc.DeleteObject(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteObjectResponse{}, nil
}

// InitiateMultipartUpload initiates a multipart upload.
func (s *ObjectStoreService) InitiateMultipartUpload(ctx context.Context, req *pb.InitiateMultipartUploadRequest) (*pb.InitiateMultipartUploadResponse, error) {
	uploadID, objectID, err := s.uc.InitiateMultipartUpload(ctx, req.Name, req.ContentType)
	if err != nil {
		return nil, err
	}
	return &pb.InitiateMultipartUploadResponse{
		UploadId: uploadID,
		ObjectId: objectID,
	}, nil
}

// GetMultipartUploadUrl gets a URL for uploading a part.
func (s *ObjectStoreService) GetMultipartUploadUrl(ctx context.Context, req *pb.GetMultipartUploadUrlRequest) (*pb.GetMultipartUploadUrlResponse, error) {
	// Default expiration: 1 hour
	url, err := s.uc.GetMultipartUploadURL(ctx, req.ObjectId, req.UploadId, req.PartNumber, time.Hour)
	if err != nil {
		return nil, err
	}
	return &pb.GetMultipartUploadUrlResponse{
		UploadUrl: url,
	}, nil
}

// ListParts lists the parts that have been uploaded for a specific multipart upload.
func (s *ObjectStoreService) ListParts(ctx context.Context, req *pb.ListPartsRequest) (*pb.ListPartsResponse, error) {
	parts, err := s.uc.ListParts(ctx, req.ObjectId, req.UploadId)
	if err != nil {
		return nil, err
	}
	return &pb.ListPartsResponse{
		Parts: parts,
	}, nil
}

// CompleteMultipartUpload completes a multipart upload.
func (s *ObjectStoreService) CompleteMultipartUpload(ctx context.Context, req *pb.CompleteMultipartUploadRequest) (*pb.CompleteMultipartUploadResponse, error) {
	obj, err := s.uc.CompleteMultipartUpload(ctx, req.ObjectId, req.UploadId, req.Parts)
	if err != nil {
		return nil, err
	}
	return &pb.CompleteMultipartUploadResponse{
		Object: obj,
	}, nil
}

// AbortMultipartUpload aborts a multipart upload.
func (s *ObjectStoreService) AbortMultipartUpload(ctx context.Context, req *pb.AbortMultipartUploadRequest) (*pb.AbortMultipartUploadResponse, error) {
	if err := s.uc.AbortMultipartUpload(ctx, req.ObjectId, req.UploadId); err != nil {
		return nil, err
	}
	return &pb.AbortMultipartUploadResponse{}, nil
}
