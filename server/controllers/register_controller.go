package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"forum/server/models"
	"forum/server/utils"
)

func GetRegisterPage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	if _, _, valid = models.ValidSession(r, db); valid {
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
		utils.RenderError(db, w, r, http.StatusInternalServerError, false, "")
		return
	}
}

func Signup(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordConfirmation := r.FormValue("password-confirmation")

	if len(strings.TrimSpace(username)) < 4 || len(strings.TrimSpace(password)) < 6 || email == "" || password != passwordConfirmation {
		w.WriteHeader(400)
		return
	}

	_, err := models.StoreUser(db, email, username, password)
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
