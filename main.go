package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/gopher", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			/* 範例一 */
			// b, err := ioutil.ReadFile("./static/image/gopher.jpg")
			// if err != nil {
			// 	panic(err)
			// }
			// w.Header().Set("Content-Type", "image/jpeg")
			// w.Write(b)

			/* 範例二 */
			http.ServeFile(w, r, "./static/image/gopher.jpg")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	/* 範例三 */
	fs := http.FileServer(http.Dir("./static/image"))
	http.Handle("/image/", http.StripPrefix("/image/", fs))

	http.ListenAndServe(":8080", nil)
}
