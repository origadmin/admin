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
	"entgo.io/ent/schema/mixin"
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

func (Resource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_resources"),
		entsql.WithComments(true),
		schema.Comment("Resource table"),
	}
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("ID").
			Immutable().
			Unique(),
		field.String("name").MaxLen(128).Default("").Comment("Name"),
		field.String("keyword").MaxLen(64).Unique().Comment("Keyword"),
		field.String("type").MaxLen(2).Default(ResourceTypeMenu).Comment("Type"),
		field.Int8("status").Default(ResourceStatusEnabled).Comment("Status"),
		field.String("path").MaxLen(256).Default("").Comment("Path"),
		field.String("component").MaxLen(128).Default("").Comment("Component"),
		field.String("icon").MaxLen(64).Default("").Comment("Icon"),
		field.Int("sequence").Default(0).Comment("Sequence"),
		field.Bool("visible").Default(true).Comment("Visible"),
		field.Int8("level").Default(0).Comment("Level"),
		field.String("tree_path").MaxLen(256).Default("").Comment("Tree path"),
		field.JSON("properties", map[string]string{}).Optional().Comment("Properties"),
		field.String("description").MaxLen(1024).Default("").Comment("Description"),
		field.Int64("parent_id").Optional().Comment("Parent ID"),
	}
}

// Mixin of the Resource.
func (Resource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Indexes of the Resource.
func (Resource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("parent_id"),
		index.Fields("level"),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Resource.Type).From("parent").Field("parent_id").Unique(),
		edge.From("permissions", Permission.Type).
			Ref("resources").
			Through("permission_resources", PermissionResource.Type),
	}
}
