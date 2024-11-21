package models

import "database/sql"

type Category struct {
	ID         int
	Label      string
	PostsCount int
}

func FetchCategories(db *sql.DB) ([]Category, error) {
	var categories []Category
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
		var category Category
		rows.Scan(&category.ID, &category.Label, &category.PostsCount)
		categories = append(categories, category)
	}
	return categories, nil
}
