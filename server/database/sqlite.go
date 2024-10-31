package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

// CreateTables executes all queries from schema.sql
func createTables(db *sql.DB) error {
	content, err := os.ReadFile("../server/database/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema.sql file: %v", err)
	}

	queries := strings.TrimSpace(string(content))

	_, err = db.Exec(queries)
	if err != nil {
		log.Printf("failed to create tables %q: %v\n", queries, err)
		return err
	}

	return nil
}

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

	err1 := createTables(db)
	if err1 != nil {
		log.Fatal(err1)
	}
	return db, nil
}
