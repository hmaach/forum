package models

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func StoreUser(db *sql.DB, email, username, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	task := `INSERT INTO users (email,username,password) VALUES (?,?,?)`
	result, err := db.Exec(task, email, username, hashedPassword)
	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}

	userID, _ := result.LastInsertId()

	return userID, nil
}

//////////////////////////////////////////////////////////

func AddSession(db *sql.DB, user_id int, session_id string, expires_at time.Time) error {
	task := `INSERT OR REPLACE INTO sessions (user_id,session_id,expires_at) VALUES (?,?,?)`

	_, err := db.Exec(task, user_id, session_id, expires_at)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

// ////////////////////////////////////////////////////////
func AddUser(db *sql.DB, email, username, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	task := `INSERT INTO users (email,username,password) VALUES (?,?,?)`
	result, err := db.Exec(task, email, username, hashedPassword)
	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}

	userID, _ := result.LastInsertId()

	return userID, nil
}
