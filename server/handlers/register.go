package handlers

import (
	"log"
	"net/http"

	"forum/server/utils"
)

func GetRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed)
		return
	}
	err := utils.RenderTemplate(w, "register", http.StatusOK, nil)
	if err != nil {
		log.Println(err)
		utils.RenderError(w, http.StatusInternalServerError)
	}
}
