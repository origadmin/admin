/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal_test

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	objdal "origadmin/application/admin/internal/features/objectstore/dal"
)

// TestLocalStorageConfig tests the LocalStorageConfig creation
func TestLocalStorageConfig(t *testing.T) {
	// Test default config
	cfg := &objdal.LocalStorageConfig{
		BasePath: "./tmp/test-objects",
		BaseURL:  "http://localhost:8080/objects",
	}

	if cfg.BasePath != "./tmp/test-objects" {
		t.Errorf("BasePath = %v, want %v", cfg.BasePath, "./tmp/test-objects")
	}

	if cfg.BaseURL != "http://localhost:8080/objects" {
		t.Errorf("BaseURL = %v, want %v", cfg.BaseURL, "http://localhost:8080/objects")
	}
}

// TestLocalStorageLifecycle tests the full lifecycle of storing, retrieving, and deleting files
func TestLocalStorageLifecycle(t *testing.T) {
	ctx := context.Background()
	tmpDir := filepath.Join(os.TempDir(), "objectstore-test")
	defer os.RemoveAll(tmpDir)

	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
		BaseURL:  "http://localhost:8080/objects",
	}

	storage, err := objdal.NewLocalStorage(cfg)
	if err != nil {
		t.Fatalf("NewLocalStorage() error = %v", err)
	}

	// Test Put
	testData := "This is test content for object storage"
	testName := "test-file.txt"
	reader := strings.NewReader(testData)

	info, err := storage.Put(ctx, testName, reader, int64(len(testData)))
	if err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	if info == nil {
		t.Fatal("Put() returned nil info")
	}

	if info.Id == "" {
		t.Error("Object.Id should not be empty")
	}

	if info.Name != testName {
		t.Errorf("Object.Name = %v, want %v", info.Name, testName)
	}

	if info.Size != int64(len(testData)) {
		t.Errorf("Object.Size = %v, want %v", info.Size, len(testData))
	}

	expectedURL := "http://localhost:8080/objects/" + info.Id
	if info.Url != expectedURL {
		t.Errorf("Object.Url = %v, want %v", info.Url, expectedURL)
	}

	// Test Get
	rc, gotInfo, err := storage.Get(ctx, info.Id)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if rc == nil {
		t.Fatal("Get() returned nil ReadCloser")
	}

	// IMPORTANT: We must read and close the reader BEFORE deleting the file.
	// Windows locks open files, preventing deletion.
	content, err := io.ReadAll(rc)
	rc.Close() // Explicitly close here

	if err != nil {
		t.Fatalf("Failed to read file content: %v", err)
	}

	if gotInfo == nil {
		t.Fatal("Get() returned nil info")
	}

	if gotInfo.Id != info.Id {
		t.Errorf("Get() Object.Id = %v, want %v", gotInfo.Id, info.Id)
	}

	// Note: LocalStorage implementation currently returns ID as Name on Get
	// because it doesn't store metadata separately.
	if gotInfo.Name != info.Id {
		t.Errorf("Get() Object.Name = %v, want %v", gotInfo.Name, info.Id)
	}

	// Verify file exists on disk
	filePath := filepath.Join(tmpDir, info.Id)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("File was not created on disk: %v", filePath)
	}

	if string(content) != testData {
		t.Errorf("File content = %v, want %v", string(content), testData)
	}

	// Test Delete
	err = storage.Delete(ctx, info.Id)
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	// Verify file is deleted from disk
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Errorf("File was not deleted from disk: %v", filePath)
	}

	// Test Get after deletion
	_, _, err = storage.Get(ctx, info.Id)
	if err == nil {
		t.Error("Get() should return error after deletion, but got nil")
	}
}

// TestLocalStorageGetNotFound tests Get() with a non-existent ID
func TestLocalStorageGetNotFound(t *testing.T) {
	ctx := context.Background()
	tmpDir := filepath.Join(os.TempDir(), "objectstore-test-notfound")
	defer os.RemoveAll(tmpDir)

	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
		BaseURL:  "http://localhost:8080/objects",
	}

	storage, _ := objdal.NewLocalStorage(cfg)

	rc, info, err := storage.Get(ctx, "non-existent-id")
	if err == nil {
		t.Error("Get() with non-existent ID should return error")
	}

	if rc != nil {
		rc.Close()
		t.Error("Get() with non-existent ID should return nil ReadCloser")
	}

	if info != nil {
		t.Error("Get() with non-existent ID should return nil info")
	}
}

// TestLocalStorageDirectoryCreation tests that storage directory is created
func TestLocalStorageDirectoryCreation(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "objectstore-test-dir-creation")
	defer os.RemoveAll(tmpDir)

	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
		BaseURL:  "http://localhost:8080/objects",
	}

	_, err := objdal.NewLocalStorage(cfg)
	if err != nil {
		t.Fatalf("NewLocalStorage() error = %v", err)
	}

	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		t.Errorf("Storage directory was not created: %v", tmpDir)
	}
}

// TestLocalStorageLargeFile tests storing a larger file
func TestLocalStorageLargeFile(t *testing.T) {
	ctx := context.Background()
	tmpDir := filepath.Join(os.TempDir(), "objectstore-test-large")
	defer os.RemoveAll(tmpDir)

	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
		BaseURL:  "http://localhost:8080/objects",
	}

	storage, _ := objdal.NewLocalStorage(cfg)

	// Create 1MB test data
	testData := strings.Repeat("A", 1024*1024)
	reader := strings.NewReader(testData)

	info, err := storage.Put(ctx, "large-file.bin", reader, int64(len(testData)))
	if err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	if info.Size != int64(len(testData)) {
		t.Errorf("Object.Size = %v, want %v", info.Size, len(testData))
	}
}

// TestLocalStorageDeleteNonExistent tests deleting a non-existent file
func TestLocalStorageDeleteNonExistent(t *testing.T) {
	ctx := context.Background()
	tmpDir := filepath.Join(os.TempDir(), "objectstore-test-del")
	defer os.RemoveAll(tmpDir)

	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
		BaseURL:  "http://localhost:8080/objects",
	}

	storage, _ := objdal.NewLocalStorage(cfg)

	// Deleting a non-existent file should not return an error
	err := storage.Delete(ctx, "non-existent-id")
	if err != nil {
		t.Errorf("Delete() of non-existent file should not return error, got: %v", err)
	}
}

// BenchmarkLocalStoragePut benchmarks the Put operation
func BenchmarkLocalStoragePut(b *testing.B) {
	ctx := context.Background()
	tmpDir := filepath.Join(os.TempDir(), "objectstore-bench-put")
	defer os.RemoveAll(tmpDir)

	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
		BaseURL:  "http://localhost:8080/objects",
	}

	storage, _ := objdal.NewLocalStorage(cfg)
	testData := strings.Repeat("test data ", 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = storage.Put(ctx, "bench-file.txt", strings.NewReader(testData), int64(len(testData)))
	}
}

// BenchmarkLocalStorageGet benchmarks the Get operation
func BenchmarkLocalStorageGet(b *testing.B) {
	ctx := context.Background()
	tmpDir := filepath.Join(os.TempDir(), "objectstore-bench-get")
	defer os.RemoveAll(tmpDir)

	cfg := &objdal.LocalStorageConfig{
		BasePath: tmpDir,
		BaseURL:  "http://localhost:8080/objects",
	}

	storage, _ := objdal.NewLocalStorage(cfg)
	testData := strings.Repeat("test data ", 1000)
	info, _ := storage.Put(ctx, "bench-file.txt", strings.NewReader(testData), int64(len(testData)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rc, _, _ := storage.Get(ctx, info.ID)
		if rc != nil {
			io.Copy(io.Discard, rc)
			rc.Close()
		}
	}
}
