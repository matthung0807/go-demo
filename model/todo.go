package model

import "time"

type Todo struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
