package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"forum/server/utils"
)

func GetRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}
	err := utils.RenderTemplate(w, r, "register", http.StatusOK, nil)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func Register(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}
}
