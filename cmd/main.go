package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"forum/server/database"
	"forum/server/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	// Handle command-line flags
	if len(os.Args) > 1 {
		if err := utils.HandleFlags(os.Args[1:], db); err != nil {
			fmt.Println(err)
			utils.Usage()
			os.Exit(1)
		}
		return
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: routes(db),
	}

	// Start the HTTP server
	log.Println("Server starting on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
