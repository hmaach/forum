package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"forum/server/models"
	"forum/server/utils"

	"golang.org/x/crypto/bcrypt"
)

func GetLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool

	if _, _, valid = models.ValidSession(r, db); valid {
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
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func Signin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool

	if _, _, valid = models.ValidSession(r, db); valid {
		w.WriteHeader(302)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 4 || len(password) < 6 {
		w.WriteHeader(400)
		return
	}

	// Retrieve user information from SQLite
	var passwordHash string
	var user_id int
	err := db.QueryRow("SELECT id,password FROM users WHERE username = ?", username).Scan(&user_id, &passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		w.WriteHeader(401)
		return
	}
	////////////////////////////////////////

	sessionID, err := utils.GenerateSessionID()
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	err = models.StoreSession(db, user_id, sessionID, time.Now().Add(10*time.Hour))
	if err != nil {
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
