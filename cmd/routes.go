package main

import (
	"database/sql"
	"net/http"

	"forum/server/handlers"
)

func routes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/assets/", handlers.ServeStaticFiles)
	mux.HandleFunc("/", handlers.GetHome)
	mux.HandleFunc("/login", handlers.GetLogin)
	mux.HandleFunc("/about", handlers.GetAbout)
	mux.HandleFunc("/topics", handlers.GetTopics)
	mux.HandleFunc("/register", handlers.GetRegister)

	return mux
}
