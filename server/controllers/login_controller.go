package controllers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
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
		return
	}
	if valid {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Retrieve user information from SQLite
	var passwordHash string
	var user_id int
	err := db.QueryRow("SELECT id,password FROM users WHERE username = ?", username).Scan(&user_id, &passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(401)
			//utils.RenderError(db, w, r, http.StatusNotFound, false, "")
			return
		}
		utils.RenderError(db, w, r, http.StatusInternalServerError, false, "")
		return
	}
	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		w.WriteHeader(401)
		//http.Error(w, "Invalid username or password", http.StatusUnauthorized)
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
	// Return the new count as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{})
	//http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
	// w.Write([]byte("Logged in successfully"))
}

func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
