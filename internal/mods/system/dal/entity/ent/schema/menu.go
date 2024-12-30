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

const (
	MenuStatusActivated = 0
	MenuStatusFrozen    = 1
)

const (
	MenuTypeAction rune = 'A'
	MenuTypeButton rune = 'B'
	MenuTypeLink   rune = 'L'
	MenuTypeMenu   rune = 'M'
	MenuTypePage   rune = 'P'
)

// Menu holds the schema definition for the Menu domain.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").
			MaxLen(32).
			Default("").
			Comment(i18n.Text("menu:field:keyword")), // Unique keyword for the menu item
		field.String("name").
			MaxLen(128).
			Default("").
			Comment(i18n.Text("menu:field:name")), // Display name of the menu item
		field.String("i18n_key").
			MaxLen(128).
			Default("").
			Comment(i18n.Text("menu:field:i18n_key")), // I18n key for the menu item
		field.String("description").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("menu:field:description")), // Description of the menu item
		field.String("type").
			Default("page").
			Comment(i18n.Text("menu:field:type")), // Type of the menu item (e.g., Page,Button,Action,Menu,Link)
		field.String("icon").
			MaxLen(32).Default("").
			Comment(i18n.Text("menu:field:icon")), // Icon for the menu item
		field.String("path").
			MaxLen(255).
			Default("").
			Comment(i18n.Text("menu:field:path")), // Path associated with the menu item
		field.Int8("status").
			Default(MenuStatusActivated).
			Comment(i18n.Text("menu:field:status")), // Status of the menu item (e.g., activated, deactivated)
		field.String("parent_path").
			MaxLen(255).
			Default("").
			Comment(i18n.Text("menu:field:parent_path")), // Parent path of the menu item
		field.Int("sequence").
			Default(0).
			Comment(i18n.Text("menu:field:sequence")), // Sequence for sorting the menu item
		field.Text("properties").
			Default("").
			Optional().
			Comment(i18n.Text("menu:field:properties")), // Additional properties of the menu item
		mixin.OP("parent_id", "menu:field:parent_id"),   // Parent ID of the menu item

	}
}

// Mixin of the Menu.
func (Menu) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Indexes of the Menu.
func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		// Unique index
		index.Fields("keyword"),
		index.Fields("name"),
		index.Fields("type"),
		index.Fields("status"),
		index.Fields("parent_id"),
		index.Fields("parent_path"),
	}
}

// Annotations of the Menu.
func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_menus"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("menu:table:comment")),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		// Child menus
		edge.To("children", Menu.Type),
		// Parent menus
		edge.From("parent", Menu.Type).
			Ref("children").
			Field("parent_id").
			Unique(),
		edge.To("resources", Resource.Type),
		edge.From("roles", Role.Type).
			Ref("menus").
			Through("role_menus", RoleMenu.Type),
		edge.To("permissions", Permission.Type).
			StorageKey(edge.Columns("menu_id", "permission_id")).
			Through("menu_permissions", MenuPermission.Type),
	}
}
