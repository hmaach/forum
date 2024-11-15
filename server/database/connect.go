package database

import (
	"database/sql"
	"fmt"
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
	return db, nil
}
