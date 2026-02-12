/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/timestamppb"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/features/objectstore/dto"
)

// LocalStorageConfig holds the configuration for local file storage.
type LocalStorageConfig struct {
	BasePath string
	BaseURL  string
}

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(
	NewLocalStorage,
	NewLocalStorageConfig,
	wire.Bind(new(dto.ObjectRepo), new(*LocalStorage)),
)

// NewLocalStorageConfig creates a local storage config from the main config.
// For now, returns default config as the full protobuf integration is complex.
func NewLocalStorageConfig(c *conf.Config) (*LocalStorageConfig, error) {
	// Return default config for now
	// TODO: Integrate with full protobuf configuration when needed
	return &LocalStorageConfig{
		BasePath: "./tmp/objects",
		BaseURL:  "http://localhost:8080/objects",
	}, nil
}

// LocalStorage implements ObjectRepo for the local file system.
type LocalStorage struct {
	basePath string
	baseURL  string
}

// NewLocalStorage creates a new LocalStorage instance.
func NewLocalStorage(cfg *LocalStorageConfig) (*LocalStorage, error) {
	if err := os.MkdirAll(cfg.BasePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}
	// Create a directory for multipart uploads
	multipartPath := filepath.Join(cfg.BasePath, "multipart")
	if err := os.MkdirAll(multipartPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create multipart directory: %w", err)
	}

	return &LocalStorage{
		basePath: cfg.BasePath,
		baseURL:  cfg.BaseURL,
	}, nil
}

// Put stores data to the local file system.
func (s *LocalStorage) Put(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error) {
	id := uuid.New().String()
	filePath := filepath.Join(s.basePath, id)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	written, err := io.Copy(file, data)
	if err != nil {
		return nil, fmt.Errorf("failed to write file content: %w", err)
	}

	return &types.Object{
		Id:          id,
		Name:        name,
		Size:        written,
		Url:         fmt.Sprintf("%s/%s", s.baseURL, id),
		CreatedTime: timestamppb.New(time.Now()),
		ContentType: "application/octet-stream", // Default content type
	}, nil
}

// Get retrieves data from the local file system.
func (s *LocalStorage) Get(ctx context.Context, id string) (io.ReadCloser, *types.Object, error) {
	filePath := filepath.Join(s.basePath, id)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, fmt.Errorf("object not found: %s", id)
		}
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, nil, fmt.Errorf("failed to stat file: %w", err)
	}

	info := &types.Object{
		Id:          id,
		Name:        id, // Local storage doesn't store original name, use ID as name
		Size:        stat.Size(),
		Url:         fmt.Sprintf("%s/%s", s.baseURL, id),
		CreatedTime: timestamppb.New(stat.ModTime()),
		ContentType: "application/octet-stream", // Default content type
	}

	return file, info, nil
}

// Delete removes data from the local file system.
func (s *LocalStorage) Delete(ctx context.Context, id string) error {
	filePath := filepath.Join(s.basePath, id)
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

// GetPresignedURL generates a presigned URL for downloading the object.
func (s *LocalStorage) GetPresignedURL(ctx context.Context, id string, expires time.Duration) (string, error) {
	// For local storage, we just return the direct URL.
	// In a real implementation, this would generate a signed URL with expiration.
	return fmt.Sprintf("%s/%s?expires=%d", s.baseURL, id, time.Now().Add(expires).Unix()), nil
}

// InitiateMultipartUpload initiates a multipart upload.
func (s *LocalStorage) InitiateMultipartUpload(ctx context.Context, name string, contentType string) (string, string, error) {
	uploadID := uuid.New().String()
	objectID := uuid.New().String() // Pre-allocate object ID

	// Create a directory for this upload
	uploadPath := filepath.Join(s.basePath, "multipart", uploadID)
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	return uploadID, objectID, nil
}

// GetMultipartUploadURL generates a presigned URL for uploading a specific part.
func (s *LocalStorage) GetMultipartUploadURL(ctx context.Context, objectID string, uploadID string, partNumber int32, expires time.Duration) (string, error) {
	// For local storage simulation, we return a URL that points to where the part *should* be uploaded.
	// Note: This requires an HTTP handler to actually accept the PUT request and write to this path.
	return fmt.Sprintf("%s/multipart/%s/%d", s.baseURL, uploadID, partNumber), nil
}

// CompleteMultipartUpload completes a multipart upload by assembling the parts.
func (s *LocalStorage) CompleteMultipartUpload(ctx context.Context, objectID string, uploadID string, parts []*types.PartInfo) (*types.Object, error) {
	uploadPath := filepath.Join(s.basePath, "multipart", uploadID)
	finalPath := filepath.Join(s.basePath, objectID)

	// Sort parts by part number to ensure correct order
	sort.Slice(parts, func(i, j int) bool {
		return parts[i].PartNumber < parts[j].PartNumber
	})

	// Create the final file
	finalFile, err := os.Create(finalPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create final file: %w", err)
	}
	defer finalFile.Close()

	var totalSize int64

	// Append each part to the final file
	for _, part := range parts {
		// In a real S3 implementation, we don't read local files.
		// But for this local simulation, we assume the parts have been uploaded to the uploadPath.
		// The filename of the part is assumed to be just the part number.
		partPath := filepath.Join(uploadPath, fmt.Sprintf("%d", part.PartNumber))

		partFile, err := os.Open(partPath)
		if err != nil {
			// If part file is missing, it means the upload failed or wasn't done correctly
			return nil, fmt.Errorf("missing part %d: %w", part.PartNumber, err)
		}

		n, err := io.Copy(finalFile, partFile)
		partFile.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to append part %d: %w", part.PartNumber, err)
		}
		totalSize += n
	}

	// Cleanup multipart directory
	if err := os.RemoveAll(uploadPath); err != nil {
		// Log error but don't fail the operation
		fmt.Printf("failed to cleanup multipart directory: %v\n", err)
	}

	return &types.Object{
		Id:          objectID,
		Name:        objectID, // Use ID as name
		Size:        totalSize,
		Url:         fmt.Sprintf("%s/%s", s.baseURL, objectID),
		CreatedTime: timestamppb.New(time.Now()),
		ContentType: "application/octet-stream",
	}, nil
}

// AbortMultipartUpload aborts a multipart upload and cleans up resources.
func (s *LocalStorage) AbortMultipartUpload(ctx context.Context, objectID string, uploadID string) error {
	uploadPath := filepath.Join(s.basePath, "multipart", uploadID)
	if err := os.RemoveAll(uploadPath); err != nil {
		return fmt.Errorf("failed to remove upload directory: %w", err)
	}
	return nil
}
