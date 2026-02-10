/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// UserSetting holds the schema definition for the UserSetting entity.
type UserSetting struct {
	ent.Schema
}

// Fields of the UserSetting.
func (UserSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("theme").
			Default("light").
			Comment(i18n.Text("entity.user_setting.field.theme")), // e.g., "light", "dark", "system"
		field.String("language").
			Default("en-US").
			Comment(i18n.Text("entity.user_setting.field.language")), // e.g., "en-US", "zh-CN"
		field.String("timezone").
			Default("UTC").
			Comment(i18n.Text("entity.user_setting.field.timezone")), // e.g., "UTC", "Asia/Shanghai"
		field.JSON("preferences", map[string]interface{}{}).
			Optional().
			Comment(i18n.Text("entity.user_setting.field.preferences")), // Flexible JSON field for various settings
	}
}

// Mixin of the UserSetting.
func (UserSetting) Mixin() []ent.Mixin {
	return append(mixin.AuditHookModelMixin, mixin.SoftDeleteMixin{})
}

// Annotations of the UserSetting.
func (UserSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_user_settings"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.user_setting.table.comment")),
	}
}

// Edges of the UserSetting.
func (UserSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("setting").
			Unique().
			Required(),
	}
}
