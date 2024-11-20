package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"forum/server/database/queries"
	"forum/server/utils"
)

func HandleHome(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" {
		utils.RenderError(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	posts, statusCode, err := queries.FetchPosts(db)
	if err != nil {
		log.Println("Error fetching posts from the database:", err)
		utils.RenderError(w, r, statusCode)
		return
	}

	err = utils.RenderTemplate(w, r, "home", statusCode, posts)
	if err != nil {
		log.Println(err)
		utils.RenderError(w, r, http.StatusInternalServerError)
	}
}
