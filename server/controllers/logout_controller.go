package controllers

import (
	"database/sql"
	"net/http"

	"forum/server/models"
)

func Logout(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if userID, _, valid := models.ValidSession(r, db); valid {
		// Use the new model function
		err := models.DeleteUserSession(db, userID)
		if err != nil {
			http.Error(w, "Error while logging out!", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}
