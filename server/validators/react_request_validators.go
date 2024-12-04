package validators

import (
	"database/sql"
	"forum/server/config"
	"net/http"
	"strconv"
)

func ReactToPost_Request(r *http.Request, db *sql.DB) (int, string, bool, int, int,string) {
	userId, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, username, valid, 0,0,""
	}
	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, username, valid,0,0,""
	}

	reaction := r.FormValue("reaction")
	id := r.FormValue("post_id")
	post_id, err := strconv.Atoi(id)
	if err != nil{
		return http.StatusBadRequest, username, valid,0,0,""
	}
	return http.StatusOK, username, valid, userId,post_id,reaction
}

///////////////////////////////////////////////////////////////
func ReactToComment_Request(r *http.Request, db *sql.DB) (int, string, bool, int, int,string) {
	userId, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, username, valid, 0,0,""
	}
	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, username, valid,0,0,""
	}

	userReaction := r.FormValue("reaction")
	id := r.FormValue("comment_id")
	comment_id, err := strconv.Atoi(id)
	if err != nil{
		return http.StatusBadRequest, username, valid,0,0,""
	}
	return http.StatusOK, username, valid, userId,comment_id,userReaction
}