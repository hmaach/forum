package validators

import (
	"database/sql"
	"net/http"
	"strconv"

	"forum/server/config"
)

func IndexPosts_Request(r *http.Request, db *sql.DB) (int, string, bool, int) {
	_, username, valid := config.ValidSession(r, db)
	if r.URL.Path != "/" {
		return http.StatusNotFound, username, valid, 0
	}

	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, username, valid, 0
	}

	id := r.FormValue("PageID")
	page, er := strconv.Atoi(id)
	if er != nil && id != "" {
		return http.StatusBadRequest, username, valid, 0
	}
	page = (page - 1) * 10
	if page < 0 {
		page = 0
	}
	return http.StatusOK, username, valid, page
}

// ////////////////////////////////////////////////////////////////
func IndexPostsByCategory_Request(r *http.Request, db *sql.DB) (int, string, bool, int, int) {
	_, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, username, valid, 0, 0
	}

	// check categoryID if can be converted to int
	idStr := r.PathValue("id")
	categorieId, err := strconv.Atoi(idStr)
	if err != nil {
		return http.StatusBadRequest, username, valid, 0, 0
	}

	page := r.FormValue("PageID")
	pageId, er := strconv.Atoi(page)

	if er != nil && page != "" {
		return http.StatusBadRequest, username, valid, 0, 0
	}

	pageId = (pageId - 1) * 10
	if pageId < 0 {
		pageId = 0
	}
	return http.StatusOK, username, valid, categorieId, pageId
}

// /////////////////////////////////////////////////////////////////////////////
func ShowPost_Request(r *http.Request, db *sql.DB) (int, string, bool, int) {
	_, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, username, valid, 0
	}

	// check categoryID if can be converted to int
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return http.StatusBadRequest, username, valid, 0
	}

	return http.StatusOK, username, valid, id
}

// /////////////////////////////////////////////////////////////////////////////
func GetPostCreationForm_Request(r *http.Request, db *sql.DB) (int, string, bool) {
	_, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, username, valid
	}

	return http.StatusOK, username, valid
}

// /////////////////////////////////////////////////////////////////////////////
func CreatePost_Request(r *http.Request, db *sql.DB) (int, string, bool, int, string, string, []string) {
	userid, username, valid := config.ValidSession(r, db)
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, username, valid, userid, "", "", nil
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"]

	// Validate the title
	if title == "" {
		return http.StatusBadRequest, username, valid, userid, "", "", nil
	} else if len(title) > 100 {
		return http.StatusBadRequest, username, valid, userid, "", "", nil
	}

	// Validate the content
	if content == "" {
		return http.StatusBadRequest, username, valid, userid, "", "", nil
	} else if len(content) > 1000 {
		return http.StatusBadRequest, username, valid, userid, "", "", nil
	}

	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, username, valid, userid, "", "", nil
	}

	return http.StatusOK, username, valid, userid, title, content, categories
}
