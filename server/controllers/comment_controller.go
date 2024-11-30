package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/server/config"
	"forum/server/utils"
)

func CreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user_id int
	var valid bool
	var username string

	if user_id, username, valid = config.ValidSession(r, db); !valid {
		w.WriteHeader(401)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if r.Method != http.MethodPost {
		utils.RenderError(db,w, r, http.StatusMethodNotAllowed,valid,username)
		return
	}

	

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	content := r.FormValue("comment")
	id := r.FormValue("postid")
	postid, err := strconv.Atoi(id)
	if err != nil || strings.TrimSpace(content) == "" {
		w.WriteHeader(400)
		utils.RenderError(db,w, r, http.StatusBadRequest,valid,username)
		return
	}
	comm_id, err := AddComment(db, user_id, postid, content)
	if err != nil {
		http.Error(w, "Cannot add comment, try again!", http.StatusBadRequest)
		return
	}
	// http.Redirect(w, r, "/post/"+strconv.Itoa(postid), http.StatusFound)

	var commentscount int
	err2 := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = ?", postid).Scan(&commentscount)
	if err != nil || err2 != nil {
		fmt.Println(err)
		utils.RenderError(db,w, r, 500,valid,username)
		return
	}

	// Return the new count as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ID":            comm_id,
		"username":      username,
		"created_at":    time.Now().Format("15:04 02/01/2006"),
		"content":       content,
		"likes":         0,
		"dislikes":      0,
		"commentscount": commentscount,
	})
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
