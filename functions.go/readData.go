package forum

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id            int
	Email         string
	Username      string
	password_hash string
}

type Post struct {
	Id            int
	User_Name     string
	User_id       int
	Title         string
	Content       string
	Created_at    time.Time
	Category      string
	Comments      []Comment
	LikesCount    int
	DislikesCount int
	CommentsCount int
}

type Category struct {
	Id         int
	Name       string
	PostsCount int
}

func GetUserData(db *sql.DB, userid int) (*User, error) {
	task := `SELECT id,email,username,password_hash FROM users WHERE id=?`
	row := db.QueryRow(task, userid)

	var user User

	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.password_hash)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User Not found in db!")
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetPostData(db *sql.DB, postid int) (*Post, error) {
	task := `SELECT id,user_id,title,content,created_at,category_id FROM posts WHERE id=?`
	row := db.QueryRow(task, postid)

	var post Post

	err := row.Scan(&post.Id, &post.User_id, &post.Title, &post.Content, &post.Created_at, &post.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Post Not found in db!")
			return nil, err
		}
		return nil, err
	}
	err = db.QueryRow("SELECT username FROM users WHERE id=?", post.User_id).Scan(&post.User_Name)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetAllPosts(db *sql.DB) (*[]Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.User_id, &post.Title, &post.Content, &post.Created_at, &post.Category)
		if err != nil {
			return nil, err
		}
		err = db.QueryRow("SELECT username FROM users WHERE id=?", post.User_id).Scan(&post.User_Name)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
	}

	return &posts, nil
}

func GetCategories(db *sql.DB) (*[]Category, error) {
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.Id, &category.Name, &category.PostsCount)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	for i, j := 0, len(categories)-1; i < j; i, j = i+1, j-1 {
		categories[i], categories[j] = categories[j], categories[i]
	}

	return &categories, nil
}
