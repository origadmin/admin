/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz_test

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	"origadmin/application/admin/api/v1/services/types"
	objbiz "origadmin/application/admin/internal/features/objectstore/biz"
	"origadmin/application/admin/internal/features/objectstore/dto"
)

// MockObjectRepo is a mock implementation of ObjectRepo for testing.
type MockObjectRepo struct{}

func (s *MockObjectRepo) Put(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error) {
	id := "test-id"
	if size > 0 && data != nil {
		_, _ = io.Copy(io.Discard, data)
	}
	return &types.Object{
		Id:          id,
		Name:        name,
		Size:        size,
		ContentType: "application/octet-stream",
		Url:         "http://localhost:8080/objects/" + id,
		CreatedTime: timestamppb.New(time.Now()),
	}, nil
}

func (s *MockObjectRepo) Get(ctx context.Context, id string) (io.ReadCloser, *types.Object, error) {
	if id == "non-existent-id" {
		return nil, nil, fmt.Errorf("object not found: %s", id)
	}
	return io.NopCloser(strings.NewReader("test content")), &types.Object{
		Id:          id,
		Name:        "test-name",
		Size:        12,
		ContentType: "text/plain",
		Url:         "http://localhost:8080/objects/" + id,
		CreatedTime: timestamppb.New(time.Now()),
	}, nil
}

func (s *MockObjectRepo) Delete(ctx context.Context, id string) error {
	if id == "non-existent-id" {
		return nil // Idempotent
	}
	return nil
}

// TestObjectUseCase tests the ObjectUseCase logic
func TestObjectUseCase(t *testing.T) {
	ctx := context.Background()
	repo := &MockObjectRepo{}
	logger := log.DefaultLogger
	uc := objbiz.NewObjectUseCase(repo, logger)

	// Test UploadObject
	testData := "test data content"
	info, err := uc.UploadObject(ctx, "test-file.txt", strings.NewReader(testData), int64(len(testData)))
	if err != nil {
		t.Fatalf("UploadObject() error = %v", err)
	}

	if info == nil {
		t.Fatal("UploadObject() returned nil info")
	}

	if info.Id == "" {
		t.Error("Object.Id should not be empty")
	}

	if info.Name != "test-file.txt" {
		t.Errorf("Object.Name = %v, want %v", info.Name, "test-file.txt")
	}

	// Test GetObject
	rc, gotInfo, err := uc.GetObject(ctx, info.Id)
	if err != nil {
		t.Fatalf("GetObject() error = %v", err)
	}

	if rc == nil {
		t.Fatal("GetObject() returned nil ReadCloser")
	}
	defer rc.Close()

	if gotInfo == nil {
		t.Fatal("GetObject() returned nil info")
	}

	if gotInfo.Id != info.Id {
		t.Errorf("GetObject() Object.Id = %v, want %v", gotInfo.Id, info.Id)
	}

	// Test DeleteObject
	err = uc.DeleteObject(ctx, info.Id)
	if err != nil {
		t.Fatalf("DeleteObject() error = %v", err)
	}
}

// TestObjectUseCaseInterfaceCompliance tests that MockObjectRepo satisfies the interface
func TestObjectUseCaseInterfaceCompliance(t *testing.T) {
	var _ dto.ObjectRepo = (*MockObjectRepo)(nil)
}
