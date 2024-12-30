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

// UserPosition holds the schema definition for the UserPosition entity.
type UserPosition struct {
	ent.Schema
}

// Fields of the UserPosition.
func (UserPosition) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("user_id"),     // From User.ID
		mixin.FK("position_id"), // From Position.ID
	}
}

// Mixin of the UserPosition.
func (UserPosition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

// Indexes of the UserPosition.
func (UserPosition) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),     // From User.ID
		index.Fields("position_id"), // From Position.ID
		index.Fields("user_id", "position_id").
			Unique(),
	}
}

// Annotations of the UserPosition
func (UserPosition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_user_positions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("user_position:table:comment")),
	}
}

// Edges of the UserPosition.
func (UserPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			//Ref("users").
			Field("user_id").
			Unique().
			Required(),
		edge.To("position", Position.Type).
			//Ref("positions").
			Field("position_id").
			Unique().
			Required(),
	}
}
