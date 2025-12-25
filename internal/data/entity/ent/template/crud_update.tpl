{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "update/additional/crud/update" }}

    {{ $builder := .UpdateName }}
    {{ $receiver := receiver $builder }}
    {{ $fields := .Fields }}
    {{- if or (hasSuffix $builder "Update") (hasSuffix $builder "UpdateOne") }}
        {{ $fields = .MutableFields }}
    {{- end }}

    {{ print "// Set" .Name " set the " .Name ". This method includes zero values in the update." }}
    func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
        {{- $const := print .Package}}
        m := {{ $receiver }}.mutation
        if len(fields) == 0 {
            fields =  {{$const}}.OmitColumns({{$const}}.FieldID)
        }
        _ = m.SetFields(input, fields...)
        return {{ $receiver }}
    }

    {{ print "// Set" .Name "SkipZero set the " .Name ", skipping zero values." }}
    func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}SkipZero(input *{{ .Name }}, fields ...string) *{{ $builder }} {
        {{- $const := print .Package}}
        m := {{ $receiver }}.mutation
        if len(fields) == 0 {
            fields =  {{$const}}.OmitColumns({{$const}}.FieldID)
        }
        _ = m.SetFieldsSkipZero(input, fields...)
        return {{ $receiver }}
    }
{{- end -}}
