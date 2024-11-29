package routes

import (
	"database/sql"
	"net/http"

	"forum/server/controllers"
)

func Routes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// serve static files
	mux.HandleFunc("/assets/", controllers.ServeStaticFiles)

	// routes for GET requests
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexPosts(w, r, db)
	})
	mux.HandleFunc("/user-posts", func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexPostsByUser(w, r, db)
	})
	mux.HandleFunc("/liked-posts", func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexPostsLikedByUser(w, r, db)
	})
	mux.HandleFunc("/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.IndexPostsByCategory(w, r, db)
	})
	mux.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.ShowPost(w, r, db)
	})
	mux.HandleFunc("/post/create", controllers.GetPostCreationForm)
	mux.HandleFunc("/login", controllers.GetLogin)
	mux.HandleFunc("/register", controllers.GetRegister)
	mux.HandleFunc("/500", controllers.InternalServerError)

	// routes for POST requests
	mux.HandleFunc("/post/store", func(w http.ResponseWriter, r *http.Request) {
		controllers.StorePost(w, r, db)
	})

	return mux
}
