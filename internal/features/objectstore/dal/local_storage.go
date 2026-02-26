/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/wire"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/origadmin/runtime/extensions/configutil"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/conf"
	"origadmin/application/admin/internal/features/objectstore/dto"
	"origadmin/application/admin/internal/helpers/idutil"
)

type uploadSession struct {
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
	ObjectID    string `json:"object_id"`
	Visibility  string `json:"visibility"`
}

type LocalStorageConfig struct {
	BasePath, MultipartPath string
	ExpirationAge           time.Duration
}

var ProviderSet = wire.NewSet(NewLocalStorage, NewLocalStorageConfig, wire.Bind(new(dto.ObjectRepo), new(*LocalStorage)))

func NewLocalStorageConfig(c *conf.Config) (*LocalStorageConfig, error) {
	o, m := filepath.Join("tmp", "objects"), filepath.Join("tmp", "multipart")
	objectStores := c.GetBootstrap().GetData().GetObjectStores()
	if objectStores != nil {
		store, _, err := configutil.Normalize(objectStores.GetActive(), objectStores.GetDefault(), objectStores.GetConfigs())
		if err == nil && store != nil && store.GetLocal() != nil {
			if root := store.GetLocal().GetRoot(); root != "" {
				o = root
			}
		}
	}
	if envObj := os.Getenv("OBJECTSTORE_OBJECTS_ROOT"); envObj != "" {
		o = envObj
	}
	if envMulti := os.Getenv("OBJECTSTORE_MULTIPART_ROOT"); envMulti != "" {
		m = envMulti
	}
	absO, _ := filepath.Abs(o)
	absM, _ := filepath.Abs(m)
	return &LocalStorageConfig{BasePath: absO, MultipartPath: absM, ExpirationAge: 24 * time.Hour}, nil
}

type LocalStorage struct {
	basePath, multipartPath string
	expiration              time.Duration
}

func NewLocalStorage(cfg *LocalStorageConfig) (*LocalStorage, error) {
	_ = os.MkdirAll(cfg.BasePath, 0755)
	_ = os.MkdirAll(cfg.MultipartPath, 0755)
	s := &LocalStorage{basePath: cfg.BasePath, multipartPath: cfg.MultipartPath, expiration: cfg.ExpirationAge}
	go s.runInternalCleanup(context.Background())
	return s, nil
}

func (s *LocalStorage) Cleanup(ctx context.Context) error {
	entries, _ := os.ReadDir(s.multipartPath)
	now := time.Now()
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		p := filepath.Join(s.multipartPath, entry.Name())
		if info, err := entry.Info(); err == nil {
			_, errAbort := os.Stat(filepath.Join(p, ".aborted"))
			if errAbort == nil || now.Sub(info.ModTime()) > s.expiration {
				_ = os.RemoveAll(p)
			}
		}
	}
	return nil
}

func (s *LocalStorage) runInternalCleanup(ctx context.Context) {
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = s.Cleanup(ctx)
		}
	}
}

func (s *LocalStorage) Put(ctx context.Context, name string, data io.Reader, size int64) (*types.Object, error) {
	id := idutil.GenUUID()
	file, err := os.Create(filepath.Join(s.basePath, id))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	h := sha256.New()
	mw := io.MultiWriter(file, h)
	n, _ := io.Copy(mw, data)
	return &types.Object{Id: id, Name: name, Size: n, Url: fmt.Sprintf("/objects/%s?sha256=%s", id, hex.EncodeToString(h.Sum(nil))), CreateTime: timestamppb.New(time.Now())}, nil
}

func (s *LocalStorage) Get(ctx context.Context, id string) (io.ReadCloser, *types.Object, error) {
	file, err := os.Open(filepath.Join(s.basePath, id))
	if err != nil {
		return nil, nil, err
	}
	stat, _ := file.Stat()
	return file, &types.Object{
		Id:          id,
		Size:        stat.Size(),
		ContentType: "application/octet-stream", // Default fallback
	}, nil
}

func (s *LocalStorage) Delete(ctx context.Context, id string) error {
	return os.Remove(filepath.Join(s.basePath, id))
}

func (s *LocalStorage) GetPresignedURL(ctx context.Context, id string, expires time.Duration) (string, error) {
	return fmt.Sprintf("/objects/%s", id), nil
}

func (s *LocalStorage) InitiateMultipartUpload(ctx context.Context, name string, contentType string) (string, string, error) {
	uID, oID := idutil.GenUUID(), idutil.GenUUID()
	vis := "private"
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if val := md.Get("x-file-visibility"); len(val) > 0 {
			vis = val[0]
		}
	}
	p := filepath.Join(s.multipartPath, uID)
	_ = os.MkdirAll(p, 0755)
	info, _ := json.Marshal(uploadSession{Name: name, ContentType: contentType, ObjectID: oID, Visibility: vis})
	_ = os.WriteFile(filepath.Join(p, "session.json"), info, 0644)
	return uID, uID, nil
}

func (s *LocalStorage) GetMultipartUploadURL(ctx context.Context, objectID string, uploadID string, partNumber int32, expires time.Duration) (string, error) {
	return fmt.Sprintf("/multipart/%s/%d", uploadID, partNumber), nil
}

func (s *LocalStorage) ListParts(ctx context.Context, objectID string, uploadID string) ([]*types.PartInfo, error) {
	p := filepath.Join(s.multipartPath, uploadID)
	entries, err := os.ReadDir(p)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var parts []*types.PartInfo
	for _, entry := range entries {
		if !entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") && entry.Name() != "session.json" {
			partNum, _ := strconv.Atoi(entry.Name())
			content, _ := os.ReadFile(filepath.Join(p, entry.Name()))
			h := md5.Sum(content)
			parts = append(parts, &types.PartInfo{PartNumber: int32(partNum), Etag: hex.EncodeToString(h[:])})
		}
	}
	sort.Slice(parts, func(i, j int) bool { return parts[i].PartNumber < parts[j].PartNumber })
	return parts, nil
}

func (s *LocalStorage) CompleteMultipartUpload(ctx context.Context, objectID string, uploadID string, parts []*types.PartInfo) (*types.Object, error) {
	p := filepath.Join(s.multipartPath, uploadID)
	if _, err := os.Stat(filepath.Join(p, ".aborted")); err == nil {
		return nil, fmt.Errorf("aborted")
	}
	var info uploadSession
	infoD, _ := os.ReadFile(filepath.Join(p, "session.json"))
	_ = json.Unmarshal(infoD, &info)
	fID := objectID
	if fID == "" {
		fID = info.ObjectID
	}
	if fID == "" {
		fID = uploadID
	}
	fPath := filepath.Join(s.basePath, fID)
	fFile, err := os.Create(fPath)
	if err != nil {
		return nil, err
	}
	defer fFile.Close()

	h := sha256.New()
	mw := io.MultiWriter(fFile, h)
	var tSize int64
	for _, part := range parts {
		pPath := filepath.Join(p, fmt.Sprintf("%d", part.PartNumber))
		pContent, err := os.ReadFile(pPath)
		if err != nil {
			_ = fFile.Close()
			_ = os.Remove(fPath)
			return nil, fmt.Errorf("part %d missing", part.PartNumber)
		}
		pMD5 := md5.Sum(pContent)
		if hex.EncodeToString(pMD5[:]) != part.Etag && part.Etag != "" {
			_ = fFile.Close()
			_ = os.Remove(fPath)
			return nil, fmt.Errorf("part %d corrupted", part.PartNumber)
		}
		n, _ := mw.Write(pContent)
		tSize += int64(n)
	}
	_ = os.RemoveAll(p)
	return &types.Object{
		Id: fID, Name: info.Name, Size: tSize, ContentType: info.ContentType,
		Url:        fmt.Sprintf("/objects/%s?visibility=%s&sha256=%s", fID, info.Visibility, hex.EncodeToString(h.Sum(nil))),
		CreateTime: timestamppb.New(time.Now()),
	}, nil
}

func (s *LocalStorage) AbortMultipartUpload(ctx context.Context, objectID string, uploadID string) error {
	p := filepath.Join(s.multipartPath, uploadID)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil
	}
	return os.WriteFile(filepath.Join(p, ".aborted"), []byte(time.Now().Format(time.RFC3339)), 0644)
}
