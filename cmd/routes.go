package main

import (
	"forum/server/handlers"
	"net/http"
)

func routes() http.Handler {
	
	mux := http.NewServeMux()

	mux.HandleFunc("/assets/", handlers.ServeStaticFiles)
	mux.HandleFunc("/", handlers.GetHome)

	return mux
}
