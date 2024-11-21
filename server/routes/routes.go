package routes

import (
	"database/sql"
	"forum/server/controllers"
	"net/http"
)

func Routes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// serve static files
	mux.HandleFunc("/assets/", controllers.ServeStaticFiles)

	// routes to get pages
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandleHome(w, r, db)
	})
	mux.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.HandlePost(w, r, db)
	})
	mux.HandleFunc("/login", controllers.GetLogin)
	mux.HandleFunc("/register", controllers.GetRegister)
	mux.HandleFunc("/500", controllers.InternalServerError)
	mux.HandleFunc("/about", controllers.GetAbout)

	return mux
}
