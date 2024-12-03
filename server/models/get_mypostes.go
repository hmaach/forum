package models

import (
	"database/sql"
)

func getmyposts(db *sql.DB) {
	query := "SELECT id FROM posts WHERE user_id = ?"

	// rows, err := db.Query(query, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	var postIDs []int

	for rows.Next() {
		var postID int
		if err := rows.Scan(&postID); err != nil {
			return
		}
		postIDs = append(postIDs, postID)
	}

	// Check for any error during iteration
	if err = rows.Err(); err != nil {
		return
	}

	return postIDs
}
