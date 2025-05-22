{{/* The line below tells Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "crud" }}
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

    {{ range $n := $.Nodes }}
        {{- /* Support adding create methods by global templates. */}}
        {{- with $tmpls := matchTemplate "crud/helper/*" }}
            {{- range $tmpl := $tmpls }}
                {{ xtemplate $tmpl $n }}
            {{- end }}
        {{- end }}
    {{ end }}

{{ end }}


