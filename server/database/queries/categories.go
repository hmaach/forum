package queries

import (
	"database/sql"

	"forum/server/models"
)

func GetCategories(db *sql.DB) ([]models.Category, error) {
	var categories []models.Category
	rows, err := db.Query("SELECT id, label FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var category models.Category
		rows.Scan(&category.ID, &category.Label)
		categories = append(categories, category)
	}
	return categories, nil
}
