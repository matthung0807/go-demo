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
	http.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // get a websocket connettion
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage() // read a amessage from client
		if err != nil {
			panic(err)
		}

		fmt.Printf("receive: %s", message)
		err = conn.WriteMessage(mt, message) // write a message to client
		if err != nil {
			panic(err)
		}
	}
}
