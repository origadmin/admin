{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "update/additional/crud_one" }}
    {{ $builder := $.UpdateOneName }}
    {{- if hasSuffix $builder "UpdateOne" }}
        {{ $receiver := receiver $builder }}
        {{ $const := print .Package}}

        // Set{{ .Name }} sets the {{ .Name }} fields from input struct.
        // If no fields are specified, all fields except ID will be set.
        // Zero values are included in the update.
        func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
        m := {{ $receiver }}.mutation
        if len(fields) == 0 {
        fields = {{ $const }}.OmitColumns({{ $const }}.FieldID)
        }
        _ = m.SetFields(input, fields...)
        return {{ $receiver }}
        }

        // Set{{ .Name }}SkipZero sets the {{ .Name }} fields from input struct.
        // If no fields are specified, all fields except ID will be set.
        // Zero values are skipped and not updated.
        func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}SkipZero(input *{{ .Name }}, fields ...string) *{{ $builder }} {
        m := {{ $receiver }}.mutation
        if len(fields) == 0 {
        fields = {{ $const }}.OmitColumns({{ $const }}.FieldID)
        }
        _ = m.SetFieldsSkipZero(input, fields...)
        return {{ $receiver }}
        }

        {{ $onebuilder := $.UpdateOneName }}
        {{ $receiver = receiver $onebuilder }}
        // Omit excludes the specified fields from the update operation.
        // By default, all fields are updated. Use this method to exclude specific fields.
        func ({{ $receiver }} *{{ $onebuilder }}) Omit(fields ...string) *{{ $onebuilder }} {
        omits := make(map[string]struct{}, len(fields))
        for i := range fields {
				omits[fields[i]] = struct{}{}
        }
        {{ $receiver }}.fields = []string(nil)
        for _, col := range {{ .Package }}.Columns {
				if _, ok := omits[col]; !ok {
        {{ $receiver }}.fields = append({{ $receiver }}.fields, col)
				}
        }
        return {{ $receiver }}
        }
    {{- end }}

{{- end -}}
