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
	if _, _,valid = ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}


	if r.Method != http.MethodGet {
		utils.RenderError(db,w, r, http.StatusMethodNotAllowed,false,"")
		return
	}
	

	err := utils.RenderTemplate(db,w, r, "register", http.StatusOK, nil,false,"")
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func Signup(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	if _, _,valid = ValidSession(r, db); valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method != http.MethodPost {
		utils.RenderError(db,w, r, http.StatusMethodNotAllowed,false,"")
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

	_, err := AddUser(db, email, username, password)
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
