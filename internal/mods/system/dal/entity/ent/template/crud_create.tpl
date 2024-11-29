{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "create/additional/crud" }}

    {{ $builder := .CreateName }}
    {{ $receiver := .CreateReceiver }}
    {{ $fields := .Fields }}
    {{- if .ID.UserDefined }}
        {{ $fields = append $fields .ID }}
    {{- end }}

    {{ print "// Set" .Name " set the " .Name }}
		func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
		m := {{ $receiver }}.mutation
		_ = m.SetFields(input, fields...)
		return {{ $receiver }}
		}

{{- end -}}
