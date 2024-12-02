package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"forum/server/models"
	"forum/server/utils"
	"forum/server/validators"
)

func GetRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, username, valid := validators.GetRegister_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}
	if valid {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	err := utils.RenderTemplate(db, w, r, "register", http.StatusOK, nil, false, "")
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func Signup(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, username, valid, email, newUserName, password := validators.Signup_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}

	_, err := models.AddUser(db, email, newUserName, password)
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
