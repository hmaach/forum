package validators

import (
	"database/sql"
	"net/http"

	"forum/server/config"
)

func GetRegister_Request(r *http.Request, db *sql.DB) (int, string, bool) {
	_, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, username, valid
	}

	return http.StatusOK, username, valid
}

// /////////////////////////////////////////////////////////////
func Signup_Request(r *http.Request, db *sql.DB) (int, string, bool, string, string, string) {
	_, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, username, valid, "", "", ""
	}
	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, username, valid, "", "", ""
	}

	email := r.FormValue("email")
	newUserName := r.FormValue("username")
	password := r.FormValue("password")
	passwordConfirmation := r.FormValue("password-confirmation")

	if len(newUserName) < 4 || len(password) < 6 || email == "" || password != passwordConfirmation {
		return http.StatusBadRequest, username, valid, "", "", ""
	}
	return http.StatusOK, username, valid, email, newUserName, password
}
