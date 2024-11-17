package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"forum/server/utils"
)

func GetCategories(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, r, http.StatusMethodNotAllowed)
		return
	}
	err := utils.RenderTemplate(w, r, "categories", http.StatusOK, nil)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}
