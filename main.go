package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseForm()
			name := r.Form.Get("name")
			email := r.Form.Get("email")
			age := r.Form.Get("age")
			birthday := r.Form.Get("birthday")
			gender := r.Form.Get("gender")
			langs := r.Form["lang"]

			fmt.Fprint(w, fmt.Sprintf(
				"name=%v\nemail=%v\nage=%v\nbirtyday=%v\ngender=%v\nlanguages=%v",
				name, email, age, birthday, gender, langs))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
