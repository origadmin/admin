/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").
			Default("").
			Comment(i18n.Text("entity.notification.field.subject")),
		field.String("content").
			Default("").
			Comment(i18n.Text("entity.notification.field.content")),
		field.Int8("status").
			GoType(enums.Status(0)).              // Tell entc to generate the Go type as enums.Status
			Default(int8(enums.StatusUnknown)). // Provide the underlying type (int8) to the builder method
			Comment(i18n.Text("entity.notification.field.status")),
		mixin.FK("category_id", i18n.Text("entity.notification.field.category_id")),
	}
}

// Annotations of the Notification.
func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("msg_notifications"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.notification.table.comment")),
	}
}

// Mixin of the Notification.
func (Notification) Mixin() []ent.Mixin {
	return mixin.AuditModelMixin
}
