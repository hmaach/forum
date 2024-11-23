package controllers

import (
	"forum/server/utils"
	"log"
	"net/http"
)

func GetAbout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}
	err := utils.RenderTemplate(w, r, "about", http.StatusOK, nil)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}
