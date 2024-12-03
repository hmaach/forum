package validators

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"forum/server/config"
)

func CreateComment_Request(r *http.Request, db *sql.DB) (int, string, bool, string, int, int) {
	user_id, username, valid := config.ValidSession(r, db)

	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, username, valid, "", 0, 0
	}
	if err := r.ParseForm(); err != nil {
		return http.StatusBadRequest, username, valid, "", 0, 0
	}
	content := r.FormValue("comment")
	idstr := r.FormValue("postid")
	postId, err := strconv.Atoi(idstr)
	if err != nil || (strings.TrimSpace(content) == "" && valid) {
		return http.StatusBadRequest, username, valid, "", 0, 0
	}
	return http.StatusOK, username, valid, content, postId, user_id
}
