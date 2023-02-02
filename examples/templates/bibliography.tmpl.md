---
sidebar_position: 1000
---

# 📚 Bibliography

The page contains an organized list of all papers used by this course.
The papers are organized by topic.

**To cite this course, use the provided citation in the Github repository.**

🔵 = Paper directly cited in this course. Other papers have informed my understanding of the topic.

Note: since [neither the GPT-3 nor the GPT-3 Instruct paper correspond to davinci models](https://twitter.com/janleike/status/1584618242756132864), I attempt not to
cite them as such.

{{ $sections := list "Prompt Engineering Strategies" }}

{{ range $section := $sections -}}
## {{ $section }}
{{ range $row := $.rows -}}
{{ if eq $row.lpSection $section }}
#### {{ $row.title }}(@{{ (index $row.keys 0) }}) {{ if $row.directlyCited }}🔵{{ end }}
{{ end }}
{{- end }}
{{- end }}
