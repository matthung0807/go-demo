package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connect()

	id := 1
	rows, err := db.Query("SELECT * FROM employee where id=?", id)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		fmt.Printf("id=%d, name=%s, age=%d\n", id, name, age)
	}
}

func connect() *sql.DB {
	driver := "mysql"
	user := "mysqladmin"
	password := "mysqladmin"
	endpoint := "rds-mysql-free-001.c9u62rqfjvs8.us-east-2.rds.amazonaws.com"
	port := "3306"
	dbName := "mydb"
	charset := "charset=utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, endpoint, port, dbName, charset)
	fmt.Println(dsn) // mysqladmin:mysqladmin@tcp(rds-mysql-free-001.c9u62rqfjvs8.us-east-2.rds.amazonaws.com:3306)/mydb?charset=utf8mb4

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
