package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow CORS request's Origin header
	},
}

func main() {
	http.HandleFunc("/echo", echoHandler)   // websocket
	http.HandleFunc("/hello", helloHandler) // http

	http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		fmt.Printf("receive: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			panic(err)
		}
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	content := fmt.Sprintf("hello, %s", name)
	fmt.Fprint(w, content)
}
