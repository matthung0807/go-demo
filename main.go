package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func main() {
	ctx := context.Background()
	driver := "postgres"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, PASSWORD, DATABASE, SSL)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Panic("open database error")
	}

	queries := sqlcDb.New(db)
	employees, err := queries.GetAll(ctx)
	if err != nil {
		log.Panic("query employee error")
	}
	fmt.Println(employees)

}
