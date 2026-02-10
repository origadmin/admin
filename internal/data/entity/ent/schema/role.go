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

	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// Role holds the schema definition for the Role domain.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").
			MaxLen(255).
			Unique().
			Comment(i18n.Text("entity.role.field.keyword")), // keyword of role (unique)
		field.String("name").
			MaxLen(128).
			Default("").
			Comment(i18n.Text("entity.role.field.name")), // Display name of role
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("entity.role.field.description")), // Details about role
		field.Int8("type").
			GoType(enums.RoleType(0)).
			Default(int8(enums.RoleTypeUser)).
			Comment(i18n.Text("entity.role.field.type")), // Role type: 1 - System role 2 - User role
		field.Int("sequence").
			Default(0).
			Comment(i18n.Text("entity.role.field.sequence")), // Sequence for sorting
		field.Int8("status").
			GoType(enums.Status(0)).
			Default(int8(enums.StatusActive)).
			Comment(i18n.Text("entity.role.field.status")),
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return mixin.AuditModelMixin
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
		schema.Comment(i18n.Text("entity.role.table.comment")),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("menus", Menu.Type).
		//	StorageKey(edge.Columns("role_id", "menu_id")).
		//	Through("role_menus", RoleMenu.Type),
		edge.From("users", User.Type).
			Ref("roles").
			Through("user_roles", UserRole.Type),
		edge.To("permissions", Permission.Type).
			StorageKey(edge.Columns("role_id", "permission_id")).
			Through("role_permissions", RolePermission.Type),
	}
}
