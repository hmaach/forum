package requests

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"forum/server/config"
)

type Post_Request struct {
	UserID      int
	UserName    string
	Isvalid     bool
	Title       string
	Content     string
	Categorie   []string
	CategorieId int
	PostId      int
}

var RequestPost Post_Request

func IndexPosts_Request(r *http.Request, db *sql.DB) (int, error) {
	if r.URL.Path != "/" {
		return http.StatusBadRequest, fmt.Errorf("Page Not Found!")
	}
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}

	var valid bool
	var username string
	_, username, valid = config.ValidSession(r, db)

	RequestPost.UserName = username
	RequestPost.Isvalid = valid

	return http.StatusOK, nil
}

// ////////////////////////////////////////////////////////////////
func IndexPostsByCategory_Request(r *http.Request, db *sql.DB) (int, error) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}

	// check categoryID if can be converted to int
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid category ID %s", idStr)
	}

	var valid bool
	var username string
	_, username, valid = config.ValidSession(r, db)

	RequestPost.UserName = username
	RequestPost.Isvalid = valid
	RequestPost.CategorieId = id

	return http.StatusOK, nil
}

// /////////////////////////////////////////////////////////////////////////////
func ShowPost_Request(r *http.Request, db *sql.DB) (int, error) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}

	// check categoryID if can be converted to int
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid category ID %s", idStr)
	}

	var valid bool
	var username string
	_, username, valid = config.ValidSession(r, db)

	RequestPost.UserName = username
	RequestPost.Isvalid = valid
	RequestPost.PostId = id

	return http.StatusOK, nil
}

// /////////////////////////////////////////////////////////////////////////////
func GetPostCreationForm_Request(r *http.Request, db *sql.DB) (int, error) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}

	var valid bool
	var username string
	_, username, valid = config.ValidSession(r, db)

	RequestPost.UserName = username
	RequestPost.Isvalid = valid

	return http.StatusOK, nil
}

// /////////////////////////////////////////////////////////////////////////////
func CreatePost_Request(r *http.Request, db *sql.DB) (int, error) {
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("failed to parse form data: %v", err)
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"]

	// Validate the title
	if title == "" {
		return http.StatusBadRequest, fmt.Errorf("post's title cannot be empty")
	} else if len(title) > 100 {
		return http.StatusBadRequest, fmt.Errorf("post's title cannot be over 100 characters")
	}

	// Validate the content
	if content == "" {
		return http.StatusBadRequest, fmt.Errorf("post's content cannot be empty")
	} else if len(content) > 1000 {
		return http.StatusBadRequest, fmt.Errorf("post's content cannot be over 1000 characters")
	}

	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, fmt.Errorf("Invalid form data")
	}
	var valid bool
	var username string
	var userid int
	userid, username, valid = config.ValidSession(r, db)

	RequestPost.UserName = username
	RequestPost.Isvalid = valid
	RequestPost.Content = content
	RequestPost.Title = title
	RequestPost.Categorie = categories
	RequestPost.UserID = userid

	return http.StatusOK, nil
}
