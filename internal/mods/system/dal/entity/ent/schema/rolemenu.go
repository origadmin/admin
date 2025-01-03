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

// RoleMenu holds the schema definition for the RoleMenu domain.
type RoleMenu struct {
	ent.Schema
}

// Fields of the RoleMenu.
func (RoleMenu) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("role_id", "role_menu:field:role_id"), // From Role.ID
		mixin.FK("menu_id", "role_menu:field:menu_id"), // From Menu.ID
	}
}

// Mixin of the RoleMenu.
func (RoleMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

// Indexes of the RoleMenu.
func (RoleMenu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id"), // From Role.ID
		index.Fields("menu_id"), // From Menu.ID
		index.Fields("role_id", "menu_id").
			Unique(),
	}
}

// Annotations of the Role.
func (RoleMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_role_menus"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("role_menu:table:comment")),
	}
}

// Edges of the RoleMenu.
func (RoleMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).
			Field("role_id").
			//Through("role_menus", RoleMenu.Type).
			Required().
			Unique(),
		edge.To("menu", Menu.Type).
			Field("menu_id").
			Required().
			Unique(),
		//edge.From("role", Role.Type).
		//	Ref("role_menus").
		//	Field("role_id").Unique(),
		//edge.From("menu", Menu.Type).
		//	Ref("role_menus").
		//	Field("menu_id").Unique(),
	}
}
