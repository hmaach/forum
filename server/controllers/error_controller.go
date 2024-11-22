package controllers

import (
	"net/http"

	"forum/server/utils"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w,r, http.StatusMethodNotAllowed)
		return
	}
	utils.RenderError(w,r, http.StatusInternalServerError)
}
