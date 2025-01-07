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

	"origadmin/application/admin/helpers/ent/mixin"
	"origadmin/application/admin/helpers/i18n"
)

// Department holds the schema definition for the Department domain.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").
			MaxLen(32).
			Default("").
			Comment(i18n.Text("department.field.keyword")),
		field.String("name").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("department.field.name")),
		field.String("description").
			MaxLen(256).
			Default("").
			Comment(i18n.Text("department.field.description")),
		field.Int("sequence").
			Comment(i18n.Text("department.field.sequence")),
		field.Int8("status").
			Default(0).
			Comment(i18n.Text("department.field.status")),
		field.Int("level").
			Default(1).
			Comment(i18n.Text("department.field.level")),
		mixin.OP("parent_id", "department.field.parent_id"),
	}
}

// Mixin of the Department.
func (Department) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Department.
func (Department) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("keyword"),
		index.Fields("name"),
		index.Fields("sequence"),
		index.Fields("status"),
	}
}

// Annotations of the Department.
func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_departments"),
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		schema.Comment(i18n.Text("department.table.comment")),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("departments").
			Through("user_departments", UserDepartment.Type),
		edge.To("positions", Position.Type),
		edge.To("roles", Role.Type).
			StorageKey(edge.Columns("department_id", "role_id")).
			Through("department_roles", DepartmentRole.Type),
		// 添加部门层级关系
		edge.To("children", Department.Type),
		edge.From("parent_id", Department.Type).
			Ref("children").
			Field("parent_id").
			Unique(),
	}
}
