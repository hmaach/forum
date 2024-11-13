package utils

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"forum/server/common"
	"forum/server/models"
)

// RenderError handles error responses
func RenderError(w http.ResponseWriter, r *http.Request, statusCode int) {
	typeError := models.Error{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	}
	if err := RenderTemplate(w, r, "error", statusCode, typeError); err != nil {
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

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, statusCode int, data any) error {
	t, err := ParseTemplates(tmpl)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return err
	}

	// Limit categories to the first 6
	Categories := common.Categories
	limitedCategories := Categories
	if len(Categories) > 6 {
		limitedCategories = Categories[:6]
	}

	globalData := models.GlobalData{
		IsAuthenticated: common.IsAuthenticated,
		Data:            data,
		Categories:      limitedCategories,
	}
	w.WriteHeader(statusCode)
	// Execute the template with the provided data
	err = t.ExecuteTemplate(w, tmpl+".html", globalData)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
