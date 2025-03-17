{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Type*/}}

{{ define "meta/additional/fields" }}

	// SelectColumns returns all selected fields.
	func SelectColumns(fields []string) []string  {
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
	return omitColumns(Columns,fields, true)
	}

	// OmitCustomColumns returns all fields that are not in the list of fields.
	func OmitCustomColumns(src []string,fields ...string) []string {
	if len(src) == 0 {
	src= Columns
	}
	// Default removal FieldID
	return omitColumns(src,fields, true)
	}

	// OmitColumnsWithID returns all fields that are not in the list of fields.
	func OmitColumnsWithID(fields ...string) []string {
	// Not remove FieldID
	return omitColumns(Columns,fields, false)
	}

	// OmitCustomColumns returns all fields that are not in the list of fields.
	func OmitCustomColumnsWithID(src []string,fields ...string) []string {
	if len(src) == 0 {
	src= Columns
	}
	// Not remove FieldID
	return omitColumns(src,fields, false)
	}

	func omitColumns(src []string,fields []string,omitID bool) []string {
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
{{ end }}
