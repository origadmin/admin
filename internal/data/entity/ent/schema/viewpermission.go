package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/internal/helpers/ent/mixin"
	"origadmin/application/admin/internal/helpers/i18n"
)

// ViewPermission holds the schema definition for the ViewPermission entity.
// It's a through-table for the M-N relationship between View and Permission.
type ViewPermission struct {
	ent.Schema
}

// Fields of the ViewPermission.
func (ViewPermission) Fields() []ent.Field {
	return []ent.Field{
		mixin.FK("view_id", i18n.Text("view_permission.view_id.comment")),
		mixin.FK("permission_id", i18n.Text("view_permission.permission_id.comment")),
	}
}

// Indexes of the ViewPermission.
func (ViewPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("view_id", "permission_id").
			Unique(),
	}
}

// Annotations of the ViewPermission.
func (ViewPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("sys_view_permissions"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.view_permission.table.comment")),
	}
}

// Edges of the ViewPermission.
func (ViewPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("view", View.Type).
			Field("view_id").
			Unique().
			Required().
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("permission", Permission.Type).
			Field("permission_id").
			Unique().
			Required().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
