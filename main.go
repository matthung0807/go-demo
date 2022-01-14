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
		log.Panic("open database error")
	}
	return db
}

func main() {
	db := connect()
	emps, err := GetAllEmployees(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(emps)

	emp, err := GetEmployeeByID(db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(*emp)
}

func GetAllEmployees(db *sql.DB) ([]Employee, error) {
	rows, err := db.Query("SELECT id, name, age, created_on FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emps []Employee
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Age, &e.CreatedOn)
		if err != nil {
			return nil, err
		}
		emps = append(emps, e)
	}
	return emps, nil
}

func GetEmployeeByID(db *sql.DB, id int64) (*Employee, error) {
	row := db.QueryRow("SELECT * FROM employee WHERE id = $1 LIMIT 1", id)
	var emp Employee
	err := row.Scan(
		&emp.ID,
		&emp.Name,
		&emp.Age,
		&emp.CreatedOn,
	)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}
