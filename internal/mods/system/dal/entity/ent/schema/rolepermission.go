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
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/helpers/ent/mixin"
	"origadmin/application/admin/helpers/i18n"
)

// RolePermission holds the schema definition for the RolePermission entity.
type RolePermission struct {
	ent.Schema
}

// Fields of the RolePermission.
func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("role_id"),       // From Role.ID
		mixin.FK("permission_id"), // From Permission.ID
	}
}

// Mixin of the RolePermission.
func (RolePermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		//mixin.ID{},
	}
}

// Indexes of the RolePermission.
func (RolePermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id"),       // From Role.ID
		index.Fields("permission_id"), // From Permission.ID
		index.Fields("role_id", "permission_id").
			Unique(),
	}
}

// Annotations of the RolePermission.
func (RolePermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_role_permissions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.role_permission.table.comment")),
	}
}

// Edges of the RolePermission.
func (RolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).
			Field("role_id").
			Unique().
			Required(),
		edge.To("permission", Permission.Type).
			Field("permission_id").
			Unique().
			Required(),
	}
}
