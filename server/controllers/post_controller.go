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
	fmt.Println(statucode)
	posts, statusCode, err := models.FetchPosts(db)
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

	pid, err := AddPost(db, user_id, title, content)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Cannot create post, try again", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(catids); i++ {
		catid, err := strconv.Atoi(catids[i])
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		_, err = AddPostCat(db, pid, catid)
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

func AddPost(db *sql.DB, user_id int, title, content string) (int64, error) {
	task := `INSERT INTO posts (user_id,title,content) VALUES (?,?,?)`

	result, err := db.Exec(task, user_id, title, content)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	postID, _ := result.LastInsertId()

	return postID, nil
}

func AddPostCat(db *sql.DB, post_id int64, category_id int) (int64, error) {
	task := `INSERT INTO post_category (post_id,category_id) VALUES (?,?)`

	result, err := db.Exec(task, post_id, category_id)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	postcatID, _ := result.LastInsertId()

	return postcatID, nil
}
