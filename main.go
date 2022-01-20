package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID   int64
	Name string
	Age  int
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
	ctx := context.Background()
	db := connect()
	emps, err := GetAllEmployees(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Println(emps)
}

func GetAllEmployees(ctx context.Context, db *sql.DB) ([]Employee, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 3*time.Second) // create timeout Context
	defer cancel()                                              // make sure release resources held by Context when exiting function

	rows, err := db.QueryContext(queryCtx, "SELECT *, pg_sleep(5) FROM employee")
	if err != nil {
		fmt.Println("query timeout")
		return nil, err
	}
	defer rows.Close()

	var emps []Employee
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Age)
		if err != nil {
			return nil, err
		}
		emps = append(emps, e)
	}
	return emps, nil
}
