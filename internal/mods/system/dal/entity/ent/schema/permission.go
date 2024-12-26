/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"origadmin/application/admin/helpers/ent/mixin"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(64).Unique(),
		field.String("description").MaxLen(256),
	}
}

// Mixin of the Permission.
func (Permission) Mixin() []ent.Mixin {
	return mixin.ModelMixin
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("permissions").
			Through("role_permissions", RolePermission.Type),
		edge.From("menus", Menu.Type).
			Ref("permissions").
			Through("menu_permissions", MenuPermission.Type),
		edge.To("resources", Resource.Type).
			Through("permission_resources", PermissionResource.Type),
	}
}
