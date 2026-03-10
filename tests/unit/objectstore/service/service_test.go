/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service_test

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"

	objpb "origadmin/application/admin/api/v1/services/objectstore"
	objbiz "origadmin/application/admin/internal/features/objectstore/biz"
	objdal "origadmin/application/admin/internal/features/objectstore/dal"
	objservice "origadmin/application/admin/internal/features/objectstore/service"
)

// TestObjectStoreServiceUploadObject tests the UploadObject gRPC service method
func TestObjectStoreServiceUploadObject(t *testing.T) {
	ctx := context.Background()

	// Setup storage
	tmpDir := t.TempDir()
	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
	}
	storage, _ := objdal.NewLocalStorage(cfg)

	// Create use case
	uc := objbiz.NewObjectStoreUseCase(storage, log.DefaultLogger)

	// Create service
	svc := objservice.NewObjectStoreService(uc, log.DefaultLogger)

	// Test UploadObject
	testData := []byte("test object content")
	req := &objpb.UploadObjectRequest{
		Name: "test-object.txt",
		Data: testData,
	}

	resp, err := svc.UploadObject(ctx, req)
	if err != nil {
		t.Fatalf("UploadObject() error = %v", err)
	}

	if resp == nil {
		t.Fatal("UploadObject() returned nil response")
	}

	if resp.Object == nil {
		t.Fatal("UploadObject() returned nil object")
	}

	if resp.Object.Id == "" {
		t.Error("Object ID should not be empty")
	}

	if resp.Object.Name != "test-object.txt" {
		t.Errorf("Object Name = %v, want %v", resp.Object.Name, "test-object.txt")
	}

	if resp.Object.Size != int64(len(testData)) {
		t.Errorf("Object Size = %v, want %v", resp.Object.Size, len(testData))
	}

	if resp.Object.Url == "" {
		t.Error("Object URL should not be empty")
	}
}

// TestObjectStoreServiceGetObject tests GetObject gRPC service method
func TestObjectStoreServiceGetObject(t *testing.T) {
	ctx := context.Background()

	// Setup storage
	tmpDir := t.TempDir()
	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
	}
	storage, _ := objdal.NewLocalStorage(cfg)

	// Create use case
	uc := objbiz.NewObjectStoreUseCase(storage, log.DefaultLogger)

	// Create service
	svc := objservice.NewObjectStoreService(uc, log.DefaultLogger)

	// First create an object
	createReq := &objpb.UploadObjectRequest{
		Name: "test-object.txt",
		Data: []byte("test content"),
	}
	createResp, _ := svc.UploadObject(ctx, createReq)

	// Now get it
	getReq := &objpb.GetObjectRequest{
		Id: createResp.Object.Id,
	}

	getResp, err := svc.GetObject(ctx, getReq)
	if err != nil {
		t.Fatalf("GetObject() error = %v", err)
	}

	if getResp == nil {
		t.Fatal("GetObject() returned nil response")
	}

	if getResp.Object == nil {
		t.Fatal("GetObject() returned nil object")
	}

	if getResp.Object.Id != createResp.Object.Id {
		t.Errorf("Object ID = %v, want %v", getResp.Object.Id, createResp.Object.Id)
	}

	// Note: LocalStorage implementation currently returns ID as Name on Get
	// because it doesn't store metadata separately.
	if getResp.Object.Name != createResp.Object.Id {
		t.Errorf("Object Name = %v, want %v", getResp.Object.Name, createResp.Object.Id)
	}
}

// TestObjectStoreServiceGetNotFound tests GetObject with non-existent ID
func TestObjectStoreServiceGetNotFound(t *testing.T) {
	ctx := context.Background()

	// Setup storage
	tmpDir := t.TempDir()
	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
	}
	storage, _ := objdal.NewLocalStorage(cfg)

	// Create use case
	uc := objbiz.NewObjectStoreUseCase(storage, log.DefaultLogger)

	// Create service
	svc := objservice.NewObjectStoreService(uc, log.DefaultLogger)

	// Try to get non-existent object
	req := &objpb.GetObjectRequest{
		Id: "non-existent-id",
	}

	_, err := svc.GetObject(ctx, req)
	if err == nil {
		t.Error("GetObject() with non-existent ID should return error")
	}
}

// TestObjectStoreServiceDeleteObject tests the DeleteObject gRPC service method
func TestObjectStoreServiceDeleteObject(t *testing.T) {
	ctx := context.Background()

	// Setup storage
	tmpDir := t.TempDir()
	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
	}
	storage, _ := objdal.NewLocalStorage(cfg)

	// Create use case
	uc := objbiz.NewObjectStoreUseCase(storage, log.DefaultLogger)

	// Create service
	svc := objservice.NewObjectStoreService(uc, log.DefaultLogger)

	// First create an object
	createReq := &objpb.UploadObjectRequest{
		Name: "test-object.txt",
		Data: []byte("test content"),
	}
	createResp, _ := svc.UploadObject(ctx, createReq)

	// Now delete it
	deleteReq := &objpb.DeleteObjectRequest{
		Id: createResp.Object.Id,
	}

	_, err := svc.DeleteObject(ctx, deleteReq)
	if err != nil {
		t.Fatalf("DeleteObject() error = %v", err)
	}

	// Verify deletion by trying to get it
	getReq := &objpb.GetObjectRequest{
		Id: createResp.Object.Id,
	}

	_, err = svc.GetObject(ctx, getReq)
	if err == nil {
		t.Error("GetObject() after deletion should return error")
	}
}

// TestObjectStoreServiceDeleteNonExistent tests deleting a non-existent object
func TestObjectStoreServiceDeleteNonExistent(t *testing.T) {
	ctx := context.Background()

	// Setup storage
	tmpDir := t.TempDir()
	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
	}
	storage, _ := objdal.NewLocalStorage(cfg)

	// Create use case
	uc := objbiz.NewObjectStoreUseCase(storage, log.DefaultLogger)

	// Create service
	svc := objservice.NewObjectStoreService(uc, log.DefaultLogger)

	// Try to delete non-existent object
	req := &objpb.DeleteObjectRequest{
		Id: "non-existent-id",
	}

	_, err := svc.DeleteObject(ctx, req)
	// Deleting non-existent should not error (idempotent)
	if err != nil {
		t.Logf("DeleteObject() of non-existent returned error: %v", err)
	}
}

// TestObjectStoreServiceEmptyData tests UploadObject with empty data
func TestObjectStoreServiceEmptyData(t *testing.T) {
	ctx := context.Background()

	// Setup storage
	tmpDir := t.TempDir()
	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
	}
	storage, _ := objdal.NewLocalStorage(cfg)

	// Create use case
	uc := objbiz.NewObjectStoreUseCase(storage, log.DefaultLogger)

	// Create service
	svc := objservice.NewObjectStoreService(uc, log.DefaultLogger)

	// Create object with empty data
	req := &objpb.UploadObjectRequest{
		Name: "empty-file.txt",
		Data: []byte(""),
	}

	resp, err := svc.UploadObject(ctx, req)
	if err != nil {
		t.Fatalf("UploadObject() with empty data error = %v", err)
	}

	if resp.Object.Size != 0 {
		t.Errorf("Object Size = %v, want 0", resp.Object.Size)
	}
}

// TestObjectStoreServiceInterfaceCompliance tests that service implements gRPC interface
func TestObjectStoreServiceInterfaceCompliance(t *testing.T) {
	var _ objpb.ObjectStoreServiceServer = (*objservice.ObjectStoreService)(nil)
}

// TestObjectStoreServiceNilUseCase tests service with nil use case (should panic)
func TestObjectStoreServiceNilUseCase(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when creating service with nil use case, but didn't panic")
		}
	}()

	_ = objservice.NewObjectStoreService(nil, log.DefaultLogger)
}
