package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"forum/server/config"
	"forum/server/models"
	"forum/server/utils"
)

func GetRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	if _, _, valid = config.ValidSession(r, db); valid {
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
	if _, _, valid = config.ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, false, "")
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordConfirmation := r.FormValue("password-confirmation")

	if len(username) < 4 || len(password) < 6 || email == "" || password != passwordConfirmation {
		http.Error(w, "Please verify your data and try again!", http.StatusBadRequest)
		return
	}

	_, err := models.StoreUser(db, email, username, password)
	if err != nil {
		w.Write([]byte("Cannot create user, try again later!"))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
	   <html>
	   <body>
		  <p>User ` + username + ` has been created successfully. Redirecting to the login page in 2 seconds...</p>
		  <script>
			 setTimeout(function() {
				window.location.href = "/login";
			 }, 2000);
		  </script>
	   </body>
	   </html>
	`))
}
