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

    {{ print "// Set" .Name " set the " .Name }}
		func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
		m := {{ $receiver }}.mutation
		if len(fields) == 0 {
		fields = {{ $const }}.Columns
		}
		_ = m.SetFields(input, fields...)
		return {{ $receiver }}
		}

    {{ print "// Set" .Name "WithZero set the " .Name }}
		func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}WithZero(input *{{ .Name }}, fields ...string) *{{ $builder }} {
		m := {{ $receiver }}.mutation
		if len(fields) == 0 {
		fields = {{ $const }}.Columns
		}
		_ = m.SetFieldsWithZero(input, fields...)
		return {{ $receiver }}
		}

{{- end -}}
