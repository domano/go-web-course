package main

import (
	"os"
	"html/template"
)

var htmlUserTemplate = `
<html>
<body>
<h1>{{. | len}} User(s)</h1>
{{if .}}
<table>
<tr>
	<th>Name</th>
	<th>Status</th>
</tr>
{{ range . }}
<tr>
	<td>{{.Name}}</td>
	<td>{{ .IsAdmin | showStatus}}</td>
</tr>
</table>
{{end}}
{{end}}
<body>
</html>
`

var htmlFunctionMap = template.FuncMap{
	"showStatus": func(isAdmin bool) string {
		if isAdmin {
			return "Admin"
		} else {
			return "User"
		}
	},
}

func main() {

	users := []struct {
		Name    string
		IsAdmin bool
	}{
		{"Dino Omanovic", true},
		{"Max Mustermann", false},
		{"Herr Bärmann<script>alert('schlimm!');</script>",false},
	}

	t := template.New("template")
	t.Funcs(htmlFunctionMap)

	template.Must(t.Parse(htmlUserTemplate))

	err := t.Execute(os.Stdout, users)
	if err != nil {
		panic(err)
	}
}
