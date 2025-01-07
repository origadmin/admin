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

type PermissionResource struct {
	ent.Schema
}

func (PermissionResource) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("permission_id"),
		mixin.FK("resource_id"),
		field.String("actions").
			Default("*").
			Comment(i18n.Text("permission_resource.field.actions")),
		//// 使用JSON字段存储更复杂的权限控制
		//field.JSON("access_control", struct {
		//	Actions    []string          `json:"actions"`     // 允许的操作
		//	Conditions map[string]string `json:"conditions"`  // 条件约束
		//	ValidFrom  *time.Time        `json:"valid_from"`  // 生效时间
		//	ValidUntil *time.Time        `json:"valid_until"` // 失效时间
		//	Attributes map[string]any    `json:"attributes"`  // 其他属性
		//}{}).
		//	Optional().
		//	Comment(i18n.Text("permission_resource.field.access_control")),
	}
}

func (PermissionResource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

func (PermissionResource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("permission_id", "resource_id").
			Unique(),
	}
}

func (PermissionResource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_permission_resources"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("permission_resource.table.comment")),
	}
}

func (PermissionResource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permission", Permission.Type).
			Field("permission_id").
			Unique().
			Required(),
		edge.To("resource", Resource.Type).
			Field("resource_id").
			Unique().
			Required(),
	}
}
