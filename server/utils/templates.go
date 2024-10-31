package utils

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"forum/server/models"
)

// RenderError handles error responses
func RenderError(w http.ResponseWriter, statusCode int) {
	typeError := models.Error{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	}
	if err := RenderTemplate(w, "error", statusCode, typeError); err != nil {
		http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, statusCode int, data any) error {
	// Set the status code before writing the response
	w.WriteHeader(statusCode)

	// Parse the template files
	t, err := template.ParseFiles(
		"../web/templates/partials/header.html",
		"../web/templates/partials/footer.html",
		"../web/templates/"+tmpl+".html",
	)
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return fmt.Errorf("error parsing template files: %w", err)
	}

	// Execute the template with the provided data
	err = t.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
