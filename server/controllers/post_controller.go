package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"forum/server/models"
	"forum/server/utils"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusNotFound)
		return
	}
	var (
		posts      []models.Post
		statusCode int
		err        error
	)

	categoryID := r.FormValue("category_id")

	if categoryID == "" {
		posts, statusCode, err = models.FetchPosts(db)
	} else {
		id, errConv := strconv.Atoi(categoryID)
		if errConv != nil {
			utils.RenderError(w, r, http.StatusBadRequest)
		}
		posts, statusCode, err = models.FetchPostsByCategory(db, id)
	}

	if err != nil {
		log.Println("Error fetching posts:", err)
		utils.RenderError(w, r, statusCode)
		return
	}

	if err := utils.RenderTemplate(w, r, "home", statusCode, posts); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(w, r, http.StatusInternalServerError)
	}
}

func ShowPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}
}
