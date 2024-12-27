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

// DepartmentRole holds the schema definition for the DepartmentRole domain.
type DepartmentRole struct {
	ent.Schema
}

// Fields of the DepartmentRole.
func (DepartmentRole) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("department_id", "department:department_id"), // From Department.ID
		mixin.FK("role_id", "role:role_id"),                   // From Role.ID
	}
}

// Mixin of the DepartmentRole.
func (DepartmentRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

// Indexes of the DepartmentRole.
func (DepartmentRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("department_id"), // From Department.ID
		index.Fields("role_id"),       // From Role.ID
		index.Fields("department_id", "role_id").
			Unique(),
	}
}

// Annotations of the DepartmentRole.
func (DepartmentRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_department_roles"),
	}
}

// Edges of the DepartmentRole.
func (DepartmentRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("department", Department.Type).
			Field("department_id").
			Required().
			Unique(),
		edge.To("role", Role.Type).
			Field("role_id").
			Required().
			Unique(),
	}
}
