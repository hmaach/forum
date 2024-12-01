package controllers

// func Pagination(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// 	var valid bool
// 	var username string
// 	_, username, valid = ValidSession(r, db)

// 	posts, statusCode, err := models.FetchPosts(db,currentPage)
// 	if err != nil {
// 		log.Println("Error fetching posts:", err)
// 		utils.RenderError(db, w, r, statusCode, valid, username)
// 		return
// 	}

// 	if err := utils.RenderTemplate(db, w, r, "home", statusCode, posts, valid, username); err != nil {
// 		log.Println("Error rendering template:", err)
// 		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
// 	}
// }
