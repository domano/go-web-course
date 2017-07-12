package main

import (
	"os"
	"text/template"
)

var userTemplate = `
{{. | len}} User
{{ range . }}
=================
Name:   {{.Name}}
Status: {{ .IsAdmin | showStatus}}
=================
{{end}}
`

var functionMap = template.FuncMap{
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
		Name   string
		IsAdmin bool
	}{
		{"Dino Omanovic", true},
		{"Max Mustermann", false},
	}

	t := template.New("template")
	t.Funcs(functionMap)

	template.Must(t.Parse(userTemplate))

	err := t.Execute(os.Stdout, users)
	if err != nil {
		panic(err)
	}
}
