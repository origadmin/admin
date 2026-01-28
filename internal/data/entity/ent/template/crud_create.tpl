{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "create/additional/crud" }}

    {{ $builder := .CreateName }}
    {{ $receiver := .CreateReceiver }}
    {{ $fields := .Fields }}
    {{- $const := print .Package}}
    {{- if .ID.UserDefined }}
        {{ $fields = append $fields .ID }}
    {{- end }}

    // Set{{ .Name }} sets the {{ .Name }} fields from input struct.
    // If no fields are specified, all fields from Columns will be set.
    // Zero values are included in the update.
    func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
		m := {{ $receiver }}.mutation
		if len(fields) == 0 {
    fields = {{ $const }}.Columns
		}
		_ = m.SetFields(input, fields...)
		return {{ $receiver }}
    }

    // Set{{ .Name }}SkipZero sets the {{ .Name }} fields from input struct.
    // If no fields are specified, all fields from Columns will be set.
    // Zero values are skipped and not updated.
    func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}SkipZero(input *{{ .Name }}, fields ...string) *{{ $builder }} {
		m := {{ $receiver }}.mutation
		if len(fields) == 0 {
    fields = {{ $const }}.Columns
		}
		_ = m.SetFieldsSkipZero(input, fields...)
		return {{ $receiver }}
    }

{{- end -}}
