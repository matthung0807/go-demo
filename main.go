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
		ID:   1,
		Name: "john",
		Age:  28,
	}

	err := UpdateEmployee(db, &emp)
	if err != nil {
		panic(err)
	}

	emp.Age = 33
	id, err := UpdateEmployeeReturnID(db, &emp)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

func UpdateEmployee(db *sql.DB, emp *Employee) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `UPDATE employee
			SET 
				name = $1, 
                age = $2
			WHERE id = $3`

	_, err = tx.Exec(query, &emp.Name, &emp.Age, &emp.ID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateEmployeeReturnID(db *sql.DB, emp *Employee) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `UPDATE employee
			SET 
				name = $1, 
                age = $2
			WHERE id = $3
			RETURNING id`

	id := int64(0)
	err = tx.QueryRow(query, &emp.Name, &emp.Age, &emp.ID).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}
