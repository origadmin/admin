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
)

// Department holds the schema definition for the Department domain.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").MaxLen(32).Default(""),      // Code of Department (unique)
		field.String("name").MaxLen(64).Default(""),         // Display name of Department
		field.String("description").MaxLen(256).Default(""), // Details about Department
		field.Int("sequence"),                               // Sequence for sorting
		field.Int8("status").Default(0),
		// Department menu list
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
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("menus", Menu.Type).
		//	StorageKey(edge.Columns("department_id", "menu_id")).
		//	Through("Department_menus", DepartmentMenu.Type),
		// edge.To("menus", Menu.Dialect).StorageKey(edge.Column("menu_id")),
		// edge.To("users", User.Type).
		//     StorageKey(edge.Columns("user_id", "Department_id")).
		//     Through("user_Department", UserDepartment.Type), // .Field("user_id"),
		edge.From("users", User.Type).
			Ref("departments").
			Through("user_departments", UserDepartment.Type), // .Field("user_id"),
		//edge.To("user_departments", UserDepartment.Type),
		//	From("department").
		//	Ref("department"),
		// edge.From("users", User.Dialect).Ref("Departments").StorageKey(edge.Table(config.c.FormatTableName("user_Department")), edge.Column("Department_id")).Unique(),
		edge.To("positions", Position.Type),
	}
}
