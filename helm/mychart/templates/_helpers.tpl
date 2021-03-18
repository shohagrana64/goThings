{{- define "labels" -}}
chart-name: {{ quote .Chart.Name}}
release-name: {{ .Release.Name | quote }}
{{- range $key, $val := .Values.labels }}
{{ $key }}: {{ $val | quote }}
{{- end }}
{{- end }}