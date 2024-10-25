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
		log.Println(err)
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, statusCode int, data any) error {
	tmpl = "../web/templates/" + tmpl + ".html"
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
		return fmt.Errorf("template '%s' not found", tmpl)
	}
	w.WriteHeader(statusCode)
	return t.Execute(w, data)
}
