# This template showcases the list of Projects

## Using context

{{ range .Values.projects }}
  *  {{ . -}}
{{ end }}

## Declaring a variable value Name

{{range $element := .Values.projects}}
  *  {{ $element -}}
{{ end }}
