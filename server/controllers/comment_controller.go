package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user_id int
	var valid bool
	var username string

	if user_id, username, valid = ValidSession(r, db); !valid {
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
	content := r.FormValue("comment")
	id := r.FormValue("postid")
	postid, err := strconv.Atoi(id)
	if err != nil || strings.TrimSpace(content) == "" {
		w.WriteHeader(400)
		return
	}
	comm_id, err := AddComment(db, user_id, postid, content)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var commentscount int
	err2 := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = ?", postid).Scan(&commentscount)
	if err2 != nil {
		w.WriteHeader(500)
		return
	}

	var commentTime string
	err2 = db.QueryRow("SELECT strftime('%m/%d/%Y %I:%M %p', created_at) AS formatted_created_at FROM comments WHERE id = ?", comm_id).Scan(&commentTime)
	if err2 != nil {
		w.WriteHeader(500)
		return
	}

	// Return the new count as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ID":            comm_id,
		"username":      username,
		"created_at":    commentTime,
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
