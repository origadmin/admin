{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Type */}}

{{ define "query/additional/crud" }}

    {{ $pkg := .Package }}
    {{ $fields := .Fields }}
    {{ $builder := .QueryName }}
    {{ $receiver := receiver $builder }}
    {{ $selectBuilder := pascal .Name | printf "%sSelect" }}

    // Omit excludes the specified fields/columns from the query result.
    // By default, all fields are selected. Use this method to select only specific fields.
    //
    // Example:
    //  var v []struct {
    {{- range $f := $fields }}
	    //    {{ $f.StructField }} {{ $f.Type }} `{{ $f.StructTag }}`
    {{- end }}
    //  }
    //
    //  client.{{ pascal $.Name }}.Query().
    //    Omit(
    {{- range $f := $fields }}
	    //    {{ $pkg }}.{{ $f.Constant }},
    {{- end }}
    //    ).
    //    Scan(ctx, &v)
    func ({{ $receiver }} *{{ $builder }}) Omit(fields ...string) *{{ $selectBuilder }} {
		omits := make(map[string]struct{}, len(fields))
		for i := range fields {
    omits[fields[i]] = struct{}{}
		}
		for _, col := range {{ $pkg }}.Columns {
    if _, ok := omits[col]; !ok {
    {{ $receiver }}.ctx.Fields = append({{ $receiver }}.ctx.Fields, col)
    }
		}

		sbuild := &{{ $selectBuilder }}{ {{ $builder }}: {{ $receiver }} }
		sbuild.label = {{ $pkg }}.Label
		sbuild.flds, sbuild.scan = &{{ $receiver }}.ctx.Fields, sbuild.Scan
		return sbuild
    }

{{- end -}}
