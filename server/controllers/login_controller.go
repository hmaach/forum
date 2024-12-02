package controllers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"forum/server/models"
	"forum/server/utils"
	"forum/server/validators"

	"golang.org/x/crypto/bcrypt"
)

func GetLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, valid := validators.GetLogin_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, false, "")
		return
	}

	if valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err := utils.RenderTemplate(db, w, r, "login", http.StatusOK, nil, false, "")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func Signin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, valid, username, password := validators.Signin_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, false, "")
		return
	}
	if valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Retrieve user information from SQLite
	var passwordHash string
	var user_id int
	err := db.QueryRow("SELECT id,password FROM users WHERE username = ?", username).Scan(&user_id, &passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RenderError(db, w, r, http.StatusNotFound, false, "")
			return
		}
		utils.RenderError(db, w, r, http.StatusInternalServerError, false, "")
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	////////////////////////////////////////

	sessionID, err := generateSessionID()
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	err = models.AddSession(db, user_id, sessionID, time.Now().Add(10*time.Hour))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Set session ID as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Expires: time.Now().Add(10 * time.Hour),
		Path:    "/",
	})
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
	// w.Write([]byte("Logged in successfully"))
}

func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
