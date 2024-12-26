/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"origadmin/application/admin/helpers/ent/mixin"
)

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(64).Unique(),
		field.String("description").MaxLen(256),
		mixin.OP("department_id"),
	}
}

// Mixin of the Position.
func (Position) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("department", Department.Type).
			Ref("positions").
			Field("department_id").
			Unique(),
		//Required(),
		edge.To("user_positions", UserPosition.Type),
	}
}
