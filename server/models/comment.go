package models

import (
	"database/sql"
)

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	UserName  string
	Content   string
	Likes     int
	Dislikes  int
	CreatedAt string
}

func FetchCommentsByPostID(postID int, db *sql.DB) ([]Comment, error) {
	var comments []Comment
	query := `
	SELECT
		c.id,
		c.user_id,
		u.username,
		c.content,
		c.created_at,
		(
			SELECT
				COUNT(*)
			FROM
				comment_reactions AS cr
			WHERE
				cr.comment_id = c.id
				AND cr.reaction = 'like'
		) AS likes_count,
		(
			SELECT
				COUNT(*)
			FROM
				comment_reactions AS cr
			WHERE
				cr.comment_id = c.id
				AND cr.reaction = 'dislike'
		) AS dislikes_count
	FROM
		comments c
	INNER JOIN users u 
	ON c.user_id = u.id
	WHERE
		c.post_id = ?
	ORDER BY
		c.created_at DESC
	`

	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.UserName,
			&comment.Content,
			&comment.CreatedAt,
			&comment.Likes,
			&comment.Dislikes,
		)
		if err != nil {
			return nil, err
		}

		// Assign the post ID and format the created_at field
		comment.PostID = postID
		// comment.CreatedAt = utils.FormatTime(comment.CreatedAt)

		// Append the comment to the slice
		comments = append(comments, comment)
	}

	return comments, nil
}
