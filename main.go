package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseMultipartForm(1 << 20) // 1MB

			name := r.Form.Get("name")
			email := r.Form.Get("email")
			age := r.Form.Get("age")
			birthday := r.Form.Get("birthday")
			gender := r.Form.Get("gender")
			langs := r.Form["lang"]

			file, header, err := r.FormFile("photo")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			buf := bytes.NewBuffer(nil)
			_, err = io.Copy(buf, file)
			if err != nil {
				panic(err)
			}
			filename := header.Filename
			err = ioutil.WriteFile(filename, buf.Bytes(), 0666)
			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, fmt.Sprintf(
				"name=%v\nemail=%v\nage=%v\nbirtyday=%v\ngender=%v\nlanguages=%v\nphoto=%v",
				name, email, age, birthday, gender, langs, filename))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
