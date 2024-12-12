package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"forum/server/models"
)

func CreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Validate session
	userID, username, valid := models.ValidSession(r, db)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Validate method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content := strings.TrimSpace(r.FormValue("comment"))
	postIDStr := r.FormValue("postid")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil || content == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Store the comment using the models package
	commentID, err := models.StoreComment(db, userID, postID, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Fetch additional details using the models package
	commentsCount, err := models.CountCommentsByPostID(db, postID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commentTime, err := models.FetchCommentTimeByID(db, commentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the new comment details as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ID":            commentID,
		"username":      username,
		"created_at":    commentTime,
		"content":       content,
		"likes":         0,
		"dislikes":      0,
		"commentscount": commentsCount,
	})
}
