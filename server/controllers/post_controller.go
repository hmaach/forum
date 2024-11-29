package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"forum/server/models"
	"forum/server/requests"
	"forum/server/utils"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusNotFound)
		return
	}

	posts, statusCode, err := models.FetchPosts(db)
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

func IndexPostsByCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	status, err := requests.IndexPostsByCategoryRequest(r)
	if status != http.StatusOK {
		log.Println("Error creating post:", err)
		utils.RenderError(w, r, status)
	}

	id, _ := strconv.Atoi(r.PathValue("id"))

	posts, statusCode, err := models.FetchPostsByCategory(db, id)
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
	status, err := requests.CreatePostRequest(r)
	if status != http.StatusOK {
		log.Println("Error creating post:", err)
		utils.RenderError(w, r, status)
	}

	postID, _ := strconv.Atoi(r.PathValue("id"))
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

func GetPostCreationForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	if err := utils.RenderTemplate(w, r, "post-form", http.StatusOK, nil); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(w, r, http.StatusInternalServerError)
	}
}

func StorePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	status, err := requests.CreatePostRequest(r)
	if status != http.StatusOK {
		log.Println("Error creating post:", err)
		utils.RenderError(w, r, status)
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"]

	log.Println("Title:", title)
	log.Println("Content:", content)
	log.Println("Categories:", categories)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
