package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"forum/server/models"
	"forum/server/utils"
)

func ReactToPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user_id int
	var valid bool

	if user_id, _, valid = models.ValidSession(r, db); !valid {
		w.WriteHeader(401)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		return
	}

	userReaction := r.FormValue("reaction")
	id := r.FormValue("post_id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var dbreaction string
	db.QueryRow("SELECT reaction FROM post_reactions WHERE user_id=? AND post_id=?", user_id, post_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = models.StorePostReaction(db, user_id, post_id, userReaction)
	} else {
		if userReaction == dbreaction {
			query := "DELETE FROM post_reactions WHERE user_id = ? AND post_id = ?"
			_, err = db.Exec(query, user_id, post_id)
		} else {
			query := "UPDATE post_reactions SET reaction = ? WHERE user_id = ? AND post_id = ?"
			_, err = db.Exec(query, userReaction, user_id, post_id)
		}
	}

	if err != nil {
		w.WriteHeader(500)
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user_id int
	var valid bool

	if user_id, _, valid = models.ValidSession(r, db); !valid {
		w.WriteHeader(401)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(400)
		return
	}

	userReaction := r.FormValue("reaction")
	id := r.FormValue("comment_id")
	comment_id, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var dbreaction string
	db.QueryRow("SELECT reaction FROM comment_reactions WHERE user_id=? AND comment_id=?", user_id, comment_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = models.StoreCommentReaction(db, user_id, comment_id, userReaction)
	} else {
		if userReaction == dbreaction {
			query := "DELETE FROM comment_reactions WHERE user_id = ? AND comment_id = ?"
			_, err = db.Exec(query, user_id, comment_id)

		} else {
			query := "UPDATE comment_reactions SET reaction = ? WHERE user_id = ? AND comment_id = ?"
			_, err = db.Exec(query, userReaction, user_id, comment_id)
		}
	}

	if err != nil {
		w.WriteHeader(500)
		return
	}

	// Fetch the new count of reactions for this post
	likeCount, dislikeCount, err := models.FetchCountReactionsOfPost(db, comment_id)
	if err != nil {
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, "")
		return
	}

	// Return the new count as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"commentlikesCount": likeCount, "commentdislikesCount": dislikeCount})
}
