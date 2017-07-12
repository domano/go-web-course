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
Age:    {{.Age}}
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
		Age    int
		IsAdmin bool
	}{
		{"Dino Omanovic", 27, true},
		{"Max Mustermann", 54, false},
	}

	t := template.New("template")
	t.Funcs(functionMap)

	template.Must(t.Parse(userTemplate))

	err := t.Execute(os.Stdout, users)
	if err != nil {
		panic(err)
	}
}
