package main

import (
	"os"
	"text/template"
)

type Employee struct {
	Name string
	Age  int
}

func main() {
	text := `Employees:
{{range .}}    Name:{{.Name}}, Age:{{.Age}}
{{end}}`
	t := template.Must(template.New("demo").Parse(text)) // 解析模板內容來建立名稱為'demo'的模板

	data := []Employee{
		{"John", 33},
		{"Mary", 28},
	}
	err := t.Execute(os.Stdout, data) // 將資料物件套用在模板，並將結果輸出到標準輸出
	if err != nil {
		panic(err)
	}

}
