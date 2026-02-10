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

	"origadmin/application/admin/internal/data/enums"
	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// UserProfile holds the schema definition for the UserProfile entity.
type UserProfile struct {
	ent.Schema
}

// Fields of the UserProfile.
func (UserProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user_profile.field.nickname")), // Nickname display name of user
		field.String("avatar").
			MaxLen(256).
			Default("").
			Comment("entity.user_profile.field.avatar"), // Avatar display avatar of user
		field.String("name").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user_profile.field.name")), // Name of user
		field.Enum("gender").
			Values(
				string(enums.GenderMale),
				string(enums.GenderFemale),
				string(enums.GenderUnknown),
			).
			Default(string(enums.GenderUnknown)).
			Comment(i18n.Text("entity.user_profile.field.gender")), // Gender of user
		field.String("department").
			MaxLen(64).
			Default("").
			Comment(i18n.Text("entity.user_profile.field.department")), // Department of user
		field.String("remark").
			MaxLen(1024).
			Default("").
			Comment(i18n.Text("entity.user_profile.field.remark")), // Remark of user
	}
}

// Mixin of the UserProfile.
func (UserProfile) Mixin() []ent.Mixin {
	return append(mixin.AuditHookModelMixin, mixin.SoftDeleteMixin{})
}

// Annotations of the UserProfile.
func (UserProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_user_profiles"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.user_profile.table.comment")),
	}
}

// Edges of the UserProfile.
func (UserProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("profile").
			Unique().
			Required(),
	}
}
