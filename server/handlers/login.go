package handlers

import (
	"log"
	"net/http"

	"forum/server/utils"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed)
		return
	}
	err := utils.RenderTemplate(w, "login", http.StatusOK, nil)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}
