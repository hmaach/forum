package controllers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"forum/server/common"
	"forum/server/utils"

	"golang.org/x/crypto/bcrypt"
)

func GetLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	var valid bool

	if _,valid = ValidSession(r,db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	
	err := utils.RenderTemplate(w, r, "login", http.StatusOK, nil)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
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
	if r.Method != http.MethodPost {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Retrieve user information from SQLite
	var passwordHash string
	var user_id int
	err := db.QueryRow("SELECT id,password FROM users WHERE username = ?", username).Scan(&user_id, &passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
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

	err = AddSession(db, user_id, sessionID, time.Now().Add(10*time.Hour))
	common.IsAuthenticated = true
	common.UserName = username
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

func ValidSession(r *http.Request, db *sql.DB) (int, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		return -1, false
	}
	var expiration time.Time
	var user_id int
	err = db.QueryRow("SELECT user_id,expires_at FROM sessions WHERE session_id = ?", cookie.Value).Scan(&user_id, &expiration)
	if err != nil || expiration.Before(time.Now()) {
		return -1, false
	}
	return user_id, true
}


func AddSession(db *sql.DB, user_id int, session_id string, expires_at time.Time ) error {
	task := `INSERT INTO sessions (user_id,session_id,expires_at) VALUES (?,?,?)`

	_, err := db.Exec(task, user_id, session_id,expires_at)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}