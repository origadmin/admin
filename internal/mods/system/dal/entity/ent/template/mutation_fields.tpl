{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "mutation_fields" }}
    {{- $pkg := base $.Config.Package -}}
    {{- template "header" $ -}}

    {{/* Additional dependencies injected to config. */}}
    {{ $deps := list }}{{ with $.Config.Annotations }}{{ $deps = $.Config.Annotations.Dependencies }}{{ end }}

		import (
		"log"

		"entgo.io/ent/dialect"

    {{- range $n := $.Nodes }}
        {{ $n.PackageAlias }} "{{ $n.Config.Package }}/{{ $n.PackageDir }}"
    {{- end }}
    {{- range $dep := $deps }}
        {{ $dep.Type.PkgName }} "{{ $dep.Type.PkgPath }}"
    {{- end }}
		"{{ $.Config.Package }}/migrate"
    {{- range $import := $.Storage.Imports }}
			"{{ $import }}"
    {{- end -}}
    {{- template "import/additional" $ }}
		)

    {{ range $n := $.MutableNodes }}
        {{ $fields := $n.Fields }}
        {{ $mutation := $n.MutationName }}
				// SetFields sets the values of the fields with the given names. It returns an
				// error if the field is not defined in the schema, or if the type mismatched the
				// field type.
				func (m *{{ $mutation }}) SetFields(input *{{ .Name }}, fields ...string) error {
				for i := range fields {
				switch fields[i] {
        {{- range $f := $fields }}
            {{- $const := print $n.Package "." $f.Constant }}
						case {{ $const }}:
            {{- if $f.Nillable}}
                {{"if"}} input.{{ $f.StructField }} {{"!= nil {"}}
                {{- $setter := print "Set" $f.StructField }}
                {{ $mutation }}.{{ $setter }}(*input.{{ $f.StructField }})
                {{"}"}}
            {{- else}}
                {{- if $f.IsTime }}
                    {{- $setter := print "Set" $f.StructField }}
                    {{"if !"}}input.{{ $f.StructField }}.IsZero() {{"{"}}
										m.{{ $setter }}(input.{{ $f.StructField }})
                    {{"}"}}
                {{- else }}
                    {{- $setter := print "Set" $f.StructField }}
										m.{{ $setter }}(input.{{ $f.StructField }})
                {{- end}}
            {{- end}}
        {{- end }}
				default:
				return fmt.Errorf("unknown {{ .Name }} field %s", fields[i])
				}
				}
				return nil
				}
    {{- end }}

{{ end }}

