package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Gambitier/grandchattutorial/internal/server"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	srv := server.New(db)
	srv.SetupRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
