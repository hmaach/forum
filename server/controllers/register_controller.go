package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"forum/server/utils"

	"golang.org/x/crypto/bcrypt"
)

func GetRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	if _, _, valid = ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}

	err := utils.RenderTemplate(db, w, r, "register", http.StatusOK, nil, false, "")
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func Signup(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	if _, _, valid = ValidSession(r, db); valid {
		w.WriteHeader(302)
		return
	}

	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordConfirmation := r.FormValue("password-confirmation")

	if len(username) < 4 || len(password) < 6 || email == "" || password != passwordConfirmation {
		w.WriteHeader(400)
		return
	}

	_, err := AddUser(db, email, username, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			w.WriteHeader(304)
			return
		}

		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
}

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
