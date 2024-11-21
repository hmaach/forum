package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"forum/server/models"
	"forum/server/utils"
)

func fetchPosts(r *http.Request, db *sql.DB) ([]models.Post, int, error) {
	categoryID := r.FormValue("category_id")
	if categoryID == "" {
		return models.FetchPosts(db)
	}

	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return models.FetchPostsByCategory(db, id)
}

func HandleHome(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusNotFound)
		return
	}

	posts, statusCode, err := fetchPosts(r, db)
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
