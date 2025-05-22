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

type PermissionResource struct {
	ent.Schema
}

func (PermissionResource) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("permission_id"),
		mixin.FK("resource_id"),
	}
}

func (PermissionResource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		//mixin.ID{},
	}
}

func (PermissionResource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("permission_id", "resource_id").
			Unique(),
	}
}

func (PermissionResource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_permission_resources"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.permission_resource.table.comment")),
	}
}

func (PermissionResource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permission", Permission.Type).
			Field("permission_id").
			Unique().
			Required(),
		edge.To("resource", Resource.Type).
			Field("resource_id").
			Unique().
			Required(),
	}
}
