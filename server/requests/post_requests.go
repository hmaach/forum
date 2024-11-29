package requests

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"forum/server/common"
)

// CreatePostRequest processes the incoming POST request for creating a new post.
func CreatePostRequest(r *http.Request) (int, error) {
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}

	// Parse form values
	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("failed to parse form data: %v", err)
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"]

	// Validate the title
	if title == "" {
		return http.StatusBadRequest, fmt.Errorf("post's title cannot be empty")
	} else if len(title) > 100 {
		return http.StatusBadRequest, fmt.Errorf("post's title cannot be over 100 characters")
	}
	// Validate the content
	if content == "" {
		return http.StatusBadRequest, fmt.Errorf("post's content cannot be empty")
	} else if len(content) > 1000 {
		return http.StatusBadRequest, fmt.Errorf("post's content cannot be over 1000 characters")
	}

	// Validate categories
	if len(categories) == 0 {
		return http.StatusBadRequest, fmt.Errorf("at least one category must be selected")
	}
	var (
		categoriesIDs         []int
		databaseCategoriesIDs []int
	)

	for _, category := range common.Categories {
		databaseCategoriesIDs = append(databaseCategoriesIDs, category.ID)
	}

	for _, category := range categories {
		id, err := strconv.Atoi(category)
		if err != nil {
			return http.StatusBadRequest, fmt.Errorf("invalid category ID: '%s'", category)
		}
		if !slices.Contains(databaseCategoriesIDs, id) {
			return http.StatusBadRequest, fmt.Errorf("category not found: '%s'", category)
		}
		// Check if the ID is already in categoriesIDs (i.e., it's a duplicate)
		if slices.Contains(categoriesIDs, id) {
			return http.StatusBadRequest, fmt.Errorf("duplicate category ID: '%s'", category)
		}
		categoriesIDs = append(categoriesIDs, id)
	}

	// Return success
	return http.StatusOK, nil
}

func ShowPostRequest(r *http.Request) (int, error) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}
	id := r.PathValue("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid post ID %s", id)
	}
	return http.StatusOK, nil
}

func IndexPostsByCategoryRequest(r *http.Request) (int, error) {
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed")
	}
	// check categoryID if can be converted to int
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid category ID %s", idStr)
	}

	// check category if exists
	var categoriesIDs []int
	for _, category := range common.Categories {
		categoriesIDs = append(categoriesIDs, category.ID)
	}

	if !slices.Contains(categoriesIDs, id) {
		return http.StatusNotFound, fmt.Errorf("category not found: '%s'", idStr)
	}

	return http.StatusOK, nil
}
