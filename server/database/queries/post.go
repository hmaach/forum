package queries

import (
	"database/sql"
	"log"
	"strings"

	"forum/server/models"
	"forum/server/utils"
)

func GetPosts(db *sql.DB) ([]models.Post, error) {
	var posts []models.Post

	// Query to fetch posts
	query := `SELECT
		p.id,
		p.user_id,
		u.username,
		p.title,
		p.content,
		p.created_at,
		(
			SELECT
				COUNT(*)
			FROM
				posts_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.type = 'like'
		) AS likes_count,
		(
			SELECT
				COUNT(*)
			FROM
				posts_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.type = 'dislike'
		) AS dislikes_count,
		(
			SELECT
				COUNT(*)
			FROM
				comments c
			WHERE
				c.post_id = p.id
		) AS comments_count,
		(
			SELECT
				GROUP_CONCAT(c.label)
			FROM
				categories c
			INNER JOIN post_category pc ON c.id = pc.category_id
			WHERE
				pc.post_id = p.id
		) AS categories
	FROM
		posts p
		INNER JOIN users u ON p.user_id = u.id
	ORDER BY
		p.created_at;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var post models.Post
		// Scan the data into the Post struct
		err := rows.Scan(&post.ID,
			&post.UserID,
			&post.UserName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&post.CategoriesStr)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		// it came from the  database as "technology,sports...", so we need to split it
		post.Categories = strings.Split(post.CategoriesStr, ",")

		// Format the created_at field to a more readable format
		post.CreatedAt = utils.FormatTime(post.CreatedAt)

		// Append the Post struct to the posts slice
		posts = append(posts, post)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return posts, nil
}
