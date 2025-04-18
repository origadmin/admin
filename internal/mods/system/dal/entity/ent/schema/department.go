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
			MaxLen(64).
			Unique().
			Comment(i18n.Text("entity.department.field.keyword")),
		field.String("name").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.department.field.name")),
		// use materialized path model to store the tree structure
		// Parent path of the menu item
		field.String("tree_path").
			MaxLen(256).
			Default("").
			Comment(i18n.Text("entity.menu.field.tree_path")),
		field.Int("sequence").
			Comment(i18n.Text("entity.department.field.sequence")),
		field.Int8("status").
			Default(0).
			Comment(i18n.Text("entity.department.field.status")),
		field.Int("level").
			Default(1).
			Comment(i18n.Text("entity.department.field.level")),
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("entity.department.field.description")),
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
		schema.Comment(i18n.Text("entity.department.table.comment")),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("departments").
			Through("user_departments", UserDepartment.Type),
		edge.To("positions", Position.Type),
		edge.To("children", Department.Type).
			From("parent").Unique().Field("parent_id"),
		//edge.To("children", Department.Type),
		//edge.From("parent", Department.Type).
		//	Ref("children").
		//	Field("parent_id").
		//	Unique(),
	}
}
