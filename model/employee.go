package model

import "time"

type Employee struct {
	ID        int64
	Name      string
	Age       int
	CreatedAt time.Time
}
