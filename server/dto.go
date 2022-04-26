package server

import "time"

type WebhooksRegisterDto struct {
	Name      string
	URL       string
	CreatedAt time.Time
}
