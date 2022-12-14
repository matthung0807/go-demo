package ws

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Notifier interface {
	Register(userId string, conn *websocket.Conn)
	Unregister(userId string) error
	Notify(userId, message string) error
}

type WebSocketManger struct {
	connMap map[string]*websocket.Conn
	rwmu    sync.RWMutex
}

func NewWebSocketManager() WebSocketManger {
	return WebSocketManger{
		connMap: make(map[string]*websocket.Conn),
	}
}

func (wm *WebSocketManger) Register(userId string, conn *websocket.Conn) {
	wm.rwmu.Lock()
	defer wm.rwmu.Unlock()
	wm.connMap[userId] = conn
	log.Printf("userId=[%s] websocket connection registered", userId)
}

func (wm *WebSocketManger) Unregister(userId string) error {
	wm.rwmu.Lock()
	defer wm.rwmu.Unlock()
	conn, ok := wm.connMap[userId]
	if !ok {
		return errors.New("unregister connection failed")
	}
	wm.connMap[userId] = nil
	log.Printf("userId=[%s] websocket connection unregistered", userId)
	return conn.Close()
}

func (wm *WebSocketManger) GetConn(userId string) *websocket.Conn {
	wm.rwmu.RLock()
	defer wm.rwmu.RUnlock()

	conn, ok := wm.connMap[userId]
	if !ok {
		return nil
	}
	return conn
}

func (wm *WebSocketManger) Notify(userId, message string) error {
	conn := wm.GetConn(userId)
	if conn == nil {
		return fmt.Errorf("userId=[%s]'s websocket connection not found", userId)
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func OpenConn(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // diable CORS check
		},
	}

	wsConn, err := upgrader.Upgrade(w, r, nil) // get a websocket connection
	if err != nil {
		return nil, err
	}
	return wsConn, nil
}
