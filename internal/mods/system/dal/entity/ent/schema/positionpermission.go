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

// PositionPermission holds the schema definition for the PositionPermission domain.
type PositionPermission struct {
	ent.Schema
}

// Fields of the PositionPermission.
func (PositionPermission) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("position_id", "position_permission.field.position_id"),     // From Position.ID
		mixin.FK("permission_id", "position_permission.field.permission_id"), // From Permission.ID
	}
}

// Mixin of the PositionPermission.
func (PositionPermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		//mixin.ID{},
	}
}

// Indexes of the PositionPermission.
func (PositionPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("permission_id"), // From Permission.ID
		index.Fields("position_id"),   // From Position.ID
		index.Fields("position_id", "permission_id").
			Unique(),
	}
}

// Annotations of the PositionPermission.
func (PositionPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_position_permissions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("position_permission.table.comment")),
	}
}

// Edges of the PositionPermission.
func (PositionPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("position", Position.Type).
			Field("position_id").
			Unique().
			Required(),
		edge.To("permission", Permission.Type).
			Field("permission_id").
			Unique().
			Required(),
	}
}
