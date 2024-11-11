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
	mux.HandleFunc("/", handlers.GetHome)
	mux.HandleFunc("/login", handlers.GetLogin)
	mux.HandleFunc("/about", handlers.GetAbout)
	mux.HandleFunc("/topics", handlers.GetTopics)
	mux.HandleFunc("/register", handlers.GetRegister)
	mux.HandleFunc("/500", handlers.InternalServerError)

	return mux
}
