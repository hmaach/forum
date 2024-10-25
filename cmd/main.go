// File: cmd/main.go
package main

import (
    "log"
    "net/http"

    "forum/server/handlers"
)

func main() {
    mux := http.NewServeMux()

    // Register static file handler
    mux.HandleFunc("/assets/", handlers.ServeStaticFiles())

    log.Println("Server starting on http://localhost:8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal("Server error:", err)
    }
}