package server

import (
	"sync"
	"time"
)

type WebhooksService struct {
	registered map[string]WebhooksRegisterDto
	rwmu       sync.RWMutex
}

func NewWebhooksService() *WebhooksService {
	return &WebhooksService{
		registered: make(map[string]WebhooksRegisterDto),
		rwmu:       sync.RWMutex{},
	}
}

func (ws *WebhooksService) Save(req WebhooksRegisterRequest) {
	ws.rwmu.Lock()
	defer ws.rwmu.Unlock()
	ws.registered[req.URL] = WebhooksRegisterDto{
		Name:      req.Name,
		URL:       req.URL,
		CreatedAt: time.Now(),
	}
}

func (ws *WebhooksService) GetRegisteredDtos() []WebhooksRegisterDto {
	ws.rwmu.RLock()
	defer ws.rwmu.RUnlock()
	dtos := make([]WebhooksRegisterDto, 0)
	for _, v := range ws.registered {
		dtos = append(dtos, v)
	}
	return dtos
}

func (ws *WebhooksService) GetRegisteredUrls() []string {
	urls := make([]string, 0)
	for _, dto := range ws.GetRegisteredDtos() {
		urls = append(urls, dto.URL)
	}
	return urls
}
