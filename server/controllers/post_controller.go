package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"forum/server/models"
	"forum/server/utils"
)

func HandlePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RenderError(w, r, http.StatusBadRequest)
		return
	}
	post, statusCode, err := models.FetchPost(db, postID)
	if err != nil {
		log.Println("Error fetching posts from the database:", err)
		utils.RenderError(w, r, statusCode)
		return
	}

	err = utils.RenderTemplate(w, r, "post", statusCode, post)
	if err != nil {
		log.Println(err)
		utils.RenderError(w, r, http.StatusInternalServerError)
	}
}
