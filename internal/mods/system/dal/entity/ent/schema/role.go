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

// Role type constant
const (
	RoleTypeSystem     int8 = 1 // System roles (e.g., Super Administrator, Guest)
	RoleTypeUser       int8 = 2 // User roles (e.g., general user, operation, customer service)
	RoleTypeDepartment int8 = 3 // Department roles (e.g., department heads, department members)
)

// Role holds the schema definition for the Role domain.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").
			MaxLen(32).
			Default("").
			Comment("role.field.keyword"), // Code of role (unique)
		field.String("name").
			MaxLen(128).
			Default("").
			Comment("role.field.name"), // Display name of role
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment("role.field.description"), // Details about role
		field.Int8("type").
			Default(RoleTypeUser).
			Comment("role.field.type"), //("Role type: 1 - System role 2 - User role 3 - Department role"),
		field.Int("sequence").
			Default(0).
			Comment("role.field.sequence"), // Sequence for sorting
		field.Int8("status").
			Default(0).
			Comment("role.field.status"),
		field.Bool("is_system").
			Default(false).
			Comment("role.field.is_system"), //("Whether the system is built-in (the built-in role cannot be deleted)"),
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Role.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("keyword"),
		index.Fields("name"),
		index.Fields("sequence"),
		index.Fields("status"),
	}
}

// Annotations of the Role.
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_roles"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("role.table.comment")),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menus", Menu.Type).
			StorageKey(edge.Columns("role_id", "menu_id")).
			Through("role_menus", RoleMenu.Type),
		edge.From("users", User.Type).
			Ref("roles").
			Through("user_roles", UserRole.Type),
		edge.To("permissions", Permission.Type).
			StorageKey(edge.Columns("role_id", "permission_id")).
			Through("role_permissions", RolePermission.Type),
		edge.From("departments", Department.Type).
			Ref("roles").
			Through("department_roles", DepartmentRole.Type),
	}
}
