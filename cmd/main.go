package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"forum/server/config"
	"forum/server/routes"
	"forum/server/utils"

	_ "github.com/mattn/go-sqlite3"
)

func clearSession(db *sql.DB) {
	_, err := db.Exec("DELETE FROM sessions")
	if err != nil {
		fmt.Println("Error clearing sessions:", err)
	} else {
		fmt.Println("Sessions cleared successfully")
	}
}

func main() {
	// Connect to the database
	db, err := config.Connect()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	err = config.CreateTables(db)

	// 	for i := 1; i <= 50; i++ {
	// 		db.Exec(`INSERT INTO posts (user_id, title, content) VALUES
	// (1, ?, ?)`, "title "+strconv.Itoa(i), "content"+strconv.Itoa(i))
	// 		_, err := db.Exec(`INSERT INTO post_category (post_id, category_id) VALUES
	// (?,1)`, i)
	// 		fmt.Println(err, i)
	// 	}
	// Handle command-line flags
	if len(os.Args) > 1 {
		if err := utils.HandleFlags(os.Args[1:], db); err != nil {
			fmt.Println(err)
			utils.Usage()
			os.Exit(1)
		}
		return
	}

	// Fetch categories from the database to display in the navbar
	// common.Categories, err = models.FetchCategories(db)
	if err != nil {
		log.Println("Error fetching categories from the database:", err)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(db),
	}

	// Start the HTTP server
	log.Println("Server starting on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
