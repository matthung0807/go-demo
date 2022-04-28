package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req HelloRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s\n", req.Name)
		rw.Write([]byte("success"))
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
