/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/helpers/ent/mixin"
)

// UserRole holds the schema definition for the UserRole domain.
type UserRole struct {
	ent.Schema
}

// Fields of the UserRole.
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").MaxLen(36),                // From User.ID
		field.String("role_id").MaxLen(36),                // From Role.ID
		field.String("role_name").MaxLen(128).Default(""), // From Role.Name
	}
}

// Mixin of the UserRole.
func (UserRole) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the UserRole.
func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("role_id"),
	}
}

// Annotations of the UserRole.
func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_user_roles"),
	}
}

// Edges of the UserRole.
func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Required().
			Unique(),
		edge.To("role", Role.Type).
			Field("role_id").
			Required().
			Unique(),
	}
}
