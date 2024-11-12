package forum

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func LikeDislike(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method)
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	userReaction := r.FormValue("reaction")
	id := r.FormValue("post_id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid form data!", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		return
	}

	session, exists := sessions[cookie.Value]
	if !exists || session.ExpiresAt.Before(time.Now()) {
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		return
	}
	user_id := session.UserID

	// var reacted int
	var dbreaction string
	Db.QueryRow("SELECT reaction FROM Reactions WHERE user_id=? AND post_id=?", user_id, post_id).Scan(&dbreaction)

	if dbreaction == "" {
		_, err = AddReaction(Db, user_id, post_id, 0, userReaction)
		if err != nil {
			http.Error(w,"failed to Add reaction",http.StatusInternalServerError)
			return
		}
	} else {
			query := "UPDATE Reactions SET reaction = ? WHERE user_id = ? AND post_id = ?"
			_, err := Db.Exec(query, userReaction, user_id, post_id)
			if err != nil {
				http.Error(w,"failed to update reaction",http.StatusInternalServerError)
				return
			}
	}

	http.Redirect(w,r,"/post?id="+strconv.Itoa(post_id),http.StatusFound)
}
