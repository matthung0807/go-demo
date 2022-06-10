package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			file, header, err := r.FormFile("file") // get form file input
			if err != nil {
				panic(err)
			}
			defer file.Close()

			buf := bytes.NewBuffer(nil)
			_, err = io.Copy(buf, file)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(header.Filename, buf.Bytes(), 0666)
			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, fmt.Sprintf("%v uploaded", header.Filename))

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
