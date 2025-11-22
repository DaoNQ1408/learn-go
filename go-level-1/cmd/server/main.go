package main

import (
	"database/sql"
	"log"
	"net/http"

	"com/daonq1408/learn-go/internal/store"

	_ "github.com/lib/pq"
)

func main() {
	// 1. tạo connection string tới DB
	connStr := "postgres://postgres:12345678@localhost/goLevel1?sslmode=disable"

	// 2. open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 3. check connection
	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	// 4. create table if not yet created
	storage := store.NewStorage(db)
	if err := storage.InitDB(); err != nil {
		log.Fatal(err)
	}

	// 5. start server
	mux := http.NewServeMux()

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
