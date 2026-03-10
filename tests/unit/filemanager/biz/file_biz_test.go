/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz_test

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	objclient "origadmin/application/admin/api/v1/services/objectstore"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/filemanager/biz"
	"origadmin/application/admin/internal/features/filemanager/dto"
)

// MockFileRepo is a mock implementation of FileRepo.
type MockFileRepo struct {
	mock.Mock
}

func (m *MockFileRepo) Create(ctx context.Context, file *types.FileMetadata, opts ...*dto.FileCreateOption) (*types.FileMetadata, error) {
	args := m.Called(ctx, file)
	return args.Get(0).(*types.FileMetadata), args.Error(1)
}

func (m *MockFileRepo) Get(ctx context.Context, id int64, opts ...*dto.FileQueryOption) (*types.FileMetadata, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*types.FileMetadata), args.Error(1)
}

func (m *MockFileRepo) List(ctx context.Context, opts ...*dto.FileQueryOption) ([]*types.FileMetadata, int, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).([]*types.FileMetadata), args.Int(1), args.Error(2)
}

func (m *MockFileRepo) Update(ctx context.Context, file *types.FileMetadata, opts ...*dto.FileUpdateOption) (*types.FileMetadata, error) {
	args := m.Called(ctx, file)
	return args.Get(0).(*types.FileMetadata), args.Error(1)
}

func (m *MockFileRepo) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockObjectStoreClient is a mock implementation of ObjectStoreServiceClient.
type MockObjectStoreClient struct {
	mock.Mock
	objclient.ObjectStoreServiceClient
}

func (m *MockObjectStoreClient) UploadObject(ctx context.Context, in *objclient.UploadObjectRequest, opts ...grpc.CallOption) (*objclient.UploadObjectResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*objclient.UploadObjectResponse), args.Error(1)
}

func TestFileUseCase_UploadFile(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockFileRepo)
	mockObj := new(MockObjectStoreClient)
	logger := log.DefaultLogger

	uc := biz.NewFileUseCase(mockRepo, mockObj, logger)

	// Test data
	fileName := "test.txt"
	fileData := []byte("test content")
	objID := "obj-123"
	fileID := int64(123)
	ownerID := int64(1001)

	// Mock ObjectStore response
	mockObj.On("UploadObject", ctx, mock.Anything, mock.Anything).Return(&objclient.UploadObjectResponse{
		Object: &types.Object{Id: objID},
	}, nil)

	// Mock FileRepo response
	mockRepo.On("Create", ctx, mock.Anything).Return(&types.FileMetadata{
		Id:       fileID,
		Name:     fileName,
		ObjectId: objID,
		OwnerId:  ownerID,
	}, nil)

	// Execute
	fileMeta, err := uc.UploadFile(ctx, fileName, fileData, ownerID, "text/plain", "private")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, fileMeta)
	assert.Equal(t, fileID, fileMeta.Id)
	assert.Equal(t, fileName, fileMeta.Name)
	assert.Equal(t, objID, fileMeta.ObjectId)
	assert.Equal(t, ownerID, fileMeta.OwnerId)

	mockObj.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
