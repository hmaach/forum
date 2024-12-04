package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum/server/models"
	"forum/server/utils"
	"forum/server/validators"
)

func CreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, username, valid, content, postid, user_id := validators.CreateComment_Request(r, db)
	if statuscode != http.StatusOK {
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}
	if !valid {
		w.WriteHeader(401)
		return
	}

	comm_id, err := models.AddComment(db, user_id, postid, content)
	if err != nil {
		http.Error(w, "Cannot add comment, try again!", http.StatusBadRequest)
		return
	}

	var commentscount int
	err2 := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = ?", postid).Scan(&commentscount)
	if err2 != nil {
		fmt.Println(err)
		utils.RenderError(db, w, r, 500, valid, username)
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
