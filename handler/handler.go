package handler

import (
	"log"
	"net/http"
	"time"

	"abc.com/demo/mq"
	"abc.com/demo/ws"
)

type NotificationHandler struct {
	notifier ws.Notifier
}

func NewNotificationHandler(notifier ws.Notifier) NotificationHandler {
	return NotificationHandler{
		notifier: notifier,
	}
}

func (h *NotificationHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userId")
		if userId == "" {
			log.Print("refuse websocket connection, userId is required")
			return
		}

		wsConn, err := ws.OpenConn(w, r)
		if err != nil {
			log.Printf("open websocket connection error, err=%v", err)
		}

		h.notifier.Register(userId, wsConn)
		defer h.notifier.Unregister(userId)

		ticker := time.NewTicker(time.Second * 5)
		for {
			<-ticker.C
			_, _, err = wsConn.ReadMessage()
			if err != nil {
				log.Printf("read message from websocket client error, err=%v", err)
				break
			}
		}
	}
}

func (h *NotificationHandler) MessageHandler() mq.MessageHandler {
	return func(userId, message string) error {
		return h.notifier.Notify(userId, message)
	}
}
