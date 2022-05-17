package main

import (
	"context"
	"database/sql"
	"fmt"

	sqlcDb "abc.com/demo/db"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	DATABASE = "postgres"
	USER     = "admin"
	PASSWORD = "12345"
	SSL      = "disable"
)

func OpenDB(ctx context.Context) *sql.DB {
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
	db := OpenDB(ctx)
	queries := sqlcDb.New(db)
	employees, err := queries.GetAll(ctx)
	if err != nil {
		panic("query employee error")
	}
	fmt.Println(employees)
}
