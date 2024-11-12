package forum

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	UserID    int
	ExpiresAt time.Time
}

var sessions = make(map[string]Session) // In-memory session store

func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	////////////////////////////////////////
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
	err := Db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&passwordHash)
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

	var user_id int
	err = Db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&user_id)
	_ = err
	// Create a session and store it
	sessions[sessionID] = Session{
		UserID:    user_id,
		ExpiresAt: time.Now().Add(10 * time.Hour), // Session expiration time
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

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil || cookie == nil {
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
			return
		}

		session, exists := sessions[cookie.Value]
		if !exists || session.ExpiresAt.Before(time.Now()) {
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
