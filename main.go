package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow CORS request's Origin header
	},
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// websocket handler
func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // upgrade to a websocket connection
	if err != nil {
		panic(err)
	}

	go ping(conn)
	handlePong(conn)
	read(conn)
}

const (
	PING = "ping"
	PONG = "pong"

	pongWait   = 5 * time.Second
	pingPeriod = pongWait - 1
)

func ping(conn *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	fmt.Println("websocket connection openned")

	for range ticker.C {
		err := conn.WriteMessage(websocket.PingMessage, []byte(PING))
		if err != nil {
			return
		}
	}
}

func handlePong(conn *websocket.Conn) {
	conn.SetPongHandler(func(appData string) error {
		if appData == PING {
			fmt.Println(PONG)
		}
		return nil
	})
}

func read(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("websocket connection closed")
	}()

	for {
		_, message, err := conn.ReadMessage() // read message from client
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Printf("error: %v", err)
			}
			break
		}
		fmt.Printf("receive: %s\n", message)
	}
}
