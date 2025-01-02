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

// MenuPermission holds the schema definition for the MenuPermission domain.
type MenuPermission struct {
	ent.Schema
}

// Fields of the MenuPermission.
func (MenuPermission) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("menu_id", "menu_permission.field.menu_id"),             // From Menu.ID
		mixin.FK("permission_id", "menu_permission.field.permission_id"), // From Permission.ID
	}
}

// Mixin of the MenuPermission.
func (MenuPermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

// Indexes of the MenuPermission.
func (MenuPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("menu_id"),
		index.Fields("permission_id"),
		index.Fields("menu_id", "permission_id").
			Unique(),
	}
}

// Annotations of the MenuPermission.
func (MenuPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_menu_permissions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("menu_permission.table.comment")),
	}
}

// Edges of the MenuPermission.
func (MenuPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menu", Menu.Type).
			Field("menu_id").
			Unique().
			Required(),
		edge.To("permission", Permission.Type).
			Field("permission_id").
			Unique().
			Required(),
	}
}
