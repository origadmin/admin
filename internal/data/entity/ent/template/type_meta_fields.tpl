{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Type*/}}

{{ define "meta/additional/fields" }}

	// SelectColumns returns all selected fields excluding the ID field.
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

	// OmitColumns returns all fields from Columns that are not in the provided list.
	// The ID field is always excluded.
	func OmitColumns(fields ...string) []string {
	// Default removal FieldID
	return omitColumns(Columns, fields, true)
	}

	// OmitCustomColumns returns all fields from src that are not in the provided list.
	// If src is empty, Columns will be used as the source.
	// The ID field is always excluded.
	func OmitCustomColumns(src []string, fields ...string) []string {
	if len(src) == 0 {
	src = Columns
	}
	// Default removal FieldID
	return omitColumns(src, fields, true)
	}

	// OmitColumnsWithID returns all fields from Columns that are not in the provided list.
	// The ID field is included in the result.
	func OmitColumnsWithID(fields ...string) []string {
	// Not remove FieldID
	return omitColumns(Columns, fields, false)
	}

	// OmitCustomColumnsWithID returns all fields from src that are not in the provided list.
	// If src is empty, Columns will be used as the source.
	// The ID field is included in the result.
	func OmitCustomColumnsWithID(src []string, fields ...string) []string {
	if len(src) == 0 {
	src = Columns
	}
	// Not remove FieldID
	return omitColumns(src, fields, false)
	}

	// omitColumns returns all fields from src that are not in the fields list.
	// If omitID is true, the ID field is excluded from the result.
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

	// contains checks if the item exists in the slice.
	func contains(slice []string, item string) bool {
	for _, s := range slice {
	if s == item {
	return true
	}
	}
	return false
	}
{{ end }}
