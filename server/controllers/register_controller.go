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
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}
	var valid bool
	if _, valid = ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err := utils.RenderTemplate(w, r, "register", http.StatusOK, nil)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func Signup(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
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

	if username == "" || password == "" || email == "" || password != passwordConfirmation {
		http.Error(w, "Please verify your data and try again!", http.StatusBadRequest)
		return
	}

	_, err := AddUser(db, email, username, password)
	if err != nil {
		w.Write([]byte("Cannot create user, try again later!"))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
	   <html>
	   <body>
		  <p>User ` + username + ` has been created successfully. Redirecting to the login page in 5 seconds...</p>
		  <script>
			 setTimeout(function() {
				window.location.href = "/login";
			 }, 5000);
		  </script>
	   </body>
	   </html>
	`))
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
