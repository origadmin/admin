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

	"origadmin/application/admin/internal/data/entity/ent/schema/types"
)

// Role type constant
const (
	RoleTypeSystem int8 = 1 // System roles (e.g., Super Admin)
	RoleTypeUser   int8 = 2 // User roles (e.g., general user, operation, customer service)
)

// Role holds the schema definition for the Role domain.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_roles"),
		entsql.WithComments(true),
		schema.Comment("Role table"),
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("ID").
			Immutable().
			Unique(),
		field.String("keyword").
			MaxLen(32).
			Unique().
			Comment("keyword of role (unique)"),
		field.String("name").
			MaxLen(128).
			Default("").
			Comment("Display name of role"),
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment("Details about role"),
		field.Int8("type").
			Default(RoleTypeUser).
			Comment("Role type: 1 - System role 2 - User role 3 - Department role"),
		field.Int("sequence").
			Default(0).
			Comment("Sequence for sorting"),
		field.Int8("status").
			Default(types.Active).
			Comment("status"),
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Indexes of the Role.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("keyword"),
		index.Fields("name"),
		index.Fields("sequence"),
		index.Fields("status"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("roles").
			Through("user_roles", UserRole.Type),
		edge.To("permissions", Permission.Type).
			StorageKey(edge.Columns("role_id", "permission_id")).
			Through("role_permissions", RolePermission.Type),
	}
}
