// Code generated by ent, DO NOT EDIT.

package casbinrule

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the casbinrule type in the database.
	Label = "casbin_rule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPtype holds the string denoting the ptype field in the database.
	FieldPtype = "ptype"
	// FieldV0 holds the string denoting the v0 field in the database.
	FieldV0 = "v0"
	// FieldV1 holds the string denoting the v1 field in the database.
	FieldV1 = "v1"
	// FieldV2 holds the string denoting the v2 field in the database.
	FieldV2 = "v2"
	// FieldV3 holds the string denoting the v3 field in the database.
	FieldV3 = "v3"
	// FieldV4 holds the string denoting the v4 field in the database.
	FieldV4 = "v4"
	// FieldV5 holds the string denoting the v5 field in the database.
	FieldV5 = "v5"
	// Table holds the table name of the casbinrule in the database.
	Table = "casbin_rules"
)

// Columns holds all SQL columns for casbinrule fields.
var Columns = []string{
	FieldID,
	FieldPtype,
	FieldV0,
	FieldV1,
	FieldV2,
	FieldV3,
	FieldV4,
	FieldV5,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPtype holds the default value on creation for the "Ptype" field.
	DefaultPtype string
	// DefaultV0 holds the default value on creation for the "V0" field.
	DefaultV0 string
	// DefaultV1 holds the default value on creation for the "V1" field.
	DefaultV1 string
	// DefaultV2 holds the default value on creation for the "V2" field.
	DefaultV2 string
	// DefaultV3 holds the default value on creation for the "V3" field.
	DefaultV3 string
	// DefaultV4 holds the default value on creation for the "V4" field.
	DefaultV4 string
	// DefaultV5 holds the default value on creation for the "V5" field.
	DefaultV5 string
)

// OrderOption defines the ordering options for the CasbinRule queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPtype orders the results by the Ptype field.
func ByPtype(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPtype, opts...).ToFunc()
}

// ByV0 orders the results by the V0 field.
func ByV0(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV0, opts...).ToFunc()
}

// ByV1 orders the results by the V1 field.
func ByV1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV1, opts...).ToFunc()
}

// ByV2 orders the results by the V2 field.
func ByV2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV2, opts...).ToFunc()
}

// ByV3 orders the results by the V3 field.
func ByV3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV3, opts...).ToFunc()
}

// ByV4 orders the results by the V4 field.
func ByV4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV4, opts...).ToFunc()
}

// ByV5 orders the results by the V5 field.
func ByV5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV5, opts...).ToFunc()
}

// SelectColumns returns all selected fields.
func SelectColumns(fields []string) []string {
	// Default removal FieldID
	filteredFields := make([]string, 0, len(fields))
	for _, field := range fields {
		if field != FieldID {
			filteredFields = append(filteredFields, field)
		}
	}
	return filteredFields
}

// OmitColumns returns all fields that are not in the list of fields.
func OmitColumns(fields ...string) []string {
	// Default removal FieldID
	return omitColumns(Columns, fields, true)
}

// OmitCustomColumns returns all fields that are not in the list of fields.
func OmitCustomColumns(src []string, fields ...string) []string {
	if len(src) == 0 {
		src = Columns
	}
	// Default removal FieldID
	return omitColumns(src, fields, true)
}

// OmitColumnsWithID returns all fields that are not in the list of fields.
func OmitColumnsWithID(fields ...string) []string {
	// Not remove FieldID
	return omitColumns(Columns, fields, false)
}

// OmitCustomColumns returns all fields that are not in the list of fields.
func OmitCustomColumnsWithID(src []string, fields ...string) []string {
	if len(src) == 0 {
		src = Columns
	}
	// Not remove FieldID
	return omitColumns(src, fields, false)
}

func omitColumns(src []string, fields []string, omitID bool) []string {
	// Default removal FieldID
	filteredFields := make([]string, 0, len(src))
	for _, field := range src {
		if !(omitID && field == FieldID) && !contains(fields, field) {
			filteredFields = append(filteredFields, field)
		}
	}
	return filteredFields
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
