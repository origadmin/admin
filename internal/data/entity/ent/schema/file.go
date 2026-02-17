/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment(i18n.Text("entity.filemanager.field.name")),
		field.String("object_id").
			Comment(i18n.Text("entity.filemanager.field.object_id")),
		field.String("visibility").
			Comment(i18n.Text("entity.filemanager.field.visibility")).
			Default("private"),
		field.String("mime_type").
			Comment(i18n.Text("entity.filemanager.field.mime_type")).
			Optional(),
		field.Int64("size").
			Comment(i18n.Text("entity.filemanager.field.size")).
			Default(0),
		field.Bool("is_permanent").
			Comment("Whether the file access is permanent").
			Default(true),
		field.Time("expires_at").
			Comment("Access expiration time (nullable, means no expiration)").
			Optional(),
		field.Int("max_downloads").
			Comment("Maximum download count (0 means unlimited)").
			Default(0),
		field.Int("download_count").
			Comment("Current download count").
			Default(0),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}

// Indexes of the File.
func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("object_id").Unique(),
	}
}

// Mixin of the File.
func (File) Mixin() []ent.Mixin {
	return append(
		mixin.AuditHookModelMixin,
		mixin.SoftDeleteMixin{},
		mixin.DefaultOwnerMixin(),
	)
}

// Annotations of the File.
func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("fm_files"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.filemanager.table.comment")),
	}
}

// Interceptors of the File.
func (File) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		mixin.SoftDeleteInterceptor(mixin.SoftDeleteMixin{}),
	}
}

// Hooks of the File.
func (File) Hooks() []ent.Hook {
	return []ent.Hook{
		mixin.SoftDeleteHook(mixin.SoftDeleteMixin{}),
	}
}
