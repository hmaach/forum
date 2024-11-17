package queries

import (
	"database/sql"

	"forum/server/models"
)

func GetCommentsByPostID(postID int, db *sql.DB) ([]models.Comment, error) {
	var comments []models.Comment
	query := `
		SELECT 
			id, 
			user_id, 
			content, 
			created_at 
		FROM comments 
		WHERE post_id = ?
		`
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		rows.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.CreatedAt)
		comment.PostID = postID
		comments = append(comments, comment)
	}

	return comments, nil
}
