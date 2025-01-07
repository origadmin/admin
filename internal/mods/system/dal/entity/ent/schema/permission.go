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

	"origadmin/application/admin/helpers/ent/mixin"
	"origadmin/application/admin/helpers/i18n"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// DataScope 数据范围
const (
	DataScopeSelf string = "self" // 仅本人数据
	DataScopeDept string = "dept" // 本部门数据
	DataScopeAll  string = "all"  // 所有数据
)

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("permission.field.name")),
		field.String("keyword").
			MaxLen(64).
			Unique().
			Comment(i18n.Text("permission.field.keyword")),
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("permission.field.description")),
		// 数据权限相关字段
		field.String("data_scope").
			Default(DataScopeSelf).
			Comment(i18n.Text("permission.field.data_scope")),
		field.JSON("data_rules", map[string]string{}).
			Optional().
			Comment(i18n.Text("permission.field.data_rules")),
	}
}

// Annotations of the Permission.
func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_permissions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("permission.table.comment")),
	}
}

// Mixin of the Permission.
func (Permission) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("permissions").
			Through("role_permissions", RolePermission.Type),
		edge.To("resources", Resource.Type).
			Through("permission_resources", PermissionResource.Type),
		edge.From("positions", Position.Type).
			Ref("permissions").
			Through("position_permissions", PositionPermission.Type),
	}
}
