package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Employee struct {
	ID        uuid.UUID
	Name      string
	Age       int
	CreatedOn time.Time
}

func main() {
	db := connect()

	rows, err := db.Query("SELECT id, name, age, created_on FROM employee")
	if err != nil {
		panic(err)
	}

	employees := []Employee{}
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Age, &e.CreatedOn)
		if err != nil {
			panic("scan error")
		}
		employees = append(employees, e)
	}

	fmt.Println(employees) // f3f6148e-1ac3-4ae4-b96e-16f937393a10
}

const (
	HOST     = "localhost"
	PORT     = "5432"
	DATABASE = "postgres"
	USER     = "matt" // 使用者名稱
	SSL      = "disable"
)

func connect() *sql.DB {
	driver := "postgres"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, DATABASE, SSL)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Panic("open database error")
	}
	return db
}
