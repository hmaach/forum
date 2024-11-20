package main

import (
	"database/sql"
	"net/http"

	"forum/server/handlers"
)

func routes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// serve static files
	mux.HandleFunc("/assets/", handlers.ServeStaticFiles)

	// routes to get pages
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleHome(w, r, db)
	})
	mux.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlePost(w, r, db)
	})
	mux.HandleFunc("/login", handlers.GetLogin)
	mux.HandleFunc("/register", handlers.GetRegister)
	mux.HandleFunc("/500", handlers.InternalServerError)
	mux.HandleFunc("/about", handlers.GetAbout)

	return mux
}
