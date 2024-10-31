package main

import (
	"log"
	"net/http"
	// _sqlite "github.com/mattn/go-sqlite3"
)

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	log.Println("Server starting on http://localhost:8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
