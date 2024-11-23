package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Post struct {
	ID            int
	UserID        int
	UserName      string
	Title         string
	Content       string
	CreatedAt     string
	Likes         int
	Dislikes      int
	Comments      int
	CategoriesStr string
	Categories    []string
}

type PostDetail struct {
	Post     Post
	Comments []Comment
}

func FetchPosts(db *sql.DB) ([]Post, int, error) {
	var posts []Post

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
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'like'
		) AS likes_count,
		(
			SELECT
				COUNT(*)
			FROM
				post_reactions AS pr
			WHERE
				pr.post_id = p.id
				AND pr.reaction = 'dislike'
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
		p.created_at DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, 500, err
	}
	defer rows.Close()

	// Iterate through the rows
	for rows.Next() {
		var post Post
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
			return nil, 500, err
		}
		// it came from the  database as "technology,sports...", so we need to split it
		post.Categories = strings.Split(post.CategoriesStr, ",")

		// Format the created_at field to a more readable format
		// post.CreatedAt = utils.FormatTime(post.CreatedAt)

		// Append the Post struct to the posts slice
		posts = append(posts, post)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, 500, err
	}

	return posts, 200, nil
}

func FetchPost(db *sql.DB, postID int) (PostDetail, int, error) {
	var post Post
	post.ID = postID

	// Query to fetch the post
	query := `SELECT
		p.user_id,
		u.username,
		p.title,
		p.content,
		p.created_at,
		(
			SELECT COUNT(*)
			FROM post_reactions AS pr
			WHERE pr.post_id = p.id
			AND pr.reaction = 'like'
		) AS likes_count,
		(
			SELECT COUNT(*)
			FROM post_reactions AS pr
			WHERE pr.post_id = p.id
			AND pr.reaction = 'dislike'
		) AS dislikes_count,
		(
			SELECT COUNT(*)
			FROM comments c
			WHERE c.post_id = p.id
		) AS comments_count,
		(
			SELECT GROUP_CONCAT(c.label)
			FROM categories c
			INNER JOIN post_category pc ON c.id = pc.category_id
			WHERE pc.post_id = p.id
		) AS categories
	FROM
		posts p
		INNER JOIN users u ON p.user_id = u.id
	WHERE p.id = ?`

	// Use QueryRow for a single result
	row := db.QueryRow(query, postID)

	// Scan the data into the Post struct
	err := row.Scan(
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
		if err == sql.ErrNoRows {
			return PostDetail{}, 404, fmt.Errorf("post not found: %w", err)
		}
		log.Println("Error scanning row:", err)
		return PostDetail{}, 500, err
	}

	// Process categories
	post.Categories = strings.Split(post.CategoriesStr, ",")

	// Format the created_at field
	// post.CreatedAt = utils.FormatTime(post.CreatedAt)
	comments, err := FetchCommentsByPostID(postID, db)
	if err != nil {
		log.Println("Error fetching comments from the database:", err)
	}

	return PostDetail{
		Post:     post,
		Comments: comments,
	}, 200, nil
}

func FetchPostsByCategory(db *sql.DB, categoryID int) ([]Post, int, error) {
	var posts []Post
	query := `
		SELECT
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
					post_reactions AS pr
				WHERE
					pr.post_id = p.id
					AND pr.reaction = 'like'
			) AS likes_count,
			(
				SELECT
					COUNT(*)
				FROM
					post_reactions AS pr
				WHERE
					pr.post_id = p.id
					AND pr.reaction = 'dislike'
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
			INNER JOIN post_category pc ON p.id = pc.post_id
		WHERE pc.category_id = ?
		ORDER BY
			p.created_at
	`
	rows, err := db.Query(query, categoryID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, 500, err
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
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
			return nil, 500, err
		}

		// it came from the  database as "technology,sports...", so we need to split it
		post.Categories = strings.Split(post.CategoriesStr, ",")

		// post.CreatedAt = utils.FormatTime(post.CreatedAt)

		posts = append(posts, post)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, 500, err
	}

	return posts, 200, nil
}
