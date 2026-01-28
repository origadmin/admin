{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "update/additional/crud/update" }}

    {{ $builder := .UpdateName }}
    {{ $receiver := receiver $builder }}
    {{ $fields := .Fields }}
    {{- if or (hasSuffix $builder "Update") (hasSuffix $builder "UpdateOne") }}
        {{ $fields = .MutableFields }}
    {{- end }}

    // Set{{ .Name }} sets the {{ .Name }} fields from input struct.
    // If no fields are specified, all mutable fields except ID will be set.
    // Zero values are included in the update.
    func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
    {{- $const := print .Package}}
    m := {{ $receiver }}.mutation
    if len(fields) == 0 {
    fields = {{ $const }}.OmitColumns({{ $const }}.FieldID)
    }
    _ = m.SetFields(input, fields...)
    return {{ $receiver }}
    }

    // Set{{ .Name }}SkipZero sets the {{ .Name }} fields from input struct.
    // If no fields are specified, all mutable fields except ID will be set.
    // Zero values are skipped and not updated.
    func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}SkipZero(input *{{ .Name }}, fields ...string) *{{ $builder }} {
    {{- $const := print .Package}}
    m := {{ $receiver }}.mutation
    if len(fields) == 0 {
    fields = {{ $const }}.OmitColumns({{ $const }}.FieldID)
    }
    _ = m.SetFieldsSkipZero(input, fields...)
    return {{ $receiver }}
    }
{{- end -}}
