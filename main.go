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
	emps, err := GetAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(emps)
}

func GetAll(ctx context.Context) ([]sqlcDb.Employee, error) {
	db := OpenDB(ctx)
	queries := sqlcDb.New(db)
	emps, err := queries.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return emps, nil
}

func GetById(ctx context.Context, id int64) (*sqlcDb.Employee, error) {
	db := OpenDB(ctx)
	qr := sqlcDb.New(db)
	emp, err := qr.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func Insert(ctx context.Context, name string, age int) error {
	db := OpenDB(ctx)
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	params := sqlcDb.InsertParams{
		Name: name,
		Age:  sql.NullInt32{int32(age), true},
	}
	qr := sqlcDb.New(tx)
	err = qr.Insert(ctx, params)
	if err != nil {
		return err
	}

	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func Update(ctx context.Context, id int64, name string, age int) (int64, error) {
	db := OpenDB(ctx)
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	params := sqlcDb.UpdateParams{
		ID:   id,
		Name: name,
		Age:  sql.NullInt32{int32(33), true},
	}
	qr := sqlcDb.New(tx)
	rows, err := qr.Update(ctx, params)
	if err != nil {
		return 0, err
	}

	tx.Commit()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
