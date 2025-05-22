{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Type*/}}

{{ define "where/additional/with" }}
{{/*    {{- $type := $.Name }}*/}}
{{/*    {{- range $edge := $.Edges }}*/}}
{{/*        {{- if $edge.StructField }}*/}}
{{/*            {{ $func := print "With" $edge.StructField }}*/}}
{{/*						// With{{ $edge.StructField }} tells the query-builder to eager-load the nodes that are connected to*/}}
{{/*						// the "{{ $edge.StructField }}" edge. The optional arguments are used to configure the query builder of the edge.*/}}
{{/*						func {{$func}}(query *{{ $edge.StructField }}Query) {*/}}
{{/*						query.{{$func}}()*/}}
{{/*						}*/}}
{{/*        {{- end }}*/}}
{{/*    {{- end }}*/}}
{{ end }}