/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/helpers/ent/mixin"
	"origadmin/application/admin/helpers/i18n"
)

const (
	ResourceStatusEnabled  int8 = 1 // 启用
	ResourceStatusDisabled int8 = 2 // 禁用
)
const (
	ResourceTypeUnknown  = "U"    // 未知
	ResourceTypeRoot     = "ROOT" // 根目录
	ResourceTypeGroup    = "G"    // 分组
	ResourceTypeMenu     = "M"    // 目录
	ResourceTypePage     = "P"    // 页面
	ResourceTypeButton   = "B"    // 按钮
	ResourceTypeAPI      = "A"    // API接口
	ResourceTypeRedirect = "R"    // 重定向

)

// Resource holds the schema definition for the Resource domain.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(128).
			Default("").
			Comment(i18n.Text("resource.field.name")),
		field.String("keyword").
			MaxLen(64).
			Unique().
			Comment(i18n.Text("resource.field.keyword")),
		field.String("i18n_key").
			MaxLen(128).
			Default("").
			Comment(i18n.Text("resource.field.i18n_key")),
		field.String("type").
			MaxLen(2).
			Default(ResourceTypeMenu).
			Comment(i18n.Text("resource.field.type")),
		field.Int8("status").
			Default(ResourceStatusEnabled).
			Comment(i18n.Text("resource.field.status")),
		// fields that are unique to api resources
		field.String("uri").
			MaxLen(256).
			Default("").
			Comment(i18n.Text("resource.field.uri")),
		// fields that are unique to grpc resources
		field.String("operation").
			MaxLen(128).
			Default("").
			Comment(i18n.Text("resource.field.operation")),
		field.String("method").
			MaxLen(16).
			Default("").
			Comment(i18n.Text("resource.field.method")),
		// fields specific to ui resources
		field.String("component").
			MaxLen(128).
			Default("").
			Comment(i18n.Text("resource.field.component")),
		// fields specific to ui resources
		field.String("icon").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("resource.field.icon")),
		// menu sort field
		field.Int("sequence").
			Default(0).
			Comment(i18n.Text("resource.field.sequence")),
		// menu specific fields
		field.Bool("visible").
			Default(true).
			Comment(i18n.Text("resource.field.visible")),
		field.String("tree_path").
			MaxLen(256).
			Default("").
			Comment(i18n.Text("resource.field.tree_path")),
		// extended properties
		field.JSON("properties", map[string]string{}).
			Optional().
			Comment(i18n.Text("resource.field.properties")),
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("resource.field.description")),
		mixin.OP("parent_id", "resource.field.parent_id"),
	}
}

// Mixin of the Resource.
func (Resource) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Resource.
func (Resource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("parent_id"),
	}
}

// Annotations of the Menu.
func (Resource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_resources"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("resource.table.comment")),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Resource.Type),
		edge.From("parent", Resource.Type).
			Ref("children").
			Field("parent_id").
			Unique(),
		edge.From("permissions", Permission.Type).
			Ref("resources").
			Through("permission_resources", PermissionResource.Type),
	}
}
