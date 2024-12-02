package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/server/models"
	"forum/server/utils"
	"forum/server/validators"
	"net/http"
)

func ReactToPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, username, valid, user_id, post_id, userReaction := validators.ReactToPost_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}
	var dbreaction string
	var err error
	db.QueryRow("SELECT reaction FROM post_reactions WHERE user_id=? AND post_id=?", user_id, post_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = models.AddPostReaction(db, user_id, post_id, userReaction)
	} else {
		if userReaction == dbreaction {
			query := "DELETE FROM post_reactions WHERE user_id = ? AND post_id = ?"
			_, err1 := db.Exec(query, user_id, post_id)
			_ = err1
		} else {
			query := "UPDATE post_reactions SET reaction = ? WHERE user_id = ? AND post_id = ?"
			_, err = db.Exec(query, userReaction, user_id, post_id)
		}
	}

	if err != nil {
		http.Error(w, "failed to update reaction", http.StatusInternalServerError)
		return
	}

	// Fetch the new count of reactions for this post
	var likeCount, dislikeCount int
	db.QueryRow("SELECT COUNT(*) FROM post_reactions WHERE post_id=? AND reaction=?", post_id, "like").Scan(&likeCount)
	db.QueryRow("SELECT COUNT(*) FROM post_reactions WHERE post_id=? AND reaction=?", post_id, "dislike").Scan(&dislikeCount)

	// Return the new count as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"likesCount": likeCount, "dislikesCount": dislikeCount})
}

func ReactToComment(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statuscode, username, valid, user_id, comment_id, userReaction := validators.ReactToPost_Request(r, db)

	if statuscode != http.StatusOK {
		w.WriteHeader(statuscode)
		utils.RenderError(db, w, r, statuscode, valid, username)
		return
	}

	var dbreaction string
	var err error
	db.QueryRow("SELECT reaction FROM comment_reactions WHERE user_id=? AND comment_id=?", user_id, comment_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = models.AddCommentReaction(db, user_id, comment_id, userReaction)
	} else {
		query := "UPDATE comment_reactions SET reaction = ? WHERE user_id = ? AND comment_id = ?"
		_, err = db.Exec(query, userReaction, user_id, comment_id)
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to update reaction", http.StatusInternalServerError)
		return
	}

	// Fetch the new count of reactions for this post
	var likeCount, dislikeCount int
	db.QueryRow("SELECT COUNT(*) FROM comment_reactions WHERE comment_id=? AND reaction=?", comment_id, "like").Scan(&likeCount)
	db.QueryRow("SELECT COUNT(*) FROM comment_reactions WHERE comment_id=? AND reaction=?", comment_id, "dislike").Scan(&dislikeCount)

	// Return the new count as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"commentlikesCount": likeCount, "commentdislikesCount": dislikeCount})
}
