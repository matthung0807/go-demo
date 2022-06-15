package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			body := new(bytes.Buffer)
			mw := multipart.NewWriter(body)

			fw, err := mw.CreateFormFile("file", "gopher.jpg")
			if err != nil {
				panic(err)
			}

			f, err := os.Open("./static/image/gopher.jpg")
			if err != nil {
				panic(err)
			}
			defer f.Close()

			_, err = io.Copy(fw, f)
			if err != nil {
				panic(err)
			}
			mw.Close()

			req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
			if err != nil {
				panic(err)
			}
			req.Header.Set("Content-Type", mw.FormDataContentType())

			client := &http.Client{}
			_, err = client.Do(req)
			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, "success")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseMultipartForm(1 << 20) // 1MB

			file, header, err := r.FormFile("file")
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
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "%v uploaded", filename)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
