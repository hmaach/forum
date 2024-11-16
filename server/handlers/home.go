package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"forum/server/database/queries"
	"forum/server/utils"
)

func GetHome(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" {
		utils.RenderError(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	posts, err := queries.GetPosts(db)
	if err != nil {
		log.Println("Error fetching posts from the database:", err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	err = utils.RenderTemplate(w, r, "home", http.StatusOK, posts)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}
