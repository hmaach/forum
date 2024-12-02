package controllers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"forum/server/config"
	"forum/server/utils"

	"golang.org/x/crypto/bcrypt"
)

func GetLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool

	if _, _, valid = config.ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}

	err := utils.RenderTemplate(db, w, r, "login", http.StatusOK, nil, false, "")
	if err != nil {
		log.Println(err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, false, "")
	}
}

func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Signin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool

	if _, _, valid = config.ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}

	if err := r.ParseForm(); err != nil {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 4 || len(password) < 6 {
		utils.RenderError(db, w, r, http.StatusNotFound, false, "")
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
		// http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		utils.RenderError(db, w, r, http.StatusUnauthorized, false, "")
		return
	}
	////////////////////////////////////////

	sessionID, err := generateSessionID()
	if err != nil {
		// http.Error(w, "Failed to create session", http.StatusInternalServerError)
		utils.RenderError(db, w, r, http.StatusInternalServerError, false, "")
		return
	}

	err = config.AddSession(db, user_id, sessionID, time.Now().Add(10*time.Hour))
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Failed to create session", http.StatusInternalServerError)
		utils.RenderError(db, w, r, http.StatusInternalServerError, false, "")
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
