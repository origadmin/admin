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

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(64).
			Unique().
			Comment(i18n.Text("position.field.name")),
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("position.field.description")),
		mixin.FK("department_id", i18n.Text("department.field.department_id")),
	}
}

// Mixin of the Position.
func (Position) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Annotations of the Position.
func (Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_positions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("position.table.comment")),
	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("department", Department.Type).
			Ref("positions").
			Field("department_id").
			Unique().
			Required(), // Required() must set with Field Optional(), comment it if the field is not Optional() Field
		edge.To("user_positions", UserPosition.Type),
	}
}
