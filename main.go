package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID        int64
	Name      string
	Age       int
	CreatedOn time.Time
}

const (
	HOST     = "localhost"
	PORT     = "5432"
	DATABASE = "postgres"
	USER     = "admin" // 使用者名稱
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
		log.Panic("open database error")
	}
	return db
}

func main() {
	db := connect()

	rows, err := db.Query("SELECT id, name, age, created_on FROM employee")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Age, &e.CreatedOn)
		if err != nil {
			panic("scan error")
		}
		employees = append(employees, e)
	}

	fmt.Println(employees)
}
