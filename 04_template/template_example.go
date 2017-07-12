package main

import (
	"os"
	"text/template"
)

var theTemplate = `
{{. | len}} Personen
{{ range . }}
{{.Given}} {{.Name}}
=================
Name:   {{.Name}}
Given:  {{.Given}} 
Age:    {{.Age}}
Gender: {{ .Male | genderToText }}
State: {{ if lt .Age 18 }}Child{{ else }}Adult{{end}}
{{end}}
`

var functionMap = template.FuncMap{
	"genderToText": func(male bool) string {
		if male {
			return "M"
		} else {
			return "F"
		}
	},
}

func main() {

	data := []struct {
		Name  string
		Given string
		Age   int
		Male  bool
	}{
		{"Sebastian", "Mancke", 37, true},
		{"Sabrina", "Mancke", 37, false},
		{"Felix", "Mancke", 12, true},
		{"Nils", "Mancke", 9, true},
		{"Emils", "Mancke", 6, true},
		{"Linus", "Mancke", 2, true},
	}

	t := template.New("template")
	t.Funcs(functionMap)

	template.Must(t.Parse(theTemplate))

	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
