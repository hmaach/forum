package validators

import (
	"database/sql"
	"forum/server/config"
	"html"
	"net/http"
)

func GetLogin_Request(r *http.Request, db *sql.DB) (int, bool) {
	_, _, valid := config.ValidSession(r, db)

	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, valid
	}
	return http.StatusOK, valid
}

// /////////////////////////////////////////////////////////////////////////////
func Signin_Request(r *http.Request, db *sql.DB) (int, bool, string, string) {

	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, false, "", ""
	}
	_, _, valid := config.ValidSession(r, db)

	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, valid, "", ""
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	username = html.EscapeString(username)
	password = html.EscapeString(password)
	if len(username) < 4 || len(password) < 6 {
		return http.StatusBadRequest, valid, "", ""
	}
	return http.StatusOK, valid, username, password
}
