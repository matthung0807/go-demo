package main

import (
	"os"
	"text/template"
)

func main() {
	text := "hello {{.}}"                                // 模板內容。{{.}}會插入資料物件
	t := template.Must(template.New("demo").Parse(text)) // 解析模板內容來建立名稱為'demo'的模板

	data := "john"
	err := t.Execute(os.Stdout, data) // 將資料物件套用在模板，並將結果輸出到標準輸出
	if err != nil {
		panic(err)
	}
}
