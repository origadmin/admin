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
	"entgo.io/ent/schema/mixin"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// DataScope 数据范围
const (
	DataScopeSelf string = "self" // 仅本人数据
	DataScopeDept string = "dept" // 部门数据
	DataScopeRole string = "role" // 角色数据
	DataScopeAll  string = "all"  // 所有数据
)

func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_permissions"),
		entsql.WithComments(true),
		schema.Comment("Permission table"),
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("ID").
			Immutable().
			Unique(),
		field.String("name").MaxLen(64).Default("").Comment("Name"),
		field.String("keyword").MaxLen(64).Unique().Comment("Keyword"),
		field.String("description").MaxLen(1024).Default("").Comment("Description"),
		field.String("data_scope").Default(DataScopeSelf).Comment("Data scope"),
		field.JSON("data_rules", map[string]string{}).Optional().Comment("Data rules"),
		field.Enum("actions").Values("read", "write", "delete", "manage").Default("read").Comment("Actions"),
	}
}

// Mixin of the Permission.
func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("permissions").
			Through("role_permissions", RolePermission.Type),
		edge.To("resources", Resource.Type).
			Through("permission_resources", PermissionResource.Type),
	}
}
