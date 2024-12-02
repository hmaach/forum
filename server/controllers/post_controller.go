package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"forum/server/models"

	"forum/server/requests"

	"forum/server/utils"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statucode, er := requests.IndexPosts_Request(r, db)
	if statucode != http.StatusOK {
		fmt.Println(er)
		utils.RenderError(db, w, r, http.StatusNotFound, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}
	// fmt.Println(statucode)
	posts, statusCode, err := models.FetchPosts(db)
	if err != nil {
		log.Println("Error fetching posts:", err)
		utils.RenderError(db, w, r, statusCode, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}
	// log.Println("hi", statusCode)
	if db == nil {
		// http.Error(w, "Status  code: 500 | Internal Server Error", http.StatusInternalServerError)
		utils.RenderError(db, w, r, http.StatusInternalServerError, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
	}
	if err := utils.RenderTemplate(db, w, r, "home", statusCode, posts, requests.RequestPost.Isvalid, requests.RequestPost.UserName); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
	}
}

func IndexPostsByCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statucode, er := requests.IndexPosts_Request(r, db)

	if statucode != http.StatusOK {
		fmt.Println(er)
		utils.RenderError(db, w, r, http.StatusNotFound, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}

	posts, statusCode, err := models.FetchPostsByCategory(db, requests.RequestPost.CategorieId)
	if err != nil {
		log.Println("Error fetching posts:", err)
		utils.RenderError(db, w, r, statusCode, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}

	if err := utils.RenderTemplate(db, w, r, "home", statusCode, posts, requests.RequestPost.Isvalid, requests.RequestPost.UserName); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
	}
}

func ShowPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statucode, er := requests.ShowPost_Request(r, db)

	if statucode != http.StatusOK {
		fmt.Println(er)
		utils.RenderError(db, w, r, http.StatusNotFound, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}
	post, statusCode, err := models.FetchPost(db, requests.RequestPost.PostId)
	if err != nil {
		log.Println("Error fetching posts from the database:", err)
		utils.RenderError(db, w, r, statusCode, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}

	err = utils.RenderTemplate(db, w, r, "post", statusCode, post, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
	if err != nil {
		log.Println(err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
	}
}

func GetPostCreationForm(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statucode, er := requests.GetPostCreationForm_Request(r, db)

	if statucode != http.StatusOK {
		fmt.Println(er)
		utils.RenderError(db, w, r, http.StatusNotFound, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}

	if !requests.RequestPost.Isvalid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := utils.RenderTemplate(db, w, r, "post-form", http.StatusOK, nil, requests.RequestPost.Isvalid, requests.RequestPost.UserName); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statucode, er := requests.CreatePost_Request(r, db)
	if statucode != http.StatusOK {
		fmt.Println(er)
		utils.RenderError(db, w, r, http.StatusNotFound, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}

	if !requests.RequestPost.Isvalid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	title := requests.RequestPost.Title
	content := requests.RequestPost.Content
	catids := requests.RequestPost.Categorie
	user_id := requests.RequestPost.UserID

	pid, err := models.StorePost(db, user_id, title, content)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Cannot create post, try again", http.StatusBadRequest)
		utils.RenderError(db, w, r, http.StatusBadRequest, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
		return
	}

	for i := 0; i < len(catids); i++ {
		catid, err := strconv.Atoi(catids[i])
		if err != nil {
			// http.Error(w, "Internal server error", 500)
			utils.RenderError(db, w, r, http.StatusInternalServerError, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
			return
		}
		_, err = models.StorePostCategory(db, pid, catid)
		if err != nil {
			fmt.Println(err)
			// http.Error(w, "Cannot create post, try again", http.StatusBadRequest)
			utils.RenderError(db, w, r, http.StatusBadRequest, requests.RequestPost.Isvalid, requests.RequestPost.UserName)
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
