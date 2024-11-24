package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"forum/server/common"
	"forum/server/utils"
)

func CreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}

	var user_id int
	var valid bool

	if user_id,valid = ValidSession(r,db); !valid {
		common.IsAuthenticated = false
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	content := r.FormValue("comment")
	id := r.FormValue("postid")
	postid, err := strconv.Atoi(id)
	if err != nil {
		utils.RenderError(w, r, http.StatusBadRequest)
		return
	}
	_, err = AddComment(db, user_id, postid, content)
	if err != nil {
		http.Error(w, "Cannot add comment, try again!", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/post/"+strconv.Itoa(postid), http.StatusFound)

}

func AddComment(db *sql.DB, user_id, post_id int, content string) (int64, error) {
	task := `INSERT INTO comments (user_id,post_id,content) VALUES (?,?,?)`

	result, err := db.Exec(task, user_id, post_id, content)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	commentID, _ := result.LastInsertId()

	return commentID, nil
}
