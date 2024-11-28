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
		w.Write([]byte(`
	   <html>
	   <body>
		  <p>You are loged out successfully. Redirecting to the main page in 2 seconds...</p>
		  <script>
			 setTimeout(function() {
				window.location.href = "/";
			 }, 2000);
		  </script>
	   </body>
	   </html>
	`))
	} else {
		http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
		return
	}
}
