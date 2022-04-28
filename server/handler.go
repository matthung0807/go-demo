package server

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WebhooksHandler struct {
	WebhooksService *WebhooksService
}

func NewWebhooksHandler(ws *WebhooksService) *WebhooksHandler {
	return &WebhooksHandler{
		WebhooksService: ws,
	}
}

//go:embed resources/events.yaml
var events string

func (wh *WebhooksHandler) Register(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req WebhooksRegisterRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}
		wh.WebhooksService.Save(req)
		fmt.Fprint(rw, events)
	case http.MethodGet:
		json.NewEncoder(rw).Encode(wh.WebhooksService.GetRegisteredUrls())
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (wh *WebhooksHandler) Greeting(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Good day")
		dtos := wh.WebhooksService.GetRegisteredDtos()
		if len(dtos) == 0 {
			fmt.Println("webhooks has no registered urls")
			return
		}

		for _, dto := range wh.WebhooksService.GetRegisteredDtos() {
			data := fmt.Sprintf("{\"name\": \"%s\"}", dto.Name)
			resp, err := http.Post(dto.URL, "application/json", bytes.NewBuffer([]byte(data)))
			if err != nil {
				fmt.Printf("send event failed")
			}

			b, err := io.ReadAll(resp.Body)
			msg := string(b)
			if msg == "success" {
				fmt.Println("send event success")
			}
		}
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
