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

type PermissionMenu struct {
	ent.Schema
}

func (PermissionMenu) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("permission_id"),
		mixin.FK("menu_id"),
	}
}

func (PermissionMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.ID{},
	}
}

func (PermissionMenu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("permission_id"),
		index.Fields("menu_id"),
		index.Fields("permission_id", "menu_id").
			Unique(),
	}
}

func (PermissionMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_permission_menus"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("permission_menu.table.comment")),
	}
}

func (PermissionMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permission", Permission.Type).
			Field("permission_id").
			Unique().
			Required(),
		edge.To("menu", Menu.Type).
			Field("menu_id").
			Unique().
			Required(),
	}
}
