package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if user_id, _, valid := ValidSession(r, db); valid {
		_, err := db.Exec(`DELETE FROM sessions WHERE user_id = ?;`, user_id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error while loging out!", http.StatusSeeOther)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
		return
	} else {
		http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
		return
	}
}
