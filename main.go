package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// employee.contact要轉換的型態
type Contact struct {
	Name  string
	Phone string
}
type Employee struct {
	ID        int64
	Name      string
	Age       int
	Contact   Contact // mapping to column `employee.contact`
	CreatedAt time.Time
}

// column type to field type
//
// 讀取時employee.contact(jsonb) -> Employee.Contact
func (c *Contact) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to unmarshal JSONB value=%v", value)
	}
	err := json.Unmarshal(bytes, c)
	return err
}

// field type to column type
//
// 插入時Employee.Contact -> employee.contact(jsonb)
func (c Contact) Value() (driver.Value, error) {
	json.Marshal(c)
	return json.Marshal(c)
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

	newEmp := Employee{
		Name: "john",
		Age:  33,
		Contact: Contact{
			Name:  "mary",
			Phone: "0912345678",
		},
	}

	db.Create(&newEmp)

	emp := Employee{}
	db.Last(&emp)

	fmt.Println(emp) // {5 tony 45 {mary 0912345678} 2022-12-02 16:53:35.170276 +0000 UTC}
}
