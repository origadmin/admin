/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	httptransport "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"google.golang.org/genproto/googleapis/api/httpbody"

	pb "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/internal/features/objectstore/biz"
)

// ... (ObjectStoreService struct definition)

func (s *ObjectStoreService) RegisterHandlers(srv *httptransport.Server) {
	// FIX: Use '/obs/multipart' to align with physical directory separation
	srv.Route("/").PUT("/obs/multipart/{object_id}/uploads/{upload_id}/parts/{part_number}", s.handleUploadPart)
}

func (s *ObjectStoreService) handleUploadPart(ctx httptransport.Context) error {
	oID := ctx.Vars().Get("object_id")
	uID := ctx.Vars().Get("upload_id")
	pNumStr := ctx.Vars().Get("part_number")
	pNum, _ := strconv.Atoi(pNumStr)

	s.log.Infof("[ObjectStore] handleUploadPart: oID=%s, uID=%s, pNum=%d, Length=%d", oID, uID, pNum, ctx.Request().ContentLength)

	// Stream the body directly to the usecase (which calls LocalStorage.UploadPart)
	etag, err := s.uc.UploadPart(ctx, oID, uID, int32(pNum), ctx.Request().Body)
	if err != nil {
		s.log.Errorf("[ObjectStore] UploadPart execution failed: %v", err)
		return err
	}

	// Set ETag header required for successful multipart completion
	if etag != "" {
		ctx.Response().Header().Set("ETag", etag)
	}

	return ctx.Result(http.StatusOK, nil)
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewObjectStoreService)

// ObjectStoreService implements the ObjectStoreService gRPC service.
type ObjectStoreService struct {
	pb.UnimplementedObjectStoreServiceServer
	uc  *biz.ObjectStoreUseCase
	log *log.Helper
}

func NewObjectStoreService(uc *biz.ObjectStoreUseCase, logger log.Logger) *ObjectStoreService {
	return &ObjectStoreService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "objectstore.service")),
	}
}

func (s *ObjectStoreService) UploadObject(ctx context.Context, req *pb.UploadObjectRequest) (*pb.UploadObjectResponse, error) {
	s.log.Infof("UploadObject: name=%s, size=%d", req.Name, len(req.Data))
	obj, err := s.uc.UploadObject(ctx, req.Name, bytes.NewReader(req.Data), int64(len(req.Data)))
	if err != nil {
		return nil, err
	}
	return &pb.UploadObjectResponse{Object: obj}, nil
}

func (s *ObjectStoreService) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	obj, err := s.uc.GetObject(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetObjectResponse{Object: obj}, nil
}

func (s *ObjectStoreService) ListObjects(ctx context.Context, req *pb.ListObjectsRequest) (*pb.ListObjectsResponse, error) {
	objs, total, err := s.uc.ListObjects(ctx, req.Prefix, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &pb.ListObjectsResponse{Objects: objs, Total: total}, nil
}

func (s *ObjectStoreService) DownloadObject(ctx context.Context, req *pb.DownloadObjectRequest) (*httpbody.HttpBody, error) {
	rc, obj, err := s.uc.DownloadObject(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	data, _ := io.ReadAll(rc)

	if tr, ok := transport.FromServerContext(ctx); ok {
		tr.ReplyHeader().Set("Content-Disposition", "attachment; filename="+obj.Name)
	}

	return &httpbody.HttpBody{
		ContentType: obj.ContentType,
		Data:        data,
	}, nil
}

func (s *ObjectStoreService) DeleteObject(ctx context.Context, req *pb.DeleteObjectRequest) (*pb.DeleteObjectResponse, error) {
	err := s.uc.DeleteObject(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteObjectResponse{}, nil
}

func (s *ObjectStoreService) InitiateMultipartUpload(ctx context.Context, req *pb.InitiateMultipartUploadRequest) (*pb.InitiateMultipartUploadResponse, error) {
	uID, oID, err := s.uc.InitiateMultipartUpload(ctx, req.Name, req.ContentType)
	if err != nil {
		return nil, err
	}
	return &pb.InitiateMultipartUploadResponse{UploadId: uID, ObjectId: oID}, nil
}

func (s *ObjectStoreService) GetMultipartUploadUrl(ctx context.Context, req *pb.GetMultipartUploadUrlRequest) (*pb.GetMultipartUploadUrlResponse, error) {
	url, err := s.uc.GetMultipartUploadUrl(ctx, req.ObjectId, req.UploadId, req.PartNumber, time.Hour)
	if err != nil {
		return nil, err
	}
	return &pb.GetMultipartUploadUrlResponse{UploadUrl: url}, nil
}

func (s *ObjectStoreService) ListParts(ctx context.Context, req *pb.ListPartsRequest) (*pb.ListPartsResponse, error) {
	parts, err := s.uc.ListParts(ctx, req.ObjectId, req.UploadId)
	if err != nil {
		return nil, err
	}
	return &pb.ListPartsResponse{Parts: parts}, nil
}

func (s *ObjectStoreService) CompleteMultipartUpload(ctx context.Context, req *pb.CompleteMultipartUploadRequest) (*pb.CompleteMultipartUploadResponse, error) {
	obj, err := s.uc.CompleteMultipartUpload(ctx, req.ObjectId, req.UploadId, req.Parts)
	if err != nil {
		return nil, err
	}
	return &pb.CompleteMultipartUploadResponse{Object: obj}, nil
}

func (s *ObjectStoreService) AbortMultipartUpload(ctx context.Context, req *pb.AbortMultipartUploadRequest) (*pb.AbortMultipartUploadResponse, error) {
	err := s.uc.AbortMultipartUpload(ctx, req.ObjectId, req.UploadId)
	if err != nil {
		return nil, err
	}
	return &pb.AbortMultipartUploadResponse{}, nil
}
