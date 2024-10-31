package main

import (
	"log"
	"net/http"

	"forum/server/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, dbErr := database.Connect()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	srv := http.Server{
		Addr:    ":8080",
		Handler: routes(db),
	}

	log.Println("Server starting on http://localhost:8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
