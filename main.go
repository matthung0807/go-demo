package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

func connect() *sql.DB {
	driver := "postgres"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Println(err)
	}
	return db
}

func main() {
	db := connect()
	if err := createEmployee(db, "john", 33); err != nil {
		fmt.Println(err)
	}
}

func createEmployee(db *sql.DB, name string, age int) error {
	tx, err := db.Begin() // 開始一個交易

	if err != nil {
		return err
	}
	defer tx.Rollback() // 確保錯誤發生時rollback.

	sql := "INSERT INTO employee (name, age, created_on) VALUES ($1, $2, $3)"

	_, err = tx.Exec(sql, "john", 33, time.Now()) // 執行DML sql並帶入參數
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil { // 提交
		return err
	}

	return nil
}
