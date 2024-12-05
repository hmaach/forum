package config

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../server/database/database.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign keys:", err)
	}

	return db, nil
}
