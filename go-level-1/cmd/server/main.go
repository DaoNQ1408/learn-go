package main

import (
	"database/sql"
	"log"
	"net/http"

	"com/daonq1408/learn-go/internal/store"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string: user=postgres password=yourpassword dbname=yourdb sslmode=disable
	connStr := "postgres://postgres:12345678@localhost/goLevel1?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	storage := store.NewStorage(db)
	if err := storage.InitDB(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
