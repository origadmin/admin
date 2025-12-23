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
)

type PermissionResource struct {
	ent.Schema
}

func (PermissionResource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_permission_resources"),
		entsql.WithComments(true),
		schema.Comment("Permission-Resource mapping table"),
	}
}

func (PermissionResource) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("permission_id"),
		field.Int64("resource_id"),
	}
}

func (PermissionResource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("permission_id", "resource_id").
			Unique(),
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
