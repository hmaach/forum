package models

type Post struct {
	ID       int
	UserName string
	Content  string
	Time     string
	Likes    int
	Dislikes int
	Comments int
}
