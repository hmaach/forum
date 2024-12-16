package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

// CreateTables executes all queries from schema.sql
func CreateTables(db *sql.DB) error {
	// read file that contains all queries  to create tables for database schema
	content, err := os.ReadFile(BasePath + "server/database/sql/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema.sql file: %v", err)
	}

	queries := strings.TrimSpace(string(content))

	// execute all queries to create database schema
	_, err = db.Exec(queries)
	if err != nil {
		log.Printf("failed to create tables %q: %v\n", queries, err)
		return err
	}
	log.Println("Database schema created successfully")
	return nil
}

// CreateFakeData generates and inserts fake data into the database
func CreateDemoData(db *sql.DB) error {
	// create database schema before creating demo data
	if err := CreateTables(db); err != nil {
		return err
	}

	// read file that contains all queries  to create demo data
	content, err := os.ReadFile(BasePath + "server/database/sql/seed.sql")
	if err != nil {
		return fmt.Errorf("failed to read seed.sql file: %v", err)
	}

	queries := strings.TrimSpace(string(content))

	// execute all queries to create demo data in the database
	_, err = db.Exec(queries)
	if err != nil {
		log.Printf("failed to isert demo data %q: %v\n", queries, err)
		return err
	}

	log.Println("Demo data created successfully")
	return nil
}

// Drop all tables in the database.
func Drop() error {
	err := os.Remove(BasePath + "server/database/database.db")
	if err != nil {
		log.Printf("failed to drop tables: %v\n", err)
		return err
	}

	log.Println("Database schema dropped successfully")
	return nil
}
