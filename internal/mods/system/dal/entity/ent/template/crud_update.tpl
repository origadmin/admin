{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "update/additional/crud/update" }}

    {{ $builder := .UpdateName }}
    {{ $receiver := receiver $builder }}
    {{ $fields := .Fields }}
    {{- if or (hasSuffix $builder "Update") (hasSuffix $builder "UpdateOne") }}
        {{ $fields = .MutableFields }}
    {{- end }}

    {{ print "// Set" .Name " set the " .Name }}
		func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
		m := {{ $receiver }}.mutation
		_ = m.SetFields(input, fields...)
		return {{ $receiver }}
		}

{{- end -}}
