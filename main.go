package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Employee struct {
	ID        int64     // primary key, column name is `id`
	Name      string    // column name is `name`
	Age       int       // column name is `age`
	CreatedAt time.Time // column name is `created_at`
}

const (
	HOST     = "localhost"
	PORT     = "5432"
	DATABASE = "postgres"
	USER     = "admin"
	PASSWORD = "12345"
	SSL      = "disable"
)

func getGormDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic("open gorm db error")
	}

	return gormDB
}

func main() {
	db := getGormDB()

	emp := Employee{
		ID: 1, // primary key
	}

	result := db.Model(&emp).Update("age", "34") // UPDATE employee SET age = 34 WHERE id = 1;
	fmt.Println(result.Error)                    // nil
	fmt.Println(result.RowsAffected)             // 1

	db.First(&emp)   // SELECT * FROM employee WHERE id = 1;
	fmt.Println(emp) // {1 john 34 2022-12-22 21:56:37.061419 +0000 UTC}
}
