package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"forum/server/config"
	"forum/server/routes"
	"forum/server/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Check if running in Docker
	isDocker := os.Getenv("BASE_PATH") != ""
	if isDocker {
		config.BasePath = os.Getenv("BASE_PATH")
	}

	// Connect to the database
	db, err := config.Connect()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	// Handle database setup based on environment
	if isDocker {
		// Create the database schema and demo data
		err := config.CreateDemoData(db)
		if err != nil {
			log.Fatalf("Error creating the database schema and demo data: %v", err)
		}
		log.Println("Database setup complete.")
	} else {
		// Handle command-line flags for database setup
		if len(os.Args) > 1 {
			if err := utils.HandleFlags(os.Args[1:], db); err != nil {
				fmt.Println(err)
				utils.Usage()
				os.Exit(1)
			}
			return
		}
	}
	

	// Start the HTTP server
	server := http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(db),
	}

	log.Println("Server starting on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
