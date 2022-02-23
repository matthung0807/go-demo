package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int64
	Name      string
	Age       int
	CreatedAt time.Time
}

const (
	HOST     = "localhost"
	PORT     = "5432"
	DATABASE = "postgres"
	USER     = "admin"
	PASSWORD = "12345"
	SSL      = "disable"
)

func connect() *sql.DB {
	driver := "postgres"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic("open database error")
	}
	return db
}

func main() {
	db := connect()
	emp := Employee{
		Name: "tony",
		Age:  23,
	}
	id, err := CreateEmployee(db, &emp)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

func CreateEmployee(db *sql.DB, emp *Employee) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	id := int64(0)
	query := `INSERT INTO employee (name, age, created_at) 
			VALUES ($1, $2, $3) 
			RETURNING id`

	err = tx.QueryRow(query, &emp.Name, &emp.Age, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}
