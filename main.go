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

	// select with where conidtions
	emp := Employee{}                       // use struct model for single result
	db.Where("name = ?", "john").Find(&emp) // SELECT * FROM employee WHERE name = 'john';
	fmt.Println(emp)

	db.Where("name = ? AND age = ?", "john", 33).Find(&emp) // SELECT * FROM employee WHERE name = 'john' AND age = 33;
	fmt.Println(emp)                                        // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	// use map as where conditions
	conditionMap := map[string]interface{}{
		"name": "john",
	}
	db.Where(conditionMap).Find(&emp) // SELECT * FROM employee WHERE name = 'john';
	fmt.Println(emp)                  // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	// use struct model as where conditions
	conditionStruct := Employee{
		Name: "john",
	}
	db.Where(conditionStruct).Find(&emp) // SELECT * FROM employee WHERE name = 'john';
	fmt.Println(emp)                     // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	// inline select with where conidtion
	db.Find(&emp, "name = ?", "john") // SELECT * FROM employee WHERE name = 'john';
	fmt.Println(emp)                  // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	db.Find(&emp, "name = ? AND age = ?", "john", 33) // SELECT * FROM employee WHERE name = 'john' AND age = 33;
	fmt.Println(emp)                                  // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	// inline select with map conditions
	db.Find(&emp, conditionMap) // SELECT * FROM employee WHERE name = 'john';
	fmt.Println(emp)            // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	// inline select with struct model conditions
	db.Find(&emp, conditionStruct) // SELECT * FROM employee WHERE name = 'john';
	fmt.Println(emp)               // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	// select with where condition and primary key
	emp = Employee{ID: 2}                   // model has primary key field value
	db.Where("name = ?", "john").Find(&emp) // SELECT * FROM employee WHERE name = 'john' AND id = 2 ORDER BY id LIMIT 1;
	fmt.Println(emp)                        // {1 john 33 2022-11-29 18:44:54.114161 +0000 UTC}

	// select multiple results
	var emps []Employee                 // use model slice for multiple results
	db.Where("age > ?", 30).Find(&emps) // SELECT * FROM employee WHERE age > 30;
	fmt.Println(emps)                   // [{3 tony 45 2022-12-22 22:05:09.83327 +0000 UTC} {1 john 33 2022-12-22 21:56:37.061419 +0000 UTC}]

	// select with where between
	db.Where("age BETWEEN ? AND ? ", 18, 30).Find(&emps) // SELECT * FROM employee WHERE age BETWEEN 18 AND 30;
	fmt.Println(emps)                                    // [{2 mary 28 2022-12-22 21:56:37.061419 +0000 UTC}]

	// select with where order by
	db.Where("age > ?", 30).Order("age asc").Find(&emps) // SELECT * FROM employee WHERE age > 30 ORDERY BY age DESC;
	fmt.Println(emps)                                    // [{1 john 33 2022-12-22 21:56:37.061419 +0000 UTC} {3 tony 45 2022-12-22 22:05:09.83327 +0000 UTC}]

}
