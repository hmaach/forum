package utils

import (
	"database/sql"
	"fmt"
	"slices"

	"forum/server/config"
)

var ValidFlags = []string{"--migrate", "--seed", "--drop"}

func HandleFlags(flags []string, db *sql.DB) error {
	if len(flags) != 1 {
		return fmt.Errorf("expected a single flag, got %d", len(flags))
	}

	flag := flags[0]
	if !slices.Contains(ValidFlags, flag) {
		return fmt.Errorf("invalid flag: '%s'", flag)
	}

	switch flag {
	case "--migrate":
		return config.CreateTables(db)
	case "--seed":
		return config.CreateDemoData(db)
	case "--drop":
		return config.Drop()
	}
	return nil
}

func Usage() {
	fmt.Println(`Usage: go run main.go [option]
Options:
  --migrate   Create database tables
  --seed      Insert demo data into the database
  --drop      Drop all tables`)
}
