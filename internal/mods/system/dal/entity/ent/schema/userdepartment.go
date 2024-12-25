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
)

// UserDepartment holds the schema definition for the UserDepartment domain.
type UserDepartment struct {
	ent.Schema
}

// Fields of the UserDepartment.
func (UserDepartment) Fields() []ent.Field {
	return []ent.Field{
		mixin.FieldFK("user_id"),       // From User.ID
		mixin.FieldFK("department_id"), // From Department.ID
	}
}

// Mixin of the UserDepartment.
func (UserDepartment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

// Indexes of the UserDepartment.
func (UserDepartment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("department_id"),
		index.Fields("user_id", "department_id").
			Unique(),
	}
}

// Edges of the UserDepartment.
func (UserDepartment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			//Ref("departments").
			Field("user_id").
			Unique().
			Required(),
		edge.To("department", Department.Type).
			//Ref("users").
			Field("department_id").
			Unique().
			Required(),
	}
}

// Annotations of the UserDepartment.
func (UserDepartment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_user_departments"),
	}
}
