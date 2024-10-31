package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"forum/server/utils"
)

// ServeStaticFiles returns a handler function for serving static files
func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Get clean file path and prevent directory traversal
	filePath := filepath.Clean("../web/assets" + strings.TrimPrefix(r.URL.Path, "/assets"))

	// block access to dirictories
	if info, err := os.Stat(filePath); err != nil || info.IsDir() {
		utils.RenderError(w, http.StatusNotFound)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
