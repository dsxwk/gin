package model{{ if .Imports }}

{{ .Imports }}{{ end }}

const {{ .TableConst }} = "{{ .Table }}"

type {{ .Struct }} struct {
{{- range .Fields }}
	{{ . }}
{{- end }}
}

func (*{{ .Struct }}) TableName() string {
	return {{ .TableConst }}
}
