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

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").
			Default("").
			Comment(i18n.Text("entity.message.field.subject")),
		field.String("content").
			Default("").
			Comment(i18n.Text("entity.message.field.content")),
		field.Int8("status").
			GoType(enums.Status(0)).
			Default(int8(enums.StatusUnknown)).
			Comment(i18n.Text("entity.message.field.status")),
		mixin.FK("category_id", i18n.Text("entity.message.field.category_id")),
	}
}

// Annotations of the Message.
func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("ntf_messages"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.message.table.comment")),
	}
}

// Mixin of the Message.
func (Message) Mixin() []ent.Mixin {
	return mixin.AuditModelMixin
}
