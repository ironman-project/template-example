This file shows how to use a map field values

Managers using an explicit reference

Infrastructure: {{.Values.projectManagers.infrastructure}}
Platform:      {{.Values.projectManagers.platform}}
Tooling:       {{.Values.projectManagers.tooling}}


Managers using  a range

{{range $key, $element := .Values.projectManagers}}
{{ $key }}:  {{ $element -}}
{{ end }}
