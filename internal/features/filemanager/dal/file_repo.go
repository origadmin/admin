/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/google/wire"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/file"

	"origadmin/application/admin/internal/features/filemanager/dto"
	"origadmin/application/admin/internal/helpers/repo"
)

// ProviderSet is dal providers.
var ProviderSet = wire.NewSet(NewFileRepo)

// FileRepo implements dto.FileRepo using Ent.
type FileRepo struct {
	db *ent.Database
}

// NewFileRepo creates a new FileRepo.
func NewFileRepo(db *ent.Database) dto.FileRepo {
	return &FileRepo{
		db: db,
	}
}

// Create creates a new file metadata record.
func (r *FileRepo) Create(ctx context.Context, in *types.FileMetadata, opts ...*dto.FileCreateOption) (*types.FileMetadata, error) {
	created, err := r.db.File(ctx).
		Create().
		SetOwnerID(in.OwnerId).
		SetName(in.Name).
		SetObjectID(in.ObjectId).
		SetVisibility(in.Visibility).
		SetMimeType(in.MimeType).
		SetSize(in.Size).
		SetSha256(in.Sha256).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertFileToFileMetadataPB(created), nil
}

// Get retrieves a file metadata record by its ID.
func (r *FileRepo) Get(ctx context.Context, id int64, opts ...*dto.FileQueryOption) (*types.FileMetadata, error) {
	found, err := r.db.File(ctx).Query().
		Where(
			file.ID(id),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertFileToFileMetadataPB(found), nil
}

// List retrieves a list of file metadata records.
func (r *FileRepo) List(ctx context.Context, opts ...*dto.FileQueryOption) ([]*types.FileMetadata, int, error) {
	opt := repo.FirstOrDefault(opts...)
	q := r.db.File(ctx).Query()

	if opt.OwnerID > 0 {
		q.Where(file.OwnerID(opt.OwnerID))
	}
	if opt.Visibility != "" {
		q.Where(file.Visibility(opt.Visibility))
	}

	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if opt.Page > 0 && opt.PageSize > 0 {
		q.Limit(opt.PageSize).Offset((opt.Page - 1) * opt.PageSize)
	}

	results, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	pbResults := make([]*types.FileMetadata, 0, len(results))
	for _, result := range results {
		pbResults = append(pbResults, dto.ConvertFileToFileMetadataPB(result))
	}

	return pbResults, total, nil
}

// Update updates a file metadata record.
func (r *FileRepo) Update(ctx context.Context, in *types.FileMetadata, opts ...*dto.FileUpdateOption) (*types.FileMetadata, error) {
	updater := r.db.File(ctx).
		Update().
		Where(file.ID(in.Id))

	if in.Name != "" {
		updater.SetName(in.Name)
	}
	if in.Visibility != "" {
		updater.SetVisibility(in.Visibility)
	}
	if in.ObjectId != "" {
		updater.SetObjectID(in.ObjectId)
	}
	if in.Size > 0 {
		updater.SetSize(in.Size)
	}
	if in.Sha256 != "" {
		updater.SetSha256(in.Sha256)
	}

	if err := updater.Exec(ctx); err != nil {
		return nil, err
	}

	updated, err := r.db.File(ctx).Query().Where(file.ID(in.Id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ConvertFileToFileMetadataPB(updated), nil
}

// Delete removes a file metadata record by its ID.
func (r *FileRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.File(ctx).Delete().Where(file.ID(id)).Exec(ctx)
	return err
}
