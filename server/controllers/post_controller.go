package controllers

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"forum/server/models"
	"forum/server/utils"
)

func IndexPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	var username string
	_, username, valid = ValidSession(r, db)

	if r.URL.Path != "/" || r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusNotFound, valid, username)
		return
	}
	id := r.FormValue("PageID")
	page, er := strconv.Atoi(id)
	if er != nil && id != "" {
		utils.RenderError(db, w, r, http.StatusBadRequest, valid, username)
		return
	}
	page = (page - 1) * 10
	if page < 0 {
		page = 0
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
	var valid bool
	var username string
	_, username, valid = ValidSession(r, db)

	if r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, valid, username)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RenderError(db, w, r, http.StatusBadRequest, valid, username)
	}

	pid := r.FormValue("PageID")
	page, _ := strconv.Atoi(pid)
	page = (page - 1) * 10
	if page < 0 {
		page = 0
	}

	posts, statusCode, err := models.FetchPostsByCategory(db, id, page)
	if err != nil {
		log.Println("Error fetching posts:", err)
		utils.RenderError(db, w, r, statusCode, valid, username)
		return
	}

	if posts == nil && page > 0 {
		utils.RenderError(db, w, r, 404, valid, username)
	}

	if err := utils.RenderTemplate(db, w, r, "home", statusCode, posts, valid, username); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
	}
}

func ShowPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	var username string
	_, username, valid = ValidSession(r, db)

	if r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, valid, username)
		return
	}
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RenderError(db, w, r, http.StatusBadRequest, valid, username)
		return
	}
	post, statusCode, err := models.FetchPost(db, postID)
	if err != nil {
		log.Println("Error fetching posts from the database:", err)
		utils.RenderError(db, w, r, statusCode, valid, username)
		return
	}

	err = utils.RenderTemplate(db, w, r, "post", statusCode, post, valid, username)
	if err != nil {
		log.Println(err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
	}
}

func GetPostCreationForm(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	var username string

	if _, username, valid = ValidSession(r, db); !valid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, valid, username)
		return
	}

	if err := utils.RenderTemplate(db, w, r, "post-form", http.StatusOK, nil, valid, username); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		return
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user_id int
	var valid bool
	var username string

	if user_id, username, valid = ValidSession(r, db); !valid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method != http.MethodPost {
		utils.RenderError(db, w, r, http.StatusMethodNotAllowed, valid, username)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	catids := r.Form["categories"]

	title = html.EscapeString(title)
	content = html.EscapeString(content)

	if catids == nil || strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
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
