package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"origadmin/application/admin/internal/helpers/i18n"
)

// CasbinRule holds the schema definition for the CasbinRule entity.
type CasbinRule struct {
	ent.Schema
}

// Fields of the CasbinRule.
func (CasbinRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("ptype").Default("").Comment(i18n.Text("entity.casbin_rule.field.ptype")),
		field.String("v0").Default("").Comment(i18n.Text("entity.casbin_rule.field.v0")),
		field.String("v1").Default("").Comment(i18n.Text("entity.casbin_rule.field.v1")),
		field.String("v2").Default("").Comment(i18n.Text("entity.casbin_rule.field.v2")),
		field.String("v3").Default("").Comment(i18n.Text("entity.casbin_rule.field.v3")),
		field.String("v4").Default("").Comment(i18n.Text("entity.casbin_rule.field.v4")),
		field.String("v5").Default("").Comment(i18n.Text("entity.casbin_rule.field.v5")),
	}
}

// Edges of the CasbinRule.
func (CasbinRule) Edges() []ent.Edge {
	return nil
}

func (CasbinRule) Index() []ent.Index {
	return []ent.Index{
		index.Fields("ptype", "v0", "v1", "v2", "v3", "v4", "v5").Unique(),
	}
}

// Annotations of the Notification.
func (CasbinRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("casbin_rule"),
		entsql.WithComments(true),
		schema.Comment(i18n.Text("entity.casbin_rule.table.comment")),
	}
}
