package forum

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(db *sql.DB, email, username, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	task := `INSERT INTO users (email,username,password_hash) VALUES (?,?,?)`
	result, err := db.Exec(task, email, username, hashedPassword)
	if err != nil {
		return -1, fmt.Errorf("error inserting user data -> ")
	}

	userID, _ := result.LastInsertId()

	return userID, nil
}

func AddPost(db *sql.DB, user_id int, title, content, category string, time time.Time) (int64, error) {
	task := `INSERT INTO posts (user_id,title,content,created_at,category_id) VALUES (?,?,?,?,?)`

	result, err := db.Exec(task, user_id, title, content, time, category)
	if err != nil {
		return 0, fmt.Errorf("error inserting post data -> ")
	}

	postID, _ := result.LastInsertId()

	return postID, nil
}

func AddComment(db *sql.DB, post_id, user_id int, content string) (int64, error) {
	task := `INSERT INTO comments (post_id,user_id,content) VALUES (?,?,?)`

	result, err := db.Exec(task, post_id, user_id, content)
	if err != nil {
		return 0, fmt.Errorf("error inserting post data -> ")
	}

	commentID, _ := result.LastInsertId()

	return commentID, nil
}

func AddCategory(db *sql.DB, name string, postsCount int) (int64, error) {
	task := `INSERT INTO categories (name,postscount) VALUES (?,?)`

	result, err := db.Exec(task, name, postsCount)
	if err != nil {
		return 0, fmt.Errorf("error inserting post data -> ")
	}

	categoryID, _ := result.LastInsertId()

	return categoryID, nil
}

func AddLike(db *sql.DB, user_id, post_id int) (int64, error) {
	task := `INSERT INTO Likes (user_id,post_id) VALUES (?,?)`
	result, err := db.Exec(task, user_id, post_id)
	if err != nil {
		return 0, fmt.Errorf("error inserting post data -> ")
	}
	likeID, _ := result.LastInsertId()

	return likeID, nil
}


func AddDislike(db *sql.DB, user_id string, post_id int) (int64, error) {
	task := `INSERT INTO Dislikes (user_id,post_id) VALUES (?,?)`
	result, err := db.Exec(task, user_id, post_id)
	if err != nil {
		return 0, fmt.Errorf("error inserting post data -> ")
	}
	likeID, _ := result.LastInsertId()

	return likeID, nil
}


func AddReaction(db *sql.DB, user_id, post_id, comment_id int, reaction string) (int64, error) {
	task := `INSERT INTO Reactions (user_id,post_id,comment_id,reaction) VALUES (?,?,?,?)`
	result, err := db.Exec(task, user_id, post_id,comment_id,reaction)
	if err != nil {
		return 0, fmt.Errorf("error inserting reaction data -> ")
	}
	reactionID, _ := result.LastInsertId()

	return reactionID, nil
}