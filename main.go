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

			body := new(bytes.Buffer) // the body content writed by multipart.Writer
			mw := multipart.NewWriter(body)

			// wirte field "name"
			fw, err := mw.CreateFormField("name")
			if err != nil {
				panic(err)
			}
			_, err = fw.Write([]byte("John"))
			if err != nil {
				panic(err)
			}

			// write field "age"
			fw, err = mw.CreateFormField("age")
			if err != nil {
				panic(err)
			}
			_, err = fw.Write([]byte("33"))
			if err != nil {
				panic(err)
			}

			// write file field "file"
			fw, err = mw.CreateFormFile("file", "gopher.jpg")
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
			mw.Close() // close multipart.Writer

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

			name := r.Form.Get("name")
			age := r.Form.Get("age")

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

			fmt.Printf("name=%v\nage=%v\nfilename=%v\n", name, age, filename)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "%v uploaded", filename)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
