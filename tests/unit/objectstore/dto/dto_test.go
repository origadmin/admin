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

	"google.golang.org/protobuf/types/known/timestamppb"

	"origadmin/application/admin/api/v1/services/types"
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
	return nil, nil, nil
}

func (s *MockObjectRepo) Delete(ctx context.Context, id string) error {
	return nil
}

// TestObjectRepoInterfaceCompliance tests that MockObjectRepo satisfies the interface
func TestObjectRepoInterfaceCompliance(t *testing.T) {
	var _ dto.ObjectRepo = (*MockObjectRepo)(nil)
}
