package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("hi")

	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(20 * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
