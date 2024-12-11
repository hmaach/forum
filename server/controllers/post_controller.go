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
		return
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
		return
	}

	if err := utils.RenderTemplate(db, w, r, "home", statusCode, posts, valid, username); err != nil {
		log.Println("Error rendering template:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		return
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

	if user_id, _, valid = ValidSession(r, db); !valid {
		w.WriteHeader(401)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	catids := r.Form["categories"]

	catids = strings.Split(catids[0], ",")

	title = html.EscapeString(title)
	content = html.EscapeString(content)

	if catids == nil || strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		w.WriteHeader(400)
		return
	}

	var catidsInt []int
	for i := range catids {
		id, e := strconv.Atoi(catids[i])
		if e != nil {
			w.WriteHeader(400)
			return
		}
		catidsInt = append(catidsInt, id)
	}

	err := checkCategories(db, catidsInt)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	pid, err := AddPost(db, user_id, title, content)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	for i := 0; i < len(catidsInt); i++ {

		_, err = AddPostCat(db, pid, catidsInt[i])
		if err != nil {
			w.WriteHeader(400)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
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

func checkCategories(db *sql.DB, ids []int) error {
	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1]

	query := fmt.Sprintf(`
        SELECT id
        FROM categories
        WHERE id IN (%s);
    `, placeholders)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return err
		}
		count++
	}
	if count != len(ids) {
		return fmt.Errorf("categories does not exists in db")
	}

	return nil
}

func MyCreatedPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	var username string
	var user_id int
	if user_id, username, valid = ValidSession(r, db); !valid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method != http.MethodGet {
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
	posts, statusCode, err := models.FetchCreatedPostsByUser(db, user_id, page)
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

func MyLikedPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	var username string
	var user_id int
	if user_id, username, valid = ValidSession(r, db); !valid {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method != http.MethodGet {
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
	posts, statusCode, err := models.FetchLikedPostsByUser(db, user_id, page)
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
