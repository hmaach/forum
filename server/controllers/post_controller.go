package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"forum/server/common"
	"forum/server/models"
	"forum/server/utils"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusNotFound)
		return
	}

	var user_id int
	var valid bool
	if user_id,valid = ValidSession(r,db); valid {
		common.IsAuthenticated = true
		db.QueryRow("SELECT username FROM users WHERE id = ?",user_id).Scan(&common.UserName)
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
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RenderError(w, r, http.StatusBadRequest)
	}

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

func GetPostCreationForm(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	var valid bool

	if _, valid = ValidSession(r, db); !valid {
		common.IsAuthenticated = false
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := utils.RenderTemplate(w, r, "post-form", http.StatusOK, nil); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(w, r, http.StatusInternalServerError)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	var user_id int
	var valid bool

	if user_id, valid = ValidSession(r, db); !valid {
		common.IsAuthenticated = false
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	catids := r.Form["categories[]"]

	if catids == nil || title == "" || content == "" {
		http.Error(w, "Please verify your entries and try again!", http.StatusBadRequest)
		return
	}

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
