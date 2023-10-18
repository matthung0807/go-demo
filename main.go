package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"os"
)

type Table struct {
	Name      string
	Employees []Employee
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

//go:embed template
var htmlTemplate string

func main() {
	t := template.Must(template.New("htmlTemplate").Parse(htmlTemplate))

	employees := []Employee{
		{"1", "John", 33},
		{"2", "Mary", 28},
		{"3", "Tony", 44},
		{"4", "Bill", 22},
	}

	data := Table{
		Name:      "Employee List",
		Employees: employees,
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
		panic(err)
	}

	os.WriteFile("result.html", buf.Bytes(), 0666)

}
