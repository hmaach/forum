// File: server/handlers/assets.go
package handlers

import (
	"net/http"
	"os"
	"strings"
)

// ServeStaticFiles returns a handler function for serving static files
func ServeStaticFiles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Block direct access to the /assets/ directory
		if r.URL.Path == "/assets/" || strings.HasSuffix(r.URL.Path, "/") {
		    RenderError(w, http.StatusForbidden, "403 | Access to this resource is forbidden!")
		    return
		}

		// Clean the file path
		filePath := "../web/assets" + strings.TrimPrefix(r.URL.Path, "/assets")

		// Check if file exists
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			RenderError(w, http.StatusNotFound, "404 | Page Not Found")
			return
		}

		// Serve the file
		http.ServeFile(w, r, filePath)
	}
}

// RenderError handles error responses
func RenderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
