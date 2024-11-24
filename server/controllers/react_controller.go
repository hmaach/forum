package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/server/common"
	"net/http"
	"strconv"
)

func ReactToPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user_id int
	var valid bool

	if user_id, valid = ValidSession(r, db); !valid {
		w.WriteHeader(401)
		common.IsAuthenticated = false
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	userReaction := r.FormValue("reaction")
	id := r.FormValue("post_id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid form data!", http.StatusBadRequest)
		return
	}

	var dbreaction string
	db.QueryRow("SELECT reaction FROM post_reactions WHERE user_id=? AND post_id=?", user_id, post_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = AddPostReaction(db, user_id, post_id, userReaction)
	} else {
		query := "UPDATE post_reactions SET reaction = ? WHERE user_id = ? AND post_id = ?"
		_, err = db.Exec(query, userReaction, user_id, post_id)
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
}

func AddPostReaction(db *sql.DB, user_id, post_id int, reaction string) (int64, error) {
	task := `INSERT INTO post_reactions (user_id,post_id,reaction) VALUES (?,?,?)`
	result, err := db.Exec(task, user_id, post_id, reaction)
	if err != nil {
		return 0, fmt.Errorf("error inserting reaction data -> ")
	}
	preactionID, _ := result.LastInsertId()

	return preactionID, nil
}
