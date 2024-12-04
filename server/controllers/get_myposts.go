package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"forum/server/config"
	"forum/server/models"
	"forum/server/utils"
)

func GetuserPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var valid bool
	var username string
	_, username, valid = config.ValidSession(r, db)

	if r.URL.Path != "/user-posts" || r.Method != http.MethodGet {
		utils.RenderError(db, w, r, http.StatusNotFound, valid, username)
		return
	}

	id := r.FormValue("PageID")
	page, err := strconv.Atoi(id)
	if err != nil && id != "" {
		utils.RenderError(db, w, r, http.StatusBadRequest, valid, username)
		return
	}

	offset := page * 10
	if offset < 0 {
		offset = 0
	}

	posts, statusCode, err := FetchuserPosts(db, username, offset)
	if err != nil {
		log.Println("Error fetching posts:", err)
		utils.RenderError(db, w, r, statusCode, valid, username)
		return
	}

	if len(posts) == 0 && page > 0 {
		utils.RenderError(db, w, r, http.StatusNotFound, valid, username)
		return
	}

	if err := utils.RenderTemplate(db, w, r, "home", http.StatusOK, posts, valid, username); err != nil {
		log.Println("Template render error:", err)
		utils.RenderError(db, w, r, http.StatusInternalServerError, valid, username)
		return
	}
}

func FetchuserPosts(db *sql.DB, username string, offset int) ([]models.Post, int, error) {
	if username == "" {
		return nil, 400, fmt.Errorf("invalid username")
	}

	query := `
    SELECT 
        p.id,
        p.user_id,
        u.username,
        p.title,
        p.content,
        p.created_at,
        (SELECT COUNT(*) FROM post_reactions pr WHERE pr.post_id = p.id AND pr.reaction = 'like') AS likes_count,
        (SELECT COUNT(*) FROM post_reactions pr WHERE pr.post_id = p.id AND pr.reaction = 'dislike') AS dislikes_count,
        (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS comments_count,
        (SELECT GROUP_CONCAT(c.label)
         FROM categories c
         INNER JOIN post_category pc ON c.id = pc.category_id
         WHERE pc.post_id = p.id) AS categories
    FROM 
        posts p
    INNER JOIN 
        users u ON p.user_id = u.id
    WHERE 
        u.username = ?
    ORDER BY 
        p.created_at DESC
    LIMIT 10 OFFSET ?`

	rows, err := db.Query(query, username, offset)
	if err != nil {
		return nil, 500, fmt.Errorf("database error: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		var categoriesStr string

		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.UserName,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.Comments,
			&categoriesStr,
		)
		if err != nil {
			return nil, 500, fmt.Errorf("scan error: %w", err)
		}

		if categoriesStr != "" {
			post.Categories = strings.Split(categoriesStr, ",")
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, 500, fmt.Errorf("rows error: %w", err)
	}

	if len(posts) == 0 {
		return nil, 404, fmt.Errorf("no posts found for user %s", username)
	}

	return posts, 200, nil
}