/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"bytes"
	"context"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/genproto/googleapis/api/httpbody"

	pb "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/features/objectstore/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewObjectStoreService)

// ObjectStoreService implements the ObjectStoreService gRPC service.
type ObjectStoreService struct {
	pb.UnimplementedObjectStoreServiceServer
	uc  *biz.ObjectUseCase
	log *log.Helper
}

// NewObjectStoreService creates a new ObjectStoreService.
func NewObjectStoreService(uc *biz.ObjectUseCase, logger log.Logger) *ObjectStoreService {
	if uc == nil {
		panic("use case cannot be nil")
	}
	return &ObjectStoreService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "objectstore.service")),
	}
}

// UploadObject uploads a new object.
func (s *ObjectStoreService) UploadObject(ctx context.Context, req *pb.UploadObjectRequest) (*pb.UploadObjectResponse, error) {
	dataLen := len(req.Data)
	s.log.Infof("UploadObject called: name=%s, size=%d", req.Name, dataLen)

	if dataLen > 0 {
		preview := req.Data
		if dataLen > 20 {
			preview = req.Data[:20]
		}
		s.log.Infof("UploadObject data preview: %x", preview)
	} else {
		s.log.Warn("UploadObject received empty data!")
	}

	// In a real scenario, streaming might be preferred, but for this simple example we use bytes.
	reader := bytes.NewReader(req.Data)
	size := int64(dataLen)

	// Directly use the PB object returned by Biz layer
	obj, err := s.uc.UploadObject(ctx, req.Name, reader, size)
	if err != nil {
		s.log.Errorf("UploadObject failed: %v", err)
		return nil, err
	}

	s.log.Infof("UploadObject success: obj.Size=%d", obj.Size)

	return &pb.UploadObjectResponse{
		Object: obj,
	}, nil
}

// GetObject gets an object metadata.
// Note: This API currently returns metadata. Downloading content would typically be a separate stream API or HTTP download.
func (s *ObjectStoreService) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	s.log.Infof("GetObject called: id=%s", req.Id)
	// For metadata only, we might need a separate method in StorageBackend or optimize Get.
	// Here we use Get but close the stream immediately.
	rc, obj, err := s.uc.GetObject(ctx, req.Id)
	if err != nil {
		s.log.Errorf("GetObject failed: %v", err)
		return nil, err
	}
	defer rc.Close()

	return &pb.GetObjectResponse{
		Object: obj,
	}, nil
}

// DownloadObject downloads the object content.
func (s *ObjectStoreService) DownloadObject(ctx context.Context, req *pb.DownloadObjectRequest) (*httpbody.HttpBody, error) {
	s.log.Infof("DownloadObject called: id=%s, range=%s", req.Id, req.Range)
	rc, obj, err := s.uc.GetObject(ctx, req.Id)
	if err != nil {
		s.log.Errorf("DownloadObject failed to get object: %v", err)
		return nil, err
	}
	defer rc.Close()

	var data []byte
	var contentType = obj.ContentType

	// Handle Range Request
	if req.Range != "" {
		// Parse Range header: bytes=start-end
		// Simple parser for demonstration. Production code should be more robust.
		rangeStr := strings.TrimPrefix(req.Range, "bytes=")
		parts := strings.Split(rangeStr, "-")
		if len(parts) == 2 {
			start, err1 := strconv.ParseInt(parts[0], 10, 64)
			end, err2 := strconv.ParseInt(parts[1], 10, 64)

			if err1 == nil && err2 == nil && start <= end {
				// Seek to start
				if seeker, ok := rc.(io.Seeker); ok {
					_, err = seeker.Seek(start, io.SeekStart)
					if err != nil {
						s.log.Errorf("DownloadObject failed to seek: %v", err)
						return nil, err
					}
				} else {
					// Fallback: Read and discard if not seekable (inefficient but works)
					// In a real implementation, the storage backend should support ReadAt or Seek.
					// For now, assuming rc is a file or supports seeking is better, or we just read everything and slice (bad for memory).
					// Given the context of avoiding OOM, we should ideally use a backend that supports range reads.
					// But since we are modifying the service layer, let's try to read only the requested amount.
					// NOTE: io.ReadCloser from local file system IS seekable if it's an *os.File.
					// Let's check if we can cast it.
					// If not, we might have to read until start.
					_, err = io.CopyN(io.Discard, rc, start)
					if err != nil {
						s.log.Errorf("DownloadObject failed to skip to start: %v", err)
						return nil, err
					}
				}

				// Read length
				length := end - start + 1
				// Limit reading to length
				limitReader := io.LimitReader(rc, length)
				data, err = io.ReadAll(limitReader)
				if err != nil {
					s.log.Errorf("DownloadObject failed to read range content: %v", err)
					return nil, err
				}
			}
		}
	} else {
		// No range, read all (Warning: still OOM risk if called without range for large files)
		// But the Gateway is expected to always call with Range for large files.
		data, err = io.ReadAll(rc)
		if err != nil {
			s.log.Errorf("DownloadObject failed to read content: %v", err)
			return nil, err
		}
	}

	return &httpbody.HttpBody{
		ContentType: contentType,
		Data:        data,
	}, nil
}

// DeleteObject deletes an object.
func (s *ObjectStoreService) DeleteObject(ctx context.Context, req *pb.DeleteObjectRequest) (*pb.DeleteObjectResponse, error) {
	s.log.Infof("DeleteObject called: id=%s", req.Id)
	if err := s.uc.DeleteObject(ctx, req.Id); err != nil {
		s.log.Errorf("DeleteObject failed: %v", err)
		return nil, err
	}
	return &pb.DeleteObjectResponse{}, nil
}

// InitiateMultipartUpload initiates a multipart upload.
func (s *ObjectStoreService) InitiateMultipartUpload(ctx context.Context, req *pb.InitiateMultipartUploadRequest) (*pb.InitiateMultipartUploadResponse, error) {
	s.log.Infof("InitiateMultipartUpload called: name=%s", req.Name)
	uploadID, objectID, err := s.uc.InitiateMultipartUpload(ctx, req.Name, req.ContentType)
	if err != nil {
		s.log.Errorf("InitiateMultipartUpload failed: %v", err)
		return nil, err
	}
	return &pb.InitiateMultipartUploadResponse{
		UploadId: uploadID,
		ObjectId: objectID,
	}, nil
}

// GetMultipartUploadUrl gets a URL for uploading a part.
func (s *ObjectStoreService) GetMultipartUploadUrl(ctx context.Context, req *pb.GetMultipartUploadUrlRequest) (*pb.GetMultipartUploadUrlResponse, error) {
	s.log.Infof("GetMultipartUploadUrl called: objectID=%s, uploadID=%s, partNumber=%d", req.ObjectId, req.UploadId, req.PartNumber)
	// Default expiration: 1 hour
	url, err := s.uc.GetMultipartUploadURL(ctx, req.ObjectId, req.UploadId, req.PartNumber, time.Hour)
	if err != nil {
		s.log.Errorf("GetMultipartUploadUrl failed: %v", err)
		return nil, err
	}
	return &pb.GetMultipartUploadUrlResponse{
		UploadUrl: url,
	}, nil
}

// ListParts lists the parts that have been uploaded for a specific multipart upload.
func (s *ObjectStoreService) ListParts(ctx context.Context, req *pb.ListPartsRequest) (*pb.ListPartsResponse, error) {
	s.log.Infof("ListParts called: objectID=%s, uploadID=%s", req.ObjectId, req.UploadId)
	parts, err := s.uc.ListParts(ctx, req.ObjectId, req.UploadId)
	if err != nil {
		s.log.Errorf("ListParts failed: %v", err)
		return nil, err
	}
	return &pb.ListPartsResponse{
		Parts: parts,
	}, nil
}

// CompleteMultipartUpload completes a multipart upload.
func (s *ObjectStoreService) CompleteMultipartUpload(ctx context.Context, req *pb.CompleteMultipartUploadRequest) (*pb.CompleteMultipartUploadResponse, error) {
	s.log.Infof("CompleteMultipartUpload called: objectID=%s, uploadID=%s, partsCount=%d", req.ObjectId, req.UploadId, len(req.Parts))
	obj, err := s.uc.CompleteMultipartUpload(ctx, req.ObjectId, req.UploadId, req.Parts)
	if err != nil {
		s.log.Errorf("CompleteMultipartUpload failed: %v", err)
		return nil, err
	}
	return &pb.CompleteMultipartUploadResponse{
		Object: obj,
	}, nil
}

// AbortMultipartUpload aborts a multipart upload.
func (s *ObjectStoreService) AbortMultipartUpload(ctx context.Context, req *pb.AbortMultipartUploadRequest) (*pb.AbortMultipartUploadResponse, error) {
	s.log.Infof("AbortMultipartUpload called: objectID=%s, uploadID=%s", req.ObjectId, req.UploadId)
	if err := s.uc.AbortMultipartUpload(ctx, req.ObjectId, req.UploadId); err != nil {
		s.log.Errorf("AbortMultipartUpload failed: %v", err)
		return nil, err
	}
	return &pb.AbortMultipartUploadResponse{}, nil
}
