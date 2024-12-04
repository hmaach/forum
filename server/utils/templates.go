package utils

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"forum/server/models"

)

// RenderError handles error responses
func RenderError(db *sql.DB, w http.ResponseWriter, r *http.Request, statusCode int, isauth bool, username string) {
	typeError := models.Error{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	}
	if err := RenderTemplate(db, w, r, "error", statusCode, typeError, isauth, username); err != nil {
		http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}

func ParseTemplates(tmpl string) (*template.Template, error) {
	// Parse the template files
	t, err := template.ParseFiles(
		"../web/templates/partials/header.html",
		"../web/templates/partials/footer.html",
		"../web/templates/partials/navbar.html",
		"../web/templates/"+tmpl+".html",
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing template files: %w", err)
	}
	return t, nil
}

func RenderTemplate(db *sql.DB, w http.ResponseWriter, r *http.Request, tmpl string, statusCode int, data any, isauth bool, username string) error {
	t, err := ParseTemplates(tmpl)
	if err != nil {
		return err
	}
	categories, err := models.FetchCategories(db)
	if err != nil {
		categories = nil
	}

	globalData := models.GlobalData{
		IsAuthenticated: isauth,
		Data:            data,
		UserName:        username,
		Categories:      categories,
	}
	w.WriteHeader(statusCode)
	// Execute the template with the provided data
	err = t.ExecuteTemplate(w, tmpl+".html", globalData)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
