/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto_test

import (
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/features/objectstore/dto"
)

// MockObjectRepo is a mock implementation of the ObjectRepo interface
type MockObjectRepo struct {
	storage map[string][]byte
}

func NewMockObjectRepo() *MockObjectRepo {
	return &MockObjectRepo{
		storage: make(map[string][]byte),
	}
}

func (m *MockObjectRepo) Put(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error) {
	content, _ := io.ReadAll(data)
	id := name // In this simple mock, name is the ID
	m.storage[id] = content
	return &types.Object{
		Id:          id,
		Name:        name,
		Size:        int64(len(content)),
		Url:         "/objects/" + id,
		ContentType: "text/plain",
		CreateTime:  timestamppb.Now(),
	}, nil
}

func (m *MockObjectRepo) Get(ctx context.Context, id string) (io.ReadCloser, *types.Object, error) {
	content, ok := m.storage[id]
	if !ok {
		return nil, nil, io.EOF
	}
	return io.NopCloser(strings.NewReader(string(content))), &types.Object{
		Id:          id,
		Name:        id,
		Size:        int64(len(content)),
		Url:         "/objects/" + id,
		ContentType: "text/plain",
		CreateTime:  timestamppb.Now(),
	}, nil
}

func (m *MockObjectRepo) Delete(ctx context.Context, id string) error {
	delete(m.storage, id)
	return nil
}

func (m *MockObjectRepo) List(ctx context.Context, prefix string, page, pageSize int32) ([]*types.Object, int32, error) {
	return nil, 0, nil
}

func (m *MockObjectRepo) GetPresignedURL(ctx context.Context, id string, expires time.Duration) (string, error) {
	return "", nil
}

func (m *MockObjectRepo) InitiateMultipartUpload(ctx context.Context, name string, contentType string) (string, string, error) {
	return "", "", nil
}

func (m *MockObjectRepo) GetMultipartUploadURL(ctx context.Context, objectID string, uploadID string, partNumber int32, expires time.Duration) (string, error) {
	return "", nil
}

func (m *MockObjectRepo) UploadPart(ctx context.Context, objectID string, uploadID string, partNumber int32, data io.Reader) (string, error) {
	return "", nil
}

func (m *MockObjectRepo) ListParts(ctx context.Context, objectID string, uploadID string) ([]*types.PartInfo, error) {
	return nil, nil
}

func (m *MockObjectRepo) CompleteMultipartUpload(ctx context.Context, objectID string, uploadID string, parts []*types.PartInfo) (*types.Object, error) {
	return nil, nil
}

func (m *MockObjectRepo) AbortMultipartUpload(ctx context.Context, objectID string, uploadID string) error {
	return nil
}

func TestObjectRepo(t *testing.T) {
	var repo dto.ObjectRepo = NewMockObjectRepo()
	ctx := context.Background()

	t.Run("InterfaceCheck", func(t *testing.T) {
		assert.NotNil(t, repo)
		obj, err := repo.Put(ctx, "test.bin", strings.NewReader("data"), 8)
		assert.NoError(t, err)
		assert.NotNil(t, obj)
	})
}
