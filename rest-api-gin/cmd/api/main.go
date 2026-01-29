package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://postgres:12345678@localhost:5432/gin_db?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Optional: kiểm tra kết nối thật sự
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
