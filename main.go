package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Employee struct {
	ID        int64     // primary key, column name is `id`"
	Name      string    // column name is `name`"
	Age       int       // column name is `age`"
	CreatedAt time.Time // column name is `created_at`"
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

	emp := Employee{}
	db.First(&emp) // SELECT * FROM employee ORDER BY id LIMIT 1;

	fmt.Println(emp) // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}
}
