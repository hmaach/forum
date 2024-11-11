package handlers

import (
	"net/http"

	"forum/server/utils"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed)
		return
	}
	utils.RenderError(w, http.StatusInternalServerError)
}
