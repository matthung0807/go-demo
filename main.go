package main

import (
	"os"
	"text/template"
)

type AddValues struct {
	X int
	Y int
}

var funcMap = template.FuncMap{
	"add": add,
}

func add(a, b int) int {
	return a + b
}

func main() {
	text := "{{.X}} + {{.Y}} = {{add .X .Y}}\n" // template content
	t := template.Must(template.New("demo").Funcs(funcMap).Parse(text))

	data := AddValues{
		X: 1,
		Y: 2,
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
