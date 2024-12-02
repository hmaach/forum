package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"forum/server/models"
	"forum/server/utils"
	"forum/server/validators"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, username, valid, page := validators.IndexPosts_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}

	posts, statusCode, err := models.FetchPosts(db, page)
	if err != nil {
		log.Println("Error fetching posts:", err)
		utils.RenderError(db, w, r, statusCode, valid, username)
		return
	}
	if posts == nil && page > 0 {
		utils.RenderError(db, w, r, 404, valid, username)
		return
	}

	if err := utils.RenderTemplate(db, w, r, "home", statusCode, posts, valid, username); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		return
	}
}

func IndexPostsByCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	statuscode, username, valid, categorieId, pageId := validators.IndexPostsByCategory_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}

	posts, statusCode, err := models.FetchPostsByCategory(db, categorieId, pageId)
	if err != nil {
		log.Println("Error fetching posts:", err)
		utils.RenderError(db, w, r, statusCode, valid, username)
		return
	}

	if err := utils.RenderTemplate(db, w, r, "home", statusCode, posts, valid, username); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		return
	}
}

func ShowPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, username, valid, postId := validators.ShowPost_Request(r, db)
	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}

	post, statusCode, err := models.FetchPost(db, postId)
	if err != nil {
		log.Println("Error fetching posts from the database:", err)
		utils.RenderError(db, w, r, statusCode, valid, username)
		return
	}

	err = utils.RenderTemplate(db, w, r, "post", statusCode, post, valid, username)
	if err != nil {
		log.Println(err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		return
	}
}

func GetPostCreationForm(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	statuscode, username, valid := validators.GetPostCreationForm_Request(r, db)
	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}
	if !valid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := utils.RenderTemplate(db, w, r, "post-form", http.StatusOK, nil, valid, username); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		return
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	statuscode, username, valid, userid, title, content, categories := validators.CreatePost_Request(r, db)
	if statuscode != http.StatusOK {
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}

	if !valid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	pid, err := models.AddPost(db, userid, title, content)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Cannot create post, try again", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(categories); i++ {
		catid, err := strconv.Atoi(categories[i])
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		_, err = models.AddPostCat(db, pid, catid)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Cannot create post, try again", http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
			   <html>
			   <body>
				  <p>Post created successfully. Redirecting to the main page in 2 seconds...</p>
				  <script>
					 setTimeout(function() {
						window.location.href = "/";
					 }, 2000);
				  </script>
			   </body>
			   </html>
			`))
}
