// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Employee struct {
	ID        int64
	Name      string
	Age       sql.NullInt32
	CreatedOn time.Time
}
