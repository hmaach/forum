package queries

import (
	"database/sql"

	"forum/server/models"
)

func FetchCategories(db *sql.DB) ([]models.Category, error) {
	var categories []models.Category
	query := `
		SELECT
			c.id,
			c.label,
			(
				SELECT
					COUNT(id)
				FROM
					post_category pc
				WHERE
					pc.category_id = c.id
			) as posts_count
		FROM categories c
		ORDER BY posts_count DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var category models.Category
		rows.Scan(&category.ID, &category.Label, &category.PostsCount)
		categories = append(categories, category)
	}
	return categories, nil
}
