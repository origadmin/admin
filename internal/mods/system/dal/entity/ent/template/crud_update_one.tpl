{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "update/additional/crud_one" }}
    {{ $builder := $.UpdateOneName }}
    {{- if hasSuffix $builder "UpdateOne" }}
        {{ $receiver := receiver $builder }}
        {{ print "// Set" .Name " set the " .Name }}
				func ({{ $receiver }} *{{ $builder }}) Set{{ .Name }}(input *{{ .Name }}, fields ...string) *{{ $builder }} {
				m := {{ $receiver }}.mutation
				_ = m.SetFields(input, fields...)
				return {{ $receiver }}
				}

        {{ $onebuilder := $.UpdateOneName }}
        {{ $receiver = receiver $onebuilder }}
				// Omit allows the unselect one or more fields/columns for the given query,
				// instead of selecting all fields in the entity.
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
