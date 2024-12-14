package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetUserInfo(db *sql.DB, username string) (int, string, error) {
	var user_id int
	var hashedPassword string
	err := db.QueryRow("SELECT id,password FROM users WHERE username = ?", username).Scan(&user_id, &hashedPassword)
	if err != nil {
		return 0, "", err
	}
	return user_id, hashedPassword, nil
}

func StoreUser(db *sql.DB, email, username, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	query := `INSERT INTO users (email,username,password) VALUES (?,?,?)`
	result, err := db.Exec(query, email, username, hashedPassword)
	if err != nil {
		return -1, fmt.Errorf("%v", err)
	}

	userID, _ := result.LastInsertId()

	return userID, nil
}
